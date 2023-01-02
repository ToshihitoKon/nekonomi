# NEKONOMI

simple KVS library and command line Tool

## CLI

```
go install github.com/ToshihitoKon/nekonomi/cmd/nekonomi@latest
```

### Usage

```
% nekonomi
nekonomi is simple KVS command line tool based SQLite3

Usage:
    nekonomi [command]

Available Commands:
    completion  Generate the autocompletion script for the specified shell
    get         Read value
    help        Help about any command
    list        List stored keys
    set         Store new key and value

Flags:
    -h, --help   help for nekonomi

Use "nekonomi [command] --help" for more information about a command.
```

```bash
# set
$ nekonomi set key1 value1

# get
$ nekonomi get key1
value1

# update
$ nekonomi set key1 value1_updated
$ nekonomi get key1
value1_updated
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
