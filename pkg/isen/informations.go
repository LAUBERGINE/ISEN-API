package isen

import (
	"github.com/AYDEV-FR/ISEN-Api/pkg/aurion"
)

type SelfInfo struct {
	Titre     string `json:"date,omitempty"`
	Firstname string `json:"code,omitempty"`
	Lastname  string `json:"name,omitempty"`
}

func GetSelfInfo(token aurion.Token) (SelfInfo, error) {
	var selfInfo SelfInfo = SelfInfo{}

	_, err := aurion.MenuNavigateTo(token, SelfInfoMenuId, MainMenuPage)
	if err != nil {
		return selfInfo, err
	}

	//fmt.Println(string(currentPage))

	// reader := strings.NewReader(string(currentPage))
	// doc, err := goquery.NewDocumentFromReader(reader)
	// if err != nil {
	// 	return selfInfo, err
	// }

	// fmt.Println(doc.Text())

	// doc.Find("tr[role='row']").Each(func(i int, s *goquery.Selection) {
	// 	var note Notation
	// 	s.Find("td[role='gridcell']").Each(func(i int, s *goquery.Selection) {
	// 		switch i {
	// 		case 0:
	// 			note.Date = s.Text()
	// 		case 1:
	// 			note.Code = s.Text()
	// 		case 2:
	// 			note.Name = s.Text()
	// 		case 3:
	// 			note.Note = s.Text()
	// 		case 4:
	// 			note.AbsenceReason = s.Text()
	// 		case 5:
	// 			note.Comments = s.Text()
	// 		case 6:
	// 			note.Teachers = strings.Split(s.Text(), ",")
	// 		}
	// 	})
	// 	notationsList = append(notationsList, note)
	// })

	// reader := strings.NewReader(htmlTable)
	// doc, err := goquery.NewDocumentFromReader(reader)
	// if err != nil {
	// 	return nil, err
	// }

	return selfInfo, err
}
