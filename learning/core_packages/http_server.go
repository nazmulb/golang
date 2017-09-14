package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type", "text/html",
	)

	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>Hello Nazmul</title>
	</head>
	<body>
		Hello Nazmul!
	</body>
</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9089", nil)
}
