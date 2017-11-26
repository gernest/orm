package {{.Package}}

import (
    "{{.Type.ImportPath}}"
)

// API is the interface of the ORM object
type API interface {
    Close() error
    Create() *CreateBuilder
    Select() *SelectBuilder
    Insert() *InsertBuilder
    Update() *UpdateBuilder
    Delete() *DeleteBuilder
    Insert{{.Type.Name}}(*{{.Type.FullName}}) *InsertBuilder
    Update{{.Type.Name}}(*{{.Type.FullName}}) *UpdateBuilder

    Logger(Logger)
}

// Querier is the interface for a SELECT SQL statement
type Querier interface {
    Query() ([]{{.Type.FullName}}, error)
}

// Counter is the interface for a SELECT SQL statement for counting purposes
type Counter interface {
    Count() ([]{{.Type.Name}}Count, error)
}
