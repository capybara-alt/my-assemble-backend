package crawler

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/gocolly/colly"
)

type expansion struct{}

func NewExpansion() repository.ExternalExpansion {
	return &expansion{}
}

func (a *expansion) Fetch() (model.CrawlResultJSON, error) {
	unit_type := string(model.EXPANSION)
	collector := colly.NewCollector()
	unit_list := make(model.CrawlResultJSON)
	unit_list[unit_type] = make(model.CategoryGroupedJSON)
	unit_list[unit_type][string(model.EXPANSION)] = make(model.NameGroupedJSON)

	collector.OnHTML("div#wikibody", func(root *colly.HTMLElement) {
		root.ForEach("div.plugin_contents > ul > li > ul > li", func(i int, h *colly.HTMLElement) {
			name_links := h.DOM.ChildrenFiltered("a")

			name_links.Each(func(i int, name_link *goquery.Selection) {
				href, _ := name_link.Attr("href")
				h3 := root.DOM.ChildrenFiltered(fmt.Sprintf("h3%s", href))
				table := h3.Next().Next().Next()
				unit_list[unit_type][string(model.EXPANSION_TYPE)][h3.Text()] = a.GetUnitInfo(table)
			})
		})
	})

	err := collector.Visit(string(config.EXPANSION_PAGE))
	return unit_list, err
}

func (a *expansion) IsTargetSection(s *goquery.Selection) bool {
	return s.Text() != ""
}

func (a *expansion) GetUnitInfo(s *goquery.Selection) model.UnitInfoJSON {
	unit_info := make(model.UnitInfoJSON)
	s.ChildrenFiltered("tbody").ChildrenFiltered("tr").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered("th").First().Text()
		column := s.ChildrenFiltered("td").First().Text()
		unit_info[header] = column
	})

	return unit_info
}
