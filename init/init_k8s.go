package init

/**
*  @Author:Tristan
*  @Date: 2022/8/12
 */

import (
	"context"
	"flag"
	"github.com/shanlongpan/catgin/config"
	"github.com/shanlongpan/catgin/xlog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func init() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	configs, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		xlog.Errorln(context.TODO(), err)
	}

	// 根据指定的 config 创建一个新的 clientset
	config.K8sClient, err = kubernetes.NewForConfig(configs)
	if err != nil {
		xlog.Errorln(context.TODO(), err)
	}
}
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
