package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	metricsv1beta1 "k8s.io/metrics/pkg/client/clientset/versioned"
)

// PodMetrics holds data for a pod's resource usage
type PodMetrics struct {
	PodName string
	CPU     string
	Memory  string
}

// getClusterMetrics fetches the CPU and memory usage of each pod in the default namespace
func getClusterMetrics() []PodMetrics {
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error creating Kubernetes configuration: %v", err)
	}

	// Create a Kubernetes clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	log.Fatalf("Error creating Kubernetes clientset: %v", err)
	// }

	// Create a metrics clientset for fetching metrics
	metricsClient, err := metricsv1beta1.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Metrics clientset: %v", err)
	}

	// Fetch pod metrics from the default namespace
	podMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("Error fetching pod metrics: %v", err)
	}

	// Prepare a slice to hold the metrics data
	var podMetrics []PodMetrics

	for _, podMetricsItem := range podMetricsList.Items {
		cpuUsage := podMetricsItem.Containers[0].Usage["cpu"]
		memoryUsage := podMetricsItem.Containers[0].Usage["memory"]

		podMetrics = append(podMetrics, PodMetrics{
			PodName: podMetricsItem.Name,
			CPU:     cpuUsage.String(),
			Memory:  memoryUsage.String(),
		})
	}

	return podMetrics
}

// Utility function to get the home directory
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}
