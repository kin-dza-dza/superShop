//hello, iteration 1
//hello, iteration2
package main

import (
	"fmt"
)

import "golang-book/shop_git2/superShop/items"
import "golang-book/shop_git2/superShop/users"
import "golang-book/shop_git2/superShop/shop"

func main() {
	fmt.Println("Hello World")

	items := []shop.Itemer{
		items.New("apple", 20, 10),
		items.New("tea", 20, 50),
		items.New("salt", 0, 80),

		items.NewPrem("banana", 20, 20),
		items.NewPrem("potato", 20, 20),
	}

	users := []shop.Useer{
		users.New("dasha", 10),
		users.New("natasha", 500),
		users.New("kirill", -1000),

		users.NewPrem("masha", 500),
	}

	//создает новый шоп и тут же обновляет его,
	// убирая все товары с невозможной стоимостью
	shop2 := shop.NewShopPrepare(items, users)

	//ненужная штука для проверки CI. Можно убирать потом
	fl := true
	shop.HelloPrint(fl)

	err := shop2.Buy([]string{"apple", "banana"}, "masha")
	shop2.Intrigue(err, "masha")

	err = shop2.Buy([]string{"apple", "banana"}, "natasha")
	shop2.Intrigue(err, "natasha")
}