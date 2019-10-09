package shop

type Itemer interface {
	GetPrice() int
	GetName() string
	GetCount() int
	CountMinus()
	GetStatusItem() bool
}

type Useer interface {
	GetName() string
	GetCash() int
	CashMinus(int)
	GetStatusUser() bool
}

type newShopUpdateer interface {
	userUpdate(s *shop, users []Useer) //апдейт юзеров
	itemsUpdate(s *shop, items []Itemer) //апдейт предметов
	initBigShop(s *shop) //сносим мапу, которая с короткими заказами
	editLittleMap(s *shop, items []Itemer) //редактируем мапу с длинными заказами
}
