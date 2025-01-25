package create

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func createProjectNameInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter project name"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	return ti
}

func projectNameUpdate(m createModel, msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.projectName.Value() == "" {
				m.err = fmt.Errorf("project name cannot be empty")
				m.step = ExitStep
				return m, tea.Quit
			} else {
				m.step = TemplateListStep
				return m, nil
			}
		}
	case error:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd

	m.projectName, cmd = m.projectName.Update(msg)
	return m, cmd
}

func projectNameView(m createModel) string {
	return fmt.Sprintf(
		"Project Name?\n\n%s\n\n%s",
		m.projectName.View(),
		"(Enter to confirm, Esc to cancel)",
	) + "\n"
}
