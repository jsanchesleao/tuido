package inputbar

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var blinking = lipgloss.NewStyle().Blink(true)
const block = "\u2588"

type InputBar struct {
  Prompt string
  Text string
}

type BackspaceMsg struct {}
type SendKeyMsg struct {
  Key string
}
type PromptChangeMsg struct {
  Prompt string
}
type ClearMsg struct {}

func NewInputBar(prompt, text string) InputBar {
  return InputBar{Prompt: prompt, Text: text}
}

func (m InputBar) Init() tea.Cmd {
  return nil
}

func (m InputBar) Update(msg tea.Msg) (InputBar, tea.Cmd) {
  switch msg := msg.(type) {
    case SendKeyMsg:
      m.Text += msg.Key
    case BackspaceMsg:
      if len(m.Text) > 0 {
        m.Text = m.Text[:len(m.Text) - 1]
      }
    case PromptChangeMsg:
      m.Prompt = msg.Prompt
    case ClearMsg:
      m.Text = ""
  }
  return m, nil
}

func (m InputBar) display() string {
  return fmt.Sprintf("%s: %s", m.Prompt, m.Text)
}

func (m InputBar) View() string {
  return m.display() + blinking.Render(block)
}

