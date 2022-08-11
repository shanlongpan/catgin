package k8sresult

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/catgin/config"
	"github.com/shanlongpan/catgin/xlog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func get(ctx *gin.Context) {
	pods, err := config.K8sClient.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(200, pods)
}

func getNameSpacePod(ctx *gin.Context) {
	// 获取指定 namespace 中的 Pod 列表信息
	namespace := "kube-flannel"
	pods, err := config.K8sClient.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		xlog.Errorln(ctx, pods)
	}
	ctx.JSON(200, pods)
}

func getPod(ctx *gin.Context) {
	// 获取指定 namespaces 和 podName 的详细信息，使用 error handle 方式处理错误信息
	namespace := "kube-system"
	podName := "coredns-74586cf9b6-nnl72"
	pod, err := config.K8sClient.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if errors.IsNotFound(err) {

		ctx.JSON(200, fmt.Sprintf("Pod %s in namespace %s not found", podName, namespace))
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		ctx.JSON(200, fmt.Sprintf("Error getting pod %s in namespace %s: %v", podName, namespace, statusError.ErrStatus.Message))

	} else if err != nil {
		panic(err.Error())
	} else {

		ctx.JSON(200, pod)
	}

}
