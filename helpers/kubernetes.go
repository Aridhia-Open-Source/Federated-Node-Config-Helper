package helpers

import (
	"flag"
	"fmt"
	"fn-config-helper/components"
	"fn-config-helper/state"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/util/homedir"
)

var kubeconfig string
var kubeconfigPath *string

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}
	if flag.Lookup("kubeconfig") == nil {
		kubeconfigPath = flag.String("kubeconfig", kubeconfig, "absolute path to the kubeconfig file")
	}
}

func CreateSecret(k8s KubeClient, name string, data map[string]string, labels ...map[string]string) {
	var label map[string]string
	if len(labels) > 0 {
		label = labels[0]
	}
	for k, v := range data {
		if v == "" {
			components.ErrorBoard.SetText(
				fmt.Sprintf("%s\n%s cannot be empty", components.ErrorBoard.GetText(true), k),
			)
		}
	}
	if components.ErrorBoard.GetText(true) != "" {
		return
	} else {
		components.ErrorBoard.SetText("")
	}
	secret := &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: state.State.Namespace,
			Labels:    label,
		},
		StringData: data,
	}
	_, err := k8s.CreateSecret(state.State.Namespace, secret)
	checkErrors(err)
}

func CreateConfigMap(k8s KubeClient, name string, data map[string]string) {
	cm := &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: state.State.Namespace,
		},
		Data: data,
	}
	_, err := k8s.CreateConfigMap(state.State.Namespace, cm)
	checkErrors(err)
}

func CreateNamespace(k8s KubeClient, name string) {
	ns := &v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.NamespaceSpec{},
	}
	_, err := k8s.CreateNamespace(ns)
	checkErrors(err)
}

func checkErrors(err error) {
	if err != nil {
		components.ErrorBoard.SetText(err.Error())
	} else {
		components.ErrorBoard.SetText("")
	}
}
