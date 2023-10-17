package crawler

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/gocolly/colly"
)

type crossWeapon struct{}

func NewCrossWeapon() repository.ExternalWeapon {
	return &crossWeapon{}
}

func (a *crossWeapon) Fetch() (model.CrawlResultJSON, error) {
	unit_type := string(model.CROSS_WEAPON_UNIT_TYPE)
	collector := colly.NewCollector()
	weapon_list := make(model.CrawlResultJSON)
	weapon_list[unit_type] = make(model.CategoryGroupedJSON)
	collector.OnHTML("div#wikibody", func(root *colly.HTMLElement) {
		root.ForEach("div.plugin_contents > ul > li > a", func(i int, category *colly.HTMLElement) {
			name_links := category.DOM.NextAllFiltered("ul").ChildrenFiltered("li")

			if name_links.Size() > 0 {
				weapon_list[unit_type][category.Text] = make(model.NameGroupedJSON)
			}
			name_links.Each(func(i int, name_list *goquery.Selection) {
				name_link := name_list.Children()
				href, _ := name_link.Attr("href")
				weapon_list[unit_type][category.Text][name_link.Text()] = make(map[string]string)

				h3 := root.DOM.ChildrenFiltered(fmt.Sprintf("h3%s", href))
				tables := h3.NextAllFiltered("table:not([class=\"atwiki_plugin_region\"])")
				tbody := tables.First().ChildrenFiltered("tbody").First()
				weapon_list[unit_type][category.Text][h3.Text()] = a.getWeaponInfo(tbody)
			})
		})
	})

	err := collector.Visit(string(config.CROSS_WEAPON_PAGE))
	return weapon_list, err
}

func (a *crossWeapon) getWeaponInfo(s *goquery.Selection) model.UnitInfoJSON {
	weapon_info := make(model.UnitInfoJSON)
	s.ChildrenFiltered("tr").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered("th").First().Text()
		column := s.ChildrenFiltered("td").First().Text()
		weapon_info[header] = column
	})

	return weapon_info
}
