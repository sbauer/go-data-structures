package lists

import (
	"testing"
)

func Test_CreateLinkedList(t *testing.T) {
	list := &LinkedList{}

	if list == nil {
		t.Errorf("Nil linked list")
	}
}

func Test_InsertCreatesNewHead(t *testing.T) {
	list := &LinkedList{}

	if list.head != nil {
		t.Fatalf("Expected nil head")
	}

	list.Insert("testing")

	if list.head == nil {
		t.Fatalf("Expected non-nil head")
	}

	if list.head.data != "testing" {
		t.Errorf("Expected head node value of testing")
	}
}

func Test_InsertMovesPreviousHeadToNext(t *testing.T) {
	list := &LinkedList{}

	list.Insert("first insert")
	list.Insert("second insert")

	if list.head.data != "second insert" {
		t.Errorf("Unexpected first node value")
	}

	if list.head.next.data != "first insert" {
		t.Errorf("Unexpected second node value")
	}
}