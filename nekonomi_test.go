package nekonomi

import "testing"

func TestOptions(t *testing.T) {
	_ = OptionSQLiteFilePath("/tmp/nekonomi_test/")
	_ = OptionSchema("anothe_schema")
	_ = OptionReadOnly()
}
func TestClient(t *testing.T) {
	dbid := "db_identifire"
	opts := []Option{
		OptionSQLiteFilePath("/tmp/nekonomi_test/"),
	}

	client, err := New(dbid, opts)
	if err != nil {
		t.Error(err)
	}

	_, err = client.Read("key")
	if err != nil {
		t.Error(err)
	}

	_, err = client.Write("key", "value")
	if err != nil {
		t.Error(err)
	}

	_, err = client.Update("key", "value")
	if err != nil {
		t.Error(err)
	}

	_, err = client.Delete("key")
	if err != nil {
		t.Error(err)
	}

	_, err = client.SchemaList()
	if err != nil {
		t.Error(err)
	}

	err = client.SchemaSet("anothe_schema")
	if err != nil {
		t.Error(err)
	}
}
