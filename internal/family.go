package internal

import (
	"fmt"

	"github.com/go-rod/rod"
)

type StatsTree struct {
	Gearscore Family
	Energy    Family
	Points    Family

	Gathering  Life
	Fishing    Life
	Hunting    Life
	Cooking    Life
	Alchemy    Life
	Processing Life
	Training   Life
	Trade      Life
	Farming    Life
	Sailing    Life
	Barter     Life
}

type Family struct {
	Level string
}

type Life struct {
	Level string
	Max   string
}

func FamilyStats(cname *string) {

	familystats := map[string]Family{
		"gearscore": {
			Level: "#wrap > div > div.container > article > div > div > div.box_profile_area.box_area > div > div > div > ul > li:nth-child(3) > span.desc",
		},
		"energy": {
			Level: "#wrap > div > div.container > article > div > div > div.box_profile_area.box_area > div > div > div > ul > li:nth-child(4) > span.desc",
		},
		"points": {
			Level: "#wrap > div > div.container > article > div > div > div.box_profile_area.box_area > div > div > div > ul > li:nth-child(5) > span.desc",
		},
	}

	lifestats := map[string]Life{
		"gathering": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(1) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(1) > span.spec_stat",
		},
		"fishing": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(2) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(2) > span.spec_stat",
		},
		"hunting": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(3) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(3) > span.spec_stat",
		},
		"cooking": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(4) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(4) > span.spec_stat",
		},
		"alchemy": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(5) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(5) > span.spec_stat",
		},
		"processing": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(6) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(6) > span.spec_stat",
		},
		"training": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(7) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(7) > span.spec_stat",
		},
		"trade": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(8) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(8) > span.spec_stat",
		},
		"farming": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(9) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(9) > span.spec_stat",
		},
		"sailing": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(10) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(10) > span.spec_stat",
		},
		"barter": {
			Level: "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(11) > span.spec_level > em",
			Max:   "#wrap > div > div.container > article > div > div > div.box_game_data_wrap > div.box_spec_area.box_area > ul > li:nth-child(11) > span.spec_stat",
		},
	}

	browser := rod.New().MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	var url = fmt.Sprintf(
		"https://www.naeu.playblackdesert.com/en-US/Adventure?checkSearchText=False&region=NA&searchType=1&searchKeyword=%s",
		*cname,
	)

	page := browser.MustPage(url).MustWindowFullscreen()
	page.MustElement("#CybotCookiebotDialogBodyButtonDecline").MustClick()

	page.MustElement(
		"#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.title > a",
	).MustScrollIntoView().MustClick()

	stats := StatsTree{}

	// Map family stat names to their destination fields.
	familyFields := map[string]*Family{
		"gearscore": &stats.Gearscore,
		"energy":    &stats.Energy,
		"points":    &stats.Points,
	}

	for name, dst := range familyFields {
		selector := familystats[name].Level
		dst.Level = page.MustElement(selector).MustText()
	}

	// Map life stat names to their destination fields.
	lifeFields := map[string]*Life{
		"gathering":  &stats.Gathering,
		"fishing":    &stats.Fishing,
		"hunting":    &stats.Hunting,
		"cooking":    &stats.Cooking,
		"alchemy":    &stats.Alchemy,
		"processing": &stats.Processing,
		"training":   &stats.Training,
		"trade":      &stats.Trade,
		"farming":    &stats.Farming,
		"sailing":    &stats.Sailing,
		"barter":     &stats.Barter,
	}

	for name, dst := range lifeFields {
		selectors := lifestats[name]

		dst.Level = page.MustElement(selectors.Level).MustText()
		dst.Max = page.MustElement(selectors.Max).MustText()
	}

	fmt.Printf("Gearscore: %s | Energy: %s | Points: %s\nGathering: %s/%s | Fishing: %s/%s | Hunting: %s/%s\n",
		stats.Gearscore.Level, stats.Energy.Level, stats.Points.Level,
		stats.Gathering.Level, stats.Gathering.Max,
		stats.Fishing.Level, stats.Fishing.Max,
		stats.Hunting.Level, stats.Hunting.Max,
	)
}
