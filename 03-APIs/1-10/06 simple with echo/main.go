


/*  Link Site :

https://echo.labstack.com/docs/quick-start


*/

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

/* First : in Terminal 
	go mod tidy

Output:
	go: finding module for package github.com/labstack/echo/v4
	go: downloading github.com/labstack/echo/v4 v4.11.2
	go: downloading github.com/labstack/echo v3.3.10+incompatible
	go: found github.com/labstack/echo/v4 in github.com/labstack/echo/v4 v4.11.2
	go: downloading github.com/labstack/gommon v0.4.0
	go: downloading golang.org/x/crypto v0.14.0
	go: downloading golang.org/x/net v0.17.0
	go: downloading github.com/stretchr/testify v1.8.4
	go: downloading github.com/valyala/fasttemplate v1.2.2
	go: downloading github.com/davecgh/go-spew v1.1.1
	go: downloading github.com/pmezard/go-difflib v1.0.0
	go: downloading gopkg.in/yaml.v3 v3.0.1
	go: downloading golang.org/x/sys v0.13.0
	go: downloading golang.org/x/text v0.13.0

============================================================================================

Second in Terminal:
	go run .

Output:
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323

*/