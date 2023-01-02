package nekonomi

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrKeyNotFound = fmt.Errorf("Error: Key not found")
)

type Client struct {
	DbId   string
	Schema string
}

func New(dbId string, opts []Option) (*Client, error) {
	var err error
	client := &Client{
		DbId:   dbId,
		Schema: "default",
	}

	for _, opt := range opts {
		client, err = opt.apply(client)
		if err != nil {
			return nil, errors.Wrap(err, "Option apply failed")
		}
	}

	return client, nil
}

func (c *Client) Read(key string) (string, error) {
	return "", nil
}
func (c *Client) Write(key, value string) (string, error) {
	return "", nil
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

type Option interface {
	apply(*Client) (*Client, error)
}

type optSQLiteFilePath struct {
	path string
}

func OptionSQLiteFilePath(path string) *optSQLiteFilePath {
	return &optSQLiteFilePath{
		path: path,
	}
}
func (o *optSQLiteFilePath) apply(c *Client) (*Client, error) {
	return c, nil
}

type optSchema struct {
	schema string
}

func OptionSchema(schema string) *optSchema {
	return &optSchema{
		schema: schema,
	}
}
func (o *optSchema) apply(c *Client) (*Client, error) {
	return c, nil
}

type optReadOnly struct{}

func OptionReadOnly() *optReadOnly {
	return &optReadOnly{}
}
func (o *optReadOnly) apply(c *Client) (*Client, error) {
	return c, nil
}
