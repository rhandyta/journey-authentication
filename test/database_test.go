package test

import (
	"journey-user/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDb(t *testing.T) {
	dbSql := app.NewDb()
	assert.NotNil(t, dbSql)
}
