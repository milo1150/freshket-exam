type Order = {
  red: number
  green: number
  blue: number
  yellow: number
  pink: number
  purple: number
  orange: number
  isMember: boolean
}

const MenuSet = {
  RedSet: 'RED',
  GreenSet: 'GREEN',
  BlueSet: 'BLUE',
  YellowSet: 'YELLOW',
  PinkSet: 'PINK',
  PurpleSet: 'PURPLE',
  OrangeSet: 'ORANGE',
} as const

const MenuSetPrice = {
  Red: 50,
  Green: 40,
  Blue: 30,
  Yellow: 50,
  Pink: 80,
  Purple: 90,
  Orange: 120,
} as const

function calculateTotalDiscount(totalPrice: number, discount: number): number {
  const discountAmount = totalPrice * (discount / 100)
  return Math.round(discountAmount) // Round to the nearest integer
}

function calculateDiscountedPrice(
  totalPrice: number,
  discount: number
): number {
  return totalPrice - discount
}

function calculateSpecialPrice(
  menuPrice: number,
  discount: number,
  amount: number
): number {
  if (amount === 0) {
    return 0
  }

  const isEven = amount % 2 === 0
  const newAmount = isEven ? amount : amount - 1

  const discountPrice = calculateTotalDiscount(menuPrice, discount)
  const newItemPrice = calculateDiscountedPrice(menuPrice, discountPrice)
  let total = newItemPrice * newAmount

  if (!isEven) {
    total += menuPrice
  }

  return total
}

function calculateTotalPriceBySet(set: string, amount: number): number {
  switch (set) {
    case MenuSet.RedSet:
      return amount * MenuSetPrice.Red
    case MenuSet.GreenSet:
      return calculateSpecialPrice(MenuSetPrice.Green, 5, amount)
    case MenuSet.BlueSet:
      return amount * MenuSetPrice.Blue
    case MenuSet.YellowSet:
      return amount * MenuSetPrice.Yellow
    case MenuSet.PinkSet:
      return calculateSpecialPrice(MenuSetPrice.Pink, 5, amount)
    case MenuSet.PurpleSet:
      return amount * MenuSetPrice.Purple
    case MenuSet.OrangeSet:
      return calculateSpecialPrice(MenuSetPrice.Orange, 5, amount)
    default:
      return 0
  }
}

function calculateOrderTotal(order: Order): number {
  let total = 0
  total += calculateTotalPriceBySet(MenuSet.RedSet, order.red)
  total += calculateTotalPriceBySet(MenuSet.GreenSet, order.green)
  total += calculateTotalPriceBySet(MenuSet.BlueSet, order.blue)
  total += calculateTotalPriceBySet(MenuSet.YellowSet, order.yellow)
  total += calculateTotalPriceBySet(MenuSet.PinkSet, order.pink)
  total += calculateTotalPriceBySet(MenuSet.PurpleSet, order.purple)
  total += calculateTotalPriceBySet(MenuSet.OrangeSet, order.orange)
  return total
}

export function orderSummary(order: Order): number {
  let total = calculateOrderTotal(order)

  if (order.isMember) {
    const discount = calculateTotalDiscount(total, 10)
    total = calculateDiscountedPrice(total, discount)
  }

  return total
}

// Run
// const exOne = orderSummary({
//   red: 1,
//   blue: 2,
//   purple: 1,
//   orange: 2,
//   green: 0,
//   yellow: 0,
//   pink: 0,
//   isMember: false,
// })
// console.log(exOne)

// const exTwo = orderSummary({
//   red: 0,
//   blue: 0,
//   purple: 0,
//   orange: 4,
//   green: 0,
//   yellow: 0,
//   pink: 0,
//   isMember: true,
// })
// console.log(exTwo)
