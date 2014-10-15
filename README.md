goparse
====


[![Build Status](http://img.shields.io/travis/dogenzaka/goparse.svg?style=flat)](https://travis-ci.org/dogenzaka/goparse)
[![Coverage](http://img.shields.io/codecov/c/github/dogenzaka/goparse.svg?style=flat)](https://codecov.io/github/dogenzaka/goparse)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/dogenzaka/rotator/blob/master/LICENSE)

Goparse is a client library to access [Parse REST API](https://parse.com/docs/rest).

Installation
----

```
go get github.com/dogenzaka/goparse
```

Quick start
----

To create a session from default client,

```go
import(
  "github.com/dogenzaka/goparse"
)

func main() {

  parseSession := goparse.NewSession("SESSION_TOKEN")
  me, err := parseSession.GetMe()
  ..

}
```

Custom Client
----

To create a configured` client,

```go
parseClient := goparse.NewClient(goparse.Config({
  ApplicationId: "PARSE_APPLICATION_ID",
  RestAPIKey: "PARSE_REST_API_KEY",
  MasterKey: "PARSE_MASTER_KEY",
  EndPointURL: "PARSE_ENDPOINT_URL"
})

parseSession := parseClient.NewSession()
me, err := parseSession.GetMe()
..
```

Environment variables
----

The default client uses environment variables to access Parse REST API.

- `PARSE_APPLICATION_ID`
- `PARSE_REST_API_KEY`
- `PARSE_MASTER_KEY`
- `PARSE_ENDPOINT_URL`

License
----
Goparse is licensed under the MIT.
