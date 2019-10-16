package main

import (
	"clean-example/wrapper"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func main() {
	// do mongo connection or other driver here
	var mgo *mgo.Database

	// init gin
	r := gin.Default()
	wrapper.RunClean(mgo, r)
	r.Run()
}
