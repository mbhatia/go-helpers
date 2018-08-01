package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"

	gorp "gopkg.in/gorp.v2"
)

// Extension to the gorp.PostgresDialect to handle JSONB datatype
// JSONB is the alias so we can properly handle jsonb datatype with Postgres DB
type JSONB map[string]interface{}

func (p JSONB) Value() (driver.Value, error) {
	if p != nil {
		j, err := json.Marshal(p)
		return j, err
	}

	return nil, nil
}

func (p *JSONB) Scan(src interface{}) error {
	if src != nil {
		source, ok := src.([]byte)
		if !ok {
			return errors.New("Type assertion .([]byte) failed.")
		}

		var i interface{}
		err := json.Unmarshal(source, &i)
		if err != nil {
			return err
		}

		*p, ok = i.(map[string]interface{})
		if !ok {
			return errors.New("Type assertion .(map[string]interface{}) failed.")
		}
	}

	return nil
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
