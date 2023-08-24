package main

import (
	_ "github.com/lib/pq"
	"github.com/rootxrishabh/PasteContainer/delete/router"
	"github.com/rootxrishabh/PasteContainer/pkg/common/utils"
)

func main(){
	db, err := utils.InitDBConnection()
	if err != nil {
		panic(err)
	}
	router := router.NewRouter(db)
	router.Run()
}