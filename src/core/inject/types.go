package inject

/*
Inject Default abstract controller implementation.

	Simply integrate the default controller into your structure.

Example:

	type YourController struct {
	   mvc.Controller
	}

	// Hello
	// @GET(path="/hello") this is api method
	func (y *YourController) Hello(ctx *gin.Context) {
	   resp.Json(ctx, "Hello World")
	}

	// Access the API
	curl http://localhost:4006/hello
*/
type Inject struct{}

func (c *Inject) PostConstruct() {}

// controller Top-level interface used to declare a structure as a controller.
type inject interface {
	// PostConstruct Triggered after dependency injection is completed. You can continue to decorate the controller here.
	PostConstruct()
}

// MethodInfo Api method info
type MethodInfo struct {
	ApiMethodName  string // API method。such as: POST、GET、DELETE、PUT、OPTIONS、PATCH、HEAD
	ApiPath        string // API path
	PackPath       string // API pack
	PackName       string
	PackMethodName string            //
	Annotations    map[string]string // Annotations of the method
}
