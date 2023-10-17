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
	unitType := string(model.CROSS_WEAPON_UNIT_TYPE)
	collector := colly.NewCollector()
	weaponList := make(model.CrawlResultJSON)
	weaponList[unitType] = make(model.CategoryGroupedJSON)
	collector.OnHTML("div#wikibody", func(root *colly.HTMLElement) {
		root.ForEach("div.plugin_contents > ul > li > a", func(i int, category *colly.HTMLElement) {
			nameLinks := category.DOM.NextAllFiltered("ul").ChildrenFiltered("li")

			if nameLinks.Size() > 0 {
				weaponList[unitType][category.Text] = make(model.NameGroupedJSON)
			}
			nameLinks.Each(func(i int, name_list *goquery.Selection) {
				nameLink := name_list.Children()
				href, _ := nameLink.Attr("href")
				weaponList[unitType][category.Text][nameLink.Text()] = make(map[string]string)

				h3 := root.DOM.ChildrenFiltered(fmt.Sprintf("h3%s", href))
				tables := h3.NextAllFiltered("table:not([class=\"atwiki_plugin_region\"])")
				tbody := tables.First().ChildrenFiltered("tbody").First()
				weaponList[unitType][category.Text][h3.Text()] = a.getWeaponInfo(tbody)
			})
		})
	})

	err := collector.Visit(string(config.CROSS_WEAPON_PAGE))
	return weaponList, err
}

func (a *crossWeapon) getWeaponInfo(s *goquery.Selection) model.UnitInfoJSON {
	info := make(model.UnitInfoJSON)
	s.ChildrenFiltered("tr").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered("th").First().Text()
		column := s.ChildrenFiltered("td").First().Text()
		info[header] = column
	})

	return info
}
