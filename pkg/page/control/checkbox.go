package control

import (
	"github.com/goradd/goradd/pkg/html"
	"github.com/goradd/goradd/pkg/page"
)

// Checkbox is a basic html checkbox input form control.
type Checkbox struct {
	CheckboxBase
}

// NewCheckbox creates a new checkbox control.
func NewCheckbox(parent page.ControlI, id string) *Checkbox {
	c := &Checkbox{}
	c.Init(c, parent, id)
	return c
}

// ΩDrawingAttributes is called by the framework to set the temporary attributes that the control
// needs. Checkboxes set the grctl, name, type and value attributes automatically.
// You do not normally need to call this function.
func (c *Checkbox) ΩDrawingAttributes() *html.Attributes {
	a := c.CheckboxBase.ΩDrawingAttributes()
	a.SetDataAttribute("grctl", "checkbox")
	a.Set("name", c.ID()) // needed for posts
	a.Set("type", "checkbox")
	a.Set("value", "1") // required for html validity
	return a
}

// ΩUpdateFormValues is an internal call that lets us reflect the value of the checkbox on the form.
// You do not normally need to call this function.
func (c *Checkbox) ΩUpdateFormValues(ctx *page.Context) {
	id := c.ID()

	if v, ok := ctx.FormValue(id); ok {
		c.SetCheckedNoRefresh(v)
	} else if ctx.RequestMode() == page.Server && c.IsOnPage() {
		// We will not get a value if an item is not checked. But since this is a POST, all values on page
		// should send something if its checked, therefore we know its not checked.
		c.SetCheckedNoRefresh(false)
	}
}
