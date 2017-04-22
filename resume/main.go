package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

    "github.com/aymerick/raymond"
	"github.com/BurntSushi/toml"
	"github.com/russross/blackfriday"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cast"
)

const (
	resume_json = "data/resume.json"
	resume_style = "resume/style.css"
	resume_template = "resume/resume.hbs"
	resume_html = "resume/resume.html"
	resume_web_html = "resume/resume_web.html"
	site_config = "config.toml"
)

func loadResumeJson() (*simplejson.Json, error) {
	dat, err := ioutil.ReadFile(resume_json)
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(dat)
}

func renderToHtml(innerCtx *simplejson.Json) (string, error) {
	// set up the HTML renderer
	renderer := blackfriday.HtmlRenderer(0, "", "")
	options := blackfriday.Options{Extensions: 0}
	render := func(input string) string {
		return string(blackfriday.MarkdownOptions([]byte(input), renderer, options))
	}

	summaryPath := []string{"basics","summary"}
	summary := render(innerCtx.GetPath(summaryPath...).MustString())
	innerCtx.SetPath(summaryPath, summary)

	works := []interface{}{}
	// each work experience
	for i,_ := range innerCtx.Get("work").MustArray() {
		work := innerCtx.Get("work").GetIndex(i)
		positions := []interface{}{}
		// each position
		for j,_ := range work.Get("positions").MustArray() {
			position := work.Get("positions").GetIndex(j)
			highlights := []string{}
			// each highlight
			for _,highlight := range position.Get("highlights").MustStringArray() {
				highlights = append(highlights, render(highlight))
			}
			position.Set("highlights", highlights)
			positions = append(positions, position.Interface())
		}
		work.Set("positions", positions)
		works = append(works, work.Interface())
	}
	innerCtx.Set("work", works)

	css, err := ioutil.ReadFile(resume_style)
	ctx := map[string]interface{}{
		"resume": innerCtx.Interface(),
		"css": string(css),
	}

	work,_ := innerCtx.Get("work").Array()
	fmt.Print(len(work))
	rtpl, err := ioutil.ReadFile(resume_template)
	if err != nil {
		return "", err
	}

	return raymond.Render(string(rtpl), ctx)
}

func exportHtml(htmlString string, forWeb bool) error {
	fileName := resume_html
	if (forWeb) {
		fileName = resume_web_html
	}
	return ioutil.WriteFile(fileName, []byte(htmlString), 0644)
}

func main() {
	phoneNum := flag.String("p", os.Getenv("PHONE"), "phone #")
	flag.Parse()

	// configure raymond
	raymond.RegisterHelper("dateFormat", dateFormat)

	makeResume("")
	if (*phoneNum != "") {
		makeResume(*phoneNum)
	}
}

func makeResume(phoneNum string) error {
	forWeb := phoneNum == ""

	ctx, err := loadResumeJson()
	if err != nil {
		panic(err)
	}

	if err = extractConfigParams(ctx, phoneNum); err != nil {
		panic(err)
	}

	result, err := renderToHtml(ctx)
	if err != nil {
		panic(err)
	}

	return exportHtml(result, forWeb)
}

func extractConfigParams(ctx *simplejson.Json, phoneNum string) error {
	var foo map[string]interface{}
	if _, err := toml.DecodeFile(site_config, &foo); err != nil {
		return err
	}

	social,_ := foo["social"].(map[string]interface{})
	email := social["Email"]
	github := "github.com/" + social["Github"].(string)
	// linkedIn := "linkedin.com/in/" + social["LinkedIn"].(string)
	// twitter := "twitter.com/" + social["Twitter"].(string)
	website,_ := ctx.GetPath("basics", "website").String()

	ctx.SetPath([]string{"basics","email"}, email)
	if (phoneNum != "") {
		ctx.SetPath([]string{"basics","phone"}, phoneNum)
	}
	ctx.SetPath([]string{"basics","profiles"}, []map[string]string{
		map[string]string{"url":website},
		map[string]string{"url":github},
		// map[string]string{"url":twitter},
		// map[string]string{"url":linkedIn},
	})
	return nil
}

// dateFormat converts the textual representation of the datetime string into
// the other form or returns it of the time.Time value. These are formatted
// with the layout string
// lifted from https://github.com/spf13/hugo/blob/1cf29200b4bb0a9c006155ec76759b7f4b1ad925/tpl/tplimpl/template_funcs.go#L1649
func dateFormat(layout string, v interface{}) string {
	t, err := cast.ToTimeE(v)
	if err != nil {
		panic(err)
	}
	return t.Format(layout)
}
