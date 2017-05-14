package main

import "sync"

type Storage struct {
	sync.Mutex
	players  map[int64]*Player
	profiles map[int64]Profile
}

func NewStorage() *Storage {
	return &Storage{
		players:  make(map[int64]*Player),
		profiles: make(map[int64]Profile),
	}

}

func (s *Storage) GetPlayer(id int64) *Player {
	s.Lock()
	defer s.Unlock()
	_, ok := s.players[id]
	if !ok {
		s.players[id] = &Player{}
	}
	return s.players[id]
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
