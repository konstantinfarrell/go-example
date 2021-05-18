package main

import (
	"flag"

	"github.com/konstantinfarrell/go-example/pkg/api"
	"github.com/konstantinfarrell/go-example/pkg/util/config"
)


func main(){
	path := flag.String("c", "./cmd/go-example/config.yaml.example", "Path to config file")
	flag.Parse()

	conf, err := config.Load(*path)
	checkErr(err)

	checkErr(api.Start(conf))
}


func checkErr(err error){
	if err != nil {
		panic(err.Error())
	}
}