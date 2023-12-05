package main

import (
	"github.com/belgaied2/docker-machine-driver-ovhcloud/driver"
	"github.com/rancher/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(driver.NewDriver("", ""))
}
