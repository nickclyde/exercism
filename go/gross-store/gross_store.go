package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitSize, unitExists := units[unit]
	if !unitExists {
		return false
	}
	bill[item] += unitSize
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitSize, unitExists := units[unit]
	itemCount, itemExists := bill[item]
	if !unitExists || !itemExists || unitSize > itemCount {
		return false
	}
	if unitSize == itemCount {
		delete(bill, item)
		return true
	}
	bill[item] -= unitSize
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemCount, itemExists := bill[item]
	if !itemExists {
		return 0, false
	}
	return itemCount, true
}
