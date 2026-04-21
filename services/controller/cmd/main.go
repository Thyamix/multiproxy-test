package main

import (
	"context"
	"controller/internal/proxycontroller"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	fmt.Println("Starting controller")
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	controller := proxycontroller.ProxyController{
		Clientset: *clientset,
		Proxies:   []*corev1.Pod{},
		Namespace: "multiproxy-test",
	}

	fmt.Println("Starting proxy creation")
	err = controller.CreateProxy(context.Background(), 25565)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Starting proxy creation")
	err = controller.CreateProxy(context.Background(), 25566)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Starting proxy creation")
	err = controller.CreateProxy(context.Background(), 25567)
	if err != nil {
		fmt.Println(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	controller.ClearProxies(context.Background())
	log.Printf("Signal %v recieved\nStopping now...", sig)
}
