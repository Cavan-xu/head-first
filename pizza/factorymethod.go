package pizza

// 原料工厂
type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateVeggie() Veggie
	CreateClams() Clams
}

type Dough interface {
	GetDoughId() int
}

type Sauce interface {
	GetSauceId() int
}

type Cheese interface {
	GetCheeseId() int
}

type Veggie interface {
	GetVeggieId() int
}

type Clams interface {
	GetClamsId() int
}

type NYIngredientFactory struct {
}

type NYDough struct {
}

func (N NYDough) GetDoughId() int {
	return 0
}

type NYSauce struct {
}

func (N NYSauce) GetSauceId() int {
	return 0
}

type NYCheese struct {
}

func (N NYCheese) GetCheeseId() int {
	return 0
}

type NYVeggie struct {
}

func (N NYVeggie) GetVeggieId() int {
	return 0
}

type NYClams struct {
}

func (N NYClams) GetClamsId() int {
	return 0
}

func (N *NYIngredientFactory) CreateDough() Dough {
	return &NYDough{}
}

func (N *NYIngredientFactory) CreateSauce() Sauce {
	return &NYSauce{}
}

func (N *NYIngredientFactory) CreateCheese() Cheese {
	return &NYCheese{}
}

func (N *NYIngredientFactory) CreateVeggie() Veggie {
	return &NYVeggie{}
}

func (N *NYIngredientFactory) CreateClams() Clams {
	return &NYClams{}
}
