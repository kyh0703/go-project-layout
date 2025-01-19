// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package model

import (
	"database/sql"
)

type Edge struct {
	ID        interface{}
	SubFlowID int64
	Source    int64
	Target    int64
	Type      string
	Label     sql.NullString
	Hidden    sql.NullInt64
	MarkerEnd sql.NullString
	Points    sql.NullString
}

type Flow struct {
	ID          int64
	Name        string
	Description sql.NullString
}

type Node struct {
	ID          interface{}
	SubFlowID   int64
	Type        string
	Parent      sql.NullString
	Position    sql.NullString
	Styles      sql.NullString
	Width       sql.NullInt64
	Height      sql.NullInt64
	Hidden      sql.NullInt64
	Description sql.NullString
}

type SubFlow struct {
	ID          int64
	Name        string
	Description sql.NullString
}

type User struct {
	ID   int64
	Name string
	Bio  sql.NullString
}
