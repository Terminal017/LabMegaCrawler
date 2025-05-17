package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-003", "Cloudera Blog", "https://Blog.cloudera.com")

	engine.SetStartingURLs([]string{"https://www.cloudera.com/about/news-and-blogs.html"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".col-sm-4.col p a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".medium > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".col-sm-8.col.col_border_right > p, .col-sm-8.col.col_border_right > ul > li,.itemBody p,.non-paywall > p ,.inner-text-div  > p,.inner-text-div  > h4,.article-data p,.article-data h2,.inner-text-div  > h4", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

}
