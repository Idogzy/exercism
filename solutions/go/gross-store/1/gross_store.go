package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    m := make(map[string]int)
    m["quarter_of_a_dozen"] = 3
    m["half_of_a_dozen"] = 6
    m["dozen"] = 12
    m["small_gross"] = 120
    m["gross"] = 144
    m["great_gross"] = 1728
    return m
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return make(map[string]int)
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    u, ok := units[unit]
    if !ok {
        return false
    }    
    bill[item] += u
    return true
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    b, ok := bill[item]
    if !ok {
        return false
    }
    u, value := units[unit]
    if !value {
        return false
    }
    if (b - u) < 0 {
        return false
    }

    if (b - u) == 0 {
        delete(bill, item)
    } else {
        bill[item] = b - u
    }
    return true
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    b, ok := bill[item]
    if !ok {
        return 0, false
    }
    return b, true
	panic("Please implement the GetItem() function")
}
