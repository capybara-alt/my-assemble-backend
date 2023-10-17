package crawler_test

import (
	"testing"

	"github.com/capybara-alt/my-assemble/infrastructure/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/stretchr/testify/assert"
)

func TestExplodeWeaponCrawler(t *testing.T) {
	test := struct {
		name string
		want model.Want[model.CrawlResultJSON]
	}{
		name: "[正常系]爆発武器ページのクロールテスト",
		want: model.Want[model.CrawlResultJSON]{
			Value: model.CrawlResultJSON{
				string(model.DEFAULT_WEAPON_ARM_UNIT_TYPE): {
					"バズーカ": {
						"DF-BA-06 XUAN-GE": {
							"カテゴリー":    "BAZOOKA",
							"メーカー":     "大豊",
							"価格":       "70,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "895",
							"衝撃力":      "980",
							"衝撃残留":     "765",
							"爆発範囲":     "15",
							"直撃補正":     "185",
							"射撃反動":     "90",
							"有効射程":     "740",
							"総弾数":      "78",
							"リロード時間":   "3.4 (-0.8)",
							"弾単価":      "450",
							"重量":       "5480",
							"EN負荷":     "240",
						},
						"MAJESTIC": {
							"カテゴリー":    "BAZOOKA",
							"メーカー":     "メリニット",
							"価格":       "121,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "1109",
							"衝撃力":      "1090",
							"衝撃残留":     "850",
							"爆発範囲":     "15",
							"直撃補正":     "190",
							"射撃反動":     "85",
							"有効射程":     "800",
							"総弾数":      "52",
							"リロード時間":   "4.2 (-0.7)",
							"弾単価":      "600",
							"重量":       "4660",
							"EN負荷":     "178",
						},
						"LITTLE GEM": {
							"カテゴリー":    "BAZOOKA",
							"メーカー":     "メリニット",
							"価格":       "163,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "819",
							"衝撃力":      "910",
							"衝撃残留":     "670",
							"爆発範囲":     "15",
							"直撃補正":     "180",
							"射撃反動":     "80",
							"有効射程":     "690",
							"総弾数":      "36",
							"リロード時間":   "4.4 (-1.4)",
							"弾単価":      "450",
							"重量":       "3100",
							"EN負荷":     "192",
						},
					},
					"デトネイティングバズーカ": {
						"44-141 JVLN ALPHA": {
							"カテゴリー":    "DETONATING BAZOOKA",
							"メーカー":     "オールマインド",
							"価格":       "210,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "1075",
							"衝撃力":      "1390",
							"衝撃残留":     "905",
							"爆発範囲":     "15",
							"直撃補正":     "220",
							"射撃反動":     "80",
							"有効射程":     "760",
							"総弾数":      "44",
							"リロード時間":   "4.3 (-0.7)",
							"弾単価":      "750",
							"重量":       "7420",
							"EN負荷":     "299",
						},
					},
					"グレネード": {
						"DF-GR-07 GOU-CHEN": {
							"カテゴリー":    "GRENADE",
							"メーカー":     "大豊",
							"価格":       "140,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "1450",
							"衝撃力":      "1197",
							"衝撃残留":     "906",
							"爆発範囲":     "70",
							"直撃補正":     "140",
							"射撃反動":     "100",
							"有効射程":     "625",
							"総弾数":      "40 (+8)",
							"リロード時間":   "5.9 (-1.6)",
							"弾単価":      "1200",
							"重量":       "5460",
							"EN負荷":     "385",
						},
						"DIZZY": {
							"カテゴリー":    "GRENADE",
							"メーカー":     "メリニット",
							"価格":       "260,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "1650",
							"衝撃力":      "1278",
							"衝撃残留":     "1003",
							"爆発範囲":     "90",
							"直撃補正":     "145",
							"射撃反動":     "100",
							"有効射程":     "285 (-375)",
							"総弾数":      "38 (+12)",
							"リロード時間":   "7.1 (-1.8)",
							"弾単価":      "1500",
							"重量":       "6420",
							"EN負荷":     "455",
						},
						"IRIDIUM": {
							"カテゴリー":    "GRENADE",
							"メーカー":     "メリニット",
							"価格":       "214,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "1090",
							"衝撃力":      "991",
							"衝撃残留":     "845",
							"爆発範囲":     "60",
							"直撃補正":     "140",
							"射撃反動":     "88",
							"有効射程":     "245 (-380)",
							"総弾数":      "32 (+12)",
							"リロード時間":   "4.5 (-1.8)",
							"弾単価":      "800",
							"重量":       "2020",
							"EN負荷":     "290",
						},
					},
				},
			},
			ErrMsg: "",
		},
	}
	c := crawler.NewExplodeWeapon()
	t.Run(test.name, func(t *testing.T) {
		results, err := c.Fetch()
		assert.Equal(t, err, nil)
		assert.Equal(t, test.want.Value, results)
	})
}