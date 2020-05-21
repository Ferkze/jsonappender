# JSON Appender

## Description

Golang util package for appending json data to exinsting json files, **without bloating your machine's memory**.

### Usage

```
func main() {
  f := "foo.json"
  a, err := JSONArrayAppender(f)
  if err != nil {
    log.Fatalf("err")
  }

  data := []map[string]interface{}{
    // ...
  }

  if err = json.NewEncoder(a).Encode(&data); err != nil {
    log.Fatalf("err")
  }

  if err = a.Close(); err != nil {
    log.Fatalf("err")
  }

  // ...
}
```
