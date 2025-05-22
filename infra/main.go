package main

import (
    docker "github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
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
                    Internal: pulumi.Int(8080),
                    External: pulumi.Int(8080),
                },
            },
        })
        if err != nil {
            return err
        }

        ctx.Export("endpoint", pulumi.String("http://localhost:8080"))
        return nil
    })
}

