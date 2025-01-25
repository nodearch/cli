package create

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type createModel struct {
	projectName      textinput.Model
	templateList     list.Model
	selectedTemplate string
	step             string
	quitting         bool
	err              error
}
