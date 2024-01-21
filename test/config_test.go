package test

import (
	"journey-user/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDbConfiguration(t *testing.T) {
	config, err := app.SetDbConfiguration()
	assert.Nil(t, err)
	assert.Equal(t, "postgres", config.DBUser, "Should Be Equal")
	assert.Equal(t, "5432", config.DBPort, "PORT Should Be Equal")
}
