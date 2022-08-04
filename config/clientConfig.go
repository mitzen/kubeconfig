package config

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"istio.io/istio/pkg/kube"
)

type ClientConfig struct {
	Kubeconfig *string
}

func (c *ClientConfig) initKubeConfig() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	//flag.Parse()
	c.Kubeconfig = kubeconfig
}

func (c *ClientConfig) NewRestConfig() *rest.Config {
	c.initKubeConfig()
	config, err := clientcmd.BuildConfigFromFlags("", *c.Kubeconfig)

	if err != nil {
		panic(err)
	}
	return config
}

func (c *ClientConfig) NewClientSet(config *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func (c *ClientConfig) NewExtendedClient(config *rest.Config) (kube.ExtendedClient, error) {
	c.initKubeConfig()
	cc, err := kube.NewExtendedClient(kube.BuildClientCmd(*c.Kubeconfig, ""), "")
	return cc, err
}
