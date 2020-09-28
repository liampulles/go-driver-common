package mux_test

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	gorilla "github.com/gorilla/mux"

	"github.com/liampulles/go-driver-common/http/mux"
	"github.com/stretchr/testify/assert"
)

func TestGoRequest_Headers_GivenValidGoRequest_ShouldExtract(t *testing.T) {
	// Setup fixture
	orig := httptest.NewRequest("GET", "http://127.0.0.1/", nil)
	orig.Header.Add("Key1", "VALUE1")
	orig.Header.Add("Key2", "VALUE21")
	orig.Header.Add("Key2", "VALUE22")
	sut := mux.GoRequest{
		Original: *orig,
	}

	// Setup expectations
	expected := map[string][]string{
		"Key1": {"VALUE1"},
		"Key2": {"VALUE21", "VALUE22"},
	}

	// Exercise SUT
	actual := sut.Headers()

	// Verify results
	assert.Equal(t, actual, expected)
}

func TestGoRequest_Body_GivenValidGoRequest_ShouldExtract(t *testing.T) {
	// Setup fixture
	r := strings.NewReader("some.data")
	orig := httptest.NewRequest("GET", "http://127.0.0.1/", r)
	sut := mux.GoRequest{
		Original: *orig,
	}

	// Setup expectations
	expected := []byte("some.data")

	// Exercise SUT
	actual, err := sut.Body()

	// Verify results
	assert.NoError(t, err)
	assert.Equal(t, actual, expected)
}

func TestGoRequest_Body_GivenInvalidGoRequest_ShouldExtract(t *testing.T) {
	// Setup fixture
	r := errReader(0)
	orig := httptest.NewRequest("GET", "http://127.0.0.1/", r)
	sut := mux.GoRequest{
		Original: *orig,
	}

	// Setup expectations
	expectedErr := "ioutil failed: test error"

	// Exercise SUT
	actual, err := sut.Body()

	// Verify results
	assert.EqualError(t, err, expectedErr)
	assert.Nil(t, actual)
}

func TestGoRequest_PathParams_GivenValidGoRequest_ShouldExtract(t *testing.T) {
	// Setup fixture
	orig := httptest.NewRequest("GET", "http://127.0.0.1/users/52", nil)
	orig = gorilla.SetURLVars(orig, map[string]string{
		"id": "52",
	})
	sut := mux.GoRequest{
		Original: *orig,
	}

	// Setup expectations
	expected := map[string]string{
		"id": "52",
	}

	// Exercise SUT
	actual := sut.PathParams()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestGoRequest_QueryParams_GivenValidGoRequest_ShouldExtract(t *testing.T) {
	// Setup fixture
	orig := httptest.NewRequest("GET", "http://127.0.0.1/users?id=52&names=john&names=susan", nil)
	sut := mux.GoRequest{
		Original: *orig,
	}

	// Setup expectations
	expected := map[string][]string{
		"id":    {"52"},
		"names": {"john", "susan"},
	}

	// Exercise SUT
	actual := sut.QueryParams()

	// Verify results
	assert.Equal(t, expected, actual)
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}
