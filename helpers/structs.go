package helpers

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type DBSecret struct {
	Name string
	Key  string
}
type DB struct {
	User     string
	Name     string
	Hostname string
	Port     int
	Secret   DBSecret
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
	Azure    *AzureStorage `yaml:"aks,omitempty"`
	Local    *LocalStorage `yaml:"local,omitempty"`
}

type NginxExtraArgs struct {
	DefaultSslCertificate string `yaml:"default-ssl-certificate"`
}

type NginxClass struct {
	Name string
}

type NginxController struct {
	AllowSnippetAnnotations bool           `yaml:"allowSnippetAnnotations"`
	IngressClass            string         `yaml:"ingressClass"`
	IngressClassResource    NginxClass     `yaml:"ingressClassResource"`
	ExtraArgs               NginxExtraArgs `yaml:"extraArgs"`
}

type NginxConfig struct {
	Enabled           bool
	NamespaceOverride string `yaml:"namespaceOverride"`
	Controller        NginxController
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
	Enabled    bool
	Namespace  string
	InstallCRD bool `yaml:"installCRDs"`
}
type AzureCerts struct {
	SecretName string `yaml:"secretName"`
	Configmap  string
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
	TaskReview bool `yaml:"taskReview"`
	Host       string
}

type Config struct {
	LocalDevelopment bool `yaml:"local_development"`
	Namespaces       Namespaces
	Database         DB   `yaml:"db"`
	CleanupTime      int  `yaml:"cleanupTime"`
	OutboundMode     bool `yaml:"outboundMode"`
	TaskReview       bool `yaml:"taskReview"`
	Smoketests       bool
	Storage          Storage
	Host             string
	Keycloak         Keycloak
	OnAks            bool             `yaml:"on_aks"`
	OnEks            bool             `yaml:"on_eks"`
	ControllerConfig ControllerConfig `yaml:"fn-task-controller"`
	NginxIngress     NginxConfig      `yaml:"nginx-ingress"`
	CertManager      CertConfig       `yaml:"cert-manager"`
	Certs            Certs
	Global           GlobalConfig
}

func (conf Config) CreateYaml() {
	yamlFile, err := yaml.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("values.yaml")
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
