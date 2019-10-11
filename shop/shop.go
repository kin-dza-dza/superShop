//hello, iteration 1
//hello, iteration2
package shop

import (
	"errors"
	"fmt"
	"sort"
)

const deletekey = -1

type shop struct {
	users       map[string]Useer
	items       map[string]Itemer
	littleCache map[string]map[string]int
	bigCache    map[string]int
	updater     newShopUpdateer
}

func (s *shop) NewShopUpdate(items []Itemer, users []Useer) {
	//апдейт юзеров
	ch := make(chan struct{}, 4)

	go func() {
		s.updater.userUpdate(s, users)
		ch <- struct{}{}
	}()

	//апдейт предметов
	go func() {
		s.updater.itemsUpdate(s, items)
		ch <- struct{}{}
	}()

	//сносим мапу, которая с длинными заказами
	go func() {
		s.updater.initBigShop(s)
		ch <- struct{}{}
	}()

	//редактируем мапу, которая с короткими заказами
	go func() {
		s.updater.editLittleMap(s, items)
		ch <- struct{}{}
	}()

	i := 0
	for range ch {
		i++
		if i == 4 {
			close(ch)
		}
	}
}

func (s *shop) Buy(order []string, user string) error {
	//сначала чекаем существует ли товар, и есть ли он на складе
	err := s.checkItems(order)
	if err != nil {
		return err
	}

	//если товаров больше, чем два, делаем так
	if len(order) > 2 {
		err = s.buyBigPrice(order, user)

		return err
	}

	//если товаров меньше двух, то идем в мапу-мапу
	err = s.buyLittlePrice(order, user)
	return err
}

func (s *shop) checkItems(order []string) error {
	for _, item := range order {
		_, ok := s.items[item]
		if !ok {
			return errors.New("товара не существует")
		}
		if s.items[item].GetCount() <= 0 {
			return errors.New("товар закончился")
		}
	}

	return nil
}

func (s *shop) buyBigPrice(order []string, user string) error {
	//делаем хэш
	hash := makeHash(order)
	//если хэш найден, проводим оплату
	sum, ok := s.checkBigCache(hash)
	if ok {
		err := s.payment(sum, user, order)
		if err == nil { //если деньги есть, делаем доставку
			s.delivery(order)
		}

		return err
	}
	//если хэша нет, собираем заказ вручную
	sum, err := s.collectOrder(order, user) //собираем заказ
	if err == nil { //если нет ошибки, то
		err = s.payment(sum, user, order)
		//делаем кэш. Даже если покупка не состоится (денег нет), все равно делаем
		s.writeBigCache(hash, sum)
		//если ошибок нет, списываем товары
		if err == nil {
			s.delivery(order)
		}
	}

	return err
}

func (s *shop) checkBigCache(hash string) (int, bool) {
	sum, ok := s.bigCache[hash]

	return sum, ok
}

func (s *shop) collectOrder(order []string, user string) (int, error) {
	var sum int

	for _, item := range order {
		itemXXX, ok := s.items[item]
		if !ok {
			return 0, errors.New("товара не существует")
		}
		sum += itemXXX.GetPrice()
	}

	return sum, nil
}

func (s *shop) payment(sum int, user string, order []string) error {
	//смотрим статус пользователя - вип он или кто?
	//рассчитываем для него конечную стоимость исходя из статуса
	statusUser := s.users[user].GetStatusUser()
	//если тру, значит вип пользователь
	if statusUser == true {
		for _, item := range order {
			statusItem := s.items[item].GetStatusItem()
			if statusItem == false {
				sum = sum - s.items[item].GetPrice()
				newPriceItem := s.items[item].GetPrice() * 90 / 100
				sum = sum + newPriceItem
			}
		}
	}

	//если фолс, то пользователь обычный
	if statusUser == false {
		for _, item := range order {
			statusItem := s.items[item].GetStatusItem()
			if statusItem == true {
				sum = sum - s.items[item].GetPrice()
				newPriceItem := s.items[item].GetPrice() * 150 / 100
				sum = sum + newPriceItem
			}
		}
	}

	//проверяем, есть ли деньги у пользователя
	if s.users[user].GetCash() < sum {
		return errors.New("денег нет, но вы держитесь")
	}

	s.users[user].CashMinus(sum)
	return nil
}

func (s *shop) writeBigCache(hash string, sum int) {
	s.bigCache[hash] = sum
}

func (s *shop) delivery(order []string) {
	for _, item := range order {
		shopItem := s.items[item]
		shopItem.CountMinus()
		s.items[item] = shopItem
	}
}

func (s *shop) buyLittlePrice(order []string, user string) error {
	//проверяем кэш
	sum, ok := s.checkLittleCache(order)
	if ok { //если там что-то лежит, покупаем (чекаем, вдруг нет денег)
		err := s.payment(sum, user, order)
		if err == nil { //если деньги есть, делаем доставку
			s.delivery(order)
		}
		return err
	}
	//если хэша нет, придется собирать заказ вручную
	sum, err := s.collectOrder(order, user) //собираем заказ
	if err == nil { //если ошибки нет, то
		err = s.payment(sum, user, order)
		//делаем кэш. Даже если покупка не состоится (денег нет), все равно делаем
		s.writeLittleCache(order, sum)
		//если ошибок нет, списываем товары
		if err == nil {
			s.delivery(order)
		}
	}


	return err
}

func (s *shop) checkLittleCache(order []string) (int, bool) {
	mapStep, ok := s.littleCache[getItem(order, 0)]
	if !ok {
		return 0, false
	}

	_, ok = mapStep[getItem(order, 1)]
	if !ok {
		return 0, false
	}

	sum := mapStep[getItem(order, 1)]

	return sum, true
}

func (s *shop) writeLittleCache(order []string, sum int) {
	item := getItem(order, 0)
	mapStep1, ok := s.littleCache[item]
	if !ok {
		mapStep1 = make(map[string]int)
		s.littleCache[item] = mapStep1
	}

	item = getItem(order, 1)
	_, ok = mapStep1[item]
	if !ok {
		mapStep1[item] = sum
	}
}

func (s *shop) Intrigue(err error, user string) {
	if err == nil {
		fmt.Println("оплата прошла успешно, ")
	} else {
		fmt.Println(err)
	}

	fmt.Println("Баланс пользователя:", s.users[user])

	itemString := "Товары на складе:\n"
	for itemName, item := range s.items {
		itemString += fmt.Sprintf("\titemName %q  \tcount %d  \tprice %d\n", itemName, item.GetCount(), item.GetPrice())
	}
	fmt.Print(itemString)
}

func NewShopPrepare(items []Itemer, users []Useer) *shop {
	shop := newShop(items, users, NewShopUpdater{})

	shop.NewShopUpdate(items, users)
	return shop
}

func newShop(items []Itemer, users []Useer, updater newShopUpdateer) *shop {
	usersMap := make(map[string]Useer, len(users))
	itemsMap := make(map[string]Itemer, len(items))

	for _, user := range users {
		usersMap[user.GetName()] = user
	}
	for _, item := range items {
		itemsMap[item.GetName()] = item
	}

	littleCacheMap := make(map[string]map[string]int)
	bigCacheMap := make(map[string]int)

	return &shop{
		users:       usersMap,
		items:       itemsMap,
		littleCache: littleCacheMap,
		bigCache:    bigCacheMap,
		updater:     updater,
	}
}

func makeHash(order []string) string {
	sort.Strings(order)

	var hash string
	for _, item := range order {
		hash += item
	}

	return hash
}

func getItem(order []string, index int) string {
	item := ""
	if index < 0 {
		return item
	}
	if index < len(order) {
		item = order[index]
	}
	return item
}

func HelloPrint (fl bool) {
	if fl == true {
		fmt.Println("я все еще живой, и я все еще магазин")
	}
}

//
//func (ns *NewShopUpdater) NewShopUpdate(s *shop, items []Itemer, users []Useer) {
//	//апдейт юзеров
//	wg := &sync.WaitGroup{}
//	wg.Add(4)
//
//	i := 3
//
//	go func() {
//		s.userUpdate(users)
//		i = 5
//		wg.Done()
//	}()
//
//	//апдейт предметов
//	go func() {
//		s.itemsUpdate(items)
//		i = 6
//		wg.Done()
//	}()
//
//	//сносим мапу, которая с длинными заказами
//	go func() {
//		s.initBigShop()
//		i = 7
//		wg.Done()
//	}()
//
//	//редактируем мапу, которая с короткими заказами
//	go func() {
//		s.editLittleMap(items)
//		i++
//		wg.Done()
//	}()
//
//	wg.Wait()
//	fmt.Println("i", i)
//}