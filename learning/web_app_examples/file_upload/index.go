package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("url_long", r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}

	fmt.Fprintf(w, "Hello Nazmul!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		token := generateToken()
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		errmsg := ""
		r.ParseForm()
		username := template.HTMLEscapeString(r.Form.Get("username"))
		password := template.HTMLEscapeString(r.Form.Get("password"))
		token := template.HTMLEscapeString(r.Form.Get("token"))

		if len(username) == 0 {
			errmsg += "username is a required field \n"
		}

		if len(password) == 0 {
			errmsg += "password is a required field \n"
		}

		if token != "" {

		} else {
			errmsg += "token not found \n"
		}

		// logic part of log in
		fmt.Println("username:", username)
		fmt.Println("password:", password)

		fmt.Fprintf(w, errmsg+"username: "+username+" \npassword: "+password)
	}
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		token := generateToken()
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		maxMemory := int64(32 << 20) // 32 MB
		r.ParseMultipartForm(maxMemory)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
	}
}

func generateToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
