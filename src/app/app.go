package main

import (
	"code.google.com/p/gcfg"
	"domains"
	"flag"
	"fmt"
	"github.com/DDRBoxman/go-amazon-product-api"
)

const APP_VERSION = "0.1"

var config domains.Config

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func init() {

	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {

		return

	}

}

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	var api amazonproduct.AmazonProductAPI

	api.AccessKey = config.Aws.AccessKey
	api.SecretKey = config.Aws.SecretKey
	api.Host = config.Aws.Host
	api.AssociateTag = config.Aws.AssociateTag

	result, err := api.ItemSearchByKeyword("sgt+frog", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
