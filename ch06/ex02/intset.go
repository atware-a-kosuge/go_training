// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func (s *IntSet) debugPrint(prefix string) {
	for i, tword := range s.words {
		fmt.Printf("%v %02d(%03d): %08b(count: %v)\n", prefix, i, i*64, tword, bitCount(tword))
	}
}

// Len return length
func (s *IntSet) Len() int {
	s.debugPrint("")
	count := 0
	for _, tword := range s.words {
		count += bitCount(tword)
	}
	return count
}

func bitCount(tword uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(tword >> i & 1)
	}
	return count
}

// Remove value in set
func (s *IntSet) Remove(x int) {
	s.debugPrint("before")
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.words[word] &= ^(1 << bit)
	}
	s.debugPrint("after")
}

// Clear values
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

// Copy IntSet
func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: append([]uint64{}, s.words...)}
}

// AddAll values to set
func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}
