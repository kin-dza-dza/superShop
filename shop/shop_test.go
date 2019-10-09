package shop

import (
	"errors"
	"golang-book/go_shop/items"
	"golang-book/go_shop/users"
	"reflect"
	"testing"
)

func TestShop_userUpdate(t *testing.T) {

	t.Run("delete user", func(t *testing.T) {
		shop := newTestShop()
		user := []Useer{
			users.New("dasha", deletekey),
		}
		ns := NewShopUpdater{}
		ns.userUpdate(shop, user)

		expected := map[string]Useer{
			"masha": users.NewPrem("masha", 500),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("Не удаляет простого смертного. expected %v, got %v", expected, shop.users)
		}
	})

	t.Run("delete prem user", func(t *testing.T) {
		shop := newTestShop()
		user := []Useer{
			users.NewPrem("masha", deletekey),
		}
		ns := NewShopUpdater{}
		ns.userUpdate(shop, user)

		expected := map[string]Useer{
			"dasha": users.New("dasha", 10),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("Не удаляет прем смертного. expected %v, got %v", expected, shop.users)
		}
	})

	t.Run("update user", func(t *testing.T) {
		shop := newTestShop()
		user := []Useer{
			users.New("dasha", 999),
		}
		ns := NewShopUpdater{}
		ns.userUpdate(shop, user)

		expected := map[string]Useer{
			"dasha": users.New("dasha", 999),
			"masha": users.NewPrem("masha", 500),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("updateUser работает неправильно. expected %v, got %v", expected, shop.users)
		}
	})

	t.Run("update prem user", func(t *testing.T) {
		shop := newTestShop()
		user := []Useer{
			users.NewPrem("masha", 999),
		}
		ns := NewShopUpdater{}
		ns.userUpdate(shop, user)

		expected := map[string]Useer{
			"dasha": users.New("dasha", 10),
			"masha": users.NewPrem("masha", 999),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("updatePremUser работает неправильно. expected %v, got %v", expected, shop.users)
		}
	})

	t.Run("create mortal user", func(t *testing.T) {
		shop := newTestShop()
		user := []Useer{
			users.New("putin", 999),
		}
		ns := NewShopUpdater{}
		ns.userUpdate(shop, user)

		expected := map[string]Useer{
			"dasha": users.New("dasha", 10),
			"putin": users.New("putin", 999),
			"masha": users.NewPrem("masha", 500),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("добавление смертного работает неправильно. expected %v, got %v", expected, shop.users)
		}
	})

	t.Run("create prem User", func(t *testing.T) {
		shop := newTestShop()
		user := []Useer{
			users.NewPrem("putin", 999),
		}
		ns := NewShopUpdater{}
		ns.userUpdate(shop, user)

		expected := map[string]Useer{
			"dasha": users.New("dasha", 10),
			"putin": users.NewPrem("putin", 999),
			"masha": users.NewPrem("masha", 500),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("добавление прем смертного работает неправильно. expected %v, got %v", expected, shop.users)
		}
	})
}

func TestShop_itemsUpdate(t *testing.T)  {

	t.Run("delete item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.New("apple", 0, deletekey),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"banana": items.NewPrem("banana", 20, 20),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Не удаляет простой айтем. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("delete nonexistent item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.New("xBox", 0, deletekey),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"apple": items.New("apple", 20,10),
			"banana": items.NewPrem("banana", 20, 20),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Ошибка удаления несуществующего айтема. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("delete prem item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.NewPrem("banana", 20, deletekey),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"apple": items.New("apple", 20, 10),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Не удаляет прем айтем. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("update item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.New("apple", 20, 99),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"apple": items.New("apple", 20,99),
			"banana": items.NewPrem("banana", 20, 20),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Ошибка обновления простого айтема. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("update prem item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.NewPrem("banana", 20, 99),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"apple": items.New("apple", 20,10),
			"banana": items.NewPrem("banana", 20, 99),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Ошибка обновления прем айтема. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("create item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.New("goldBar", 1,999),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"apple": items.New("apple", 20,10),
			"goldBar": items.New("goldBar", 1,999),
			"banana": items.NewPrem("banana", 20,20),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Ошибка добавления нового предмета. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("create empty item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.New("", 1,999),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"apple": items.New("apple", 20,10),
			"": items.New("", 1,999),
			"banana": items.NewPrem("banana", 20,20),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Ошибка добавления нового предмета. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("create prem item", func(t *testing.T) {
		shop := newTestShop()
		item := []Itemer{
			items.NewPrem("goldBar", 1,999),
		}
		ns := NewShopUpdater{}
		ns.itemsUpdate(shop, item)

		expected := map[string]Itemer{
			"apple": items.New("apple", 20,10),
			"goldBar": items.NewPrem("goldBar", 1,999),
			"banana": items.NewPrem("banana", 20,20),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("Ошибка добавления нового прем предмета. expected %v, got %v", expected, shop.items)
		}
	})
}

func TestShop_newShop (t *testing.T) {

	t.Run("empty shop, nil users and nil items", func(t *testing.T) {
		shop := newShop(nil, nil, NewShopUpdater{})

		if len(shop.items) != 0 {
			t.Errorf("itemsMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.users) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.users)
		}

		if len(shop.littleCache) != 0 {
			t.Errorf("littleCache формируется неправильно. expected %v", shop.littleCache)
		}

		if len(shop.bigCache) != 0 {
			t.Errorf("bigCache формируется неправильно. expected %v", shop.bigCache)
		}
	})

	t.Run("empty shop, empty users and empty items", func(t *testing.T) {
		shop := newShop([]Itemer{}, []Useer{}, NewShopUpdater{})

		if len(shop.items) != 0 {
			t.Errorf("itemsMap формируется неправильно. expected %v", shop.items)
		}

		if len(shop.users) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.users)
		}

		if len(shop.littleCache) != 0 {
			t.Errorf("littleCache формируется неправильно. expected %v", shop.littleCache)
		}

		if len(shop.bigCache) != 0 {
			t.Errorf("bigCache формируется неправильно. expected %v", shop.bigCache)
		}
	})

	t.Run("check items", func(t *testing.T) {
		shop := newTestShop()

		expected := map[string]Itemer{
			"apple": items.New("apple", 20, 10),
			"banana": items.NewPrem("banana", 20, 20),
		}

		if !reflect.DeepEqual(expected, shop.items) {
			t.Errorf("itemsMap формируется неправильно. expected %v, got %v", expected, shop.items)
		}
	})

	t.Run("check users", func(t *testing.T) {
		shop := newTestShop()

		expected := map[string]Useer{
			"dasha": users.New("dasha", 10),
			"masha": users.NewPrem("masha", 500),
		}

		if !reflect.DeepEqual(expected, shop.users) {
			t.Errorf("usersMap формируется неправильно. expected %v, got %v", expected, shop.users)
		}
	})

	t.Run("check shop", func(t *testing.T) {
		expectedShop := &shop{
			users: map[string]Useer{
				"dasha": users.New("dasha", 10),
				"masha": users.NewPrem("masha", 500),
			},
			items: map[string]Itemer{
				"apple": items.New("apple", 20, 10),
				"banana": items.NewPrem("banana", 20, 20),
			},
		}

		shop := newTestShop()

		if !reflect.DeepEqual(expectedShop.items, shop.items) {
			t.Errorf("usersMap формируется неправильно. expected %v, got %v", expectedShop.items, shop.items)
		}

		if !reflect.DeepEqual(expectedShop.users, shop.users) {
			t.Errorf("usersMap формируется неправильно. expected %v, got %v", expectedShop.users, shop.items)
		}

		if len(shop.littleCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.littleCache)
		}

		if len(shop.bigCache) != 0 {
			t.Errorf("usersMap формируется неправильно. expected %v", shop.bigCache)
		}
	})
}

func TestShop_editLittleMap (t *testing.T) {

	t.Run("top layer cleaning", func(t *testing.T) {
		testItems := []Itemer{
			items.New("apple", 20,10),
			items.NewPrem("banana", 20,20),
		}
		ns := NewShopUpdater{}
		testShop := newTestShopWithLittleCache()

		expected := make(map[string]map[string]int)
		ns.editLittleMap(testShop, testItems)

		if !reflect.DeepEqual(expected, testShop.littleCache) {
			t.Errorf("Ошибка очистки кэша на верхнем слое. expected %v, got %v", expected, testShop.littleCache)
		}
	})

	t.Run("second layer cleaning", func(t *testing.T) {
		testItems := []Itemer{
			items.New("xBox", 20,10),
			items.NewPrem("banana", 20,20),
		}
		testShop := newTestShopWithLittleCache()
		ns := NewShopUpdater{}

		expected := make(map[string]map[string]int)
		expected["apple"] = make(map[string]int)
		ns.editLittleMap(testShop, testItems)

		if !reflect.DeepEqual(expected, testShop.littleCache) {
			t.Errorf("Ошибка очистки кэша на втором слое. expected %v, got %v", expected, testShop.littleCache)
		}
	})

	t.Run("nothing needs to be deleted here", func(t *testing.T) {
		testItems := []Itemer{
			items.New("xBox", 20,10),
			items.NewPrem("sony", 20,20),
		}
		testShop := newTestShopWithLittleCache()
		ns := NewShopUpdater{}

		expected := make(map[string]map[string]int)
		expected["apple"] = make(map[string]int)
		expected["apple"]["banana"] = 25
		ns.editLittleMap(testShop, testItems)

		if !reflect.DeepEqual(expected, testShop.littleCache) {
			t.Errorf("А вот тут ничего удалять не нужно, т.к. списка на удаление нет в мапе. expected %v, got %v", expected, testShop.littleCache)
		}
	})
}

func TestShop_getItem(t *testing.T) {

	t.Run("len order 2, index = len order = 2", func(t *testing.T) {
		item := getItem([]string{"apple", "banana"}, 2)
		expected := ""

		if item != expected {
			t.Errorf("должен вернуть %q, вернул %q", expected, item)
		}
	})

	t.Run("len order 2, index > len order = 99", func(t *testing.T) {
		item := getItem([]string{"apple", "banana"}, 2)
		expected := ""

		if item != expected {
			t.Errorf("должен вернуть %q, вернул %q", expected, item)
		}
	})

	t.Run("len order 2, index < len order > 0", func(t *testing.T) {
		item := getItem([]string{"apple", "banana"}, 1)
		expected := "banana"

		if item != expected {
			t.Errorf("должен вернуть %q, вернул %q", expected, item)
		}
	})

	t.Run("len order 2, index < len order < 0", func(t *testing.T) {
		item := getItem([]string{"apple", "banana"}, -99)
		expected := ""

		if item != expected {
			t.Errorf("должен вернуть %q, вернул %q", expected, item)
		}
	})

	t.Run("len order 0, index = 0", func(t *testing.T) {
		item := getItem([]string{}, 0)
		expected := ""

		if item != expected {
			t.Errorf("должен вернуть %q, вернул %q", expected, item)
		}
	})
}

func TestShop_NewShopUpdate_NewShopPrepare(t *testing.T) {
	//testShop := newTestShop()

	shopItems := []Itemer{
		items.New("apple", 20,10),
		items.NewPrem("banana", 20,20),
	}
	shopUsers := []Useer{
		users.New("dasha", 10),
		users.NewPrem("masha", 500),
	}
	testShop := NewShopPrepare(shopItems, shopUsers)

	mock := &spyNewShopUpdater{}
	testShop.updater = mock

	testShop.NewShopUpdate(shopItems, shopUsers)

	realMock, ok := testShop.updater.(*spyNewShopUpdater)
	if !ok {
		t.Errorf("We expected to see *spyNewShopUpdater as an updater")
	}
	if realMock.callCounter != 4 {
		t.Errorf("Неправильно работает апдейтер шопа")
	}
}

func TestShop_initBigShop(t *testing.T) {
	testShop := newTestShopWithoutMock()

	want := make(map[string]int)
	testShop.updater.initBigShop(testShop)
	got := testShop.bigCache

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Ошибка очистки кэша на верхнем слое. expected")
	}
}

func TestShop_makeHash (t *testing.T) {

	t.Run("normal hash", func(t *testing.T) {
		order := []string{"banana", "apple", "potato"}

		want := "applebananapotato"
		got := makeHash(order)

		if got != want {
			t.Errorf("неправильно делает хэш для длинных заказов")
		}
	})

	t.Run("jesuit hash", func(t *testing.T) {
		order := []string{"", "", ""}

		want := ""
		got := makeHash(order)

		if got != want {
			t.Errorf("неправильно делает хэш для длинных заказов")
		}
	})

}

func TestShop_writeBigCache (t *testing.T) {

	t.Run("normal hash", func(t *testing.T) {
		testShop := newTestShopWithoutMock()
		testShop.writeBigCache("applebanana", 20)

		want := 20
		got, _ := testShop.checkBigCache("applebanana")

		if want != got {
			t.Errorf("не работает чтение длинного хэша")
		}
	})

	t.Run("normal hash", func(t *testing.T) {
		testShop := newTestShopWithoutMock()
		testShop.writeBigCache("", 20)

		want := 20
		got, _ := testShop.checkBigCache("")

		if want != got {
			t.Errorf("не работает чтение длинного хэша")
		}
	})
}

func TestShop_collectOrder (t *testing.T) {

	t.Run("1 usual item, 1 prem item", func(t *testing.T) {
		testShop := newTestShop()
		order := []string{"apple", "banana"}

		want := 40
		got, _ := testShop.collectOrder(order, "masha")

		if want != got {
			t.Errorf("неправильно считает сумму заказа. 1 обчный, 1 прем айтем")
		}
	})

	t.Run("2 usual item", func(t *testing.T) {
		testShop := newTestShop()
		order := []string{"apple", "apple"}

		want := 20
		got, _ := testShop.collectOrder(order, "masha")

		if want != got {
			t.Errorf("неправильно считает сумму заказа, 2 обычных айтема")
		}
	})

	t.Run("2 prem item", func(t *testing.T) {
		testShop := newTestShop()
		order := []string{"banana", "banana"}

		want := 60
		got, _ := testShop.collectOrder(order, "masha")

		if want != got {
			t.Errorf("неправильно считает сумму заказа, 2 прем заказа")
		}
	})

	t.Run("non-existent item", func(t *testing.T) {
		testShop := newTestShop()
		order := []string{"potato", "potatotato"}

		_, err := testShop.collectOrder(order, "masha")
		wantErr := errors.New("товара не существует")

		if !reflect.DeepEqual(err, wantErr) {
			t.Errorf("неправильно возвращает ошибку о несуществующем товаре: %q, %q", err, wantErr)
		}
	})

	t.Run("nil list item", func(t *testing.T) {
		testShop := newTestShop()
		var order []string

		sum, _ := testShop.collectOrder(order, "masha")
		want := 0

		if sum != want {
			t.Errorf("неправильно реагирует на нил-ордер")
		}
	})

}

func TestShop_checkItems (t *testing.T) {
	
}

func newTestShop() *shop {
	shopItems := []Itemer{
		items.New("apple", 20,10),
		items.NewPrem("banana", 20,20),
	}
	shopUsers := []Useer{
		users.New("dasha", 10),
		users.NewPrem("masha", 500),
	}

	updater := &spyNewShopUpdater{}

	return newShop(shopItems, shopUsers, updater)
}

func newTestShopWithoutMock() *shop {
	shopItems := []Itemer{
		items.New("apple", 20,10),
		items.NewPrem("banana", 20,20),
	}
	shopUsers := []Useer{
		users.New("dasha", 10),
		users.NewPrem("masha", 500),
	}

	updater := &NewShopUpdater{}

	return newShop(shopItems, shopUsers, updater)
}

func newTestShopWithLittleCache () *shop {
	testShop := newTestShop()
	testShop.writeLittleCache([]string{"apple", "banana"}, 25)

	return testShop
}