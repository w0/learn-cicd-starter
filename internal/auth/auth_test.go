package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	secret := "MyCoolPassword"
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey "+secret)

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatal(err)
	}

	if key != secret {
		t.Fatalf("returned key doesn't match: %s.. Got %s\n", strings.Split(secret, "Api"), key)
	}
}
