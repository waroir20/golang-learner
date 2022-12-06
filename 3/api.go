package api

import "github.com/gin-gonic/gin"

type Aggregate struct {
	Id   string `json:"id" validate:"uuid"`
	Name string `json:"name" validate:"-"`
	Data string `json:"data" validate:"-"`
}

// RegisterRoutes - registers the API paths and rest methods used to export functionality of the app to the outside world
func RegisterRoutes(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/samples")
	_ = routerGroup // This is just here to allow the file to compile by mocking that data is used
	// TODO add routes to routerGroup below

}

// TODO create a function which returns a 200 with the phrase 'Where no Lopporit has gone before!'; accessible via GET /api/v1/samples
//  See TestGet

// TODO create a function which returns a 201 with the original request payload as the response; accessible via POST /api/v1/samples
//  run validation on the incoming data against the Aggregate struct and returns a 400 for failed validation
//  See TestPost

// TODO create a function which returns a 202 with the payload built from different parts; accessible via PUT /api/v1/samples/<id> where <id> is a path parameter
//  extract the Name and Data fields from the headers of the request with keys 'name' and 'data' respectively
//  Special case just for learning... When the path param is empty return 500
//   See TestPut

// TODO create a function which returns a 418 (it really is the best HTTP status code); accessible via DELETE /api/v1/samples/<id> where <id> is a path parameter
//  Special case just for learning... When the path param has the value 'Livingway' return 422 (because who could possible fathom that anyone would want to delete the best Lopporit ever)
//   See TestDelete
