package items

type premItem struct {
	item
}

func (pi premItem) GetPrice() int {
	return pi.Price * 150 / 100
}

func (pi premItem) GetStatusItem() bool {
	return true
}

func NewPrem(name string,  count, price int) *premItem {
	return &premItem{item{
		Name:  name,
		Count: count,
		Price: price,
	},
	}
}
