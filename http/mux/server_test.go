package mux_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/go-driver-common/http"
	"github.com/liampulles/go-driver-common/http/mux"
	"github.com/stretchr/testify/assert"
)

func TestGoServer_Start_GivenValidGoServer_ShouldListenAndServe(t *testing.T) {
	// Setup fixture
	orig := &mockListener{err: nil}
	sut := &mux.GoServer{
		Original: orig,
	}

	// Exercise SUT
	err := sut.Start()

	// Verify results
	assert.NoError(t, err)
}

func TestGoServer_Start_GivenInvalidGoServer_ShouldListenAndServe(t *testing.T) {
	// Setup fixture
	orig := &mockListener{err: fmt.Errorf("some.error")}
	sut := &mux.GoServer{
		Original: orig,
	}

	// Setup expectations
	expectedErr := "go http server error: some.error"

	// Exercise SUT
	err := sut.Start()

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestNewServer_GivenValidFixture_ShouldReturnAsExpected(t *testing.T) {
	// Setup fixture
	passingHandler := func(http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
		}
	}
	mappings := []http.HandlerMapping{
		http.HandlerMapping{
			Method:      "GET",
			PathPattern: "/api/v1/users/{id}",
			Handler:     passingHandler,
		},
	}

	// Exercise SUT
	actual := mux.NewServer(mappings, 8080)

	// Verify results
	assert.NotNil(t, actual)
}

type mockListener struct {
	err error
}

var _ mux.Listener = &mockListener{}

func (m *mockListener) ListenAndServe() error {
	return m.err
}
