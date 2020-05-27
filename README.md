# JSON Appender

## Description

Golang util package for appending json data to exinsting json files, **without bloating your machine's memory**.

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
    log.Fatalf("err")
  }

  bar := []map[string]interface{}{
    // ...
  }

  b, err := json.Marshal(&bar)
  if err != nil {
    log.Fatalf(err.Error())
  }

  if _, err = a.Write(b); err != nil {
    log.Fatalf(err.Error())
  }

  if err = a.Close(); err != nil {
    log.Fatalf("err")
  }

  // ...
}
```
