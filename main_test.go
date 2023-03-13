package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Table("test").Scopes(
		func(d *gorm.DB) *gorm.DB {
			return d.Where("a = 1")
		},
		func(d *gorm.DB) *gorm.DB {
			return d.Where(d.Or("b = 2").Or("c = 3"))
		},
	).Rows()
}
