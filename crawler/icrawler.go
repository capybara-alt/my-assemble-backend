package crawler

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/gocolly/colly"
)

type ICrawler interface {
	JobName() string
	Crawl() (model.CrawlResultJSON, error)
}

type AbstractCrawler struct {
	target_url string
	unit_type  string
}

func NewAbstractCrawler(target_url config.CrawlTargetPage, unit_type model.SecondaryUnitType) ICrawler {
	return &AbstractCrawler{
		target_url: string(target_url),
		unit_type:  string(unit_type),
	}
}

func (a *AbstractCrawler) JobName() string {
	return ""
}

func (a *AbstractCrawler) Crawl() (model.CrawlResultJSON, error) {
	collector := colly.NewCollector()
	unit_list := make(model.CrawlResultJSON)
	unit_list[a.unit_type] = make(model.CategoryGroupedJSON)
	collector.OnHTML("div#wikibody", func(root *colly.HTMLElement) {
		root.ForEach("div.plugin_contents > ul > li > a", func(i int, h *colly.HTMLElement) {
			if h.DOM.NextAllFiltered("ul").Length() != 0 {
				unit_list[a.unit_type][h.Text] = make(model.NameGroupedJSON)
			}
			h.DOM.NextAllFiltered("ul").Each(func(i int, h2list *goquery.Selection) {
				h2list.ChildrenFiltered("li").ChildrenFiltered("a").Each(func(i int, h3list *goquery.Selection) {
					id, ok := h3list.Attr("href")
					if !ok {
						return
					}

					h3WithId := fmt.Sprintf("h3%s", id)
					h3 := root.DOM.ChildrenFiltered(h3WithId)
					table := h3.NextAllFiltered("table:not([class=\"atwiki_plugin_region\"])")
					if !a.IsTargetTable(table) {
						return
					}
					unit_list[a.unit_type][h.Text][h3.Text()] = a.GetUnitInfo(table.First())
				})
			})
		})
	})

	err := collector.Visit(a.target_url)
	return unit_list, err
}

func (a *AbstractCrawler) IsTargetSection(s *goquery.Selection) bool {
	return s.Text() != ""
}

func (a *AbstractCrawler) IsTargetTable(s *goquery.Selection) bool {
	filter_table_style := "margin:0;padding:0;border:none;background-color:transparent;"
	filter_target_table := fmt.Sprintf("tr[style=\"%s\"]", filter_table_style)
	return s.ChildrenFiltered("tbody").ChildrenFiltered(filter_target_table).Text() == ""
}

func (a *AbstractCrawler) GetUnitInfo(s *goquery.Selection) model.UnitInfoJSON {
	unit_info := make(model.UnitInfoJSON)
	s.ChildrenFiltered("tbody").ChildrenFiltered("tr").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered("th").First().Text()
		column := s.ChildrenFiltered("td").First().Text()
		unit_info[header] = column
	})

	return unit_info
}
