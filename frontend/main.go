
package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/iafan/goplayspace/client/js/console"
	"github.com/bidaya0/gbatect/converter"
	batecttypes "github.com/bidaya0/gbatect/types"
	"gopkg.in/yaml.v3"
	"github.com/compose-spec/compose-go/loader"
//	"github.com/gopherjs/gopherjs/syscall/js"
//	"fmt"
//	"github.com/iafan/goplayspace/client/component/app"
//	"github.com/iafan/goplayspace/client/js/localstorage"
//  "github.com/gopherjs/gopherjs/js"
)

type Application struct {
	vecty.Core

//	editor *editor.Editor
//	log    *log.Log

	Input   string
	Topic   string
	Imports map[string]string

	// Settings
	Theme            string
	TabWidth         int
	FontWeight       string
	UseWebfont       bool
	HighlightingMode bool
	ShowSidebar      bool

	//Hash      *hash.Hash
	snippetID string

	modifierKey          string
	isLoading            bool
	isCompiling          bool
	isSharing            bool
	isDrawingMode        bool
	hasCompilationErrors bool
	needRender           bool
	showSettings         bool
	showDrawHelp         bool

	// Log properties
	hasRun bool
	err    string
//	events []*api.CompileEvent

	// Draw mode properties
	//actions draw.ActionList

	// Editor properties
	warningLines map[string]bool
	errorLines   map[string]bool
	FromText string
//	undoStack    *undo.Stack
//	changeTimer  *time.Timer
}

func (a *Application) onChange(e *vecty.Event) {
//	console.Log(e)
	a.FromText = e.Value.String()
//	console.Log(a.FromText)
//	console.Log(e.Value.String())
}

type TexA struct {
	vecty.Core
	Input string
}

func (a *TexA) Render() vecty.ComponentOrHTML {
//	c := &elem.HTML{node: a.Input}
		//	return vecty.Tag("div", 		&vecty.HTML{node: a.Input})
	return vecty.Tag("pre",vecty.Tag(
		"code",
		vecty.Markup(vecty.UnsafeHTML(a.Input)),
	))
}
type S struct {
	Name string
}

func (a *TexA) onChange(e *vecty.Event) {
	//console.Log(e)
	tmpx := e.Value.Get("target").Get("value").String()
	//console.Log(e.Value.Get("target").Get("value"))// .String()
	//console.Log(js.Global.Get("hljs").Call("highlightAuto",tmpx))
	//a.Input = js.Global.Get("hljs").Call("highlightAuto",tmpx).Get("value").String()
  k1, err := loader.ParseYAML([]byte(tmpx))
  if err != nil {
		console.Log("error: %v", err)
  }
  tmpk := k1["services"]
  tmp3, _ := tmpk.(map[string]interface{})
  services, err := converter.LoadServices(tmp3)
  containers, err := converter.TransServicesToContainer(services)
  var f1 = batecttypes.BatectConfig{
		Containers: containers,
  }
  batectyaml, err := yaml.Marshal(&f1)
	a.Input = string(batectyaml[:])
	//js.Global.Get("hljs").Call("highlight")
}
func (a *TexA) onChangeC(e *vecty.Event) {
//	console.Log(e)
	vecty.Rerender(a)
}

func (a *Application) Render() vecty.ComponentOrHTML {
		c := &TexA{
		Input: "",
	}
	return elem.Body(elem.TextArea(vecty.Markup(event.Input(c.onChange))),elem.Button(vecty.Text("button"),vecty.Markup(event.Click(c.onChangeC))),c)
	//return elem.Body(elem.TextArea(vecty.Markup(event.Input(c.onChange))),elem.Button(vecty.Text("button",)),c)
}

func main() {
	vecty.SetTitle("gbatect file online")

	a := &Application{
	}

	vecty.RenderBody(a)
}
