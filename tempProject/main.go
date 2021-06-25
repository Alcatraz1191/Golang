package main

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/datalake/store/2016-11-01/filesystem"
)


func main(){
	//n := filesystem.New()

	//file, _ := os.Open("file")
	//n.Create(context.TODO(),"mystorage1", "adl://mystorage1.azuredatalakestore.net/", file)
	v := filesystem.NewClient()
	fmt.Println(v.SendAzureGetRequest())
}