package main

import (
	"github.com/belgaied2/docker-machine-driver-ovhcloud/driver/ovhcloudpubliccloud"
	"github.com/rancher/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(ovhcloudpubliccloud.NewDriver("", ""))
}
