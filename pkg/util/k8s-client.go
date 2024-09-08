package util

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	resourcev1alpha1 "ai.plugin/scheduler/pkg/client/clientset/versioned"
	resource "ai.plugin/scheduler/pkg/client/informers/externalversions"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	K8sClient        *kubernetes.Clientset
	Resourcev1alpha1 *resourcev1alpha1.Clientset
)

const timeDruation time.Duration = time.Second * 5

func InitK8sClient() error {

	var once sync.Once
	var err error
	once.Do(func() {
		K8sClient, err = NewK8sClient()
		if err != nil {
			return
		}
		Resourcev1alpha1, err = NewResourcev1alpha1Client()
		if err != nil {
			return
		}
	})

	return err
}

func NewResourcev1alpha1Client() (Resourcev1alpha1 *resourcev1alpha1.Clientset, err error) {
	conf, err := GetK8sConfig()
	if err != nil {
		return nil, err
	}
	c, err := resourcev1alpha1.NewForConfig(conf)
	return c, err
}

func NewK8sClient() (cli *kubernetes.Clientset, err error) {
	// fetch k8s incluster configuration
	config, err := GetK8sConfig()
	if err != nil {
		return
	}

	// create the k8s clientset
	cli, err = kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	return
}

func GetK8sConfig() (*rest.Config, error) {
	var (
		config *rest.Config
		err    error
	)
	if config, err = rest.InClusterConfig(); err != nil {
		var kubeconfig string
		cfg := viper.GetString("KUBECONFIG")
		if cfg != "" {
			kubeconfig = cfg
		} else if home := HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
		// Use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func NewResourcev1alpha1Informer() (*resource.SharedInformerFactory, error) {
	conf, err := GetK8sConfig()
	if err != nil {
		return nil, err
	}
	c, err := resourcev1alpha1.NewForConfig(conf)
	informer := resource.NewSharedInformerFactory(c, timeDruation)
	return &informer, err
}
