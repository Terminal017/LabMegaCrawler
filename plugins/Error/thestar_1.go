package dev

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Article struct {
	StoryUrl string `json:"StoryUrl"`
}

type Response []Article

// 等待增加新闻数据数量
func FetchAndVisitArticles(engine *crawlers.WebsiteEngine, page int64) {
	url := fmt.Sprintf("https://cdn.thestar.com.my/Components/JustIn/JustIn-news.json?%d", page)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v\n", err)
		return
	}

	var jsonResp Response
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		log.Printf("解析 JSON 失败: %v\n", err)
		return
	}

	for _, item := range jsonResp {
		fullURL := item.StoryUrl
		if !strings.HasPrefix(fullURL, "http") {
			fullURL = "https://www.posttoday.com" + fullURL
		}
		engine.Visit(fullURL, crawlers.News)
	}
}

// AJAX请求采集
func init() {
	engine := crawlers.Register("N-0043", "Post Today", "https://www.thestar.com.my")

	engine.SetStartingURLs([]string{"https://www.thestar.com.my/news"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnLaunch(func() {
		for i := 0; i < 1000; i++ {
			timestamp := time.Now().UnixMilli() - int64(i*1000)
			FetchAndVisitArticles(engine, timestamp)
		}
	})

	engine.OnHTML("#story-body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
