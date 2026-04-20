package main

import (
	"context"
	"controller/internal/proxycontroller"
	"log"
	"os"
	"os/signal"
	"syscall"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
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
	}

	controller.CreateProxy(context.Background(), 25565)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	log.Printf("Signal %v recieved\nStopping now...", sig)
}
