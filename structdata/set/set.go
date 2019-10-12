// Package set creates a ItemSet data structure for the Item type
package set

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the Set
type Item generic.Type

// ItemSet the set of Items
type ItemSet struct {
	items map[Item]bool
	lock  sync.RWMutex
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *ItemSet) Add(t Item) *ItemSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *ItemSet) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[Item]bool)
}

// Delete removes the Item from the Set and returns Has(Item)
func (s *ItemSet) Delete(item Item) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the Item
func (s *ItemSet) Has(item Item) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.items[item]
	return ok
}

// Items returns the Item(s) stored
func (s *ItemSet) Items() []Item {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := []Item{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *ItemSet) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

// Union returns a new set with elements from both
// the given sets
func (s *ItemSet) Union(s2 *ItemSet) *ItemSet {
	s3 := ItemSet{}
	s3.items = make(map[Item]bool)
	s.lock.RLock()
	for i := range s.items {
		s3.items[i] = true
	}
	s.lock.RUnlock()
	s2.lock.RLock()
	for i := range s2.items {
		_, ok := s3.items[i]
		if !ok {
			s3.items[i] = true
		}
	}
	s2.lock.RUnlock()
	return &s3
}

// Intersection returns a new set with elements that exist in
// both sets
func (s *ItemSet) Intersection(s2 *ItemSet) *ItemSet {
	s3 := ItemSet{}
	s3.items = make(map[Item]bool)
	s.lock.RLock()
	s2.lock.RLock()
	defer s.lock.RUnlock()
	defer s2.lock.RUnlock()
	for i := range s2.items {
		_, ok := s.items[i]
		if ok {
			s3.items[i] = true
		}
	}
	return &s3
}

// Difference returns a new set with all the elements that
// exist in the first set and don't exist in the second set
func (s *ItemSet) Difference(s2 *ItemSet) *ItemSet {
	s3 := ItemSet{}
	s3.items = make(map[Item]bool)
	s.lock.RLock()
	s2.lock.RLock()
	defer s.lock.RUnlock()
	defer s2.lock.RUnlock()
	for i := range s.items {
		_, ok := s2.items[i]
		if !ok {
			s3.items[i] = true
		}
	}
	return &s3
}

// Subset returns true if s is a subset of s2
func (s *ItemSet) Subset(s2 *ItemSet) bool {
	s.lock.RLock()
	s2.lock.RLock()
	defer s.lock.RUnlock()
	defer s2.lock.RUnlock()
	for i := range s.items {
		_, ok := s2.items[i]
		if !ok {
			return false
		}
	}
	return true
}
