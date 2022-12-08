package africapi

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("africapi", "Africapi", "https://africapi.org")
	w.SetStartingUrls([]string{"https://www.africapi.org/africa-reseach-notes"})

	w.OnHTML(".bg-size-cover > .column-content-inner > .gsc-image-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.Title = element.ChildText(".title")
		subCtx.Content = element.ChildText(".desc")
		subCtx.File = append(subCtx.File, element.ChildAttr(".action > a", "href"))
		subCtx.PageType = Crawler.Report
	})
}