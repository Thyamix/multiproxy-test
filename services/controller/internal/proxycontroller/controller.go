package proxycontroller

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	DESIRED   string = "desired"
	UNDESIRED string = "undesired"
)

type ProxyController struct {
	Clientset kubernetes.Clientset
	Proxies   []*corev1.Pod
	Namespace string
}

func (c *ProxyController) CreateProxy(ctx context.Context, port int32) error {
	fmt.Printf("Creating a new proxy on port %v\n", port)
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "proxy",
			Namespace: c.Namespace,
			Labels: map[string]string{
				"app":         "proxy",
				"proxy-state": DESIRED,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "proxy",
					Image: "nginx:alpine",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: port,
						},
					},
				},
			},
		},
	}

	newProxy, err := c.Clientset.CoreV1().Pods(c.Namespace).Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	c.Proxies = append(c.Proxies, newProxy)

	return nil
}

func (c *ProxyController) GetProxy(ctx context.Context, name string) (*corev1.Pod, error) {
	return c.Clientset.CoreV1().Pods(c.Namespace).Get(ctx, name, metav1.GetOptions{})
}

func (c *ProxyController) DeleteProxy(ctx context.Context, name string) error {
	return c.Clientset.CoreV1().Pods(c.Namespace).Delete(ctx, name, metav1.DeleteOptions{})
}
