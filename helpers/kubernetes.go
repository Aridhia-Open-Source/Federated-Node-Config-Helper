package helpers

import (
	"context"
	"flag"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetClient() *kubernetes.Clientset {
	var config *rest.Config
	config, err := rest.InClusterConfig()

	// If we are not in-cluster
	if err != nil {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		// create the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		return clientset
	} else {
		// create the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		return clientset
	}
}

func CreateSecret(name string, data map[string]string) {
	secret := v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		StringData: data,
	}
	var v1 = GetClient()
	_, err := v1.CoreV1().Secrets("default").Create(context.TODO(), &secret, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

func CreateConfigMap(name string, data map[string]string) {
	secret := v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		Data: data,
	}
	var v1 = GetClient()
	_, err := v1.CoreV1().ConfigMaps("default").Create(context.TODO(), &secret, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}
