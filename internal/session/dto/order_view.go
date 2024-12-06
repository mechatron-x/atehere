package dto

type (
	TableOrderView struct {
		CustomerID       string
		CustomerFullName string
		MenuItemName     string
		Quantity         int
		UnitPrice        float64
		TotalPrice       float64
		Currency         string
	}

	ManagerOrderView struct {
		MenuItemName string
		Quantity     int
		UnitPrice    float64
		TotalPrice   float64
		Currency     string
	}
)

func (rcv TableOrderView) ToOrder() Order {
	return Order{
		CustomerFullName: rcv.CustomerFullName,
		MenuItemName:     rcv.MenuItemName,
		Quantity:         rcv.Quantity,
		UnitPrice:        rcv.UnitPrice,
		TotalPrice:       rcv.TotalPrice,
	}
}

func (rcv ManagerOrderView) ToOrder() Order {
	return Order{
		MenuItemName: rcv.MenuItemName,
		Quantity:     rcv.Quantity,
		UnitPrice:    rcv.UnitPrice,
		TotalPrice:   rcv.TotalPrice,
	}
}

func FromTableOrdersView(tableOrders []TableOrderView) []Order {
	orders := make([]Order, 0)

	for _, o := range tableOrders {
		orders = append(orders, o.ToOrder())
	}

	return orders
}

func FromTableOrdersViewWithFilter(tableOrders []TableOrderView, containsFunc func(TableOrderView) bool) []Order {
	orders := make([]Order, 0)

	for _, o := range tableOrders {
		if !containsFunc(o) {
			continue
		}

		orders = append(orders, o.ToOrder())
	}

	return orders
}

func FromManagerOrdersView(managerOrders []ManagerOrderView) []Order {
	orders := make([]Order, 0)

	for _, o := range managerOrders {
		orders = append(orders, o.ToOrder())
	}

	return orders
}
