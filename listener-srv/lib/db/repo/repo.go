package repo

import (
	"strings"

	"github.com/gocraft/dbr/v2"
)

type Repository interface {
	Session
	DB() *dbr.Session
}

type Session interface {
	Select() *dbr.SelectBuilder
	SelectBySql(string, ...interface{}) *dbr.SelectBuilder
	InsertBuilder() *dbr.InsertBuilder
	DeleteBuilder() *dbr.DeleteBuilder
	UpdateBuilder() *dbr.UpdateBuilder
}

type Repo struct {
	*dbr.Session
	*dbr.SelectBuilder
	fields []string
	table  string
}

func New(session *dbr.Session, table string, fields []string) *Repo {
	return &Repo{Session: session, table: table, fields: fields}
}

func (r *Repo) Select() *dbr.SelectBuilder {
	if r.SelectBuilder != nil {
		return r.SelectBuilder
	}

	r.SelectBuilder = r.Session.Select(strings.Join(r.fields, ", ")).From(r.table)

	return r.SelectBuilder
}

func (r *Repo) InsertBuilder() *dbr.InsertBuilder {
	return r.InsertInto(r.table)
}

func (r *Repo) DeleteBuilder() *dbr.DeleteBuilder {
	return r.DeleteFrom(r.table)
}

func (r *Repo) UpdateBuilder() *dbr.UpdateBuilder {
	return r.Update(r.table)
}

func (r *Repo) DB() *dbr.Session {
	return r.Session
}
