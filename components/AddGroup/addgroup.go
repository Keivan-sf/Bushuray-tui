package addgroup

import (
	// "fmt"
	cmds "bushuray-tui/commands"
	"bushuray-tui/utils"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width      int
	Height     int
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func InitialModel() Model {
	m := Model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Group name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.Width = 50
			t.CharLimit = 20
		case 1:
			t.Placeholder = "URL (optional)"
			t.Width = 50
			t.CharLimit = 2000
		}

		m.inputs[i] = t
	}

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			commands := make([]tea.Cmd, len(m.inputs))
			commands = append(commands, cmds.ExitAddGroupView)
			m.reset()
			commands = append(commands, m.adjustToNewFocus(commands)...)
			return m, tea.Batch(commands...)

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			commands := make([]tea.Cmd, len(m.inputs))

			if s == "enter" && m.focusIndex == len(m.inputs) {
				commands = append(commands, cmds.SubmitGroup(m.inputs[0].Value(), m.inputs[1].Value()))
				m.reset()
			} else if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else if s == "down" || s == "tab" {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			commands = append(commands, m.adjustToNewFocus(commands)...)
			return m, tea.Batch(commands...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Model) View() string {
	var views []string
	for i := range m.inputs {
		views = append(views, m.inputs[i].View())
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	views = append(views, "")
	views = append(views, *button)
	views = append(views, "")
	views = append(views, utils.GenHelp([]string{"esc"}, []string{"cancel"}))
	container := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, lipgloss.JoinVertical(lipgloss.Top, views...))
	return container

}

func (m Model) SetWH(width int, height int) Model {
	m.Height = height
	m.Width = width
	return m
}
