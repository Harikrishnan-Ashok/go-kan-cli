package models

import (
    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
)

type Column struct {
    Status Status
    List   list.Model
}

func NewColumn(status Status) Column {
    return Column{
        Status: status,
        List:   list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
    }
}

// Implement Update and View methods here
// We'll come back to these later
