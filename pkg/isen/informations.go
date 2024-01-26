package isen

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
	"github.com/PuerkitoBio/goquery"
)

type SelfInfo struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func GetSelfInfo(token aurion.Token) (SelfInfo, error) {
	var selfInfo SelfInfo = SelfInfo{}

	client := &http.Client{}

	// Get homepage
	req, err := http.NewRequest("GET", MainMenuPage, nil)
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%v", token))

	if err != nil {
		return selfInfo, err
	}

	// Do request homepage
	resp, err := client.Do(req)
	if err != nil {
		return selfInfo, err
	}
	defer resp.Body.Close()

	// Test if everithing is ok
	if resp.StatusCode != 200 {
		return selfInfo, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return selfInfo, err
	}

	// If bad token, redirect to login page. Detect it
	if strings.Contains(string(responseBody), LoginPageMessageTest) {
		return selfInfo, fmt.Errorf("bad token")
	}

	reader := strings.NewReader(string(responseBody))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return selfInfo, err
	}

	var responseName string
	doc.Find("div.menuMonCompte h3").Each(func(i int, s *goquery.Selection) {
		responseName = s.Text()
	})

	if responseName != "" && len(strings.Split(responseName, " ")) > 1 {
		selfInfo.Firstname = strings.ToLower(strings.Split(responseName, " ")[0])
		selfInfo.Lastname = strings.ToLower(strings.Split(responseName, " ")[1])

		return selfInfo, err
	}

	return selfInfo, fmt.Errorf("informations not found")
}
