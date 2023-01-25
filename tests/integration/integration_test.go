package integration

import (
	"testing"
	"gorm.io/gorm"
  _ "github.com/lib/pq"
)

func testGorm(t *testing.T) {
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  
}