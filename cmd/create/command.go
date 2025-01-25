package create

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

func initialModel() createModel {
	projectNameInput := createProjectNameInput()
	templateList := createTemplateList()
	projectNameInput.Focus()

	return createModel{
		projectName:  projectNameInput,
		templateList: templateList,
		err:          nil,
		step:         "projectName",
	}
}

func (m createModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m createModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	massage, ok := msg.(tea.KeyMsg)

	if ok {
		keypress := massage.String()

		if keypress == "ctrl+c" || keypress == "q" || keypress == "esc" {
			m.quitting = true
			return m, tea.Quit
		}
	}

	if m.step == "projectName" {
		return projectNameUpdate(m, msg)
	} else if m.step == "templateList" {
		return templateListUpdate(m, msg)
	} else {
		return m, nil
	}
}

func (m createModel) View() string {
	if m.step == "projectName" {
		return projectNameView(m)
	} else if m.step == "templateList" {
		return templateListView(m)
	} else {
		return ""
	}
}

var Command = &cobra.Command{
	Use:   "create",
	Short: "Create a new NodeArch App",
	Long:  "Create a new NodeArch App",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel())
		finalModel, err := p.Run()

		if err != nil {
			fmt.Printf("Error starting program: %v\n", err)
			os.Exit(1)
		}

		m := finalModel.(createModel)

		if m.err != nil {
			fmt.Println(quitTextStyle.Render(fmt.Sprintf("Error: %v", m.err)))
			os.Exit(1)
		}

		if m.quitting {
			fmt.Println(quitTextStyle.Render("Quitting..."))
			os.Exit(0)
		}

		handler(m)
	},
}
