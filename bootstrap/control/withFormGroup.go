//** This file was code generated by got. ***

package control

import (
	"bytes"
	"context"
	"html"

	grhtml "github.com/spekary/goradd/html"
	"github.com/spekary/goradd/page"
)

func FormGroupTmpl(ctx context.Context, wrapper *FormGroupWrapper, ctrl page.ControlI, h string, buf *bytes.Buffer) {
	var hasInnerDivAttributes bool = wrapper.HasInnerDivAttributes()
	var hasInstructions bool = (ctrl.Instructions() != "")

	if wrapper.useTooltips {
		// bootstrap requires that parent of a tooltipped object has position relative
		ctrl.WrapperAttributes().SetStyle("position", "relative")
	}

	buf.WriteString(`<div id="`)

	buf.WriteString(ctrl.ID())

	buf.WriteString(`_ctl" `)

	buf.WriteString(ctrl.WrapperAttributes().String())

	buf.WriteString(` >
`)
	if ctrl.Label() != "" {
		buf.WriteString(`  <label id="`)

		buf.WriteString(ctrl.ID())

		buf.WriteString(`_lbl" `)
		if ctrl.HasFor() {
			buf.WriteString(` for="`)

			buf.WriteString(ctrl.ID())

			buf.WriteString(`"`)
		}

		buf.WriteString(` `)
		if wrapper.HasLabelAttributes() {
			buf.WriteString(wrapper.LabelAttributes().String())
		}

		buf.WriteString(`>`)

		buf.WriteString(html.EscapeString(ctrl.Label()))

		buf.WriteString(`</label>
`)
	} else {

		buf.WriteString(`    `)
		if ctrl.HasAttribute("placeholder") {
			buf.WriteString(`  <label id="`)

			buf.WriteString(ctrl.ID())

			buf.WriteString(`_lbl" `)
			if ctrl.HasFor() {
				buf.WriteString(` for="`)

				buf.WriteString(ctrl.ID())

				buf.WriteString(`"`)
			}

			buf.WriteString(` class="sr-only">`)

			buf.WriteString(html.EscapeString(ctrl.Attribute("placeholder")))

			buf.WriteString(`</label>
    `)
		}
	}

	buf.WriteString(`
`)
	if hasInnerDivAttributes {
		buf.WriteString(`<div `)

		buf.WriteString(wrapper.InnerDivAttributes().String())

		buf.WriteString(`>`)
	}

	buf.WriteString(grhtml.Indent(h))

	buf.WriteString(`
`)
	if hasInnerDivAttributes {
		buf.WriteString(`</div>`)
	}

	buf.WriteString(`
`)
	switch ctrl.ValidationState() {
	case page.Valid:
		msg := ctrl.ValidationMessage()
		if msg == "" {
			msg = "&nbsp"
		} else {
			msg = html.EscapeString(msg)
		}

		buf.WriteString(`<div id="`)

		buf.WriteString(ctrl.ID())

		buf.WriteString(`_err" class="`)
		if wrapper.useTooltips {
			buf.WriteString(`valid-tooltip`)
		} else {

			buf.WriteString(`valid-feedback`)
		}

		buf.WriteString(`">`)

		buf.WriteString(msg)

		buf.WriteString(`</div>`)

	case page.Invalid:
		msg := ctrl.ValidationMessage()
		if msg == "" {
			msg = "&nbsp"
		} else {
			msg = html.EscapeString(msg)
		}

		buf.WriteString(`<div id="`)

		buf.WriteString(ctrl.ID())

		buf.WriteString(`_err" class="`)
		if wrapper.useTooltips {
			buf.WriteString(`invalid-tooltip`)
		} else {

			buf.WriteString(`invalid-feedback`)
		}

		buf.WriteString(`">`)

		buf.WriteString(msg)

		buf.WriteString(`</div>`)

	default:
		// Either draw instructions, or draw an empty space so that if a validation error is shown, the layout will not shift

		buf.WriteString(`<small id="`)

		buf.WriteString(ctrl.ID())

		buf.WriteString(`_inst" class="form-text text-muted" >`)
		if !hasInstructions && !wrapper.useTooltips {
			buf.WriteString(`&nbsp;`)
		} else {

			buf.WriteString(html.EscapeString(ctrl.Instructions()))
		}

		buf.WriteString(`</small>`)

	}

	buf.WriteString(`</div>
`)

	return

}
