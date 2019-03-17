package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type jsonTest struct {
	Field string `json:field`
	Value int64  `json:value`
}

func TestJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`
		{
			"field": "test",
			"value": 100
		}
		`))
	}))
	defer server.Close()

	data := jsonTest{}
	ht := HTTPClient{}
	err := ht.JSON(server.URL, &data)
	if err != nil || data.Field != "test" || data.Value != 100 {
		t.Error("Wrong response", err, data)
	}
}
