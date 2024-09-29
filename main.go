package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea" // import bubbletea with alias tea
)

// Defining the model
type model struct {
    choice   []string            // for defining the choices
    cursor   int                 // for the user cursor
    selected map[int]struct{}    // for the selected choice of the user
}

// Func to initialize the model
func InitialModel() model {
    return model{
        choice:   []string{"choice 1", "choice 2", "choice 3"},
        selected: make(map[int]struct{}),
    }
}

// Init method
func (m model) Init() tea.Cmd {
    return nil
}

// Update Method
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "up":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down":
            if m.cursor < len(m.choice)-1 {
                m.cursor++
            }
        case "enter":
            if _, ok := m.selected[m.cursor]; ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }
    return m, nil
}

// View method to render everything
func (m model) View() string {
    s := ""

    // Iterate over our choices
    for i, choice := range m.choice {
        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func main() {
    p := tea.NewProgram(InitialModel())
    if err := p.Start(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
