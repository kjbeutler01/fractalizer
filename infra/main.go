package main

import (
	"fmt"

	docker "github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		port := cfg.GetInt("port")
		if port == 0 {
			port = 8080
		}

		image, err := docker.NewImage(ctx, "fractalizer-image", &docker.ImageArgs{
			Build: docker.DockerBuildArgs{
				Context: pulumi.String("../go-fractalizer"),
			},
			ImageName: pulumi.String("fractalizer:latest"),
		})
		if err != nil {
			return err
		}

		_, err = docker.NewContainer(ctx, "fractalizer-container", &docker.ContainerArgs{
			Image: image.ImageName,
			Ports: docker.ContainerPortArray{
				docker.ContainerPortArgs{
					Internal: pulumi.Int(port),
					External: pulumi.Int(port),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("endpoint", pulumi.String(fmt.Sprintf("http://localhost:%d", port)))
		return nil
	})
}
