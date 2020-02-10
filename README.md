# Small utility to collect stats and information from kubernetes multiple contexts

it used the default kubectl config to loop over all contexts and loads all pods infos       
then exports it to different formats, like csv

## Usage
Usage of /tmp/go-build373490543/b001/exe/kubecleanup:
  -gb string
        Group output values by [none,imageName] (default "imageName")
  -kc string
        kubectl config file location. (default "/home/amaabdou/.kube/config")
  -wr string
        Write to [console,json,csv] (default "console")

## example

```
+----------------------------------------------------------------+---------------+--------------------------------------------+----------------------+----------+
|                         CONTAINER NAME                         | CONTAINER TAG |                  POD NAME                  |      NAMESPACE       | CONTEXT  |
+----------------------------------------------------------------+---------------+--------------------------------------------+----------------------+----------+
| nginxdemos/hello                                               | latest        | myapp-pod                                  | default              | minikube |
+                                                                +               +                                            +----------------------+          +
|                                                                |               |                                            | production           |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+----------------------+          +
| k8s.gcr.io/coredns                                             | 1.6.5         | coredns-6955765f44-kbxl4                   | kube-system          |          |
+                                                                +               +--------------------------------------------+                      +          +
|                                                                |               | coredns-6955765f44-s678c                   |                      |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+                      +          +
| k8s.gcr.io/etcd                                                | 3.4.3-0       | etcd-minikube                              |                      |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+                      +          +
| k8s.gcr.io/kube-apiserver                                      | v1.17.0       | kube-apiserver-minikube                    |                      |          |
+----------------------------------------------------------------+               +--------------------------------------------+                      +          +
| k8s.gcr.io/kube-proxy                                          |               | kube-proxy-d2kpn                           |                      |          |
+----------------------------------------------------------------+               +--------------------------------------------+                      +          +
| k8s.gcr.io/kube-controller-manager                             |               | kube-controller-manager-minikube           |                      |          |
+----------------------------------------------------------------+               +--------------------------------------------+                      +          +
| k8s.gcr.io/kube-scheduler                                      |               | kube-scheduler-minikube                    |                      |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+                      +          +
| cryptexlabs/minikube-ingress-dns                               | 0.2.1         | kube-ingress-dns-minikube                  |                      |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+                      +          +
| quay.io/kubernetes-ingress-controller/nginx-ingress-controller | 0.26.1        | nginx-ingress-controller-6fc5bcc8c9-pbd4b  |                      |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+                      +          +
| gcr.io/k8s-minikube/storage-provisioner                        | v1.8.1        | storage-provisioner                        |                      |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+----------------------+          +
| kubernetesui/metrics-scraper                                   | v1.0.2        | dashboard-metrics-scraper-7b64584c5c-d9pf2 | kubernetes-dashboard |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+                      +          +
| kubernetesui/dashboard                                         | v2.0.0-beta8  | kubernetes-dashboard-79d9cd965-jfzwk       |                      |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+----------------------+          +
| nginx                                                          | alpine        | nginx                                      | default              |          |
+----------------------------------------------------------------+---------------+--------------------------------------------+----------------------+----------+

```
