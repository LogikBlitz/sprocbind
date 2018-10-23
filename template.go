package main

const tFile = `
// Code generated by sprocbind, DO NOT EDIT.

package {{.packageName}}

{{if .procedures}}
import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Procedures

type Queryer interface {
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
}

{{range $procedure := .procedures}}

	// {{GoName $procedure.Name}} calls the stored procedure {{Clean $procedure.Name}}
	func {{GoName $procedure.Name}}(ctx context.Context, db Queryer{{range $param := $procedure.Params}}, {{GoNameLower $param.Name}} {{GoDataType $param.DataType}}{{end}}) (*sqlx.Rows, error) {
		// TODO: Add tracing/timing here.
		return db.QueryxContext(ctx,
			"exec {{Clean $procedure.Name}} {{range $i, $param := $procedure.Params}}{{if gt $i 0}}, {{end}}@{{Clean $param.Name}}{{end}}"{{range $i, $param := $procedure.Params}}, 
			sql.Named("{{Clean $param.Name}}", {{GoNameLower $param.Name}}){{end}})
	}
{{end}}

{{end}}


{{if .normalTables}}

// Tables

{{range $table := .normalTables}}
type {{GoName $table.Name}} struct {
	{{range $col := $table.Columns}}
		{{GoName $col.Name}} *{{GoDataType $col.DataType}} ` + "`" + `db:"{{Clean $col.Name}}"` + "`" + `
	{{end}}
}
{{end}}

{{end}}


{{if .tables}}
// Table Types

{{range $table := .tables}}
	type {{GoName $table.Name}} []{{GoName $table.Name}}Row

	func (t *{{GoName $table.Name}}) TVP() (typeName string, exampleRow []interface{}, rows [][]interface{}) {
		typeName = "{{Clean $table.Name}}"
		//columnNames = []string{ {{range $i, $col := $table.Columns}}{{if gt $i 0}}, {{end}}"{{Clean $col.Name}}"{{end}} }
		var v []{{GoName $table.Name}}Row
		if t != nil {
			v = *t
		}
		for _, r := range append(v, {{GoName $table.Name}}Row{}) {
			rows = append(rows, []interface{}{ {{range $col := $table.Columns}}
				r.{{GoName $col.Name}},{{end}}
			})
		}
		exampleRow = rows[len(rows)-1]
		rows = rows[:len(rows)-1]
		
		return
	}

	type {{GoName $table.Name}}Row struct {
		{{range $col := $table.Columns}}	{{GoName $col.Name}} {{GoDataType $col.DataType}} ` + "`" + `db:"{{Clean $col.Name}}"` + "`" + `
	{{end}}}
{{end}}

{{end}}`