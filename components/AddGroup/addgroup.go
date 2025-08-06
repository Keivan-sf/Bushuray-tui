package addgroup

import (
	// "fmt"
	"log"

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

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#7287fd"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#8caaee"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	grayStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70"))
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	// blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	blurredButton = grayStyle.Render("[ Submit ]")
)

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
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				log.Println(m.inputs[0].Value(), m.inputs[1].Value())
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
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
	container := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, lipgloss.JoinVertical(lipgloss.Top, views...))
	return container

}

func (m Model) SetWH(width int, height int) Model {
	m.Height = height
	m.Width = width
	return m
}
