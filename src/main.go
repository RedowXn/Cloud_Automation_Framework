package main
import (
    "context"
    "fmt"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil { panic(err) }
    containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
    if err != nil { panic(err) }
    for _, c := range containers {
        fmt.Println(c.ID, c.Image)
    }
}
