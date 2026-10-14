package internal

import (
	"fmt"
	"regexp"

	"github.com/go-rod/rod"
)

type Adventurer struct {
	Region string
	Guild  string
	Title  string
	Level  string
	User   string
	Type   string
}

func Selector(cname *string) {
	selectors := map[string]string{
		"region": "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.region",
		"guild":  "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.state",
		"title":  "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.title",
		"user":   "#wrap > div > div.container.adventure > article > div > div > div.box_list_area > ul > li > div.user",
	}

	browser := rod.New().MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	var url = fmt.Sprintf(
		"https://www.naeu.playblackdesert.com/en-US/Adventure?checkSearchText=False&region=NA&searchType=1&searchKeyword=%s",
		*cname,
	)

	page := browser.MustPage(url).MustWindowFullscreen()
	text := page.MustElement(selectors["user"]).MustText()

	pattern := regexp.MustCompile(`^(Lv\.\d+)(.+)\n(\w+)$`)
	matches := pattern.FindStringSubmatch(text)

	adventurer := Adventurer{
		Region: page.MustElement(selectors["region"]).MustText(),
		Guild:  page.MustElement(selectors["guild"]).MustText(),
		Title:  page.MustElement(selectors["title"]).MustText(),
		Level:  matches[1],
		User:   matches[2],
		Type:   matches[3],
	}

	fmt.Printf(
		"Region: %s, Guild: %s, Title: %s, Level: %s, User: %s, Type: %s\n",
		adventurer.Region,
		adventurer.Guild,
		adventurer.Title,
		adventurer.Level,
		adventurer.User,
		adventurer.Type,
	)
}
