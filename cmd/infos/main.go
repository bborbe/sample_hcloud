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
		pricing, _, err := client.Pricing.Get(ctx)
		if err != nil {
			glog.Exit(err)
		}
		fmt.Println("Pricing")
		fmt.Printf("%+v\n", pricing)
	}

	{
		isos, err := client.ISO.All(ctx)
		if err != nil {
			glog.Exit(err)
		}
		fmt.Println("ISOs:")
		for _, iso := range isos {
			fmt.Printf("%d %s\n", iso.ID, iso.Name)
		}
	}


	{
		actions, err := client.Action.All(ctx)
		if err != nil {
			glog.Exit(err)
		}
		fmt.Println("Actions:")
		for _, action := range actions {
			fmt.Printf("%d %s\n", action.ID, action.Command)
		}
	}


}
