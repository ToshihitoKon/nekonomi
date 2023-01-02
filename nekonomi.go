package nekonomi

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mattn/go-sqlite3"
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

	// Home Directory to Real path
	filepath := path.Join(dbDir, dbId+"_nekonomi.db")
	if strings.HasPrefix(filepath, "~/") {
		dirname, _ := os.UserHomeDir()
		filepath = path.Join(dirname, filepath[2:])
	}

	db, err := sql.Open("sqlite3", filepath)
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
		value TEXT NOT NULL,
		created_at_utc TEXT NOT NULL DEFAULT (DATETIME('now')),
		updated_at_utc TEXT NOT NULL DEFAULT (DATETIME('now'))
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

func (c *Client) ForceWrite(key, value string) (string, error) {
	result, err := c.Write(key, value)
	if err != nil {
		var sqliteErr sqlite3.Error
		if !(errors.As(err, &sqliteErr) &&
			errors.Is(sqliteErr.Code, sqlite3.ErrConstraint)) {
			return "", errors.Wrapf(err, "error: db.Write")
		}
		result, err = c.Update(key, value)
		if err != nil {
			return "", errors.Wrapf(err, "error: db.Update")
		}
	}

	return result, err
}

func (c *Client) Write(key, value string) (string, error) {
	stmt := fmt.Sprintf(`
	INSERT INTO "%s" (
		key,
		value
	) VALUES (
		"%s",
		"%s"
	);`, c.schema, key, value)
	if _, err := c.db.Exec(stmt); err != nil {
		return "", errors.Wrapf(err, "error: db.Exec")
	}

	return value, nil
}

func (c *Client) ListKeys() ([]string, error) {
	stmt := fmt.Sprintf(`SELECT key FROM "%s"`, c.schema)
	rows, err := c.db.Query(stmt)
	if err != nil {
		return []string{}, errors.Wrapf(err, "error: db.Query")
	}

	var buffer string
	var keys []string

	for rows.Next() {
		if err := rows.Scan(&buffer); err != nil {
			return []string{}, errors.Wrapf(err, "error: rows.Scan")
		}
		keys = append(keys, buffer)
	}
	if err := rows.Err(); err != nil {
		return []string{}, errors.Wrapf(err, "error: rows.Err")
	}
	return keys, nil
}

func (c *Client) Update(key, value string) (string, error) {
	stmt := fmt.Sprintf(`
		UPDATE "%s" SET value = "%s", updated_at_utc = DATETIME('now') WHERE key = "%s"
	;`, c.schema, value, key)
	if _, err := c.db.Exec(stmt); err != nil {
		return "", errors.Wrapf(err, "error: db.Exec")
	}

	return value, nil
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
