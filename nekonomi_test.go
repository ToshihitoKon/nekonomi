package nekonomi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient(t *testing.T) {
	dbdir := "/tmp"
	dbid := "nekonomi_test"
	opts := []Option{
		OptionResetDatabase(),
	}

	client, err := New(dbdir, dbid, opts)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	writtenValue, err := client.Write("key1", "value1")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(writtenValue, "value1"); diff != "" {
		t.Errorf("client.Write result mismatch (-want +got):\n%s", diff)
	}

	readValue, err := client.Read("key1")
	if err != nil {
		t.Error(err)
	}
	if diff := cmp.Diff(readValue, "value1"); diff != "" {
		t.Errorf("client.Read result mismatch (-want +got):\n%s", diff)
	}

	listedKeys, err := client.ListKeys()
	if err != nil {
		t.Error(err)
	}
	if diff := cmp.Diff(listedKeys, []string{"key1"}); diff != "" {
		t.Errorf("client.ListKeys result mismatch (-want +got):\n%s", diff)
	}

	_, err = client.Update("key1", "value1_updated")
	if err != nil {
		t.Error(err)
	}

	_, err = client.ForceWrite("key1", "value1_updated2")
	if err != nil {
		t.Error(err)
	}
	_, err = client.ForceWrite("key2", "value2")
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
