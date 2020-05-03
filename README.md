# cat-facts
[Echo](https://echo.labstack.com/) middleware that to add cat facts in headers of HTTP responses. Inspired by [cat_facts](https://github.com/fabrik42/cat_facts).
```
Response Headers:
X-Cat-Facts: A cat has 32 muscles in each ear.
```

## Example
```
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	catfacts "github.com/regentag/cat-facts"
)

func main() {
	e := echo.New()

	e.Use(catfacts.CatFactsMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

```

## Usage
Any URL will have a `X-Cat-Fact` header in the response, that will contain a random cat fact for your amusement.

