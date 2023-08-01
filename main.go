package main

import (
	"github.com/Oats87/docker-machine-driver-dop/dop"
	"github.com/rancher/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(dop.NewDriver("", ""))
}
