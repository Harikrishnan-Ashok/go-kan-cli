package models

import (
    "github.com/charmbracelet/bubbles/textarea"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type Form struct {
    Title       textinput.Model
    Description textarea.Model
    Status      Status
}

func NewForm() Form {
    return Form{
        Title:       textinput.New(),
        Description: textarea.New(),
        Status:      Todo,
    }
}

// Implement Init, Update, and View methods here
// We'll come back to these later
