package pasteprofile

import (
	cmds "github.com/Keivan-sf/Bushuray-tui/commands"
	"github.com/Keivan-sf/Bushuray-tui/components/shared"
	"github.com/Keivan-sf/Bushuray-tui/global"
	servercmds "github.com/Keivan-sf/Bushuray-tui/lib/ServerCommands"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width          int
	Height         int
	focusIndex     int
	CurrentGroupId int
	inputs         []textinput.Model
	cursorMode     cursor.Mode
}

func InitialModel() Model {
	m := Model{
		inputs: make([]textinput.Model, 1),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle.Background(global.GetBgColor())
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Paste profiles here"
			t.Focus()
			t.TextStyle = focusedStyle.Background(global.GetBgColor())
			t.PromptStyle = focusedStyle.Background(global.GetBgColor())
			t.Cursor.TextStyle = focusedStyle.Background(global.GetBgColor())
			t.PlaceholderStyle = placeHolderStyle.Background(global.GetBgColor())
			t.Width = 50
			t.CharLimit = 40000
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
			commands = append(commands, cmds.ExitPasteProfileView)
			m.reset()
			commands = append(commands, m.adjustToNewFocus(commands)...)
			return m, tea.Batch(commands...)

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			commands := make([]tea.Cmd, len(m.inputs))

			if s == "enter" && m.focusIndex == len(m.inputs) && m.inputs[0].Value() != "" {
				servercmds.AddProfiles(m.inputs[0].Value(), m.CurrentGroupId)
				commands = append(commands, cmds.ExitPasteProfileView)
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
	views = append(views, element_style.Render(bg_style.Width(54).Render(button)))
	views = append(views, "")
	views = append(views, element_style.Render(bg_style.Width(54).Render(help_text)))
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
