package main

import (
	"log"
	"net/http"

	api "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1"
	//"github.com/golang/build/kubernetes"
	"golang.org/x/build/kubernetes"
	//"golang.org/x/oauth2"
)

func main() {
	kube, err := kubernetes.NewClient("http://127.0.0.1:8080", &http.Client{
	/*Transport: &oauth2.Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "aCcessWbU3toKen"}),
	}*/})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	kube.Run(&api.Pod{
		TypeMeta: api.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: api.ObjectMeta{
			Name: "hello-world",
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name:    "hello",
					Image:   "ubuntu:14.04",
					Command: []string{"/bin/echo", "hello", "world"},
				},
			},
		},
	})
}
