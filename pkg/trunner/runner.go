package trunner

import (
	"fmt"
	"os"
	"sync"

	c "github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/k0kubun/go-ansi"
	"github.com/owasp-offat/offat/pkg/tgen"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/term"
)

// Runs API Tests
func RunApiTests(t *tgen.TGenHandler, client c.ClientInterface, apiTests []*tgen.ApiTest) {
	var wg sync.WaitGroup

	// Get the terminal size
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default width
	}

	// Adjust the progress bar width based on terminal size
	barWidth := width - 40 // Subtract 40 to account for other UI elements

	bar := progressbar.NewOptions(
		len(apiTests),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(barWidth),
		progressbar.OptionSetTheme(
			progressbar.Theme{
				Saucer:        "[green]█[reset]",
				SaucerHead:    "[green]▓[reset]",
				SaucerPadding: "░",
				BarStart:      "╢|",
				BarEnd:        "|╟",
			},
		),
	)

	for _, apiTest := range apiTests {
		wg.Add(1)
		go func(apiTest *tgen.ApiTest) {
			defer wg.Done()
			resp, err := client.Do(apiTest.Request.Uri, apiTest.Request.Method, apiTest.Request.QueryParams, apiTest.Request.Headers, apiTest.Request.Body)
			apiTest.Response = c.NewConcurrentResponse(resp, err)

			if err := bar.Add(1); err != nil {
				log.Error().Err(err).Msg("Failed to add to bar")
			}
		}(apiTest)
	}

	wg.Wait()
	fmt.Println()

}
