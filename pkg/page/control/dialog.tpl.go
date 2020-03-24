//** This file was code generated by got. DO NOT EDIT. ***

package control

import (
	"bytes"
	"context"
	"html"
)

func (d *Dialog) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {
	d.TitleBar().Draw(ctx, buf)

	l := len(d.Children())
	if l > 2 {
		for _, child := range d.Children() {
			child.Draw(ctx, buf)
		}
	} else {

		buf.WriteString(html.EscapeString(d.Text()))

	}
	d.ButtonBar().Draw(ctx, buf)
	return
}
