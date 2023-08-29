package main

import (
    "net/http"
    "log"
	"html/template"

    "github.com/julienschmidt/httprouter"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

const maxUploadSize = 1 * 1024 * 1024 // 10 mb
const uploadPath = "./uploads"
var userpicPath = "./userpic"

func main() {
    router := httprouter.New()
    router.GET("/", HasAuthCookie(Index))

	router.GET("/login", loginHandler)
    router.POST("/login", loginHandler)
    router.GET("/register", registerHandler)
    router.POST("/register", registerHandler)
    router.POST("/logout", HasAuthCookie(logoutHandler))

	
	router.GET("/fav/favicon.ico", Ignore)
	
	static := httprouter.New()
	static.ServeFiles("/userpic/*filepath", http.Dir(userpicPath))
	router.ServeFiles("/static/*filepath", http.Dir("static"))
//	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	router.NotFound = static


	log.Println("Starting Server")
    log.Fatal(http.ListenAndServe("localhost:4000", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl.ExecuteTemplate(w,"head.html",nil)
	tmpl.ExecuteTemplate(w,"index.html",nil)
}

func Ignore(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "favicon.png")
}
