package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nate-droid/go-orchestra/core/symphony"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Results struct {
	ResultsJson map[string]symphony.Song `json:"results"`
}

type Server struct {
}

// kubectl port-forward deployment/hello-music 1323:1323

func (s *Server) Run() error {
	errs, ctx := errgroup.WithContext(context.Background())
	// store, err := newStore()
	store, err := newMem()
	if err != nil {
		fmt.Println("failed to create a store")
		return err
	}
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/symphonies", func(c echo.Context) error {

		songs, err := store.FetchAllSongs()
		if err != nil {
			return c.String(http.StatusInternalServerError, "oops")
		}
		return c.JSON(http.StatusOK, Results{songs})
	})

	errs.Go(func() error {
		return store.run(ctx)
	})

	errs.Go(func() error {
		return e.Start("0.0.0.0:1323")
	})

	err = errs.Wait()
	if err != nil {
		return err
	}

	return nil
}
