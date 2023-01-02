package nekonomi

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrKeyNotFound = fmt.Errorf("Error: Key not found")
)

type Client struct {
	dbdir    string
	dbId     string
	schema   string
	readOnly bool
}

func New(Dbdir, dbId string, opts []Option) (*Client, error) {
	var err error
	client := &Client{
		dbdir: Dbdir,
		dbId:  dbId,
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
