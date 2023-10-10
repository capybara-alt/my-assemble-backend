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
	weapon_list := make(model.CrawlResultJSON)
	weapon_list[a.unit_type] = make(model.CategoryGroupedJSON)
	collector.OnHTML("div#wikibody", func(h *colly.HTMLElement) {
		h.ForEach("div.plugin_contents > ul > li > a", func(i int, h *colly.HTMLElement) {
			name_links := h.DOM.NextAllFiltered("ul").ChildrenFiltered("li")

			weapon_info := make(model.NameGroupedJSON)
			name_links.Each(func(i int, s *goquery.Selection) {
				weapon_info[s.Text()] = make(model.UnitInfoJSON)
			})
			if len(weapon_info) > 0 {
				weapon_list[a.unit_type][h.Text] = weapon_info
			}
		})

		h.ForEach("h2", func(i int, h *colly.HTMLElement) {
			if _, ok := weapon_list[a.unit_type][h.Text]; !ok {
				return
			}
			h.DOM.NextUntil("h2").Each(func(i int, s *goquery.Selection) {
				h3 := s.Filter("h3")
				if a.IsTargetSection(h3) {
					h3.NextAllFiltered("table").First().Each(func(i int, s *goquery.Selection) {
						if a.IsTargetTable(s) {
							weapon_list[a.unit_type][h.Text][h3.Text()] = a.GetWeaponInfo(s)
						}
					})
				}
			})
		})
	})

	err := collector.Visit(a.target_url)
	return weapon_list, err
}

func (a *AbstractCrawler) IsTargetSection(s *goquery.Selection) bool {
	return s.Text() != ""
}

func (a *AbstractCrawler) IsTargetTable(s *goquery.Selection) bool {
	filter_table_style := "background-color:transparent;margin:0;padding:0;border:none;"
	filter_target_table := fmt.Sprintf("tr[style=\"%s\"]", filter_table_style)
	return s.ChildrenFiltered("tbody").ChildrenFiltered(filter_target_table).Text() == ""
}

func (a *AbstractCrawler) GetWeaponInfo(s *goquery.Selection) model.UnitInfoJSON {
	weapon_info := make(model.UnitInfoJSON)
	s.ChildrenFiltered("tbody").ChildrenFiltered("tr").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered("th").First().Text()
		column := s.ChildrenFiltered("td").First().Text()
		weapon_info[header] = column
	})

	return weapon_info
}
