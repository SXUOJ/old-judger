package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sxu-Online-Judge/judger/judge"
	"gotest.tools/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

func TestSubmitRouteParameter(t *testing.T) {
	router := setupRouter()

	params := judge.Submit{
		SubmitId:       "test_1",
		ProblemId:      "1",
		ProblemType:    "1",
		CodeType:       "C",
		CodeSourcePath: "test/tmp/main.c",
		TimeLimit:      "1000",
		MemoryLimit:    "268435456", // * 1024 * 2014
	}

	paramsByte, _ := json.Marshal(params)
	// log.Printf("%s\n", paramsByte)
	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(paramsByte))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
