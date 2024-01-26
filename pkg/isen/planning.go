package isen

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
)

type Planning struct {
	Url string `json:"url,omitempty"`
}

func GetCalendar(token aurion.Token) ([]byte, error) {

	info, err := GetSelfInfo(token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://ent-toulon.isen.fr/webaurion/ICS/%s.%s.ics", info.Firstname, info.Lastname)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	planning, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return planning, nil
}
