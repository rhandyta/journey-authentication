package test

import (
	"fmt"
	"journey-user/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDbConfiguration(t *testing.T) {
	config, err := app.SetDbConfiguration()
	fmt.Println(config)
	assert.Nil(t, err)
	assert.Equal(t, "postgres", config.DBUser, "Should Be Equal")
	assert.Equal(t, "5432", config.DBPort, "PORT Should Be Equal")
}
