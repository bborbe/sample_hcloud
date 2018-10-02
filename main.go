package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/hetznercloud/hcloud-go/hcloud"
)

func main() {
	tokenPtr := flag.String("token", "", "token")
	flag.Parse()

	client := hcloud.NewClient(hcloud.WithToken(*tokenPtr))

	ctx := context.Background()
	{
		datacenters, err := client.Datacenter.All(ctx)
		if err != nil {
			glog.Exit(err)
		}
		fmt.Println("Datacenters:")
		for _, datacenter := range datacenters {
			fmt.Printf("%d %s %s\n", datacenter.ID, datacenter.Name, datacenter.Description)
			for _, s := range datacenter.ServerTypes.Available {
				fmt.Printf("%+v\n", s)
			}
		}
	}

	{
		servers, err := client.Server.All(ctx)
		if err != nil {
			glog.Exit(err)
		}
		fmt.Println("Servers:")
		for _, server := range servers {
			fmt.Printf("%d %s\n", server.ID, server.Name)
		}
	}

	{

		images, _, err := client.Image.List(ctx, hcloud.ImageListOpts{})
		if err != nil {
			glog.Exit(err)
		}
		fmt.Println("Images:")
		for _, image := range images {
			fmt.Printf("%d %s\n", image.ID, image.Name)
		}
	}

	{

		serverTypes, _, err := client.ServerType.List(ctx, hcloud.ServerTypeListOpts{})
		if err != nil {
			glog.Exit(err)
		}
		fmt.Println("ServerTypes:")
		for _, serverType := range serverTypes {
			fmt.Printf("%d %s\n", serverType.ID, serverType.Name)
		}
	}

	{
		datacenter, _, err := client.Datacenter.GetByID(ctx, 2)
		if err != nil {
			glog.Exit(err)
		}

		serverType, _, err := client.ServerType.GetByName(ctx, "cx11")
		if err != nil {
			glog.Exit(err)
		}

		image, _, err := client.Image.GetByName(ctx, "ubuntu-18.04")
		if err != nil {
			glog.Exit(err)
		}

		sshKey, _, err := client.SSHKey.GetByName(ctx, "test")
		if err != nil {
			glog.Exit(err)
		}

		fmt.Println("create server")
		start := true
		_, _, err = client.Server.Create(ctx, hcloud.ServerCreateOpts{
			Name:       "test",
			ServerType: serverType,
			Image:      image,
			SSHKeys:    []*hcloud.SSHKey{sshKey},
			Location:   datacenter.Location,
			//Datacenter:       datacenter,
			UserData:         "#cloud-config\nruncmd:\n- [touch, /root/cloud-init-worked]\n",
			StartAfterCreate: &start,
			Labels:           map[string]string{"myserver": "test"},
		})
		if err != nil {
			glog.Exit(err)
		}
	}
}
