package gorm_test

import (
	"go_tutorial/go_frame/orm/gorm"
	// "log/slog"
	// "os"
	"testing"
)

// var (
// 	db = gorm.CreateConnection("localhost", "test", "tester", "123456", 3306)
// )

func TestGormQuickStart(t *testing.T) {
	gorm.GormQuickStart()
}

// go test -v ./orm/gorm -run=^TestGormQuickStart$ -count=1
