//hello, iteration 1
//hello, iteration2
package items

type item struct {
	Name  string
	Count int
	Price int
}

func New(name string, count, price int) *item {
	return &item{
		Name:name,
		Count:count,
		Price:price,
	}
}

func (i item) GetPrice() int {
	return i.Price
}

func (i item) GetName() string {
	return i.Name
}

func (i item) GetCount() int {
	return i.Count
}

func (i *item) CountMinus() {
	i.Count--
}

func (i item) GetStatusItem() bool {
	return false
}