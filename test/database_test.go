package test

import (
	"fmt"
	"journey-user/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDb(t *testing.T) {
	dbSql := app.NewDb()
	fmt.Println(dbSql)
	assert.NotNil(t, dbSql)
}
