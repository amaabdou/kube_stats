package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type PodsData struct {
	Headers []string
	Data    [][]string
}

func GetDefaultKubectlConfigPath() string {
	return filepath.Join(os.Getenv("HOME"), ".kube", "config")
}

func GetPods(kubeconfigPath string) (PodsData, error) {
	var podsData PodsData

	dcc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	rawConfig, err := dcc.RawConfig()
	if err != nil {
		return podsData, err
	}

	for loopedContextName := range rawConfig.Contexts {
		dcc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{
				ExplicitPath: kubeconfigPath,
			},
			&clientcmd.ConfigOverrides{
				ClusterInfo:    clientcmdapi.Cluster{Server: ""},
				CurrentContext: loopedContextName,
			},
		)

		clientConfig, err := dcc.ClientConfig()
		if err != nil {
			log.Println("Failed to get pods for config ", kubeconfigPath, " , context ", loopedContextName)
			log.Println(err.Error())
			continue
		}

		clientset, err := kubernetes.NewForConfig(clientConfig)
		if err != nil {
			log.Println("Failed to get pods for config ", kubeconfigPath, " , context ", loopedContextName)
			log.Println(err.Error())
			continue
		}

		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			log.Println("Failed to get pods for config ", kubeconfigPath, " , context ", loopedContextName)
			log.Println(err.Error())
			continue
		}

		for _, pod := range pods.Items {
			for _, container := range pod.Spec.Containers {
				imageAndVersion := strings.Split(container.Image, ":")
				imageName := imageAndVersion[0]
				imageVersion := "latest"
				if len(imageAndVersion) > 1 {
					imageVersion = imageAndVersion[1]
				}
				podsData.Data = append(podsData.Data, []string{
					imageName,
					imageVersion,
					pod.Name,
					pod.Namespace,
					loopedContextName,
				})
			}
		}
	}
	podsData.Headers = []string{
		"Container name",
		"Container tag",
		"Pod Name",
		"Namespace",
		"Context",
	}

	return podsData, nil
}
