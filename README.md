[![GoDoc](https://godoc.org/github.com/flexoid/go-utils?status.svg)](https://godoc.org/github.com/flexoid/go-utils)

The collection of helper functions, utils and code snippets that can be useful
for my golang projects.

# Content

## `go-utils/logging`

[logging.RequestLogger](https://godoc.org/github.com/flexoid/go-utils/logging#RequestLogger)

A middleware function that logs request start and request end with zerolog.
A logger instance is taken from the request context.

```
2019-01-13T18:46:34.488192Z |INFO| Request started req_id=bgtof2ml0s1adiaerc70
2019-01-13T18:46:35.035226Z |INFO| Request completed bytes=3326 elapsed_ms=547.028988 http_method=GET remote_addr=[::1]:61052 req_id=bgtof2ml0s1adiaerc70 status_code=200 uri=http://localhost:4000/api/users/123 user_agent=HTTPie/0.9.9
```

```
{"level":"info","req_id":"bgtoflul0s1afn7goou0","time":"2019-01-13T18:47:51.118414Z","message":"Request started"}
{"level":"info","req_id":"bgtoflul0s1afn7goou0","status_code":200,"http_method":"GET","uri":"http://localhost:4000/api/users/123","remote_addr":"[::1]:61079","user_agent":"HTTPie/0.9.9","bytes":3326,"elapsed_ms":477.652781,"time":"2019-01-13T18:47:51.596075Z","message":"Request completed"}
```
