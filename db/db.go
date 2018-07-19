package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	gorp "gopkg.in/gorp.v2"
)

type DB struct {
	*gorp.DbMap
}

type DBCreator interface {
	AddTable(db *DB)
}

type DBInsertable interface {
	Insert(db *DB) (interface{}, error)
}

type DBSelectable interface {
	Load(db *DB, key interface{}) error
	Select(db *DB, sql string, args ...interface{}) error
}

func NewMSSQL(dsn string) (*DB, error) {
	// Connect to the DB and initialize the DB schema
	conn, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to the DB using dsn: ", dsn, err)
		return nil, err
	}

	// construct a gorp DbMap
	db := &DB{&gorp.DbMap{Db: conn, Dialect: gorp.SqlServerDialect{}}}

	return db, nil
}

func NewPostgres(dsn string) (*DB, error) {
	// Connect to the DB and initialize the DB schema
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Failed to connect to the DB using dsn: ", dsn, err)
		return nil, err
	}

	// construct a gorp DbMap
	db := &DB{&gorp.DbMap{Db: conn, Dialect: PostgresDialect{}}}

	return db, nil
}

func (db *DB) AddObject(obj DBCreator) error {
	obj.AddTable(db)
	return nil
}

func (db *DB) CreateSchemaIfNotExists() error {
	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	if err := db.CreateTablesIfNotExists(); err != nil {
		log.Println("Failed to create tables")
		return err
	}

	if err := db.CreateIndex(); err != nil {
		// TODO: Figure out if there is a true error or just the indices already exist.
		log.Println("Failed to create indices: ", err)
	}

	return nil
}
