package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	result := 0

	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}

	if order == "burger" {
		result = burger.preptime
	} else if order == "chips" {
		result = chips.preptime
	} else if order == "nuggets" {
		result = nuggets.preptime
	} else {
		return 404
	}

	return result
}
