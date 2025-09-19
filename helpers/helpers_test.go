package helpers

import (
	"errors"
	"fn-installer/components"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
)

type FakeKubeClient struct {
	ShouldFail    bool
	AlreadyExists bool
	Create        func() error
	CreatedList   []string
}

func (f *FakeKubeClient) CreateNamespace(ns *v1.Namespace) (*v1.Namespace, error) {
	if f.ShouldFail {
		if f.AlreadyExists {
			return nil, errors.New("Namespace already exists")
		}
		return nil, errors.New("some other error")
	}
	return ns, nil
}
func (f *FakeKubeClient) CreateConfigMap(namespace string, cm *v1.ConfigMap) (*v1.ConfigMap, error) {
	if f.ShouldFail {
		if f.AlreadyExists {
			return nil, errors.New("Configmap already exists")
		}
		return nil, errors.New("some other error")
	}
	return cm, nil
}
func (f *FakeKubeClient) CreateSecret(namespace string, secret *v1.Secret) (*v1.Secret, error) {
	if f.ShouldFail {
		if f.AlreadyExists {
			return nil, errors.New("secret already exists")
		}
		return nil, errors.New("some other error")
	}
	return secret, nil
}

func TestErrorNamespaceHandler(t *testing.T) {
	client := &FakeKubeClient{ShouldFail: true, AlreadyExists: true}

	CreateNamespace(client, "new")
	assert.Equal(t, components.ErrorBoard.GetText(false), "Namespace already exists")
}
func TestErrorConfigMapHandler(t *testing.T) {
	client := &FakeKubeClient{ShouldFail: true, AlreadyExists: true}

	CreateConfigMap(client, "cmname", map[string]string{"field": "value"})
	assert.Equal(t, components.ErrorBoard.GetText(false), "Configmap already exists")
}
func TestErrorSecretHandler(t *testing.T) {
	client := &FakeKubeClient{ShouldFail: true, AlreadyExists: true}

	CreateSecret(client, "sec-name", map[string]string{"field": "value"})
	assert.Equal(t, components.ErrorBoard.GetText(false), "secret already exists")
}
func TestRandomErrorHandler(t *testing.T) {
	client := &FakeKubeClient{ShouldFail: true, AlreadyExists: false}

	CreateSecret(client, "sec-name", map[string]string{"field": "value"})
	assert.Equal(t, components.ErrorBoard.GetText(false), "some other error")
}
