package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
    Up     key.Binding
    Down   key.Binding
    Left   key.Binding
    Right  key.Binding
    Enter  key.Binding
    New    key.Binding
    Edit   key.Binding
    Delete key.Binding
    Quit   key.Binding
}

var Keys = KeyMap{
    Up:     key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "move up")),
    Down:   key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "move down")),
    Left:   key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "move left")),
    Right:  key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "move right")),
    Enter:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
    New:    key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new task")),
    Edit:   key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit task")),
    Delete: key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete task")),
    Quit:   key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

func (k KeyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {k.Up, k.Down, k.Left, k.Right},
        {k.Enter, k.New, k.Edit, k.Delete},
        {k.Help, k.Quit},
    }
}
