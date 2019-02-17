package control

import (
	"github.com/goradd/goradd/pkg/html"
	"github.com/goradd/goradd/pkg/page"
)

type RadioButtonI interface {
	CheckboxI
}

type RadioButton struct {
	checkboxBase
	group string

}

func NewRadioButton(parent page.ControlI, id string) *RadioButton {
	c := &RadioButton{}
	c.Init(c, parent, id)
	return c
}

func (c *RadioButton) this() RadioButtonI {
	return c.Self.(RadioButtonI)
}

func (c *RadioButton) ΩDrawingAttributes() *html.Attributes {
	a := c.checkboxBase.ΩDrawingAttributes()
	a.SetDataAttribute("grctl", "bs-radio")
	a.Set("type", "radio")
	if c.group == "" {
		a.Set("name", c.ID()) // treat it like a checkbox if no group is specified
	} else {
		a.Set("name", c.group)
		a.Set("value", c.ID())
	}
	return a
}

// ΩUpdateFormValues is an internal call that lets us reflect the value of the checkbox on the web override
func (c *RadioButton) ΩUpdateFormValues(ctx *page.Context) {
	id := c.ID()

	if v, ok := ctx.CheckableValue(id); ok {
		c.SetCheckedNoRefresh(v)
	}
}

func (c *RadioButton) SetGroup(g string) RadioButtonI {
	c.group = g
	c.Refresh()
	return c.this()
}

func (c *RadioButton) Group() string {
	return c.group
}

// TODO: Serialize