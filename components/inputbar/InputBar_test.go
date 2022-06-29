package inputbar

import "testing"

func TestInputBarSetup(t *testing.T) {
  bar := NewInputBar("Prompt", "text")

  view := bar.View()
  expected := "Prompt: text"
  if view != expected {
    t.Fatalf("Expected view to render %q but it was %q", expected, view)
  }

}

func TextInsertKeys(t *testing.T) {
  bar := NewInputBar("Prompt", "text")

  bar, _ = bar.Update(SendKeyMsg{Key: "!"})
  bar, _ = bar.Update(SendKeyMsg{Key: " "})
  bar, _ = bar.Update(SendKeyMsg{Key: "X"})

  view := bar.View()
  expected := "Prompt: text! X"
  if view != expected {
    t.Fatalf("Expected view to render %q but it was %q", expected, view)
  }
}

func TestBackspace(t *testing.T) {
  bar := NewInputBar("Prompt", "text")
  bar, _ = bar.Update(BackspaceMsg{})

  view := bar.View()
  expected := "Prompt: tex"
  if view != expected {
    t.Fatalf("Expected view to render %q but it was %q", expected, view)
  }

  bar, _ = bar.Update(BackspaceMsg{})
  bar, _ = bar.Update(BackspaceMsg{})
  bar, _ = bar.Update(BackspaceMsg{})
  bar, _ = bar.Update(BackspaceMsg{})
  bar, _ = bar.Update(BackspaceMsg{})

  view = bar.View()
  expected = "Prompt: "
  if view != expected {
    t.Fatalf("Expected view to render %q but it was %q", expected, view)
  }
}

func TestPromptChange(t *testing.T) {
  
  bar := NewInputBar("Prompt", "text")
  bar, _ = bar.Update(PromptChangeMsg{Prompt: "$"})

  view := bar.View()
  expected := "$: text"
  if view != expected {
    t.Fatalf("Expected view to render %q but it was %q", expected, view)
  }
}

func TestClearMsg(t *testing.T) {
  bar := NewInputBar("Prompt", "text")
  bar, _ = bar.Update(ClearMsg{})

  view := bar.View()
  expected := "Prompt: "
  if view != expected {
    t.Fatalf("Expected view to render %q but it was %q", expected, view)
  }
}
