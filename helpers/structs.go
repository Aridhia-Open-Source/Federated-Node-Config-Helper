package helpers

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type DBSecret struct {
	Name string
	Key  string
}
type DB struct {
	User   string
	Name   string
	Host   string
	Port   int
	Secret DBSecret
}

type AwsStorage struct {
	FileSystemId  string `yaml:"fileSystemId"`
	AccessPointId string `yaml:"accessPointId"`
}

type AzureStorage struct {
	SecretName         string `yaml:"secretName"`
	StorageAccountName string `yaml:"storageAccountName"`
	StorageAccountKey  string `yaml:"storageAccountKey"`
	ShareName          string `yaml:"shareName"`
}

type LocalStorage struct {
	Path   string
	Dbpath string
}

type Storage struct {
	Capacity string
	Aws      *AwsStorage   `yaml:"aws,omitempty"`
	Azure    *AzureStorage `yaml:"azure,omitempty"`
	Local    *LocalStorage `yaml:"local,omitempty"`
}

type FirstUser struct {
	Name      string `yaml:"name,omitempty"`
	UserKey   string `yaml:"userKey,omitempty"`
	PassKey   string `yaml:"passKey,omitempty"`
	FirstName string `yaml:"firstName,omitempty"`
	LastName  string `yaml:"lastName,omitempty"`
	Email     string `yaml:"email,omitempty"`
}

type GatewayClass struct {
	Name string
}

type TraefikServiceSpec struct {
	ExternalTrafficPolicy string `yaml:"externalTrafficPolicy"`
}

type TraefikService struct {
	Spec TraefikServiceSpec
}

type TraefikGateway struct {
	Name string
}

type Traefik struct {
	Enabled      *bool
	Gateway      TraefikGateway
	GatewayClass GatewayClass `yaml:"gatewayClass"`
	Service      *TraefikService
}

type IdpGH struct {
	SecretName  string `yaml:"secret_name"`
	SecretKey   string `yaml:"secret_key"`
	ClientIDKey string `yaml:"clientid_key"`
}
type Idp struct {
	Github IdpGH
}
type DeliveryGH struct {
	Repository string
}
type DeliveryOther struct {
	Url      string
	AuthType string `yaml:"auth_type"`
}
type Delivery struct {
	Github *DeliveryGH    `yaml:"github,omitempty"`
	Other  *DeliveryOther `yaml:"other,omitempty"`
}

type ControllerConfig struct {
	Idp      Idp
	Storage  Storage
	Delivery Delivery
}

type Namespaces struct {
	Keycloak   string
	Controller string
	Tasks      string
}

type CertConfig struct {
	Enabled    *bool
	Namespace  string
	InstallCRD *bool `yaml:"installCRDs"`
}
type AzureCerts struct {
	Secret    string
	Configmap string
}
type Certs struct {
	RotationPolicy string `yaml:"rotationPolicy"`
	Azure          *AzureCerts
	AWS            string
}
type Keycloak struct {
	Replicas int
}
type GlobalConfig struct {
	Namespaces Namespaces
	TaskReview *bool `yaml:"taskReview"`
	Host       string
}

type Config struct {
	LocalDevelopment *bool `yaml:"local_development"`
	Namespaces       Namespaces
	Database         DB    `yaml:"db"`
	CleanupTime      int   `yaml:"cleanupTime"`
	OutboundMode     *bool `yaml:"outboundMode"`
	TaskReview       *bool `yaml:"taskReview"`
	Smoketests       *bool
	Storage          Storage
	Host             string
	Keycloak         Keycloak
	OnAks            bool             `yaml:"on_aks"`
	OnEks            bool             `yaml:"on_eks"`
	ControllerConfig ControllerConfig `yaml:"fn-task-controller"`
	Traefik          Traefik          `yaml:"traefik"`
	CertManager      CertConfig       `yaml:"cert-manager"`
	Certs            Certs
	FirstUserSecret  *FirstUser `yaml:"firstUserSecret"`
	Global           GlobalConfig
}

func (conf Config) CreateYaml(fileName string) {
	yamlFile, err := yaml.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fileName)
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

func ReadYAML(filename string) (*Config, *ArgoCD) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var result = &Config{}
	var resultArgo = &ArgoCD{}
	matched, _ := regexp.MatchString("apiVersion: argoproj.io/v1alpha1", string(data))
	if matched {
		// Unmarshal YAML into map
		err = yaml.Unmarshal(data, &resultArgo)
		if err != nil {
			panic(err)
		}
		return nil, resultArgo

	}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	return result, nil
}
