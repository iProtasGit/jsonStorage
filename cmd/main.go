package main

import (
	"jsonStorage/config"
    "jsonStorage/internal/controller"
    "jsonStorage/internal/adapter"
)


func main() {
    cfg, err := config.LoadConfig("config/config.json")
    if err != nil {
        panic(err)
    }

    storage, err := adapter.GetStorage(cfg.StorageName)
    if err != nil {
        panic(err)
    }

    controller.StartCLI(cfg, storage)
}
