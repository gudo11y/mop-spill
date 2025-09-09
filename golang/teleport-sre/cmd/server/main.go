package main

import (
    "log"
    "net/http"

    "github.com/gudo11y/mop-spill/golang/teleport-sre/internal/api"
    "github.com/gudo11y/mop-spill/golang/teleport-sre/internal/kube"
)

func main() {
    kubeClient, err := kube.New()
    if err != nil {
        log.Fatalf("Failed to create kube client: %v", err)
    }

    s := &api.Server{KubeClient: kubeClient}
    http.HandleFunc("/replicas", s.ReplicaCountHandler)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}