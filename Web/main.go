package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"

	"github.com/shorty-io/go-shorty/Shared/service"

	"github.com/shorty-io/go-shorty/web/handler"
	"github.com/shorty-io/go-shorty/web/handler/auth"
	"github.com/shorty-io/go-shorty/web/handler/flipflop"
	"github.com/shorty-io/go-shorty/web/handler/queries"
)

func main() {
	srv := service.New("Web")
	srv.Start()

	natsURL, err := srv.GetKV("NATS_URL")
	if err != nil {
		log.Println("Failed to retrieve NATS_URL from Consul key-value store:", err)
		return
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Println("Failed to Connect to nats:", err)
		return
	}
	defer nc.Flush()
	defer nc.Close()

	e := echo.New()

	e.Use(middleware.CORS())

	queries.New(e, srv.Dial)
	flipflop.New(e, srv.Dial, auth.Middleware)

	e.GET("/cmd/block-content/update/:id", func(c echo.Context) error {
		return handler.UpdateContent(c, nc)
	}, auth.Middleware)

	e.Logger.Fatal(e.Start(":8080"))
}

func maxSubArray(nums []int) int {
	var res, curr int
	for _, v := range nums {
		if curr+v > 0 {
			curr += v
		} else {
			curr = 0
		}
		res = max(curr, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
