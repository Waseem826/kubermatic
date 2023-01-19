/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package konnectivity

import (
	"fmt"

	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/registry"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var (
	defResourceRequirements = corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceMemory: resource.MustParse("10Mi"),
			corev1.ResourceCPU:    resource.MustParse("10m"),
		},
		Limits: corev1.ResourceList{
			corev1.ResourceMemory: resource.MustParse("200Mi"),
			corev1.ResourceCPU:    resource.MustParse("2"),
		},
	}
)

// ProxySidecar returns container that runs konnectivity proxy server as a sidecar in apiserver pods.
func ProxySidecar(data *resources.TemplateData, serverCount int32) (*corev1.Container, error) {
	const (
		name    = "k8s-artifacts-prod/kas-network-proxy/proxy-server"
		version = "v0.0.35"
	)

	return &corev1.Container{
		Name:            resources.KonnectivityServerContainer,
		Image:           registry.Must(data.RewriteImage(fmt.Sprintf("%s/%s:%s", resources.RegistryEUGCR, name, version))),
		ImagePullPolicy: corev1.PullIfNotPresent,
		Command: []string{
			"/proxy-server",
		},
		Args: []string{
			"--logtostderr=true",
			"-v=3",
			fmt.Sprintf("--cluster-key=/etc/kubernetes/pki/%s.key", resources.KonnectivityProxyTLSSecretName),
			fmt.Sprintf("--cluster-cert=/etc/kubernetes/pki/%s.crt", resources.KonnectivityProxyTLSSecretName),
			"--uds-name=/etc/kubernetes/konnectivity-server/konnectivity-server.socket",
			fmt.Sprintf("--kubeconfig=/etc/kubernetes/kubeconfig/%s", resources.KonnectivityServerConf),
			fmt.Sprintf("--server-count=%d", serverCount),
			"--mode=grpc",
			"--server-port=0",
			"--agent-port=8132",
			"--admin-port=8133",
			"--health-port=8134",
			"--agent-namespace=kube-system",
			fmt.Sprintf("--agent-service-account=%s", resources.KonnectivityServiceAccountName),
			"--delete-existing-uds-file=true",
			"--authentication-audience=system:konnectivity-server",
			// TODO rastislavs: switch to "--proxy-strategies=destHost,default" with "--agent-identifiers=ipv4=$(HOST_IP)"
			// once the upstream issue is resolved: https://github.com/kubernetes-sigs/apiserver-network-proxy/issues/261
			"--proxy-strategies=default",
			fmt.Sprintf("--keepalive-time=%s", data.GetKonnectivityKeepAliveTime()),
		},
		Ports: []corev1.ContainerPort{
			{
				Name:          "agentport",
				ContainerPort: 8132,
				Protocol:      corev1.ProtocolTCP,
			},
			{
				Name:          "adminport",
				ContainerPort: 8133,
				Protocol:      corev1.ProtocolTCP,
			},
			{
				Name:          "healthport",
				ContainerPort: 8134,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      resources.KonnectivityUDS,
				MountPath: "/etc/kubernetes/konnectivity-server",
			},
			{
				Name:      resources.KonnectivityKubeconfigSecretName,
				ReadOnly:  true,
				MountPath: "/etc/kubernetes/kubeconfig",
			},
			{
				Name:      resources.KonnectivityProxyTLSSecretName,
				ReadOnly:  true,
				MountPath: "/etc/kubernetes/pki/",
			},
		},
		LivenessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				Exec: nil,
				HTTPGet: &corev1.HTTPGetAction{
					Path:   "/healthz",
					Port:   intstr.IntOrString{IntVal: 8134},
					Scheme: corev1.URISchemeHTTP,
				},
				TCPSocket: nil,
			},
			InitialDelaySeconds: 15,
			TimeoutSeconds:      15,
			PeriodSeconds:       10,
			SuccessThreshold:    1,
			FailureThreshold:    3,
		},
		Resources: defResourceRequirements,
	}, nil
}
