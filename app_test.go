package main

import (
	"log"
	"os"
	"testing"
)

// TestMain initializes mock data before running tests.

func TestMain(m *testing.M) {
	if err := loadMockData(); err != nil {
		log.Fatal("Failed to load mock data: ", err)
	}
	os.Exit(m.Run())
}
