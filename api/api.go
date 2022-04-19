package api

import (
	"survayData/api/modules/surveyInformation"
	"survayData/middleware"

	"github.com/labstack/echo"
)

// Init api binding
func Init(e *echo.Echo) {
	o := e.Group("/o")
	r := e.Group("/r")  //restricted group
	c := r.Group("/c")
	ec := r.Group("/ec")

	middleware.Init(e, o, r, c, ec)
	
	surveyInformation.Init(o)

}