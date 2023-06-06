package decoration

/*
	装饰器模式
*/

func NewBeverage() *Beverage {
	return &Beverage{
		basePrice: 10,
	}
}

type Beverage struct {
	basePrice       int
	additionalPrice int
}

func (b *Beverage) Cost() int {
	return b.basePrice + b.additionalPrice
}

func (b *Beverage) WithAdditional(funcs ...AdditionalFunc) {
	for _, fn := range funcs {
		fn(b)
	}
}

type AdditionalFunc func(*Beverage)

func WithMilk(b *Beverage) {
	milkPrice := 10
	b.additionalPrice += milkPrice
}

func WithSoy(b *Beverage) {
	soyPrice := 5
	b.additionalPrice += soyPrice
}
