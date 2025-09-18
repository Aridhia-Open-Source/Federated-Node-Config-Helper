package helpers

import (
	"context"
	"flag"
	"fn-installer/components"
	"fn-installer/state"
	"path/filepath"
	"regexp"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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
		flag.Parse()
	}
}

func GetClient() *kubernetes.Clientset {
	var config *rest.Config
	config, err := rest.InClusterConfig()

	// If we are not in-cluster
	if err == nil {
		// create the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		return clientset
	}

	// use the current context in kubeconfig
	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func CreateSecret(name string, data map[string]string) {
	secret := v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: state.State.Namespace,
		},
		StringData: data,
	}
	var v1 = GetClient()
	_, err := v1.CoreV1().Secrets(state.State.Namespace).Create(context.TODO(), &secret, metav1.CreateOptions{})
	checkErrors(err)
}

func CreateConfigMap(name string, data map[string]string) {
	cm := v1.ConfigMap{
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
	var v1 = GetClient()
	_, err := v1.CoreV1().ConfigMaps(state.State.Namespace).Create(context.TODO(), &cm, metav1.CreateOptions{})
	checkErrors(err)
}

func CreateNamespace(name string) {
	ns := v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.NamespaceSpec{},
	}
	var v1 = GetClient()
	_, err := v1.CoreV1().Namespaces().Create(context.TODO(), &ns, metav1.CreateOptions{})
	checkErrors(err)
}

func checkErrors(err error) {
	if err != nil {
		msg := err.Error()
		if matched, _ := regexp.MatchString("already exists", msg); matched {
			components.ErrorBoard.SetText(msg)
		} else {
			panic(msg)
		}
	} else {
		components.ErrorBoard.SetText("")
	}
}
