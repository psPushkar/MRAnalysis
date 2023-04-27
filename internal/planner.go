package demo

import "errors"

func GetKeywordsPlannerData(keywords []string) ([]KeywordPlannerData, error) {

	var data []KeywordPlannerData

	fileData, err := CustomFileReader("", "keywordPlanner")
	if err != nil {
		return data, err
	}

	keywordData := ExtractKeywordData(fileData)

	// search for the keyword in the data
	for _, item := range keywords {

		val, ok := keywordData[item]
		if !ok {
			return nil, errors.New("no data founnd for the keyword. please check for grammatical errors")
		}

		data = append(data, val)

	}

	return data, nil

}
