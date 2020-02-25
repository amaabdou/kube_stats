package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
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

		services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
		if err != nil {
			log.Println("Failed to get services for config ", kubeconfigPath, " , context ", loopedContextName)
			log.Println(err.Error())
			continue
		}

		ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{})
		if err != nil {
			log.Println("Failed to get ingress for config ", kubeconfigPath, " , context ", loopedContextName)
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

				svcName := ""
				ExternalIPs := ""
				for _, service := range services.Items {
					for _, selector := range service.Spec.Selector {
						for _, label := range pod.ObjectMeta.Labels {
							if label == selector && label!="" {
								fmt.Println(label, selector)

								svcName = svcName + service.Name
								ExternalIPs = ExternalIPs + strings.Join(service.Spec.ExternalIPs, ",")
							}
						}
					}
				}

				ingressUrl := ""
				for _, ingress := range ingresses.Items {
					if ingress.Spec.Backend != nil {
						if ingress.Spec.Backend.ServiceName == svcName {
							log.Println("Could not process ", ingress.Name," ", ingress)
						}
					}else {
						for _, rule := range ingress.Spec.Rules {
							for _, path := range rule.IngressRuleValue.HTTP.Paths {
								if path.Backend.ServiceName == svcName {
									ingressUrl = ingressUrl + fmt.Sprint(
										rule.Host,
										path.Path,
									)
								}
							}
						}
					}
				}

				podsData.Data = append(podsData.Data, []string{
					imageName,
					imageVersion,
					pod.Name,
					pod.Namespace,
					loopedContextName,
					svcName,
					ExternalIPs,
					ingressUrl,
				})
			}
		}
	}
	podsData.Headers = []string{
		"Cont. name",
		"Cont. tag",
		"PO Name",
		"NS",
		"Context",
		"Svc-name",
		"ExternalIPs",
		"Ingress",
	}

	return podsData, nil
}
