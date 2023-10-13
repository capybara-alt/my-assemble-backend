package crawler_test

import (
	"testing"

	"github.com/capybara-alt/my-assemble/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/stretchr/testify/assert"
)

func TestAmmoWeaponCrawler(t *testing.T) {
	test := struct {
		name string
		want model.Want[model.CrawlResultJSON]
	}{
		name: "[正常系]実弾武器ページのクロールテスト",
		want: model.Want[model.CrawlResultJSON]{
			Value: model.CrawlResultJSON{
				string(model.DEFAULT_WEAPON_ARM_UNIT_TYPE): {
					"バーストライフル": {
						"MA-J-200 RANSETSU-RF": {
							"カテゴリー":    "BURST RIFLE",
							"メーカー":     "BAWS",
							"価格":       "105,000",
							"レギュレーション": "1.02",
							"攻撃力":      "224 (+30)",
							"衝撃力":      "245",
							"衝撃残留":     "91",
							"チャージ攻撃力":  "224×3 (+30)",
							"チャージ衝撃力":  "245×3",
							"チャージ衝撃残留": "91×3",
							"直撃補正":     "220",
							"射撃反動":     "12",
							"性能保証射程":   "180",
							"有効射程":     "321",
							"連射性能":     "1.5 (+0.2)",
							"チャージ時間":   "0.7",
							"マガジン弾数":   "15",
							"総弾数":      "375",
							"リロード時間":   "1.8 (-0.2)",
							"弾単価":      "100",
							"重量":       "4210",
							"EN負荷":     "158",
						},
					},
					"リニアライフル": {
						"LR-036 CURTIS": {
							"カテゴリー":     "LINEAR RIFLE",
							"メーカー":      "ベイラム",
							"価格":        "65,000",
							"レギュレーション":  "1.03.1",
							"攻撃力":       "142",
							"衝撃力":       "135",
							"衝撃残留":      "52",
							"チャージ攻撃力":   "689",
							"チャージ衝撃力":   "850",
							"チャージ衝撃残留":  "260",
							"チャージ攻撃時発熱": "550",
							"直撃補正":      "220",
							"射撃反動":      "24 (-6)",
							"性能保証射程":    "190",
							"有効射程":      "337",
							"連射性能":      "2.5",
							"チャージ時間":    "0.5 (-0.5)",
							"マガジン弾数":    "12",
							"総弾数":       "420",
							"リロード時間":    "2.2",
							"冷却性能":      "155",
							"弾単価":       "100",
							"重量":        "4150",
							"EN負荷":      "289",
						},
						"LR-037 HARRIS": {
							"カテゴリー":     "LINEAR RIFLE",
							"メーカー":      "ベイラム",
							"価格":        "135,000",
							"レギュレーション":  "1.03.1",
							"攻撃力":       "239",
							"衝撃力":       "285",
							"衝撃残留":      "109",
							"チャージ攻撃力":   "977 (+127)",
							"チャージ衝撃力":   "1250",
							"チャージ衝撃残留":  "380",
							"チャージ攻撃時発熱": "1000",
							"直撃補正":      "220",
							"射撃反動":      "35",
							"性能保証射程":    "195",
							"有効射程":      "376",
							"連射性能":      "1.3",
							"チャージ時間":    "0.8 (-0.7)",
							"マガジン弾数":    "10",
							"総弾数":       "360",
							"リロード時間":    "3.0",
							"冷却性能":      "350 (+60)",
							"弾単価":       "200",
							"重量":        "4840",
							"EN負荷":      "441",
						},
					},
					"アサルトライフル": {
						"RF-024 TURNER": {
							"カテゴリー":    "ASSAULT RIFLE",
							"メーカー":     "ベイラム",
							"価格":       "55,000",
							"レギュレーション": "1.02",
							"攻撃力":      "105 (+18)",
							"衝撃力":      "65",
							"衝撃残留":     "25",
							"直撃補正":     "185",
							"射撃反動":     "17",
							"性能保証射程":   "160",
							"有効射程":     "296",
							"連射性能":     "3.4",
							"マガジン弾数":   "18",
							"総弾数":      "540",
							"リロード時間":   "2.2 (-0.6)",
							"弾単価":      "40",
							"重量":       "3560",
							"EN負荷":     "102",
						},
						"RF-025 SCUDDER": {
							"カテゴリー":    "ASSAULT RIFLE",
							"メーカー":     "ベイラム",
							"価格":       "205,000",
							"レギュレーション": "1.02",
							"攻撃力":      "135 (+25)",
							"衝撃力":      "82",
							"衝撃残留":     "35",
							"直撃補正":     "185",
							"射撃反動":     "19",
							"性能保証射程":   "162",
							"有効射程":     "304",
							"連射性能":     "2.9",
							"マガジン弾数":   "15",
							"総弾数":      "450",
							"リロード時間":   "2.4 (-0.8)",
							"弾単価":      "50",
							"重量":       "3830",
							"EN負荷":     "153",
						},
					},
					"バーストアサルトライフル": {
						"MA-J-201 RANSETSU-AR": {
							"カテゴリー":    "BURST ASSAULT RIFLE",
							"メーカー":     "BAWS",
							"価格":       "111,000",
							"レギュレーション": "1.02",
							"攻撃力":      "77×3 (+15)",
							"衝撃力":      "64×3",
							"衝撃残留":     "17×3",
							"直撃補正":     "185",
							"射撃反動":     "7",
							"性能保証射程":   "153",
							"有効射程":     "284",
							"連射性能":     "3.2",
							"マガジン弾数":   "18",
							"総弾数":      "450",
							"リロード時間":   "1.9 (-0.5)",
							"弾単価":      "40",
							"重量":       "3620",
							"EN負荷":     "132",
						},
					},
					"マシンガン": {
						"MG-014 LUDLOW": {
							"カテゴリー":    "MACHINE GUN",
							"メーカー":     "ベイラム",
							"価格":       "45,000",
							"レギュレーション": "1.02",
							"攻撃力":      "42 (+6)",
							"衝撃力":      "41",
							"衝撃残留":     "19",
							"直撃補正":     "195",
							"射撃反動":     "4 (-1)",
							"性能保証射程":   "115",
							"有効射程":     "236",
							"連射性能":     "10.0",
							"マガジン弾数":   "30",
							"総弾数":      "720 (+180)",
							"リロード時間":   "1.5 (-0.5)",
							"弾単価":      "20",
							"重量":       "2450",
							"EN負荷":     "82",
						},
						"DF-MG-02 CHANG-CHEN": {
							"カテゴリー":    "MACHINE GUN",
							"メーカー":     "大豊",
							"価格":       "120,000",
							"レギュレーション": "1.02",
							"攻撃力":      "39 (+7)",
							"衝撃力":      "40",
							"衝撃残留":     "18",
							"直撃補正":     "195",
							"射撃反動":     "6",
							"性能保証射程":   "100",
							"有効射程":     "220",
							"連射性能":     "10.0",
							"マガジン弾数":   "45",
							"総弾数":      "990 (+270)",
							"リロード時間":   "2.2 (-0.6)",
							"弾単価":      "20",
							"重量":       "3280",
							"EN負荷":     "143",
						},
					},
					"バーストマシンガン": {
						"MA-E-210 ETSUJIN": {
							"カテゴリー":    "BURST MACHINE GUN",
							"メーカー":     "BAWS",
							"価格":       "74,000",
							"レギュレーション": "1.02",
							"攻撃力":      "46×4 (+6)",
							"衝撃力":      "48×4",
							"衝撃残留":     "22×4",
							"直撃補正":     "195",
							"射撃反動":     "3",
							"性能保証射程":   "106",
							"有効射程":     "224",
							"連射性能":     "8.1",
							"マガジン弾数":   "24",
							"総弾数":      "600 (+120)",
							"リロード時間":   "1.5 (-0.5)",
							"弾単価":      "30",
							"重量":       "2810",
							"EN負荷":     "98",
						},
					},
					"ガトリングガン": {
						"DF-GA-08 HU-BEN": {
							"カテゴリー":  "GATLING GUN",
							"メーカー":   "大豊",
							"価格":     "170,000",
							"攻撃力":    "25",
							"衝撃力":    "25",
							"衝撃残留":   "11",
							"攻撃時発熱":  "9",
							"直撃補正":   "215",
							"射撃反動":   "5",
							"性能保証射程": "130",
							"有効射程":   "226",
							"連射性能":   "20.0",
							"総弾数":    "1300",
							"冷却性能":   "220",
							"弾単価":    "30",
							"重量":     "5800",
							"EN負荷":   "425",
						},
					},
					"ショットガン": {
						"SG-026 HALDEMAN": {
							"カテゴリー":    "SHOTGUN",
							"メーカー":     "ベイラム",
							"価格":       "75,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "576",
							"衝撃力":      "360",
							"衝撃残留":     "280",
							"直撃補正":     "190 (-15)",
							"射撃反動":     "90",
							"性能保証射程":   "88",
							"有効射程":     "169",
							"マガジン弾数":   "1",
							"総弾数":      "66",
							"リロード時間":   "1.3",
							"弾単価":      "200",
							"重量":       "3660",
							"EN負荷":     "185",
						},
						"SG-027 ZIMMERMAN": {
							"カテゴリー":    "SHOTGUN",
							"メーカー":     "ベイラム",
							"価格":       "115,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "800 (-100)",
							"衝撃力":      "620 (-220)",
							"衝撃残留":     "360 (-60)",
							"直撃補正":     "180 (-25)",
							"射撃反動":     "90",
							"性能保証射程":   "102",
							"有効射程":     "184",
							"マガジン弾数":   "1",
							"総弾数":      "53",
							"リロード時間":   "2.0",
							"弾単価":      "350",
							"重量":       "4400",
							"EN負荷":     "242",
						},
						"WR-0777 SWEET SIXTEEN": {
							"カテゴリー":  "SHOTGUN",
							"メーカー":   "RaD",
							"価格":     "49,000",
							"攻撃力":    "85×13 (1105)",
							"衝撃力":    "61×13 (793)",
							"衝撃残留":   "41×13 (533)",
							"直撃補正":   "205",
							"射撃反動":   "100",
							"性能保証射程": "76",
							"有効射程":   "155",
							"マガジン弾数": "13",
							"総弾数":    "546",
							"リロード時間": "3.0",
							"弾単価":    "30×13 (390)",
							"重量":     "1640",
							"EN負荷":   "268",
						},
					},
					"ハンドガン": {
						"HG-003 COQUILLETT": {
							"カテゴリー":    "HANDGUN",
							"メーカー":     "ベイラム",
							"価格":       "35,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "166",
							"衝撃力":      "235",
							"衝撃残留":     "149",
							"直撃補正":     "125",
							"射撃反動":     "32",
							"性能保証射程":   "85",
							"有効射程":     "167",
							"連射性能":     "2.5",
							"マガジン弾数":   "7",
							"総弾数":      "196 (+91)",
							"リロード時間":   "2.1",
							"弾単価":      "100",
							"重量":       "1200",
							"EN負荷":     "122",
						},
						"HG-004 DUCKETT": {
							"カテゴリー":    "HANDGUN",
							"メーカー":     "ベイラム",
							"価格":       "112,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "235",
							"衝撃力":      "300",
							"衝撃残留":     "151",
							"直撃補正":     "125",
							"射撃反動":     "30",
							"性能保証射程":   "90",
							"有効射程":     "178",
							"連射性能":     "1.7",
							"マガジン弾数":   "7",
							"総弾数":      "182 (+84)",
							"リロード時間":   "2.5",
							"弾単価":      "120",
							"重量":       "1650",
							"EN負荷":     "158",
						},
					},
					"バーストハンドガン": {
						"MA-E-211 SAMPU": {
							"カテゴリー":    "BURST HANDGUN",
							"メーカー":     "BAWS",
							"価格":       "73,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "87×2",
							"衝撃力":      "105×2",
							"衝撃残留":     "64×2",
							"直撃補正":     "125",
							"射撃反動":     "15",
							"性能保証射程":   "80",
							"有効射程":     "165",
							"連射性能":     "5.1",
							"マガジン弾数":   "12",
							"総弾数":      "300 (+144)",
							"リロード時間":   "1.9",
							"弾単価":      "40",
							"重量":       "960",
							"EN負荷":     "62",
						},
					},
					"ニードルガン": {
						"EL-PW-00 VIENTO": {
							"カテゴリー":    "NEEDLE GUN",
							"メーカー":     "エルカノ",
							"価格":       "148,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "181",
							"衝撃力":      "195",
							"衝撃残留":     "127",
							"直撃補正":     "130",
							"射撃反動":     "15",
							"性能保証射程":   "105",
							"有効射程":     "192",
							"連射性能":     "4.0",
							"マガジン弾数":   "5",
							"総弾数":      "160 (+40)",
							"リロード時間":   "1.9",
							"弾単価":      "80",
							"重量":       "1180",
							"EN負荷":     "215",
						},
					},
				},
			},
			ErrMsg: "",
		},
	}
	c := crawler.NewAmmoWeaponCrawler()
	t.Run(test.name, func(t *testing.T) {
		results, err := c.Crawl()
		assert.Equal(t, err, nil)
		assert.Equal(t, test.want.Value, results)
	})
}
