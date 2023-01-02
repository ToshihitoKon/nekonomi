package nekonomi

import "testing"

func TestClient(t *testing.T) {
	dbdir := "/tmp/"
	dbid := "db_identifire"
	opts := []Option{}

	client, err := New(dbdir, dbid, opts)
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
