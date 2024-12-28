package view

import "gorm.io/gorm"

const (
	PastBillItems string = "past_bill_items"
)

func PastBillsItemsView(db *gorm.DB) *gorm.DB {
	return db.Table("bills").
		Select(`
			bills.id AS bill_id,
			bill_items.owner_id AS owner_id,
			restaurants.name AS restaurant_name,
			bill_items.item_name AS item_name,
			bill_items.quantity AS quantity,
			bill_items.unit_price AS unit_price,
			bill_items.unit_price * bill_items.quantity AS order_price,
			COALESCE(SUM(bill_item_payments.paid_price),0) AS paid_price,
			bill_items.currency AS currency
		`).Joins("JOIN sessions ON bills.session_id = sessions.id").
		Joins("JOIN restaurant_tables ON sessions.table_id = restaurant_tables.id").
		Joins("JOIN restaurants ON restaurant_tables.restaurant_id = restaurants.id").
		Joins("JOIN bill_items ON bills.id = bill_items.bill_id").
		Joins(`
			LEFT JOIN (
				SELECT * FROM bill_item_payments
						INNER JOIN bill_items ON bill_item_payments.bill_item_id = bill_items.id
						WHERE bill_item_payments.customer_id=bill_items.owner_id) bill_item_payments ON bill_item_payments.bill_item_id=bill_items.id
		`).Group(`
			restaurants.name,
			bills.id,
			bill_items.item_name,
			bill_items.quantity,
			bill_items.unit_price,
			bill_items.currency,
			bill_items.owner_id
		`)
}
