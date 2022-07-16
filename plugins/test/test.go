package test

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("go", "https://go.dev")
	s.UrlProcessor.OnHTML(".Hero-blurb", func(element *colly.HTMLElement) {
		s.AddUrl("https://go.dev", time.Now())
	})
	s.UrlProcessor.OnHTML(".Hero-gopherLadder", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("alt"))
	})
	s.UrlProcessor.OnHTML(".Hero-blurbList", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("content", element.Text)
	})
}
