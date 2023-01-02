package nekonomi

import (
	"database/sql"
	"fmt"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var (
	ErrKeyNotFound = fmt.Errorf("Error: Key not found")
)

type Client struct {
	db       *sql.DB
	dbDir    string
	dbId     string
	schema   string
	readOnly bool

	resetDatabase bool
}

func New(dbDir, dbId string, opts []Option) (*Client, error) {
	var err error
	client := &Client{
		dbDir:         dbDir,
		dbId:          dbId,
		schema:        "default",
		readOnly:      false,
		resetDatabase: false,
	}

	for _, opt := range opts {
		client, err = opt.apply(client)
		if err != nil {
			return nil, errors.Wrap(err, "Option apply failed")
		}
	}

	db, err := sql.Open("sqlite3", path.Join(dbDir, dbId+"_nekonomi.db"))
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open failed")
	}
	client.db = db

	if client.resetDatabase {
		dropStmt := fmt.Sprintf(`
		DROP TABLE IF EXISTS "%s";`, client.schema)

		if _, err := db.Exec(dropStmt); err != nil {
			return nil, errors.Wrapf(err, "error: db.Exec")
		}
	}

	createTableStmt := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS "%s" (
		key TEXT PRIMARY KEY,
		value TEXT,
		created_at_utc TEXT,
		updated_at_utc TEXT
	);`, client.schema)

	if _, err := db.Exec(createTableStmt); err != nil {
		return nil, errors.Wrapf(err, "error: db.Exec")
	}

	return client, nil
}
func (c *Client) Close() {
	c.db.Close()
}

func (c *Client) Read(key string) (string, error) {
	stmt := fmt.Sprintf(`
		SELECT key, value FROM "%s"
		WHERE key = ?
	;`, c.schema)

	res := struct {
		key   string
		value string
	}{}

	row := c.db.QueryRow(stmt, key)
	if err := row.Scan(&res.key, &res.value); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrKeyNotFound
		}
		return "", errors.Wrapf(err, "error: row.Scan() failed")
	}
	return res.value, nil
}

func (c *Client) Write(key, value string) (string, error) {
	stmt := fmt.Sprintf(`
	INSERT INTO "%s" (
		key,
		value,
		created_at_utc,
		updated_at_utc
	) VALUES (
		"%s",
		"%s",
		datetime('now'),
		datetime('now')
	);`, c.schema, key, value)
	if _, err := c.db.Exec(stmt); err != nil {
		return "", errors.Wrapf(err, "error: db.Exec")
	}

	return value, nil
}
func (c *Client) Update(key, value string) (string, error) {
	return "", nil
}

func (c *Client) Delete(key string) (string, error) {
	return "", nil
}

func (c *Client) SchemaList() ([]string, error) {
	return []string{}, nil
}

func (c *Client) SchemaSet(schema string) error {
	return nil
}
