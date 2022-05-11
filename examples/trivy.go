package main

import (
	"fmt"
	"log"

	"github.com/aquasecurity/trivy-kubernetes/pkg/artifacts"
	"github.com/aquasecurity/trivy-kubernetes/pkg/k8s"
	"github.com/aquasecurity/trivy-kubernetes/pkg/trivyk8s"

	"context"
)

func main() {
	ctx := context.Background()

	cluster, err := k8s.GetCluster()
	if err != nil {
		log.Fatal(err)
	}

	trivyk8s := trivyk8s.New(cluster)

	fmt.Println("Scaning cluster")
	//trivy k8s #cluster
	artifacts, err := trivyk8s.ListArtifacts(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printArtifacts(artifacts)

	fmt.Println("Scaning namespace 'default'")
	//trivy k8s --namespace default
	artifacts, err = trivyk8s.Namespace("default").ListArtifacts(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printArtifacts(artifacts)
}

func printArtifacts(artifacts []*artifacts.Artifact) {
	for _, artifact := range artifacts {
		fmt.Printf(
			"Name: %s, Kind: %s, Namespace: %s, Images: %v\n",
			artifact.Name,
			artifact.Kind,
			artifact.Namespace,
			artifact.Images,
		)
	}

}
