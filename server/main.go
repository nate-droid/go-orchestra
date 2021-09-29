package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/nate-droid/core/symphony"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Results struct {
	ResultsJson []symphony.Song `json:"results"`
}

func main() {
	errs, ctx := errgroup.WithContext(context.Background())
	store, err := newStore()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/symphonies", func(c echo.Context) error {
		res, err := store.FetchLatestSymphonies()
		var x []symphony.Song
		for _, song := range res {
			x = append(x, song[0])
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, "oops")
		}
		return c.JSON(http.StatusOK, Results{x})
	})

	errs.Go(func() error {
		return store.run(ctx)
	})
	errs.Go(func() error {
		return e.Start("0.0.0.0:1323")
	})

	err = errs.Wait()
	if err != nil {
		panic(err)
	}
}

