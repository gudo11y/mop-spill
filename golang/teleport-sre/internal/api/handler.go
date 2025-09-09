package api

import (
    "encoding/json"
    "net/http"

    "github.com/gudo11y/mop-spill/golang/teleport-sre/internal/kube"
)

type Server struct {
    KubeClient *kube.Client
}

func (s *Server) ReplicaCountHandler(w http.ResponseWriter, r *http.Request) {
    deployName := r.URL.Query().Get("deployment")
    if deployName == "" {
        http.Error(w, "Missing 'deployment' query param", http.StatusBadRequest)
        return
    }

    replicas, err := s.KubeClient.GetReplicaCount(r.Context(), deployName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    resp := map[string]interface{}{
        "deployment": deployName,
        "replicas":   replicas,
    }
    json.NewEncoder(w).Encode(resp)
}