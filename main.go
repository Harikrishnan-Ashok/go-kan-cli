package main

import (
    "fmt"
    "time"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    count string
}

type tickMsg time.Time

func tick() tea.Cmd {
    return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

func (m model) Init() tea.Cmd {
    return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    case tickMsg:
        // Rotate through the spinner states
        switch m.count {
        case "/":
            m.count = "-"
        case "-":
            m.count = "\\"
        case "\\":
            m.count = "--"
        case "--":
            m.count = "/" // Loop back to the first state
        }
        return m, tick()
    }
    return m, nil
}

func (m model) View() string {
    return fmt.Sprintf(" loading : %s\n press q to quit ", m.count)
}

func main() {
    p := tea.NewProgram(model{count: "/"})
    p.Start()
}
