# gzip

Package [gzip](https://godoc.org/github.com/zengabor/gzip) provides a gzip compressing handler.

(Original package which includes clever MIME type checking: [compress](https://godoc.org/github.com/gowww/compress). It takes care to not handle small contents, or contents that are already compressed (like JPEG, MPEG or PDF). Trying to gzip them not only wastes CPU but can potentially increase the response size.)

## Installing

1. Get package:

	```Shell
	go get -u github.com/zengabor/gzip
	```

2. Import it in your code:

	```Go
	import "github.com/zengabor/gzip"
	```

## Usage

To wrap an [http.Handler](https://golang.org/pkg/net/http/#Handler), use [Handle](https://godoc.org/github.com/zengabor/gzip#Handle):

```Go
mux := http.NewServeMux()

mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
})

http.ListenAndServe(":8080", gzip.Handle(handler))
```

To wrap an [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc), use [HandleFunc](https://godoc.org/github.com/zengabor/gzip#HandleFunc):

```Go
http.Handle("/", gzip.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}))

http.ListenAndServe(":8080", nil)
```

All in all, make sure to include this handler above any other handler that may write the response.
