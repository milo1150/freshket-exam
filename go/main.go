package main

import (
	"math"
)

// Write a Calculator class for food store (you can use any programming languages)
// This food store only have 7 items in menu
// Red set                         	50 THB/set
// Green set                    	40 THB/set
// Blue set                        	30 THB/set
// Yellow set                	        50 THB/set
// Pink set                        	80 THB/set
// Purple set                   	90 THB/set
// Orange set                  	        120 THB/set

// Customers can order multiple items
// Write a function that receives these 7 items, calculates and returns the price.
// Conditions:
// Customers can get 10% on Total, if customers have a member card.
// Order doubles of Orange, Pink or Green sets will get a 5% discount for each bundles (not in total)

// If you provide unit-tests; you will get an extra score.

// Example:
// Desk#1: Customers order Red set and Green set; price from calculation is 90
// Customers can use a member card, then the price should be deducted by discount amount 10%.
// For Orange sets, if customers order 5 items per bill. Customers will get a 5% discount for 4 items (2 pairs).

type Order struct {
	Red      uint16
	Green    uint16
	Blue     uint16
	Yellow   uint16
	Pink     uint16
	Purple   uint16
	Orange   uint16
	IsMember bool
}

type MenuSet string

const (
	RedSet    MenuSet = "RED"
	GreenSet  MenuSet = "GREEN"
	BlueSet   MenuSet = "BLUE"
	YellowSet MenuSet = "YELLOW"
	PinkSet   MenuSet = "PINK"
	PurpleSet MenuSet = "PURPLE"
	OrangeSet MenuSet = "ORANGE"
)

type MenuSetPrice uint16

const (
	RedPrice    MenuSetPrice = 50
	GreenPrice  MenuSetPrice = 40
	BluePrice   MenuSetPrice = 30
	YellowPrice MenuSetPrice = 50
	PinkPrice   MenuSetPrice = 80
	PurplePrice MenuSetPrice = 90
	OrangePrice MenuSetPrice = 120
)

func calculateSpecialPrice(menuPrice MenuSetPrice, discount uint16, amount uint16) uint16 {
	// If no items are ordered, the total price is 0.
	if amount == 0 {
		return 0
	}

	// Check whether the number of items is even.
	isEven := amount%2 == 0

	// Calculate new amount for each even, odd case
	var newAmount uint16
	if isEven {
		newAmount = amount // เลขคู่
	} else {
		newAmount = amount - 1 // เลขคี่
	}

	// Calculate total price
	var total uint16
	discountPrice := calculateTotalDiscount(uint16(menuPrice), discount)
	newItemPrice := calculateDiscountedPrice(uint16(menuPrice), uint16(discountPrice))
	total = newItemPrice * newAmount

	// If the original amount was odd, add one more item
	if !isEven {
		total += uint16(menuPrice)
	}

	return total
}

func calculateTotalPriceBySet(set MenuSet, amount uint16) uint16 {
	switch set {
	case RedSet:
		return amount * uint16(RedPrice)
	case GreenSet:
		return calculateSpecialPrice(GreenPrice, 5, amount)
	case BlueSet:
		return amount * uint16(BluePrice)
	case YellowSet:
		return amount * uint16(YellowPrice)
	case PinkSet:
		return calculateSpecialPrice(PinkPrice, 5, amount)
	case PurpleSet:
		return amount * uint16(PurplePrice)
	case OrangeSet:
		return calculateSpecialPrice(OrangePrice, 5, amount)
	default:
		return 0
	}
}

func calculateTotalDiscount(totalPrice, discount uint16) float64 {
	d := float64(totalPrice) * (float64(discount) / 100)
	return math.Round(d) // Round: Nearest integer
}

func calculateDiscountedPrice(totalPrice, discount uint16) uint16 {
	return totalPrice - discount
}

func calculateOrderTotal(order Order) uint16 {
	total := uint16(0)
	total += calculateTotalPriceBySet(RedSet, order.Red)
	total += calculateTotalPriceBySet(GreenSet, order.Green)
	total += calculateTotalPriceBySet(BlueSet, order.Blue)
	total += calculateTotalPriceBySet(YellowSet, order.Yellow)
	total += calculateTotalPriceBySet(PinkSet, order.Pink)
	total += calculateTotalPriceBySet(PurpleSet, order.Purple)
	total += calculateTotalPriceBySet(OrangeSet, order.Orange)
	return total
}

func orderSummary(order Order) uint16 {
	var total uint16
	total = calculateOrderTotal(order)

	// Check if Member then discount 10%
	if order.IsMember {
		discount := calculateTotalDiscount(total, 10)
		total = uint16(calculateDiscountedPrice(uint16(total), uint16(discount)))
	}

	return total
}

func main() {
	// orderSummary(Order{Red: 1, Blue: 2, Purple: 1, Orange: 2, IsMember: false})
	// orderSummary(Order{Orange: 4, IsMember: true})
	// orderSummary(Order{Orange: 6, IsMember: true})
}
