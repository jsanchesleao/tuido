package main

import (
	"fmt"
	"tuido/components/inputbar"
	"tuido/components/todolist"

	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	todolist todolist.TodoList
	inputbar inputbar.InputBar
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "backspace":
			input, cmd := m.inputbar.Update(inputbar.BackspaceMsg{})
			m.inputbar = input
			cmds = append(cmds, cmd)
		case "enter":
			if len(m.inputbar.Text) > 0 {
				todolist, todolistCmd := m.todolist.Update(todolist.AddMsg{Text: m.inputbar.Text})
				inputbar, inputbarCmd := m.inputbar.Update(inputbar.ClearMsg{})
				m.todolist = todolist
				m.inputbar = inputbar
				cmds = append(cmds, inputbarCmd, todolistCmd)
			}
		default:
			if len(msg.String()) == 1 {
				input, cmd := m.inputbar.Update(inputbar.SendKeyMsg{Key: msg.String()})
				m.inputbar = input
				cmds = append(cmds, cmd)
			}
		}
	}
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	todolist := m.todolist.View()
	inputbar := m.inputbar.View()
	return fmt.Sprintf("\n%s\n%s", todolist, inputbar)
}

func initMainModel() mainModel {
	return mainModel{
		todolist: todolist.NewTodoList(),
		inputbar: inputbar.NewInputBar("Add todo", ""),
	}
}

func main() {
	startApp()
}

func startApp() {

	app := tea.NewProgram(initMainModel())
	err := app.Start()
	if err != nil {
		panic(err)
	}

}
