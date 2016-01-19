package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	INCOMING_URL = "https://hooks.slack.com/services/T00000000/B00000000/xxxxxxxxxxx"
)

type Params map[string]string

func main() {
	var text = flag.String("text", "no message. from Slack.", "")
	var username = flag.String("username", "no name...whoami", "")
	var icon_emoji = flag.String("icon", ":snowman:", "")
	var channel = flag.String("channel", "#general", "")

	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*text)

	var p Params = map[string]string{
		"text":       *text,
		"username":   *username,
		"icon_emoji": *icon_emoji,
		"channel":    *channel,
	}

	jsonParam := string(jsonEncode(p))
	fmt.Println(jsonParam)

	values := url.Values{}
	values.Add("payload", jsonParam)

	resp, err := http.PostForm(INCOMING_URL, values)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}
func jsonEncode(p Params) []byte {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return b
}
