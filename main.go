package main

import (
	"fmt"
	"os"
	"path"
	"tuido/actions"
	"tuido/components/inputbar"
	"tuido/components/statusbar"
	"tuido/components/todolist"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	ListMode = iota
	InputMode
  Exit
)

type mainModel struct {
	todolist  todolist.TodoList
	inputbar  inputbar.InputBar
	statusbar statusbar.StatusBar
	mode      int
}

type SwitchModeMessage struct {
	Mode int
}

var todoFile string

func switchToMode(mode int) tea.Cmd {
  return func() tea.Msg {
    return SwitchModeMessage{Mode: mode}
  }
}

func emit(cmd tea.Msg) tea.Cmd {
  return func() tea.Msg {
    return cmd
  }
}

func (m mainModel) Init() tea.Cmd {
	return tea.Batch(
    emit(todolist.FocusMsg{}),
    switchToMode(ListMode),
  )
}

func (m mainModel) updateListMode(msg tea.Msg) (mainModel, tea.Cmd) {
	cmds := []tea.Cmd{}
  var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
    case "q":
      return m, func() tea.Msg { actions.Save(todoFile, m.todolist); return SwitchModeMessage{Mode: Exit} }
    case "a":
      cmds = append(cmds, switchToMode(InputMode), emit(todolist.UnfocusMsg{}))
    case "k":
      m.todolist, cmd = m.todolist.Update(todolist.SelectPrevMsg{})
      cmds = append(cmds, cmd)
    case "j":
      m.todolist, cmd = m.todolist.Update(todolist.SelectNextMsg{})
      cmds = append(cmds, cmd)
    case "K":
      m.todolist, cmd = m.todolist.Update(todolist.MoveUpMsg{})
      cmds = append(cmds, cmd)
    case "J":
      m.todolist, cmd = m.todolist.Update(todolist.MoveDownMsg{})
      cmds = append(cmds, cmd)
    case " ":
      m.todolist, cmd = m.todolist.Update(todolist.ToggleCompletionMsg{})
      cmds = append(cmds, cmd)
    case "r":
      m.todolist, cmd = m.todolist.Update(todolist.RevertMsg{})
      cmds = append(cmds, cmd)
    case "c":
      m.todolist, cmd = m.todolist.Update(todolist.CompleteMsg{})
      cmds = append(cmds, cmd)
    case "C":
      m.todolist, cmd = m.todolist.Update(todolist.ClearCompletedMsg{})
		}
	}
	return m, tea.Batch(cmds...)
}

func (m mainModel) updateInputMode(msg tea.Msg) (mainModel, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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
		case "esc":
      cmds = append(cmds, switchToMode(ListMode), emit(todolist.FocusMsg{}))
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

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
    case "ctrl+c":
			return m, tea.Quit
    }

  case SwitchModeMessage:
    m.mode = msg.Mode
	}

  cmds := []tea.Cmd{}
  var cmd tea.Cmd

	switch m.mode {
	case InputMode:
		m, cmd = m.updateInputMode(msg)
    cmds = append(cmds, cmd)
	case ListMode:
		m, cmd = m.updateListMode(msg)
    cmds = append(cmds, cmd)
  case Exit:
    return m, tea.Quit
	}

  m.todolist, cmd = m.todolist.Update(msg)
  cmds = append(cmds, cmd)
  m.inputbar, cmd = m.inputbar.Update(msg)
  cmds = append(cmds, cmd)

  return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
  top := m.todolist.View()
  bottom := m.statusbar.View()

  if m.mode == InputMode {
    bottom = m.inputbar.View()
  }

	return fmt.Sprintf("\n%s\n%s", top, bottom)
}

func initMainModel() mainModel {
	return mainModel{
		todolist: actions.Load(todoFile),
		inputbar: inputbar.NewInputBar("Add todo (Esc to cancel)", ""),
    statusbar: statusbar.NewStatusBar("Commands: [a]dd [c]omplete [r]evert [C]lean [q]uit. Move with [hjkl]"),
	}
}

func main() {
	startApp()
}

func startApp() {
  
  homedir, err := os.UserHomeDir()
  if err != nil {
    panic(err)
  }
  todoFile = path.Join(homedir, ".todos.json")

	app := tea.NewProgram(initMainModel())
	err = app.Start()
	if err != nil {
		panic(err)
	}

}
