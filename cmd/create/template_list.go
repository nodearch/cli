package create

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type item string

func (i item) FilterValue() string {
	return ""
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)

	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render

	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func createTemplateList() list.Model {
	items := []list.Item{
		item("Empty"),
		item("Express"),
		item("Socket IO"),
	}

	const defaultWidth = 20
	const listHeight = 14

	li := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	li.Title = "Choose a template"
	li.SetShowStatusBar(false)
	li.SetFilteringEnabled(false)
	li.Styles.Title = titleStyle
	li.Styles.PaginationStyle = paginationStyle
	li.Styles.HelpStyle = helpStyle
	li.KeyMap.Quit.SetKeys("esc")
	li.KeyMap.Quit.SetHelp("esc", "quit")

	return li
}

func templateListUpdate(m createModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.templateList.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "enter":
			i, ok := m.templateList.SelectedItem().(item)
			if ok {
				m.selectedTemplate = string(i)
				m.step = DoneStep
				return m, tea.Quit
			}

			return m, nil
		}
	}

	var cmd tea.Cmd
	m.templateList, cmd = m.templateList.Update(msg)
	return m, cmd
}

func templateListView(m createModel) string {
	return "\n" + m.templateList.View()
}
