package iterator

type Iterator interface {
	HasNext() bool
	Next() MenuItem
}

type MenuItem interface {
	GetPrice() float64
	GetName() string
	GetDescription() string
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	return &PancakeHouseMenu{
		menuItemSlice: []MenuItem{&PancakeFish{name: "fish"}},
	}
}

type PancakeHouseMenu struct {
	position      int
	menuItemSlice []MenuItem
}

type PancakeFish struct {
	price       float64
	name        string
	description string
}

func (p *PancakeFish) GetPrice() float64 {
	return p.price
}

func (p *PancakeFish) GetName() string {
	return p.name
}

func (p *PancakeFish) GetDescription() string {
	return p.description
}

func (p *PancakeHouseMenu) HasNext() bool {
	if p.position >= len(p.menuItemSlice) || p.menuItemSlice[p.position] == nil {
		return false
	}
	return true
}

func (p *PancakeHouseMenu) Next() MenuItem {
	menuItem := p.menuItemSlice[p.position]
	p.position++
	return menuItem
}

type DinnerMenu struct {
	position    int
	menuItemArr [10]MenuItem
}

func NewDinnerMenu() *DinnerMenu {
	return &DinnerMenu{
		menuItemArr: [10]MenuItem{&DinnerEgg{name: "egg"}},
	}
}

type DinnerEgg struct {
	price float64
	name  string
}

func (d *DinnerEgg) GetPrice() float64 {
	return d.price
}

func (d *DinnerEgg) GetName() string {
	return d.name
}

func (d *DinnerEgg) GetDescription() string {
	return d.name
}

func (d *DinnerMenu) HasNext() bool {
	if d.position >= len(d.menuItemArr) || d.menuItemArr[d.position] == nil {
		return false
	}
	return true
}

func (d *DinnerMenu) Next() MenuItem {
	menuItem := d.menuItemArr[d.position]
	d.position++
	return menuItem
}
