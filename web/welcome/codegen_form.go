package welcome

import (
	"context"
	"github.com/goradd/gofile/pkg/sys"
	"github.com/goradd/goradd/pkg/page"
	"github.com/goradd/goradd/pkg/page/control"
	"os/exec"
	"path/filepath"
)

const CodegenPath = "/goradd/build.g"
const CodegenID = "CodeGenerator"

const (
	CodegenRefreshAction = iota + 1
)

type CodegenForm struct {
	control.FormBase
}

func (f *CodegenForm) Init(ctx context.Context, formID string) {
	f.FormBase.Init(ctx, formID)
	f.AddRelatedFiles()
	f.createControls(ctx)
}

func (f *CodegenForm) createControls(ctx context.Context) {
	control.NewPanel(f, "infoPanel")
}

func (f *CodegenForm) Action(ctx context.Context, a page.ActionParams) {
	switch a.ID {
	case CodegenRefreshAction:
		f.Refresh()
	}
}

func (f *CodegenForm) LoadControls(ctx context.Context) {
	v, _ := page.GetContext(ctx).FormValue("cmd")
	switch v {
	case "codegen":
		result := f.startCodegen()
		control.GetPanel(f, "infoPanel").SetText(result)
	case "run":
		result := f.startApp()
		control.GetPanel(f, "infoPanel").SetText(result)

	}
}

func (f *CodegenForm) startCodegen() string {
	var result = "Running generate goradd-project/codegen/cmd/build.go ...\n"

	codegenLoc := filepath.Join("goradd-project", "codegen", "cmd", "build.go")
	cmdResult, err := sys.ExecuteShellCommand("gofile generate " + codegenLoc)
	if err != nil {
		result += string(err.(*exec.ExitError).Stderr)
	} else {
		result += string(cmdResult)
		result += "Success!"
	}
	return result
}

func (f *CodegenForm) startApp() string {
	var result = "Running go run goradd-project/main ...\n"

	app := filepath.Join("goradd-project", "main")
	cmdResult, err := sys.ExecuteShellCommand("go run " + app)
	if err != nil {
		result += string(err.(*exec.ExitError).Stderr)
	} else {
		result += string(cmdResult)
		result += "Success!"
	}
	return result
}

func init() {
	page.RegisterForm(CodegenPath, &CodegenForm{}, CodegenID)
}
