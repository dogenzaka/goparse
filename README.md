goparse
====


[![Build Status](http://img.shields.io/travis/dogenzaka/goparse.svg?style=flat)](https://travis-ci.org/dogenzaka/goparse)
[![Coverage](http://img.shields.io/codecov/c/github/dogenzaka/goparse.svg?style=flat)](https://codecov.io/github/dogenzaka/goparse)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/dogenzaka/rotator/blob/master/LICENSE)

Goparse is a client library to access Parse REST API.

Installation
----

```
go get github.com/dogenzaka/goparse
```

Quick start
----

To create a client,

```

import(
  "github.com/dogenzaka/goparse"
)

func main() {

  c := goparse.ParseConfig{}
  client, err := goparse.New(c)
  client.ApplicationId = "PARSE_APPLICATION_ID"
  client.RESTAPIKey = "PARSE REST API KEY"
  client.Url = "PARSE ENDPOINT URL"
  ..

}
```

License
----
Goparse is licensed under MIT.
