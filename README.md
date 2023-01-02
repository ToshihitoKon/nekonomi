# NEKONOMI

simple KVS library and command line Tool

## CLI

```
go install github.com/ToshihitoKon/nekonomi/cmd/nekonomi@latest
```

## Library

```go
import (
    "github.com/ToshihitoKon/nekonomi"
)

func main() {
    dbid := "db_identifire"
    opts := *nekonomi.Option[]{
        nekonomi.OptionSQLiteFilePath("dbdata/"),
        // nekonomi.OptionSchema("anothe_schema"),
        // nekonomi.OptionReadOnly(),
    }
    nekodb, err := nekonomi.New(dbid, opts)
    _ = err // error handling

    value, err := nekodb.Read("key")
    value, err := nekodb.Write("key", "value")

    value, err := nekodb.Update("key", "value")
    value, err := nekodb.Delete("key")
    // if key is not exist, return nekonomi.ErrKeyNotFound

    schemaList, err := nekodb.SchemaList()
    err := nekodb.SchemaSet("anothe_schema")
}
```
