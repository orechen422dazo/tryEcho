# tryEcho

[Quick Start](https://echo.labstack.com/docs/quick-start)

[Github](https://github.com/orechen422dazo/tryEcho)

```shell
go get github.com/labstack/echo/v4
```

example:
```go
package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

Start server:
```shell
go run main.go
```