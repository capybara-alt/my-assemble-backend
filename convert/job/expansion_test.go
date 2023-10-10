package convert_test

import (
	"sort"
	"testing"
	"time"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/stretchr/testify/assert"
)

var expansion_schema model.ValidationUnitSchemaList = []model.ValidationUnitSchema{
	{
		PropName:  "effect_range",
		NameJa:    "効果範囲",
		NameEn:    "Effect Range",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
	{
		PropName:  "duration",
		NameJa:    "持続時間",
		NameEn:    "Duration",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
	{
		PropName:  "attack_power",
		NameJa:    "攻撃力",
		NameEn:    "Attack Power",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
	{
		PropName:  "blast_radius",
		NameJa:    "爆発範囲",
		NameEn:    "Blast Radius",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
	{
		PropName:  "direct_hit_adjustment",
		NameJa:    "直撃補正",
		NameEn:    "Direct Hit Adjustment",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
	{
		PropName:  "resilience",
		NameJa:    "耐久性能",
		NameEn:    "Resilience",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
	{
		PropName:  "accumulative_impact",
		NameJa:    "衝撃残留",
		NameEn:    "Accumulative Impact",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
	{
		PropName:  "impact",
		NameJa:    "衝撃力",
		NameEn:    "Impact",
		ValueType: "INT",
		UnitType:  "EXPANSION",
	},
}

func TestConvertExpansion(t *testing.T) {
	core.SetFakeTime(time.Now())

	tests := []struct {
		name          string
		crawl_results model.CrawlResultJSON
		want          model.Want[[]model.Expansion]
	}{
		{
			name: "[正常系]EXPANSIONの変換",
			crawl_results: model.CrawlResultJSON{
				string(model.EXPANSION_TYPE): {
					"コア拡張機能": {
						"ASSAULT ARMOR": {
							"攻撃力":  "1500",
							"衝撃力":  "2000",
							"衝撃残留": "1380",
							"爆発範囲": "60",
							"効果範囲": "200",
							"直撃補正": "230",
						},
						"PULSE ARMOR": {
							"耐久性能": "3300",
							"持続時間": "10",
						},
						"PULSE PROTECTION": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "7000 (+3000)",
							"持続時間":     "25",
						},
						"TERMINAL ARMOR": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "20000",
							"持続時間":     "5 (+3)",
						},
					},
				},
			},
			want: model.Want[[]model.Expansion]{
				Value: []model.Expansion{
					{
						Name:                "ASSAULT ARMOR",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         200,
						Duration:            0,
						AttackPower:         1500,
						BlastRadius:         60,
						DirectHitAdjustment: 230,
						Resilience:          0,
						Impact:              2000,
						AccumulativeImpact:  1380,
						CreatedAt:           core.Now(),
					},
					{
						Name:                "PULSE ARMOR",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         0,
						Duration:            10,
						AttackPower:         0,
						BlastRadius:         0,
						DirectHitAdjustment: 0,
						Resilience:          3300,
						Impact:              0,
						AccumulativeImpact:  0,
						CreatedAt:           core.Now(),
					},
					{
						Name:                "PULSE PROTECTION",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         0,
						Duration:            25,
						AttackPower:         0,
						BlastRadius:         0,
						DirectHitAdjustment: 0,
						Resilience:          7000,
						Impact:              0,
						AccumulativeImpact:  0,
						CreatedAt:           core.Now(),
					},
					{
						Name:                "TERMINAL ARMOR",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         0,
						Duration:            5,
						AttackPower:         0,
						BlastRadius:         0,
						DirectHitAdjustment: 0,
						Resilience:          20000,
						Impact:              0,
						AccumulativeImpact:  0,
						CreatedAt:           core.Now(),
					},
				},
				ErrMsg: "",
			},
		},
		{
			name: "[準正常系]全角数値が入っている場合",
			crawl_results: model.CrawlResultJSON{
				string(model.EXPANSION_TYPE): {
					"コア拡張機能": {
						"ASSAULT ARMOR": {
							"攻撃力":  "１５００",
							"衝撃力":  "2000",
							"衝撃残留": "1380",
							"爆発範囲": "60",
							"効果範囲": "200",
							"直撃補正": "230",
						},
						"PULSE ARMOR": {
							"耐久性能": "3300",
							"持続時間": "10",
						},
						"PULSE PROTECTION": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "7000 (+3000)",
							"持続時間":     "25",
						},
						"TERMINAL ARMOR": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "20000",
							"持続時間":     "5 (+3)",
						},
					},
				},
			},
			want: model.Want[[]model.Expansion]{
				Value:  nil,
				ErrMsg: "(EXPANSION>attack_power) Invalid value type expect: INT, value: １５００",
			},
		},
		{
			name: "[準正常系]データが抜けている場合",
			crawl_results: model.CrawlResultJSON{
				string(model.EXPANSION_TYPE): {
					"コア拡張機能": {
						"ASSAULT ARMOR": nil,
						"PULSE ARMOR": {
							"耐久性能": "3300",
							"持続時間": "10",
						},
						"PULSE PROTECTION": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "7000 (+3000)",
							"持続時間":     "25",
						},
						"TERMINAL ARMOR": {
							"レギュレーション": "1.03.1",
							"耐久性能":     "20000",
							"持続時間":     "5 (+3)",
						},
					},
				},
			},
			want: model.Want[[]model.Expansion]{
				Value: []model.Expansion{
					{
						Name:                "ASSAULT ARMOR",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         0,
						Duration:            0,
						AttackPower:         0,
						BlastRadius:         0,
						DirectHitAdjustment: 0,
						Resilience:          0,
						Impact:              0,
						AccumulativeImpact:  0,
						CreatedAt:           core.Now(),
					},
					{
						Name:                "PULSE ARMOR",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         0,
						Duration:            10,
						AttackPower:         0,
						BlastRadius:         0,
						DirectHitAdjustment: 0,
						Resilience:          3300,
						Impact:              0,
						AccumulativeImpact:  0,
						CreatedAt:           core.Now(),
					},
					{
						Name:                "PULSE PROTECTION",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         0,
						Duration:            25,
						AttackPower:         0,
						BlastRadius:         0,
						DirectHitAdjustment: 0,
						Resilience:          7000,
						Impact:              0,
						AccumulativeImpact:  0,
						CreatedAt:           core.Now(),
					},
					{
						Name:                "TERMINAL ARMOR",
						Category:            "コア拡張機能",
						Unit:                string(model.EXPANSION_TYPE),
						EffectRange:         0,
						Duration:            5,
						AttackPower:         0,
						BlastRadius:         0,
						DirectHitAdjustment: 0,
						Resilience:          20000,
						Impact:              0,
						AccumulativeImpact:  0,
						CreatedAt:           core.Now(),
					},
				},
				ErrMsg: "",
			},
		},
	}

	for _, tt := range tests {
		c := convert.NewExpansionList(expansion_schema)
		v, err := c.Convert(tt.crawl_results)
		sort.SliceStable(v, func(i, j int) bool {
			return v[i].Name < v[j].Name
		})
		t.Run(tt.name, func(t *testing.T) {
			if err != nil {
				assert.EqualError(t, err, tt.want.ErrMsg)
			}
			assert.Equal(t, v, tt.want.Value)
		})
	}
}
