package shop

type NewShopUpdater struct {}

func (ns NewShopUpdater) userUpdate(s *shop, users []Useer) {
	for _, user := range users {
		_, ok := s.users[user.GetName()]
		if ok {
			delete(s.users, user.GetName())
			if user.GetCash() != deletekey {
				s.users[user.GetName()] = user
			}
		} else {
			s.users[user.GetName()] = user
		}
	}
}

func (ns NewShopUpdater) itemsUpdate(s *shop, items []Itemer) {
	for _, item := range items {
		_, ok := s.items[item.GetName()]
		if ok {
			delete(s.items, item.GetName())

		}
		if item.GetPrice() > deletekey {
			s.items[item.GetName()] = item
		}
	}
}

func (ns NewShopUpdater) editLittleMap(s *shop, items []Itemer) {
	for _, item := range items {
		_, ok := s.littleCache[item.GetName()]
		if ok {
			delete(s.littleCache, item.GetName())
		}
	}

	for _, item1 := range s.littleCache {
		for _, item2 := range items {
			_, ok := item1[item2.GetName()]
			if ok {
				delete(item1, item2.GetName())
			}
		}
	}
}

func (ns NewShopUpdater) initBigShop(s *shop) {
	s.bigCache = make(map[string]int)
}
