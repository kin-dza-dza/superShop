package shop

type spyNewShopUpdater struct {
	callCounter int
}

func (sn *spyNewShopUpdater) userUpdate(s *shop, users []Useer) {
	sn.callCounter++
}

func (sn *spyNewShopUpdater) itemsUpdate(s *shop, items []Itemer) {
	sn.callCounter++
}

func (sn *spyNewShopUpdater) editLittleMap(s *shop, items []Itemer) {
	sn.callCounter++
}

func (sn *spyNewShopUpdater) initBigShop(s *shop) {
	sn.callCounter++
}
