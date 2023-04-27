package demo

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
)

func getFileBaseName(key string) string {

	switch key {

	case "k_fitness":
		return "keyword_planner_fitness.csv"
	case "k_cycling":
		return "keyword_planner_cycling.csv"
	case "k_gym":
		return "keyword_planner_gym.csv"
	case "k_swimming":
		return "keyword_planner_swimming.csv"
	case "k_riding":
		return "keyword_planner_riding.csv"
	case "health insurance":
		return "trends_fitness.csv"
	case "t_cycling":
		return "trends_cycling.csv"
	case "t_gym":
		return "trends_gym.csv"
	case "t_swimming":
		return "trends_swimming.csv"
	case "t_riding":
		return "trends_riding.csv"

	}

	return "This project is under prototyping phase hence works with limited Rates. Sorry For The inconvinience!"
}

func CustomFileReader(key string, t string) (*[][]string, error) {

	var filepath string

	if t == "trends" {
		filepath = getFileBaseName(key)
	} else {
		filepath = "keyword_planner.csv"
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("some error occurred while fetching data from the api")
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	filedata, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New("some error occurred while parsing data from the api")
	}

	//ExtractKeywordData(&filedata)

	return &filedata, nil
}

func ExtractKeywordData(data *[][]string) map[string]KeywordPlannerData {

	// var list []KeywordPlannerData

	var mappedList = make(map[string]KeywordPlannerData)

	for i, line := range *data {

		if i > 0 { // omit headers in the file

			var result KeywordPlannerData

			for j, field := range line {

				if j == 0 {
					result.Index = field
				} else if j == 1 {
					result.Keyword = field
				} else if j == 2 {
					result.Country = field
				} else if j == 3 {

					result.Difficulty.IsBeneficiary = false
					result.Difficulty.Priority = 10
					result.Difficulty.Value = float64(getIntValue(field))

				} else if j == 4 {

					result.Volume.IsBeneficiary = true
					result.Volume.Priority = 30
					result.Volume.Value = float64(getIntValue(field))

				} else if j == 5 {

					result.CPC.IsBeneficiary = true
					result.CPC.Priority = 8
					result.CPC.Value = float64(getIntValue(field))

				} else if j == 6 {

					result.CPS.IsBeneficiary = true
					result.CPS.Priority = 7
					result.CPS.Value = float64(getIntValue(field))

				} else if j == 7 {
					result.ParentKeyword = field
				} else if j == 8 {
					result.LastUpdate = (field)
				} else if j == 9 {
					result.SERP = field
				} else if j == 10 {

					result.GlobalVolume.IsBeneficiary = true
					result.GlobalVolume.Priority = 30
					result.GlobalVolume.Value = float64(getIntValue(field))

				} else if j == 11 {

					result.TrafficPotential.IsBeneficiary = true
					result.CPC.Priority = 15
					result.CPC.Value = float64(getIntValue(field))

				}
			}

			if _, ok := mappedList[result.Keyword]; !ok {
				mappedList[result.Keyword] = result
			} else {
				mappedList[result.Keyword] = result
			}

			// list = append(list, result)
		}
	}
	return mappedList
}

func ExtractTrendsData(data [][]string) []GoogleTrendsResponse {

	var list []GoogleTrendsResponse

	for i, line := range data {
		if i > 0 { // omit headers in the file

			var result GoogleTrendsResponse
			for j, field := range line {
				if j == 0 {
					result.Location = field
				} else if j == 1 {
					result.Value = field
				}
			}

			list = append(list, result)
		}
	}
	return list
}

func GetNormalizedKeywordData(data []KeywordPlannerResponse) KeywordPlannerResponse {

	var res KeywordPlannerResponse

	var avgSearch int

	// get the average for the keyword
	for _, item := range data {

		avgSearch += getIntValue(item.AvgMonthlySearches)
		// res.Currency += item.Currency
		//res.ThreeMonthChange += item.ThreeMonthChange
		// res.Yoy += item.Yoy
		// res.Competition += item.Competition
		// res.CompetitionIndexed += item.CompetitionIndexed
		// res.TopPageBidHigh += item.TopPageBidHigh
		// res.TopPageBidLow += item.TopPageBidLow
	}

	res.Keyword = data[0].Keyword
	res.AvgMonthlySearches = strconv.Itoa(avgSearch)
	res.Currency = data[0].Currency
	res.ThreeMonthChange = data[0].ThreeMonthChange
	res.Yoy = data[0].Yoy
	res.Competition = data[0].Competition
	res.CompetitionIndexed = data[0].CompetitionIndexed
	res.TopPageBidHigh = data[0].TopPageBidHigh
	res.TopPageBidLow = data[0].TopPageBidLow

	return res

}

func getIntValue(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Print(err)
	}

	return v
}

func GetCalculatedResults(keywordData []KeywordPlannerData) []RankedKeywordsResponse {

	// get the data for The vars specified
	var maxVolume float64 = 0
	var maxGlobalVolume float64 = 0
	var maxTrafficPotential float64 = 0
	var minDifficulty float64 = 0
	var minCPC float64 = 0
	var minCPS float64 = 0

	for _, i := range keywordData {

		// get Max and Min based on beneficiary
		if maxVolume < i.Volume.Value {
			maxVolume = i.Volume.Value
		}
		if maxGlobalVolume < i.GlobalVolume.Value {
			maxGlobalVolume = i.GlobalVolume.Value
		}
		if maxTrafficPotential < i.TrafficPotential.Value {
			maxTrafficPotential = i.TrafficPotential.Value
		}

		if minCPC > i.CPC.Value {
			minCPC = i.CPC.Value
		}
		if minDifficulty > i.Difficulty.Value {
			minDifficulty = i.Difficulty.Value
		}
		if minCPS > i.CPS.Value {
			minCPS = i.CPS.Value
		}

	}

	// get the priority accordingly

	// use formula item/max and min/item acc to beneficiary

	for i := range keywordData {

		keywordData[i].Volume.Value = keywordData[i].Volume.Value / maxVolume
		keywordData[i].GlobalVolume.Value = keywordData[i].GlobalVolume.Value / maxGlobalVolume
		keywordData[i].TrafficPotential.Value = keywordData[i].TrafficPotential.Value / maxTrafficPotential
		keywordData[i].Difficulty.Value = minDifficulty / keywordData[i].Difficulty.Value
		keywordData[i].CPC.Value = minCPC / keywordData[i].CPC.Value
		keywordData[i].CPS.Value = minCPS / keywordData[i].CPS.Value
	}

	// mutiply the priorities
	for i := range keywordData {

		keywordData[i].Volume.Value = keywordData[i].Volume.Value * (float64(keywordData[i].Volume.Priority / 100))
		keywordData[i].GlobalVolume.Value = keywordData[i].GlobalVolume.Value * (float64(keywordData[i].GlobalVolume.Priority / 100))
		keywordData[i].TrafficPotential.Value = keywordData[i].TrafficPotential.Value * (float64(keywordData[i].TrafficPotential.Priority / 100))
		keywordData[i].Difficulty.Value = (float64(keywordData[i].Difficulty.Priority / 100)) * keywordData[i].Difficulty.Value
		keywordData[i].CPC.Value = (float64(keywordData[i].CPC.Priority / 100)) * keywordData[i].CPC.Value
		keywordData[i].CPS.Value = (float64(keywordData[i].CPS.Priority / 100)) * keywordData[i].CPS.Value
	}

	//get the summed response along with the ranks
	var keywordRanks []RankedKeywordsResponse
	// var winner int
	// var winnerScore float64 = 0

	for _, i := range keywordData {

		var ranked RankedKeywordsResponse
		ranked.Score = i.Volume.Value + i.Difficulty.Value + i.TrafficPotential.Value + i.CPC.Value + i.CPS.Value + i.Difficulty.Value
		ranked.Keyword = i.Keyword

		// if winnerScore < ranked.Score {
		// 	winnerScore = ranked.Score
		// 	winner = rank
		// }

		keywordRanks = append(keywordRanks, ranked)
	}

	// creating ranks
	sort.Slice(keywordRanks, func(i, j int) bool {
		return keywordRanks[i].Score < keywordRanks[j].Score
	})

	// return ranks and the winner of all
	for i := range keywordRanks {
		keywordRanks[i].Rank = i
	}

	return keywordRanks

}

// var shoppingList []KeywordIdeasDummyResponse
// 	for i, line := range data {
// 		if i > 0 { // omit header line
// 			var rec KeywordIdeasDummyResponse
// 			for j, field := range line {
// 				if j == 0 {
// 					rec.Keyword = field
// 				} else if j == 1 {
// 					rec.Currency = field
// 				} else if j == 2 {
// 					rec.AvgMonthlySearches = field
// 				} else if j == 3 {
// 					rec.ThreeMonthChange = field
// 				} else if j == 4 {
// 					rec.Yoy = field
// 				} else if j == 5 {
// 					rec.Competition = field
// 				} else if j == 6 {
// 					rec.CompetitionIndexed = field
// 				} else if j == 7 {
// 					rec.TopPageBidLow = field
// 				} else if j == 8 {
// 					rec.TopPageBidHigh = field
// 				}
// 			}
// 			shoppingList = append(shoppingList, rec)
// 		}
// 	}
