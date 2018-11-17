Program: Gorilla Mux Path Prefix

How to use?
- *router.PathPrefix("/Article"): return a route containing the path with prefix "/Article". For example, path of "/Article/abc/xyz/123/" contains the prefix "/Article"
- r.PathPrefix("/Article").HandlerFunc(articleHandler): the articleHandler function will be called if the incoming request contains path prefix "/Article"