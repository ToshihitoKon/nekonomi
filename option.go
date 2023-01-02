package nekonomi

type Option interface {
	apply(*Client) (*Client, error)
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
	c.schema = o.schema
	return c, nil
}

type optReadOnly struct{}

func OptionReadOnly() *optReadOnly {
	return &optReadOnly{}
}
func (o *optReadOnly) apply(c *Client) (*Client, error) {
	c.readOnly = true
	return c, nil
}
