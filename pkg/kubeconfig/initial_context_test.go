package kubeconfig

import (
	"testing"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func TestIsContextUsable(t *testing.T) {
	cfg := clientcmdapi.NewConfig()

	// good context -> cluster with server
	cfg.Clusters["cluster1"] = &clientcmdapi.Cluster{Server: "https://127.0.0.1:6443"}
	cfg.Contexts["good"] = &clientcmdapi.Context{Cluster: "cluster1"}

	// context with missing cluster
	cfg.Contexts["noCluster"] = &clientcmdapi.Context{Cluster: "missing"}

	// context with empty server
	cfg.Clusters["clusterEmpty"] = &clientcmdapi.Cluster{Server: ""}
	cfg.Contexts["emptyServer"] = &clientcmdapi.Context{Cluster: "clusterEmpty"}

	if !isContextUsable(cfg, "good") {
		t.Fatalf("expected 'good' context to be usable")
	}

	if isContextUsable(cfg, "noCluster") {
		t.Fatalf("expected 'noCluster' to be unusable because the referenced cluster is missing")
	}

	if isContextUsable(cfg, "emptyServer") {
		t.Fatalf("expected 'emptyServer' to be unusable because cluster server is empty")
	}

	if isContextUsable(cfg, "missingContext") {
		t.Fatalf("expected 'missingContext' to be unusable because it does not exist")
	}
}
