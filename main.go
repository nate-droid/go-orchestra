package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nate-droid/go-orchestra/conductor"
	"github.com/nate-droid/go-orchestra/manager"
	"github.com/nate-droid/go-orchestra/musician"
	"github.com/nate-droid/go-orchestra/server"
	"net/http"
	"os"
	"time"
)

//go:embed client/build
var res embed.FS

func main() {
	// TODO create the echo client here and just register the right routes
	err := RunService()
	if err != nil {
		fmt.Printf("failed to run the service with error: %s", err)
		return
	}

}

func RunService() error {
	time.Sleep(time.Second * 5) // TODO add contition to wait for nats or retry loop
	defaultService := "clientz"
	serviceType := getenv("SERVICE_TYPE", defaultService)
	fmt.Printf("starting service of type: %s\n", serviceType)
	switch serviceType {
	case "conductor":
		fmt.Println("conductor started")
		c, err := conductor.NewConductor()
		if err != nil {
			return err
		}
		fmt.Println("created conductor")
		err = c.Run(context.Background())
		if err != nil {
			return err
		}

		return nil
	case "manager":
		fmt.Println("Creating manager service")
		m, err := manager.NewManager()
		if err != nil {
			return err
		}
		err = m.Run(context.Background())
		if err != nil {
			return err
		}
	case "musician":
		m, err := musician.NewMusician()
		if err != nil {
			return err
		}

		err = m.Run(context.Background())
		if err != nil {
			return err
		}
		fmt.Println("Creating musician service")
	case "server":
		fmt.Println("creating an API server")
		s := &server.Server{}
		err := s.Run()
		if err != nil {
			return err
		}
	case "client":
		e := echo.New()

		e.Static("/", "/go/client/build")
		// fs := http.FileServer(http.Dir("./client/build"))
		// e.GET("/main", echo.WrapHandler(http.StripPrefix("/assets/", fs)))

		e.GET("/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World! :)")
		})
		err := e.Start(":8080")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown service type: %s", serviceType)
	}

	return nil
}

func getenv(key string, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}
