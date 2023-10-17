package crawler_test

import (
	"testing"

	"github.com/capybara-alt/my-assemble/infrastructure/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/stretchr/testify/assert"
)

func TestOrbitTaletDroneCrawler(t *testing.T) {
	test := struct {
		name string
		want model.Want[model.CrawlResultJSON]
	}{
		name: "[正常系]オービット・タレット・ドローンページのクロールテスト",
		want: model.Want[model.CrawlResultJSON]{
			Value: model.CrawlResultJSON{
				string(model.DEFAULT_WEAPON_BACK_UNIT_TYPE): {
					"実弾オービット": {
						"BO-044 HUXLEY": {
							"カテゴリー":  "BULLET ORBIT",
							"メーカー":   "ベイラム",
							"価格":     "305,000",
							"攻撃力":    "28×8",
							"衝撃力":    "39×8",
							"衝撃残留":   "15×8",
							"直撃補正":   "175",
							"性能保証射程": "130",
							"有効射程":   "205",
							"連射性能":   "5.5",
							"総弾数":    "240",
							"冷却性能":   "95",
							"弾単価":    "50",
							"重量":     "2230",
							"EN負荷":   "435",
						},
					},
					"レーザーオービット": {
						"45-091 ORBT": {
							"カテゴリー":    "LASER ORBIT",
							"メーカー":     "オールマインド",
							"価格":       "280,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "135 (+13) ×3",
							"衝撃力":      "70×3",
							"衝撃残留":     "39×3",
							"直撃補正":     "135 (+10)",
							"性能保証射程":   "198 (+68)",
							"有効射程":     "262 (+42)",
							"連射性能":     "1.0",
							"総弾数":      "165",
							"冷却性能":     "116",
							"弾単価":      "100",
							"重量":       "2010",
							"EN負荷":     "446",
						},
					},
					"レーザータレット": {
						"VP-60LT": {
							"カテゴリー":    "LASER TURRET",
							"メーカー":     "アーキバス",
							"価格":       "194,000",
							"レギュレーション": "1.03.1",
							"攻撃力":      "146×10 (+4)",
							"衝撃力":      "81×10 (+4)",
							"衝撃残留":     "39×10 (+4)",
							"直撃補正":     "135",
							"性能保証射程":   "250",
							"有効射程":     "304",
							"連射性能":     "0.9",
							"マガジン弾数":   "3",
							"総弾数":      "52",
							"リロード時間":   "5.0",
							"弾単価":      "500",
							"重量":       "2800",
							"EN負荷":     "560",
						},
					},
					"レーザードローン": {
						"Vvc-700LD": {
							"カテゴリー":    "LASER DRONE",
							"メーカー":     "VCPL",
							"価格":       "247,000",
							"攻撃力":      "288×6",
							"衝撃力":      "105×6",
							"衝撃残留":     "63×6",
							"チャージ攻撃力":  "1370×2",
							"チャージ衝撃力":  "480×2",
							"チャージ衝撃残留": "244×2",
							"直撃補正":     "135",
							"誘導性能":     "360",
							"有効射程":     "400",
							"誘導ロック時間":  "0.3",
							"最大ロック数":   "1",
							"チャージ時間":   "0.8",
							"総弾数":      "120",
							"リロード時間":   "10.0",
							"弾単価":      "150",
							"重量":       "3800",
							"EN負荷":     "570",
						},
					},
				},
			},
			ErrMsg: "",
		},
	}
	c := crawler.NewOrbitTaletDroneWeapon()
	t.Run(test.name, func(t *testing.T) {
		results, err := c.Fetch()
		assert.Equal(t, err, nil)
		assert.Equal(t, test.want.Value, results)
	})
}
