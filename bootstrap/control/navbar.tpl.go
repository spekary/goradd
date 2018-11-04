//** This file was code generated by got. ***

package control

import (
	"bytes"
	"context"
)

func (b *Navbar) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {
	b.drawToggleAndBrand(ctx, buf)

	buf.WriteString(`	<div class="collapse navbar-collapse" id="`)

	buf.WriteString(b.ID())

	buf.WriteString(`_collapse">
	    `)
	if b.brandLocation == NavbarCollapsedBrandHidden {
		b.drawBrand(ctx, buf)
	}

	buf.WriteString(`		`)

	{
		err := b.DrawChildren(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`	</div>

`)

	return
}

func (b *Navbar) drawToggleButton(ctx context.Context, buf *bytes.Buffer) {

	buf.WriteString(`  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#`)

	buf.WriteString(b.ID())

	buf.WriteString(`_collapse" aria-controls="`)

	buf.WriteString(b.ID())

	buf.WriteString(`_collapse" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
`)

}

func (b *Navbar) drawBrand(ctx context.Context, buf *bytes.Buffer) {
	if b.Text() != "" {

		buf.WriteString(`		<a class="navbar-brand" href="`)
		if b.headerAnchor == "" {
			buf.WriteString(`#`)
		} else {

			buf.WriteString(b.headerAnchor)
		}

		buf.WriteString(`">`)
		b.DrawText(ctx, buf)
		buf.WriteString(`</a>
   `)

	} else { // draw a blank brand so toggler placement still works

		buf.WriteString(`
 		<a class="navbar-brand" href="#"> </a>
    `)

	}
}

func (b *Navbar) drawToggleAndBrand(ctx context.Context, buf *bytes.Buffer) {
	switch b.brandLocation {
	case NavbarCollapsedBrandLeft:
		b.drawBrand(ctx, buf)
		b.drawToggleButton(ctx, buf)
	case NavbarCollapsedBrandRight:
		b.drawToggleButton(ctx, buf)
		b.drawBrand(ctx, buf)
	case NavbarCollapsedBrandHidden:
		b.drawToggleButton(ctx, buf)
	}
}
