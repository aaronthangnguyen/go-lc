package set_test

import (
	"github.com/aaronthangnguyen/go-lc/set"
	"testing"
)

func TestSetNew(t *testing.T) {
	s := set.New[int]()
	if len(*s) != 0 {
		t.Fatal("Set should be empty")
	}
}

func TestSetAdd(t *testing.T) {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)

	if len(*s) != 2 {
		t.Fatalf("Set should have %d elements, got %d", 2, len(*s))
	}

	s.Add(3)
	if len(*s) != 3 {
		t.Fatalf("Set should have %d elements, got %d", 3, len(*s))
	}
}

func TestSetExists(t *testing.T) {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)

	if !s.Exists(1) {
		t.Fatal("Set should have 1")
	}

	if s.Exists(3) {
		t.Fatal("Set should not have 3")
	}
}

func TestSetRemove(t *testing.T) {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.Remove(3)
	if s.Exists(3) {
		t.Fatal("3 should have been removed from Set")
	}
}
