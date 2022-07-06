package todolist

import (
	"fmt"
	"testing"
)

func TestInitList(t *testing.T) {
	list := NewTodoList()

	if len(list.Items) != 0 {
		t.Fatalf("Expected list to initialize with length 0, but was %d", len(list.Items))
	}
}

func TestAddItems(t *testing.T) {
	list := NewTodoList()
	list.add("First Item")
	list.add("Second Item")
	list.add("Third Item")

	if len(list.Items) != 3 {
		t.Fatalf("Expected list to be 3 after additions, but was %d", len(list.Items))
	}

	if list.Items[0].Text != "First Item" {
		t.Fatalf("Unexpected text on first item %q", list.Items[0].Text)
	}

	if list.Items[0].Done {
		t.Fatal("List item was not expected to be Done")
	}
}

func TestCompleteAndClean(t *testing.T) {
	list := NewTodoList()
	list.add("First")
	list.add("Second")
	list.add("Third")

	list.Items[1].Done = true

	if !list.Items[1].Done {
		t.Fatal("Item indexed [1] should be Done")
	}

	list.clearCompleted()

	if len(list.Items) != 2 {
		t.Fatalf("Expected list to clear to size 2, but was %d", len(list.Items))
	}
}

func TestMoveItem(t *testing.T) {
  list := NewTodoList()
  list.add("First")
	list.add("Second")
	list.add("Third")

  list.focus()
  list.selectNext()
  list.moveUp()

  if list.GetSelected().Text != "Second" {
    t.Fatalf("Expected selection to stay on the same item, but it was in %s", list.GetSelected().Text)
  }

  list.selectNext()
  if list.GetSelected().Text != "First" {
    t.Fatalf("Expected item with Text First to be in the second position, but the item was %s", list.GetSelected().Text)
  }

  list.moveDown()
  if list.GetSelected().Text != "First" {
    t.Fatalf("Expected item with Text First to be in the third position, but the item was %s", list.GetSelected().Text)
  }

  list.selectPrev()
  if list.GetSelected().Text != "Third" {
    t.Fatalf("Expected item with Text Third to be in the second position, but the item was %s", list.GetSelected().Text)
  }
}

func TestUpdate(t *testing.T) {

	list := NewTodoList()
	list, _ = list.Update(AddMsg{Text: "First"})
	list, _ = list.Update(AddMsg{Text: "Second"})
	list, _ = list.Update(AddMsg{Text: "Third"})

	selected := list.GetSelected()
	if selected != nil {
		t.Fatalf("List should initialize without a selected item, and unfocused")
	}
	list, _ = list.Update(FocusMsg{})
	selected = list.GetSelected()
	if selected == nil || selected.Text != "First" {
		t.Fatalf("List should begin focusing at the first item")
	}

	list, _ = list.Update(SelectNextMsg{})
	selected = list.GetSelected()
	if selected == nil || selected.Text != "Second" {
		t.Fatalf("Should have selected Second, but was %v", selected)
	}

	list, _ = list.Update(SelectNextMsg{})
	list, _ = list.Update(SelectNextMsg{})
	list, _ = list.Update(SelectNextMsg{})
	selected = list.GetSelected()
	if selected == nil || selected.Text != "Third" {
		t.Fatalf("Selection should not overflow upwards. Selection should have been Third, but was %v", selected)
	}

	list, _ = list.Update(SelectPrevMsg{})
	selected = list.GetSelected()
	if selected == nil || selected.Text != "Second" {
		t.Fatalf("Should have selected Second, but was %v", selected)
	}

	list, _ = list.Update(UnfocusMsg{})
	selected = list.GetSelected()
	if selected != nil {
		t.Fatalf("Selection should be nil after unfocus, but was %v", selected)
	}
}
