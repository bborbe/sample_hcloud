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
