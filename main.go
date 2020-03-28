package main

import (
	"ke-annotation-route/ke"
)

func main() {
	//flag.PrintDefaults()

	ke.Write(ke.GetConfig("route_file"))

	if ke.GetConfig("is_watch") == "1" {
		ke.Start(ke.GetConfig("root_path") + ke.GetConfig("app_path"))
	} else {
		println("route build complete!\r\n")
	}
}
