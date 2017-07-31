package main

import "sync"

type Home struct {
	Money int
	Size  int
}

type Market struct {
	Total int
}

func NewHome() *Home {
	return &Home{Money: 100, Size: 24}
}

type Storage struct {
	sync.Mutex
	profiles map[int64]Profile
	homes    map[int64]*Home
	market   *Market
}

func NewStorage() *Storage {
	return &Storage{
		profiles: make(map[int64]Profile),
		homes:    make(map[int64]*Home),
		market:   &Market{Total: 42},
	}
}

func (s *Storage) GetProfile(id int64) Profile {
	s.Lock()
	defer s.Unlock()
	_, ok := s.profiles[id]
	if !ok {
		s.profiles[id] = make(Profile)
	}
	return s.profiles[id]
}

func (s *Storage) GetHome(playerID int64) *Home {
	s.Lock()
	defer s.Unlock()
	_, ok := s.homes[playerID]
	if !ok {
		s.homes[playerID] = NewHome()
	}
	return s.homes[playerID]
}

func (s *Storage) GetMarket() *Market {
	return s.market
}
