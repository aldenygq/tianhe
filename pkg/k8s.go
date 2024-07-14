package pkg

import (
	"context"
	"encoding/base64"
	//"tianhe/middleware"

	//"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreV1 "k8s.io/api/core/v1"
)

type K8sClient struct {
	Client *kubernetes.Clientset
}

func NewK8sClient(base64kubeconfig string) (*K8sClient,error){
	decoded, _ := base64.StdEncoding.DecodeString(base64kubeconfig)
	// 使用kubeconfig文件来获取客户端配置
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(decoded))
	if err != nil {
		return nil,err 
	}
 
	// 根据客户端配置创建一个Kubernetes客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil,err 
	}
 
	return &K8sClient{
		Client:clientset,
	},nil 
}

func (k *K8sClient) CreateNs(ns string) error {
	// 创建namespace的spec
	var namespace coreV1.Namespace
	namespace.Name = ns
	defer k.CloseClient()
	_, err := k.Client.CoreV1().Namespaces().Create(context.TODO(),&namespace,metaV1.CreateOptions{})
    if err != nil {
        return err 
    }
	return nil 
}

func (k *K8sClient) CloseClient() {
	k.Client = nil
}
