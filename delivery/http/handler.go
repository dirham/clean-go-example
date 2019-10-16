package http

import (
	"clean-example/usecases/clean"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewCleanHTTPHandler represent the httphandler for clean
type CleanHTTPHandler struct {
	MUsecase clean.MgoUsecase
}

// NewCleanHTTPHandler for now just returning struct from usecase
func NewCleanHTTPHandler(version string, e *gin.Engine, nu clean.MgoUsecase) {
	handler := &CleanHTTPHandler{
		MUsecase: nu,
	}

	//Setting CORS
	e.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v := e.Group("api/v" + version)
	v.GET("/clean/:id", handler.FindByID)
}

// FetchAll http://{host}/api/{version}/clean/{id}
func (n *CleanHTTPHandler) FindByID(c *gin.Context) {
	id, erConv := strconv.Atoi(c.Params.ByName("id"))

	if erConv != nil {
		c.JSON(http.StatusBadRequest, "error: hmm something wrong")
		return
	}

	// be sure to handle the error,  DO NOT USING _
	object, _ := n.MUsecase.FindByID(id)
	c.JSON(http.StatusAccepted, object)
}
