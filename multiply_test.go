package simulate

import "testing"

func TestMultiply(t *testing.T) {
	var tests = []struct {
		x, y    string
		product string
	}{
		{"001101", "01011", "0010001111"},
		{"110101", "10011", "0010001111"},
		{"001101", "10101", "1101110001"},
		{"110101", "01101", "1101110001"},
	}

	for _, tt := range tests {
		product := Multiply(tt.x, tt.y)
		if product != tt.product {
			t.Errorf("%s * %s = %s, but get %s", tt.x, tt.y, tt.product, product)
		}
		if product[0] != product[1] {
			t.Errorf("double sign bit is different")
		}
	}
}
