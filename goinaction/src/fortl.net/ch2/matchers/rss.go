package matchers

import (
	"encoding/xml"
	"fortl.net/ch2/search"
	"errors"
	"net/http"
	"fmt"
	"log"
	"regexp"
)

type (
	// item 根据 item 字段的标签，将定义的字段
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image 根据 image 字段的标签，将定义的字段
	// 与 rss 文档的字段关联起来
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel 根据 channel 字段的标签，将定义的字段
	// 与 rss 文档的字段关联起来
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument 定义了与 rss 文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}


)


type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}
// retrieve 发送 HTTP Get 请求获取 rss 数据源并解码
func (m rssMatcher) retrieve(feed *search.Feed)(*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	// 从网络获得 rss 数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// 检查状态码是不是 200，这样就能知道是不是收到了正确的响应
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err




}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]* search.Result, error) {

	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n",
		feed.Type, feed.Name, feed.URI)
	document, err := m.retrieve(feed)

	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {

		// 检查标题部分是否包含搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err!= nil {
			return nil, err
		}

		// 如果找到匹配的项，将其作为结果保存
		if matched {
			results = append(results, &search.Result{Field:"Title",Content:channelItem.Title})
		}

		// 检查描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)

		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{Field:"Description", Content:channelItem.Description})
		}
	}

	return results, nil
}
