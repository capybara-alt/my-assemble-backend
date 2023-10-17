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
	unitType := string(model.EXPANSION)
	collector := colly.NewCollector()
	unitList := make(model.CrawlResultJSON)
	unitList[unitType] = make(model.CategoryGroupedJSON)
	unitList[unitType][string(model.EXPANSION)] = make(model.NameGroupedJSON)

	collector.OnHTML("div#wikibody", func(root *colly.HTMLElement) {
		root.ForEach("div.plugin_contents > ul > li > ul > li", func(i int, h *colly.HTMLElement) {
			nameLinks := h.DOM.ChildrenFiltered("a")

			nameLinks.Each(func(i int, nameLink *goquery.Selection) {
				href, _ := nameLink.Attr("href")
				h3 := root.DOM.ChildrenFiltered(fmt.Sprintf("h3%s", href))
				table := h3.Next().Next().Next()
				unitList[unitType][string(model.EXPANSION_TYPE)][h3.Text()] = a.GetUnitInfo(table)
			})
		})
	})

	err := collector.Visit(string(config.EXPANSION_PAGE))
	return unitList, err
}

func (a *expansion) IsTargetSection(s *goquery.Selection) bool {
	return s.Text() != ""
}

func (a *expansion) GetUnitInfo(s *goquery.Selection) model.UnitInfoJSON {
	info := make(model.UnitInfoJSON)
	s.ChildrenFiltered("tbody").ChildrenFiltered("tr").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered("th").First().Text()
		column := s.ChildrenFiltered("td").First().Text()
		info[header] = column
	})

	return info
}
