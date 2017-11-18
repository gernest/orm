package {{.PackageName}}

import "{{.Type.ImportPath}}"

{{ range $_, $f := .Type.Fields -}}
func (s Select) {{$f.Name}}() Select {
    return append(s, "{{$f.ColumnName}}")
}
{{ end -}}

// ScanArgs are list of fields to be given to the sql Scan command
func (s Select) scanArgs(p *{{.Type.FullName}}) []interface{} {
	if len(s) == 0 {
        // add to args all the fields of p
        return []interface{}{
            {{ range $_, $f := .Type.Fields -}}
            &p.{{$f.Name}},
            {{ end }}
        }
	}

    // select was given, choose only some fields
	m := make(map[string]int, len(s))
    for i, col := range s {
        m[col] = i + 1
    }
	args := make([]interface{}, len(s))
	{{ range $_, $f := .Type.Fields -}}
	if i := m["{{$f.ColumnName}}"]; i != 0 {
		args[i-1] = &p.{{$f.Name}}
	}
	{{ end -}}
	return args
}
