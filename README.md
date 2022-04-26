# JSON Strip Comments

This is a very simple command that will remove comments from JSON files.

_Strictly speaking this could be used on any text file not just JSON, however
the project's goal and logic is centered around JSON files._

```bash
Removes c style comments from a json file
Usage: json-strip-comments [--output OUTPUT] [--empty] [FILENAME]

Positional arguments:
  FILENAME

Options:
  --output OUTPUT, -o OUTPUT
                         Output to given filename instead of stdout
  --empty, -e            Remove empty lines after comment removals
  --help, -h             display this help and exit
```

## Caveats

Due to the simplicity of this project, only comments with whitespace preceeding them
will be removed. This prevents removing false positive comments from the inside of a
json string for example.

```jsonc
// this comment will be removed
{
  // and this one too
  "id": 1,
  /* and so will this */
  "name": "Bobby Tables",
  /*
  Multi line comments are removed as well \o/
  */
  "email": "user@example.com" // however, this comment will NOT be removed
}
```

## Usage Examples

#### Piping stdin

```bash
cat somefile.jsonc | json-strip-comments
```

#### Removing empty lines

```bash
cat somefile.jsonc | json-strip-comments -e
```

#### Specify an input file

```bash
json-strip-comments -e something.jsonc
```

#### Specify an output file

```bash
cat input.jsonc | json-strip-comments -e -o "output.json"
# or
json-strip-comments -e -o "output.json" "input.jsonc"
```

#### Using wildcards in output filename

This only workls when using the incoming filename

```bash
json-strip-comments -e -o "[folder][file].json" "/path/to/input.jsonc"
# writes to: /path/to/input.json
```

#### Perform on all files in current folder

This involves a bit of bash magic, it looks for all `.jsonc` files
and writes `.json` files with the same name in the same place.

```bash
find ./ -type f -name "*.jsonc" -exec json-strip-comments -e -o "[folder][file].json" "{}" \;
```


## Install

#### Centos

RPM hosted on [yum.jc21.com](https://yum.jc21.com)

#### Go Install

```bash
go install github.com/jc21/json-strip-comments@latest
```


#### Building

```bash
git clone https://github.com/jc21/json-strip-comments && cd json-strip-comments
go build -o bin/json-strip-comments cmd/json-strip-comments/main.go
./bin/json-strip-comments -h
```
