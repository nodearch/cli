package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type (
	errMsg error
)

type model struct {
	projectName textinput.Model
	err         error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter project name"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	return model{
		projectName: ti,
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Quit the program when the user presses Enter
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			// Quit the program without creating a project
			m.projectName.SetValue("") // Clear input if quitting
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.projectName, cmd = m.projectName.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Project Name?\n\n%s\n\n%s",
		m.projectName.View(),
		"(Enter to confirm, Esc to cancel)",
	) + "\n"
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new NodeArch App",
	Long:  "Create a new NodeArch App",
	Run: func(cmd *cobra.Command, args []string) {

		p := tea.NewProgram(initialModel())
		finalModel, err := p.Run()

		if err != nil {
			fmt.Printf("Error starting program: %v\n", err)
			os.Exit(1)
		}

		m := finalModel.(model)

		projName := m.projectName.Value()

		if projName == "" {
			fmt.Println("Project creation canceled.")
			os.Exit(0) // Exit gracefully when no project name is provided
		}

		// Print confirmation
		fmt.Printf("Creating new NodeArch app: %s\n", projName)

		// Optionally create the project directory
		projectDir := filepath.Join(".", projName)
		// if err := os.MkdirAll(projectDir, os.ModePerm); err != nil {
		// 	fmt.Printf("Failed to create project directory: %v\n", err)
		// 	os.Exit(1)
		// }

		fmt.Printf("Project directory '%s' created successfully.\n", projectDir)
		// Add more logic here to scaffold the project
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
