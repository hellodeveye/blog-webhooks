package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPullRequest(t *testing.T) {
	gitPath = "./"
	t.Run("returns pull successfully!", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		pullRequest(response, request)
		got := response.Body.String()
		want := "pull successfully!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
