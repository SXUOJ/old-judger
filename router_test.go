package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sxu-Online-Judge/judger/model"
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

	params := model.Submit{
		SubmitId:       "1",
		ProblemId:      "1",
		ProblemType:    "1",
		CodeType:       "c",
		CodeSourcePath: "main.c",
		Limit: model.Limit{
			TimeLimit:   2000,
			MemoryLimit: 256,
		},
	}

	paramsByte, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(paramsByte))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
