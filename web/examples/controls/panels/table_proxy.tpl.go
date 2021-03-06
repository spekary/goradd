//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/pkg/log"
	"github.com/goradd/goradd/pkg/orm/db"
)

func (ctrl *TableProxyPanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
`)

	buf.WriteString(`
<style>
table {
  font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

table td, table th {
  border: 1px solid #ddd;
  padding: 8px;
}

table tr:nth-child(even){background-color: #f2f2f2;}

table tr:hover {background-color: #ddd;}

table th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #4CAF50;
  color: white;
}
</style>
<h1>Tables - Proxy Column</h1>
<p>
The table below demonstrates how to combine a Proxy, a CustomColumn, and a Panel to display a list of records
that allow you to click on a record to see detail of the record.
</p>
`)
	if db.GetDatabase("goradd") == nil {
		buf.WriteString(`<h2 style="color:red">Error</h2>
<p>You have not installed the goradd example database. See the examples/readme.txt file for instructions.<p>
`)
		log.Error("goradd database not installed.")
	} else {

		buf.WriteString(`
`)
		if `` == "" {
			buf.WriteString(`    `)

			{
				err := ctrl.Page().GetControl("table1").Draw(ctx, buf)
				if err != nil {
					return err
				}
			}
		} else {

			buf.WriteString(`    `)

			{
				err := ctrl.Page().GetControl("table1").ProcessAttributeString(``).Draw(ctx, buf)
				if err != nil {
					return err
				}
			}
		}

		buf.WriteString(`
`)

		buf.WriteString(`
`)
		if `` == "" {
			buf.WriteString(`    `)

			{
				err := ctrl.Page().GetControl("personPanel").Draw(ctx, buf)
				if err != nil {
					return err
				}
			}
		} else {

			buf.WriteString(`    `)

			{
				err := ctrl.Page().GetControl("personPanel").ProcessAttributeString(``).Draw(ctx, buf)
				if err != nil {
					return err
				}
			}
		}

		buf.WriteString(`

`)
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	return
}
