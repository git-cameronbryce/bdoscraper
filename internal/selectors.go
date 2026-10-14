package internal

import (
	"fmt"

	"github.com/go-rod/rod"
)

type Adventurer struct {
	Region string
	Guild  string
	Title  string
	Class  string
	Level  string
	User   string
}

func Selector(cname *string) {
	selectors := map[string]string{
		"region": "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.region",
		"guild":  "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.state",
		"title":  "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.title",
		"class":  "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.user > div > span.character_class > span.name",
		"level":  "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.user > div > span.character_desc > span.text_area > span.level",
		"user":   "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.user > div > span.character_desc > span.text_area > span.text",
	}

	browser := rod.New().MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	var url = fmt.Sprintf(
		"https://www.naeu.playblackdesert.com/en-US/Adventure?checkSearchText=False&region=NA&searchType=1&searchKeyword=%s",
		*cname,
	)

	page := browser.MustPage(url).MustWindowFullscreen()

	adventurer := Adventurer{
		Region: page.MustElement(selectors["region"]).MustText(),
		Guild:  page.MustElement(selectors["guild"]).MustText(),
		Title:  page.MustElement(selectors["title"]).MustText(),
		Class:  page.MustElement(selectors["class"]).MustText(),
		Level:  page.MustElement(selectors["level"]).MustText(),
		User:   page.MustElement(selectors["user"]).MustText(),
	}

	fmt.Printf(
		"Region: %s, Guild: %s, Title: %s, Class: %s, Level: %s, User: %s\n",
		adventurer.Region,
		adventurer.Guild,
		adventurer.Title,
		adventurer.Class,
		adventurer.Level,
		adventurer.User,
	)
}
