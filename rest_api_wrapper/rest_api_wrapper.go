package rest_api_wrapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Check200StatusCode(t *testing.T, statusCode int) {
	assert.Equal(t, statusCode, 200, "Api req should return 200")
}
