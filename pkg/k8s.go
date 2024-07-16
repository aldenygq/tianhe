package pkg

import (
	"context"
	"encoding/base64"
	"fmt"
	"errors"
	//"net/http"
	//"io/ioutil"
	//"tianhe/middleware"
	//"k8s.io/kubectl/pkg/scheme"
	//"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	
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
func (k *K8sClient) NsInfo(ns string) (*coreV1.Namespace,error) {
	defer k.CloseClient()
	namespaceInfo,err := k.Client.CoreV1().Namespaces().Get(context.TODO(),ns,metaV1.GetOptions{})
	if err != nil {
		return namespaceInfo,err
	}
	return namespaceInfo,nil
}
func (k *K8sClient) NsList() (*coreV1.NamespaceList,error) {
	defer k.CloseClient()
	namespaceList,err := k.Client.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return namespaceList,err
	}
	return namespaceList,nil
}

func (k *K8sClient) PodList(ns string) (*coreV1.PodList,error) {
	defer k.CloseClient()
    // 列出Pods
    pods, err := k.Client.CoreV1().Pods(ns).List(context.TODO(), metaV1.ListOptions{})
    if err != nil {
        return nil,err 
    }
	return pods,nil 
}
func (k *K8sClient) PodInfo(ns,podname string) (*coreV1.Pod,error) {
	defer k.CloseClient()
	pod, err := k.Client.CoreV1().Pods(ns).Get(context.TODO(), podname, metaV1.GetOptions{})
    if err != nil {
        return nil,err 
    }
	return pod,nil 
}

func (k *K8sClient) Event(ns,podname string) (*coreV1.EventList,error) {
	defer k.CloseClient()
	events, err := k.Client.CoreV1().Events(ns).List(context.TODO(), metaV1.ListOptions{
        FieldSelector: fmt.Sprintf("involvedObject.name=%s", podname),
    })
    if err != nil {
        return nil,err 
    }
	return events,nil 
}
func (k *K8sClient) NodeList() (*coreV1.NodeList,error) {
	defer k.CloseClient()
	nodes, err := k.Client.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
    if err != nil {
        return nil,err 
    }
	return nodes,nil 
}

func (k *K8sClient) NodeInfo(nodename string) (*coreV1.Node,error) {
	defer k.CloseClient()
	node,err := k.Client.CoreV1().Nodes().Get(context.TODO(),nodename,metaV1.GetOptions{})
	if err != nil {
		return nil,err 
	}
	return node,nil 
}

func (k *K8sClient) NodeLable(nodename string) (map[string]string,error) {
	defer k.CloseClient()
	node,err := k.Client.CoreV1().Nodes().Get(context.TODO(),nodename,metaV1.GetOptions{})
	if err != nil {
		return nil,err 
	}
	return node.Labels,nil 
}

func (k *K8sClient) NodeTaint(nodename string) ([]coreV1.Taint,error) {
	defer k.CloseClient()
	node,err := k.Client.CoreV1().Nodes().Get(context.TODO(),nodename,metaV1.GetOptions{})
	if err != nil {
		return nil,err 
	}
	return node.Spec.Taints,nil 
}

func (k *K8sClient) PatchNodeLable(nodename string,labels map[string]string) error {
	defer k.CloseClient()
	node,err := k.Client.CoreV1().Nodes().Get(context.TODO(),nodename,metaV1.GetOptions{})
	if err != nil {
		return err 
	}
	// 添加或更新标签
	if node.Labels == nil {
		node.Labels = map[string]string{}
	}
	for k,v := range labels {
		node.Labels[k] = v
	}
	// 更新节点
	_, err = k.Client.CoreV1().Nodes().Update(context.TODO(), node, metaV1.UpdateOptions{})
	if err != nil {
		return err 
	}
	return nil 
}
func (k *K8sClient) PatchNodeTaint(nodename string,taints map[string]string) error {
	defer k.CloseClient()
	node,err := k.Client.CoreV1().Nodes().Get(context.TODO(),nodename,metaV1.GetOptions{})
	if err != nil {
		return err 
	}
	
	for k,v := range taints {
		taint := &coreV1.Taint{
			Key:    k,
			Value:  v,
			Effect: coreV1.TaintEffectNoSchedule,
		}
		node.Spec.Taints = append(node.Spec.Taints, *taint)
	}
	// 更新节点
	_, err = k.Client.CoreV1().Nodes().Update(context.TODO(), node, metaV1.UpdateOptions{})
	if err != nil {
		return err 
	}
	return nil 
}

func (k *K8sClient) PatchNodeSchedule(nodename,schedulerule string) error {
	defer k.CloseClient()
	node,err := k.Client.CoreV1().Nodes().Get(context.TODO(),nodename,metaV1.GetOptions{})
	if err != nil {
		return err 
	}
	switch schedulerule {
	case "disable":
		node.Spec.Unschedulable = true
	case "enable":
		node.Spec.Unschedulable = false
	default:
		return errors.New("schedule rule invalid")
	}
	_, err = k.Client.CoreV1().Nodes().Update(context.TODO(), node, metaV1.UpdateOptions{})
    if err != nil {
        return err 
    }
	return nil 
}
func (k *K8sClient) Log(ns,podname string) (runtime.Object,error) {
	defer k.CloseClient()
	podLogs,err := k.Client.CoreV1().Pods(ns).GetLogs(podname, &coreV1.PodLogOptions{}).Do(context.TODO()).Get()
	if err != nil {
		return nil,err 
	}
	/*
    logsReq, err := http.NewRequest("GET", podLogs.URL(), nil)
    if err != nil {
        return "",err 
    }
	// 执行请求并读取日志
    logsResp, err := http.DefaultClient.Do(logsReq)
    if err != nil {
        return "",err 
    }
    defer logsResp.Body.Close()
 
    if logsResp.StatusCode != http.StatusOK {
        return "",err 
    }
 
    body, err := ioutil.ReadAll(logsResp.Body)
    if err != nil {
        return "",err 
    }
	*/
	return podLogs,nil 
}

func (k *K8sClient) CloseClient() {
	k.Client = nil
}
