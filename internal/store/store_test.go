package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "aktan:aktan@localhost:5432/restapi_test?sslmode=disable"
	}
	os.Exit(m.Run())
}