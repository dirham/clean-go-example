package wrapper

import (
	"clean-example/delivery/http"
	"clean-example/repositories/clean"
	usecase "clean-example/usecases/clean"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

// RunNews to connect main with news
func RunClean(mgo *mgo.Database, r *gin.Engine) {
	// init clean repository
	repoNews := clean.NewCleanRepository(mgo)
	// init clean usecase
	uscClean := usecase.NewMgoUsecase(repoNews)

	// call the handler
	http.NewCleanHTTPHandler("1", r, *uscClean)
}
