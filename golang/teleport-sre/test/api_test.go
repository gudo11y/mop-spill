package test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gudo11y/mop-spill/golang/teleport-sre/internal/api"
)

func TestReplicaCountHandler_MissingParam(t *testing.T) {
    s := &api.Server{}
    req := httptest.NewRequest("GET", "/replicas", nil)
    w := httptest.NewRecorder()

    s.ReplicaCountHandler(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected 400 BadRequest, got %d", w.Code)
    }
}