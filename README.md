# JSON Appender

Golang util package for appending json objects to exinsting arrays in json files, **without bloating your machine's memory**.

It writes data into array json, without buffering the whole file. 

https://pkg.go.dev/github.com/ferkze/jsonappender

### Disclaimer

Note that it still only works when appending data into a root array type. 
Appending your objects to any other root data type would result in something unexpected or maybe an error.

### Download

```
go get -u github.com/ferkze/jsonappender
```

## Examples

### Usage

```
package main

import (
  "encoding/json"
  "log"

  jap "github.com/ferkze/jsonappender"
)

func main() {
  f := "foo.json"
  a, err := jap.JSONAppender(f)
  if err != nil {
    panic(err.Error())
  }

  bar := []map[string]interface{}{
    // ...
  }

  b, err := json.Marshal(&bar)
  if err != nil {
    panic(err.Error())
  }

  if _, err = a.Write(b); err != nil {
    panic(err.Error())
  }

  if err = a.Close(); err != nil {
    panic(err.Error())
  }

  // ...
}
```

### Mistakes

Don't make a mistake by passing the jsonAppender interface to a json encoder, since it inserts extra trailing lines in the file for each iteration of writes.
This may result in 16 trailing new lines in your file and an error identifying the ']' token at the end of the file.

##### Do this

```
  a, err := jap.JSONAppender(f)
  if err != nil {
    panic(err.Error())
  }

  bar := []map[string]interface{}{
    // data
  }

  b, err := json.Marshal(&bar)
  if err != nil {
    panic(err.Error())
  }

  if _, err = a.Write(b); err != nil {
    panic(err.Error())
  }

  if err = a.Close(); err != nil {
    panic(err.Error())
  }
```

##### Not this

```
  a, err := jap.JSONAppender(f)
  if err != nil {
    // There will be an error if used +16 times with json.Encoder
    panic(err.Error())
  }

  bar := []map[string]interface{}{
    // data
  }

  if err = json.NewEncoder(a).Encode(&bar); err != nil {
    panic(err.Error())
  }

  if err = a.Close(); err != nil {
    panic(err.Error())
  }
```
