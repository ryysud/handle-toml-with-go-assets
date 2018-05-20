package main

import (
	"fmt"
	"reflect"

	"github.com/ryysud/handle-toml-with-go-assets/config"
	"github.com/ryysud/handle-toml-with-go-assets/models"

	"github.com/BurntSushi/toml"
)

//go:generate go-assets-builder -p config -o config/assets.go -s /config config/config.toml

func main() {
	f, err := config.Assets.Open("/config.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var config models.Config
	_, err = toml.DecodeReader(f, &config)
	if err != nil {
		panic(err)
	}

	printWithType(config.Sample.String)
	printWithType(config.Sample.Int)
	printWithType(config.Sample.Float)
	printWithType(config.Sample.Bool)
	printWithType(config.Sample.Array)
}

func printWithType(x interface{}) {
	fmt.Printf("type: %s, value: %#v\n", reflect.TypeOf(x), x)
}
