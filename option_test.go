package nekonomi

import "testing"

func TestOption(t *testing.T) {
	_ = OptionSchema("anothe_schema")
	_ = OptionReadOnly()
}
