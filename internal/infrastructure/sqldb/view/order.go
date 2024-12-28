package view

import "gorm.io/gorm"

const (
	TableOrders   string = "table_orders"
	ManagerOrders string = "manager_orders"
	PostOrders    string = "post_orders"
)

func TableOrdersView(db *gorm.DB) *gorm.DB {
	return db.Table("session_orders").
		Select(`
			sessions.id AS session_id,
			customers.id AS customer_id,
			customers.full_name AS customer_full_name,
			menu_items.name AS menu_item_name,
			SUM(session_orders.quantity) AS quantity,
			menu_items.price_amount AS unit_price,
			menu_items.price_amount * SUM(session_orders.quantity) AS total_price,
			menu_items.price_currency AS currency
		`).Joins("INNER JOIN sessions ON session_orders.session_id = sessions.id").
		Joins("INNER JOIN menu_items ON session_orders.menu_item_id = menu_items.id").
		Joins("INNER JOIN customers ON session_orders.ordered_by = customers.id").
		Group(`
			sessions.id,
			menu_items.name,
			menu_items.price_amount,
			menu_items.price_currency,
			customers.full_name,
			customers.id
		`).Order("customers.full_name")
}

func ManagerOrdersView(db *gorm.DB) *gorm.DB {
	return db.Table("session_orders").
		Select(`
			sessions.id AS session_id,
			menu_items.name AS menu_item_name,
			SUM(session_orders.quantity) AS quantity,
			menu_items.price_amount AS unit_price,
			menu_items.price_amount * SUM(session_orders.quantity) AS total_price,
			menu_items.price_currency AS currency
		`).Joins("INNER JOIN sessions ON session_orders.session_id = sessions.id").
		Joins("INNER JOIN menu_items ON session_orders.menu_item_id = menu_items.id").
		Group(`
			sessions.id,
			menu_items.name,
			menu_items.price_amount,
			menu_items.price_currency
		`)
}

func PostOrdersView(db *gorm.DB) *gorm.DB {
	return db.Table("session_orders").
		Select(`
			sessions.id AS session_id,
			customers.id AS customer_id,
			menu_items.name AS menu_item_name,
			SUM(session_orders.quantity) AS quantity,
			menu_items.price_amount AS unit_price,
			menu_items.price_currency AS currency
		`).Joins("INNER JOIN sessions ON session_orders.session_id = sessions.id").
		Joins("INNER JOIN menu_items ON session_orders.menu_item_id = menu_items.id").
		Joins("INNER JOIN customers ON session_orders.ordered_by = customers.id").
		Where("sessions.state IN ('CHECKOUT_PENDING')").
		Group(`
			sessions.id,
			customers.id,
			menu_items.name,
			menu_items.price_amount,
			menu_items.price_currency
		`)
}
