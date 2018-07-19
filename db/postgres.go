package db

import (
	"reflect"

	gorp "gopkg.in/gorp.v2"
)

// Extension to the gorp.PostgresDialect to handle JSONB datatype
// JSONB is the alias so we can properly handle jsonb datatype with Postgres DB
type JSONB struct {
	string
}

//
// HACK: Get the correct type into the database table for JSON-B objects
//
type PostgresDialect struct {
	gorp.PostgresDialect
}

// ToSqlType returns the SQL column type for the given Go type
func (d PostgresDialect) ToSqlType(val reflect.Type, maxsize int, isAutoIncr bool) string {
	// Force JSONText to use sqlType: JSONB
	if val == reflect.TypeOf(JSONB{}) {
		return "JSONB"
	}
	return d.PostgresDialect.ToSqlType(val, maxsize, isAutoIncr)
}
