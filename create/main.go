package main

import (
	"github.com/rootxrishabh/PasteContainer/create/router"
	"github.com/rootxrishabh/PasteContainer/pkg/common/utils"
	_ "github.com/lib/pq"
)

func main(){
	db, err := utils.InitDBConnection()
	if err != nil {
		panic(err)
	}
	router := router.NewRouter(db)
	router.Run()
}