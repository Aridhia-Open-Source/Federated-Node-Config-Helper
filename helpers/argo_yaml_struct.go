package helpers

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type MD struct {
	Name      string
	Namespace string
}

type Automated struct{}

type ArgoSync struct {
	Automated   *Automated `yaml:"automated,omitempty"`
	SyncOptions []string   `yaml:"syncOptions"`
}

type ArgoHelm struct {
	Values string
}

type ArgoSource struct {
	RepoURL        string `yaml:"repoURL"`
	Path           string
	TargetRevision string `yaml:"targetRevision"`
	Helm           ArgoHelm
}

type ArgoDestination struct {
	Server    string
	Namespace string
}

type ArgoSpec struct {
	Project     string
	Source      ArgoSource
	Destination ArgoDestination
	SyncPolicy  ArgoSync `yaml:"syncPolicy"`
}

type ArgoCD struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   MD
	Spec       ArgoSpec
}

func InitArgoStruct() ArgoCD {
	var baseArgo ArgoCD
	baseArgo.ApiVersion = "argoproj.io/v1alpha1"
	baseArgo.Kind = "Application"
	baseArgo.Metadata.Name = "federatednode"
	baseArgo.Spec.Destination.Server = "https://kubernetes.default.svc"
	baseArgo.Spec.Source.RepoURL = "https://github.com/Aridhia-Open-Source/PHEMS_federated_node"
	baseArgo.Spec.Source.Path = "k8s/federated-node"

	return baseArgo
}

func (conf ArgoCD) CreateYaml() {
	yamlFile, err := yaml.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("argo-app-deployment.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Writer.Write(f, yamlFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("File created!")
}
