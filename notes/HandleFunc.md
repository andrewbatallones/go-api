## HandleFunc

This is setup where a trailing slash will tell the handler to redirect everything to that base handler. This is fine if we need to account for a lot of functions, but to surpress it, we can include this snippet:

```go
// path is replaced with whatever path you need it to check
if r.URL.Path != path {
    http.NotFound(w, r)
    return
}
```