package todolist

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TodoList struct {
	Items    []TodoItem
	hasFocus bool
	selected int
}

type TodoItem struct {
	Text string
	Done bool
}

type FocusMsg struct {}
type UnfocusMsg struct {}
type SelectNextMsg struct {}
type SelectPrevMsg struct {}
type AddMsg struct {
  Text string
}

func NewTodoList() TodoList {
	return TodoList{Items: make([]TodoItem, 0)}
}

func (list *TodoList) add(item string) {
	list.Items = append(list.Items, TodoItem{Text: item, Done: false})
}

func (list *TodoList) clearCompleted() {
	items := []TodoItem{}
	for _, item := range list.Items {
		if !item.Done {
			items = append(items, item)
		}
	}
	list.Items = items
}

func (list *TodoList) focus() {
	list.hasFocus = true
	list.selected = 0
}
func (list *TodoList) unfocus() {
	list.hasFocus = false
}
func (list *TodoList) GetSelected() *TodoItem {
	if list.hasFocus && len(list.Items) > list.selected {
		return &list.Items[list.selected]
	} else {
		return nil
	}
}
func (list *TodoList) selectNext() {
	if list.hasFocus && list.selected < (len(list.Items)-1) {
		list.selected += 1
	}
}
func (list *TodoList) selectPrev() {
	if list.hasFocus && list.selected > 0 {
		list.selected -= 1
	}

}

func (item *TodoItem) complete() {
	item.Done = true
}

func (m TodoList) Init() tea.Cmd {
	return nil
}

func (m TodoList) Update(msg tea.Msg) (TodoList, tea.Cmd) {
  switch msg := msg.(type) {
    case AddMsg:
      m.add(msg.Text)
    case FocusMsg:
      m.focus()
    case UnfocusMsg:
      m.unfocus()
    case SelectNextMsg:
      m.selectNext()
    case SelectPrevMsg:
      m.selectPrev()
  }
	return m, nil
}

var styleSelected = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4"))

var styleUnselected = lipgloss.NewStyle().
	Bold(false)

func (m TodoList) View() string {
	if len(m.Items) == 0 {
		return "There are no Todos"
	}
	s := ""
	for index, item := range m.Items {
		heading := "[ ]"
		if item.Done {
			heading = "[x]"
		}
		style := styleUnselected
		if m.hasFocus && m.selected == index {
			style = styleSelected
		}
		s += style.Render(fmt.Sprintf("%s %s", heading, item.Text))
    s += "\n"
	}

	return s
}
