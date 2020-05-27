package jsonappender

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJSONAppender(t *testing.T) {
  f := "test.json"
  a, err := JSONAppender(f)
  if err != nil {
    t.Fatalf(err.Error())
  }

  data := []map[string]interface{}{
    {
      "Asset": "ALPA4",
      "AssetType": "Shares",
      "Quantity": 1000,
      "OrderType": "Long",
      "Price": 26.42,
      "Datetime": "2020-05-26T11:33:42Z",
    },
		{
      "Asset": "ALPA4",
      "AssetType": "Shares",
      "Quantity": 1000,
      "OrderType": "Short",
      "Price": 27.42,
      "Datetime": "2020-05-26T11:42:11Z",
    },
  }

  b, err := json.Marshal(&data)
  if err != nil {
    log.Fatalf(err.Error())
  }
  
  if _, err = a.Write(b); err != nil {
    log.Fatalf(err.Error())
  }

  if err = a.Close(); err != nil {
    log.Fatalf(err.Error())
  }
}