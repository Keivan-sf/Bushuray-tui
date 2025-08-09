package addgroup

import (
	// "fmt"
	cmds "bushuray-tui/commands"
	"bushuray-tui/components/shared"
	"bushuray-tui/global"
	servercmds "bushuray-tui/lib/ServerCommands"

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
		t.Cursor.Style = cursorStyle.Background(global.GetBgColor())
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Group name"
			t.Focus()
			t.PromptStyle = focusedStyle.Background(global.GetBgColor())
			t.Cursor.TextStyle = focusedStyle.Background(global.GetBgColor())
			t.PlaceholderStyle = placeHolderStyle.Background(global.GetBgColor())
			t.Width = 50
			t.CharLimit = 20
			t.Cursor.Style = cursorStyle.Background(global.GetBgColor())
		case 1:
			t.Cursor.TextStyle = grayStyle.Background(global.GetBgColor())
			t.Placeholder = "URL (optional)"
			t.PlaceholderStyle = placeHolderStyle.Background(global.GetBgColor())
			t.Width = 50
			t.CharLimit = 2000
			t.Cursor.Style = cursorStyle.Background(global.GetBgColor())
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

			if s == "enter" && m.focusIndex == len(m.inputs) && m.inputs[0].Value() != "" {
				servercmds.AddGroup(m.inputs[0].Value(), m.inputs[1].Value())
				commands = append(commands, cmds.ExitAddGroupView)
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
	bg_style := lipgloss.NewStyle().Background(global.GetBgColor())
	element_style := bg_style.Width(m.Width).Height(1).Align(lipgloss.Center)
	var views []string
	for i := range m.inputs {
		views = append(views, element_style.Render(m.inputs[i].View()))
	}

	button := blurredButton()
	if m.focusIndex == len(m.inputs) {
		button = focusedButton()
	}

	help_text := shared.GenHelp([]string{"esc"}, []string{"cancel"})

	views = append(views, "")
	views = append(views, element_style.Render(bg_style.Width(52).Render(button)))
	views = append(views, "")
	views = append(views, element_style.Render(bg_style.Width(52).Render(help_text)))
	vertical_container := bg_style.Render(lipgloss.JoinVertical(lipgloss.Top, views...))
	content := lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, vertical_container)
	container := lipgloss.NewStyle().Background(global.GetBgColor()).Render(content)
	return container

}

func (m Model) SetWH(width int, height int) Model {
	m.Height = height
	m.Width = width
	return m
}
