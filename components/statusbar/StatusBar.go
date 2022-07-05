package statusbar

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type StatusBar struct {
  Text string
}
   
func NewStatusBar(text string) StatusBar {
  return StatusBar {Text: text}
}

func (m StatusBar) Init() tea.Cmd {
  return nil
}

func (m StatusBar) Update(msg tea.Msg) (StatusBar, tea.Cmd) {
  return m, nil
}

func (m StatusBar) View() string {
  return fmt.Sprintf("%s", m.Text)
}

