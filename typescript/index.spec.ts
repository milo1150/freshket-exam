import { describe, it, expect } from 'bun:test'
import { orderSummary } from '.'

describe('orderSummary', () => {
  it('should calculate total price for non-member simple order', () => {
    const order = {
      red: 2,
      green: 0,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: false,
    }
    const total = orderSummary(order)
    expect(total).toBe(100) // 2 * 50 = 100
  })

  it('should calculate total price for member simple order', () => {
    const order = {
      red: 2,
      green: 0,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(90) // 2 * 50 = 100 - 10% = 90
  })

  it('should calculate total price for mixed sets non-member', () => {
    const order = {
      red: 1,
      green: 0,
      blue: 2,
      yellow: 1,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: false,
    }
    const total = orderSummary(order)
    expect(total).toBe(160) // 50 + (30*2) + 50 = 160
  })

  it('should calculate total price for mixed sets with member discount', () => {
    const order = {
      red: 1,
      green: 0,
      blue: 2,
      yellow: 1,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(144) // 160 - 10% = 144
  })

  it('should calculate total price for empty order non-member', () => {
    const order = {
      red: 0,
      green: 0,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: false,
    }
    const total = orderSummary(order)
    expect(total).toBe(0)
  })

  it('should calculate total price for empty order with member discount', () => {
    const order = {
      red: 0,
      green: 0,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(0)
  })

  it('should calculate total with special discount and membership for Orange', () => {
    const order = {
      red: 0,
      green: 0,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 4,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(410) // (120*4 with 5% discount) then 10% membership
  })

  it('should calculate total for member with odd Green sets (special discount + full price)', () => {
    const order = {
      red: 0,
      green: 3,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(104) // special discount 2 + 1 full price, then membership
  })

  it('should calculate total for member with odd Pink sets (special discount + full price)', () => {
    const order = {
      red: 0,
      green: 0,
      blue: 0,
      yellow: 0,
      pink: 5,
      purple: 0,
      orange: 0,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(346) // special discount for 4 + 1 full price, then membership
  })

  it('should calculate total for non-member large mixed sets', () => {
    const order = {
      red: 5,
      green: 0,
      blue: 1,
      yellow: 0,
      pink: 2,
      purple: 0,
      orange: 3,
      isMember: false,
    }
    const total = orderSummary(order)
    expect(total).toBe(780) // calculated manually
  })

  it('should calculate total for member ordering only Orange sets (bundles)', () => {
    const order = {
      red: 0,
      green: 0,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 6,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(616) // bundles of Orange sets, then membership discount
  })

  it('should calculate total for member ordering only Blue sets', () => {
    const order = {
      red: 0,
      green: 0,
      blue: 5,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(135) // normal blue sets, then 10% membership
  })

  it('should calculate total for edge case: 1 Green set member', () => {
    const order = {
      red: 0,
      green: 1,
      blue: 0,
      yellow: 0,
      pink: 0,
      purple: 0,
      orange: 0,
      isMember: true,
    }
    const total = orderSummary(order)
    expect(total).toBe(36) // 40 - 10% = 36
  })
})
