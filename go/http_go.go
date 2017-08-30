package main

import ("net/http"; "io"
	"fmt"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, `
<doctype html>
<html>
	<head>
		<title>Hello World</title>
	</head>
	<body>
		<p>Hello World!</p>
	</body>
</html>
	`)
}

type MyServer struct{


}

func (*MyServer)ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Request:", req)
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, `
<doctype html>
<html>
	<head>
		<title>Index</title>
	</head>
	<body>
		<a href="/hello">Hello World</a>
	</body>
</html>
	`)

}


func main()  {
	http.HandleFunc("/hello", hello)
	server := new(MyServer)

	http.ListenAndServe("localhost:9000", server)
}
