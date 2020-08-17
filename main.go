package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yklaus/Go-Scrapper/scrapper"
	"log"
	"os"
	"strings"
	"time"
)

const fileName = "jobs.csv"

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func handleHome(c echo.Context) error {
	return c.File("index.html")
}

func handleScrape(c echo.Context) error {
	defer func() {
		time.Sleep(time.Second * 2)
		err := os.Remove(fileName)
		if err != nil {
			log.Fatalln(err)
		}
	}()
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}
