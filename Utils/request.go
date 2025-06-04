package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func SendNotif(playerID string, title string, messages string, jsonStringData string) (*http.Response, error) {
	appIDOneSignal := "a2362255-5b95-4908-9e3a-2da4a996946a"
	apiKeyOneSignal := "os_v2_app_ui3cevk3sveqrhr2fwsktfuunj6cjz6b53dunbu7xlpfrr6phi76tbohhpb2linxuc7zg7mzkcvfomlk5jkmofvpcxy4dkdzhuseqdy"

	r, _ := http.NewRequest("POST",
		"https://onesignal.com/api/v1/notifications",
		bytes.NewBuffer([]byte(`
		{
			"app_id":"`+appIDOneSignal+`",
			"headings":{
				"en":"`+title+`"
			},
			"contents":{
				"en":"`+messages+`"
			},
			"small_icon":"icon_isp",
			"include_player_ids":["`+playerID+`"],
			"large_icon":"",
			"android_group":"group 1",
			"data":`+jsonStringData+`
		}
		`)),
	) // URL-encoded payload - icon : https://static.thenounproject.com/png/38170-200.png
	tr := &http.Transport{
		MaxIdleConns:       100,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Basic "+apiKeyOneSignal)
	resp, err := client.Do(r)
	if err != nil {
		// Logger(ctx, "Helper", "SendNotif", "https://onesignal.com/api/v1/notifications", fmt.Sprintf("%v", r), err.Error(), "0")
		return nil, err
	}
	if resp.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		// Logger(ctx, "Helper", "SendNotif", "https://onesignal.com/api/v1/notifications", fmt.Sprintf("%v", r), string(bodyBytes), "0")
		return nil, fmt.Errorf(string(bodyBytes))
	}
	// bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Logger(ctx, "Helper", "SendNotif", "https://onesignal.com/api/v1/notifications", fmt.Sprintf("%v", r), string(bodyBytes), "2")
	return resp, nil
}

func RemoveHtmlTag(text string) string {
	result := text
	result = strings.ReplaceAll(result, "<p>", "")
	result = strings.ReplaceAll(result, "</p>", "")
	result = strings.ReplaceAll(result, "<b>", "")
	result = strings.ReplaceAll(result, "</b>", "")
	result = strings.ReplaceAll(result, "&nbsp;", " ")
	return result
}
