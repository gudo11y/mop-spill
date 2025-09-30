package kube

import (
    "context"
    "fmt"
    "os"

    // appsv1 "k8s.io/api/apps/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Client struct {
    Clientset *kubernetes.Clientset
    Namespace string
}

func New() (*Client, error) {
    var config *rest.Config
    var err error

    kubeconfig := os.Getenv("KUBECONFIG")
    if kubeconfig != "" {
        config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
    } else {
        config, err = rest.InClusterConfig()
    }
    if err != nil {
        return nil, fmt.Errorf("error building kube config: %w", err)
    }

    cs, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, fmt.Errorf("error creating clientset: %w", err)
    }

    ns := os.Getenv("NAMESPACE")
    if ns == "" {
        ns = "default"
    }

    return &Client{
        Clientset: cs,
        Namespace: ns,
    }, nil
}

func (c *Client) GetReplicaCount(ctx context.Context, deploymentName string) (int32, error) {
    deploy, err := c.Clientset.AppsV1().Deployments(c.Namespace).Get(ctx, deploymentName, metav1.GetOptions{})
    if err != nil {
        return 0, err
    }
    return *deploy.Spec.Replicas, nil
}