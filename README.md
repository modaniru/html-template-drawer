# html template drawer
## env
~~~bash
PORT=80 #port
~~~
## create new endpoints
in internal/controller/router.go you can add your html pages in GetRouter() method
~~~go
func (r *Router) GetRouter() *gin.Engine {
	// load html files
	r.router.LoadHTMLGlob("template/**/*")
	// load static files
	r.router.Static("/static", "static")
	// log middleware
	r.router.Use(middleware.JsonLoggerMiddleware())
	// routing
	r.router.GET("/first", r.LoadHtmlPage("first.html"))
	r.router.GET("/second", r.LoadHtmlPage("second.html"))
	// located in template/third/third.html but we must write just a file name
	r.router.GET("/third", r.LoadHtmlPage("third.html"))
	return r.router
}
~~~

## local run
~~~bash
git clone https://github.com/modaniru/html-template-drawer
cd html-template-drawer
make #or go run cmd/main.go
~~~
## docker
~~~bash
docker build -t image-name .
docker run -p 80:80 image-name
~~~
## file structure
* internal - go files
* cmd - contains main.go file
* static - contains static files (images, css styles, js files, fonts, etc.)
* template - contains folders which contains html files