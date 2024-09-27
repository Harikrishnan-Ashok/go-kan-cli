package models

type Status int

const (
    Todo Status = iota
    InProgress
    Done
)

type Task struct {
    Status      Status
    Title       string
    Description string
}

func NewTask(status Status, title, description string) Task {
    return Task{Status: status, Title: title, Description: description}
}

// Implement the list.Item interface
func (t Task) FilterValue() string { return t.Title }
func (t Task) Title() string       { return t.Title }
func (t Task) Description() string { return t.Description }
