package main

import "testing"

func Test_calculateSpecialPrice(t *testing.T) {
	type args struct {
		menuPrice MenuSetPrice
		discount  uint16
		amount    uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Zero items ordered",
			args: args{
				menuPrice: RedPrice,
				discount:  5,
				amount:    0,
			},
			want: 0,
		},
		{
			name: "Even number of items with discount",
			args: args{
				menuPrice: GreenPrice, // 40
				discount:  5,          // 5% discount
				amount:    4,          // 4 items
			},
			want: 152, // (40 - 2) = 38 per item, 38 * 4 = 152
		},
		{
			name: "Odd number of items with discount",
			args: args{
				menuPrice: GreenPrice, // 40
				discount:  5,          // 5% discount
				amount:    5,          // 5 items
			},
			want: 192, // (40 - 2) = 38 per item, 38 * 4 + 40 = 192
		},
		{
			name: "Only one item (no discount applied on odd last item)",
			args: args{
				menuPrice: GreenPrice,
				discount:  5,
				amount:    1,
			},
			want: 40, // Full price, because only 1 item (odd)
		},
		{
			name: "Even number of Pink sets (special price)",
			args: args{
				menuPrice: PinkPrice,
				discount:  5,
				amount:    2,
			},
			want: 152, // (80 - 4) = 76 per item, 76 * 2 = 152
		},
		{
			name: "Odd number of Orange sets",
			args: args{
				menuPrice: OrangePrice,
				discount:  5,
				amount:    3,
			},
			want: 348, // (120 - 6) = 114 per item, 114 * 2 + 120 = 348
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSpecialPrice(tt.args.menuPrice, tt.args.discount, tt.args.amount); got != tt.want {
				t.Errorf("calculateSpecialPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateTotalPriceBySet(t *testing.T) {
	type args struct {
		set    MenuSet
		amount uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Red set normal price",
			args: args{
				set:    RedSet,
				amount: 2,
			},
			want: 100, // 2 * 50
		},
		{
			name: "Green set special discount",
			args: args{
				set:    GreenSet,
				amount: 2,
			},
			want: 76, // (40 - 2) * 2 = 76
		},
		{
			name: "Pink set special discount",
			args: args{
				set:    PinkSet,
				amount: 2,
			},
			want: 152, // (80 - 4) * 2 = 152
		},
		{
			name: "Orange set odd special discount",
			args: args{
				set:    OrangeSet,
				amount: 3,
			},
			want: 348, // (120 - 6) * 2 + 120 = 348
		},
		{
			name: "Unknown set",
			args: args{
				set:    MenuSet("UNKNOWN"),
				amount: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTotalPriceBySet(tt.args.set, tt.args.amount); got != tt.want {
				t.Errorf("calculateTotalPriceBySet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateTotalDiscount(t *testing.T) {
	type args struct {
		totalPrice uint16
		discount   uint16
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "10 percent discount",
			args: args{
				totalPrice: 200,
				discount:   10,
			},
			want: 20.0,
		},
		{
			name: "5 percent discount",
			args: args{
				totalPrice: 400,
				discount:   5,
			},
			want: 20.0,
		},
		{
			name: "Zero discount",
			args: args{
				totalPrice: 100,
				discount:   0,
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTotalDiscount(tt.args.totalPrice, tt.args.discount); got != tt.want {
				t.Errorf("calculateTotalDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateDiscountedPrice(t *testing.T) {
	type args struct {
		totalPrice uint16
		discount   uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Basic discount",
			args: args{
				totalPrice: 200,
				discount:   20,
			},
			want: 180,
		},
		{
			name: "Zero discount",
			args: args{
				totalPrice: 100,
				discount:   0,
			},
			want: 100,
		},
		{
			name: "Full discount",
			args: args{
				totalPrice: 50,
				discount:   50,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateDiscountedPrice(tt.args.totalPrice, tt.args.discount); got != tt.want {
				t.Errorf("calculateDiscountedPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateOrderTotal(t *testing.T) {
	type args struct {
		order Order
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Simple order - Red only",
			args: args{
				order: Order{
					Red: 2, // 2 * 50 = 100
				},
			},
			want: 100,
		},
		{
			name: "Simple order - Green with discount",
			args: args{
				order: Order{
					Green: 2, // (40-2) * 2 = 76
				},
			},
			want: 76,
		},
		{
			name: "Mixed order",
			args: args{
				order: Order{
					Red:    1, // 50
					Blue:   2, // 2 * 30 = 60
					Purple: 1, // 90
					Orange: 2, // (120-6)*2 = 228
				},
			},
			want: 50 + 60 + 90 + 228, // 428
		},
		{
			name: "Empty order",
			args: args{
				order: Order{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateOrderTotal(tt.args.order); got != tt.want {
				t.Errorf("calculateOrderTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderSummary(t *testing.T) {
	type args struct {
		order Order
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Non-member simple order",
			args: args{
				order: Order{
					Red: 2, // 2 * 50 = 100
				},
			},
			want: 100, // no discount
		},
		{
			name: "Member simple order",
			args: args{
				order: Order{
					Red:      2,    // 2 * 50 = 100
					IsMember: true, // get 10% discount
				},
			},
			want: 90, // 100 - 10% = 90
		},
		{
			name: "Mixed sets non-member",
			args: args{
				order: Order{
					Red:    1, // 50
					Blue:   2, // 2 * 30 = 60
					Yellow: 1, // 50
				},
			},
			want: 160, // 50 + 60 + 50
		},
		{
			name: "Mixed sets member",
			args: args{
				order: Order{
					Red:      1, // 50
					Blue:     2, // 60
					Yellow:   1, // 50
					IsMember: true,
				},
			},
			want: 144, // (50+60+50) = 160 -> 160 - 10% = 144
		},
		{
			name: "Empty order non-member",
			args: args{
				order: Order{},
			},
			want: 0,
		},
		{
			name: "Empty order member",
			args: args{
				order: Order{
					IsMember: true,
				},
			},
			want: 0,
		},
		{
			name: "Special discount + membership",
			args: args{
				order: Order{
					Orange:   4, // Special discount (bundle discount)
					IsMember: true,
				},
			},
			// OrangePrice = 120
			// Discount 5% -> 120-6 = 114 per item
			// 114 * 4 = 456
			// Member discount 10% => 456 - 10% = 410.4 ~ 410
			want: 410,
		},
		{
			name: "Member orders odd Green sets (special discount + 1 full price)",
			args: args{
				order: Order{
					Green:    3, // special discount applies to 2, 1 remains full price
					IsMember: true,
				},
			},
			// Green price = 40
			// 5% of 40 = 2
			// discounted price = 38
			// 2 items discounted = 38 * 2 = 76
			// 1 item full price = 40
			// total before membership = 76 + 40 = 116
			// after 10% member discount = 116 - 11.6 ≈ 104.4
			want: 104,
		},
		{
			name: "Member orders odd Pink sets (special discount + 1 full price)",
			args: args{
				order: Order{
					Pink:     5, // 4 discounted, 1 full price
					IsMember: true,
				},
			},
			// Pink price = 80
			// 5% of 80 = 4
			// discounted price = 76
			// 4 items discounted = 76 * 4 = 304
			// 1 item full price = 80
			// total before membership = 384
			// after 10% member discount = 384 - 38.4 ≈ 345.6
			want: 346,
		},
		{
			name: "Non-member large mixed order (Red, Orange, Pink, Blue)",
			args: args{
				order: Order{
					Red:    5, // 5 * 50 = 250
					Orange: 3, // Orange special discount
					Pink:   2, // Pink special discount
					Blue:   1, // 1 * 30
				},
			},
			// Orange:
			// 5% of 120 = 6
			// discounted price = 114
			// 2 items discounted = 114 * 2 = 228
			// 1 item full price = 120
			// total Orange = 228 + 120 = 348
			// Pink:
			// 5% of 80 = 4
			// discounted price = 76
			// 2 Pink = 76 * 2 = 152
			// Blue = 30
			// Red = 250
			// total = 250 + 348 + 152 + 30 = 780
			want: 780,
		},
		{
			name: "Member orders only Orange sets (special discount bundles)",
			args: args{
				order: Order{
					Orange:   6, // all pairs
					IsMember: true,
				},
			},
			// Orange price = 120
			// 5% of 120 = 6
			// discounted price = 114
			// 6 items = 114 * 6 = 684
			// member discount 10% => 684 - 68.4 ≈ 615.6
			want: 616,
		},
		{
			name: "Member orders only Blue sets (no special discount)",
			args: args{
				order: Order{
					Blue:     5, // normal price
					IsMember: true,
				},
			},
			// Blue = 30
			// total = 5 * 30 = 150
			// 10% member discount => 150 - 15 = 135
			want: 135,
		},
		{
			name: "Edge case: Member orders 1 Green set (should not apply discount)",
			args: args{
				order: Order{
					Green:    1, // 1 item = special discount does not apply (no bundle)
					IsMember: true,
				},
			},
			// Green = 40
			// no bundle discount, normal 40
			// 10% member discount = 40 - 4 = 36
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderSummary(tt.args.order); got != tt.want {
				t.Errorf("orderSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
