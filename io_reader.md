# Empty io.Reader/io.ReadCloser

When building a web service, you will often need to access the body of an
incoming web request, typically the result of an HTTP POST. In Go, the type of
the request is usually `http.Request`. The body is accessible from this as an
[`io.ReadCloser`](https://golang.org/pkg/io/#ReadCloser), which groups the
`io.Reader` and `io.Closer` interfaces.

If the incoming request is JSON, it's typically processed with either
`json.NewDecoder().Decode()`:

```go
func processRequest(req *http.Request) (SomeType, error) {
  var t SomeType
  err := json.NewDecoder(req.Body).Decode(&t)
  //...
}
```

Or with `json.Unmarshal()`:
```go
func processRequest(req *http.Request) (SomeType, error) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    return SomeType{}, err
  }
  var t SomeType
  json.Unmarshal(body, &t)
  //...
}
```

For debugging or inspecting payloads, you may want to print out request info. If
you do, you might use some code like this to read and print the body, or return
it as a string:

```go
func printRequest(req http.Request) {
  // Print path, headers, etc.
  // ...
  buf := new(bytes.Buffer)
  buf.ReadFrom(req.Body)
  fmt.Println("Body:", buf.String())
}
```

But since `req.Body` is an `io.ReadCloser`, reading from the body empties it.
There is no way to rewind it since it doesn't support `io.Seeker`. So this
doesn't work for printing out the body and then processing it. You'll get an
empty body, which can cause confusion.

You could use an [io.TeeReader](https://golang.org/pkg/io/#example_TeeReader) to
get around this, and/or process the body after reading it into a `bytes.Buffer`.
But there's an easier way to debug -- [httputil.DumpRequest](https://golang.org/pkg/net/http/httputil/#DumpRequest)

If you pass the `body` boolean to `httputil.DumpRequest`, it will consume the
body and then replace it with a new `io.ReadCloser`, so that you can still
process it. Of course the performance implications sugget that this should be
for debugging purposes only.

```go
func processRequest(req *http.Request) (SomeType, error) {
  reqDump, err := httputil.DumpRequest(req, true)
  if err != nil {
    fmt.Println("Error calling DumpRequest: ", err)
  } else {
    fmt.Println("request:\n", string(reqDump))
  }
  // Process the req.Body as needed
  // ...
}
```

