package helpers

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type RealKubeClient struct {
	clientset *kubernetes.Clientset
}

func NewRealKubeClient() (*RealKubeClient, error) {
	// your existing GetClient logic here
	// i.e., in-cluster or kubeconfig

	var config *rest.Config
	config, err := rest.InClusterConfig()

	// If we are not in-cluster
	if err == nil {
		// create the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		return &RealKubeClient{clientset}, nil
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
	return &RealKubeClient{clientset}, nil
}

func (r *RealKubeClient) CreateSecret(namespace string, secret *v1.Secret) (*v1.Secret, error) {
	return r.clientset.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
}

func (r *RealKubeClient) CreateConfigMap(namespace string, cm *v1.ConfigMap) (*v1.ConfigMap, error) {
	return r.clientset.CoreV1().ConfigMaps(namespace).Create(context.TODO(), cm, metav1.CreateOptions{})
}

func (r *RealKubeClient) CreateNamespace(ns *v1.Namespace) (*v1.Namespace, error) {
	return r.clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
}
