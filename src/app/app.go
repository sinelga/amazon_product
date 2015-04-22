package main

import (
	"flag"
	"fmt"
	"github.com/DDRBoxman/go-amazon-product-api"
	"code.google.com/p/gcfg"
	"domains"
)

const APP_VERSION = "0.1"

var config domains.Config


// The flag package provides a default help printer via -h switch
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

//	api.AccessKey = "AKIAJ7IFGG4TVT3DO5TA"
//	api.SecretKey = "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"
	api.AccessKey = config.Aws.AccessKey
	api.SecretKey = config.Aws.SecretKey
	api.Host = config.Aws.Host
	api.AssociateTag = config.Aws.AssociateTag

	result, err := api.ItemSearchByKeyword("sgt+frog", 2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
