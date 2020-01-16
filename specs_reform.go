// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package main

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type specViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *specViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("specs").
func (v *specViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *specViewType) Columns() []string {
	return []string{"status", "filename", "line_number", "commit_id"}
}

// NewStruct makes a new struct for that view or table.
func (v *specViewType) NewStruct() reform.Struct {
	return new(Spec)
}

// SpecView represents specs view or table in SQL database.
var SpecView = &specViewType{
	s: parse.StructInfo{Type: "Spec", SQLSchema: "", SQLName: "specs", Fields: []parse.FieldInfo{{Name: "Status", Type: "int64", Column: "status"}, {Name: "Filename", Type: "string", Column: "filename"}, {Name: "LineNumber", Type: "int64", Column: "line_number"}, {Name: "CommitID", Type: "int64", Column: "commit_id"}}, PKFieldIndex: -1},
	z: new(Spec).Values(),
}

// String returns a string representation of this struct or record.
func (s Spec) String() string {
	res := make([]string, 4)
	res[0] = "Status: " + reform.Inspect(s.Status, true)
	res[1] = "Filename: " + reform.Inspect(s.Filename, true)
	res[2] = "LineNumber: " + reform.Inspect(s.LineNumber, true)
	res[3] = "CommitID: " + reform.Inspect(s.CommitID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Spec) Values() []interface{} {
	return []interface{}{
		s.Status,
		s.Filename,
		s.LineNumber,
		s.CommitID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Spec) Pointers() []interface{} {
	return []interface{}{
		&s.Status,
		&s.Filename,
		&s.LineNumber,
		&s.CommitID,
	}
}

// View returns View object for that struct.
func (s *Spec) View() reform.View {
	return SpecView
}

// check interfaces
var (
	_ reform.View   = SpecView
	_ reform.Struct = (*Spec)(nil)
	_ fmt.Stringer  = (*Spec)(nil)
)

func init() {
	parse.AssertUpToDate(&SpecView.s, new(Spec))
}