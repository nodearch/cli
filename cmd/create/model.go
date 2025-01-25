package create

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type createModel struct {
	projectName      textinput.Model
	templateList     list.Model
	selectedTemplate string
	step             StepName
	err              error
}

type StepName string

const (
	ProjectNameStep  StepName = "projectName"
	TemplateListStep StepName = "templateList"
	ExitStep         StepName = "exit"
	DoneStep         StepName = "done"
)
