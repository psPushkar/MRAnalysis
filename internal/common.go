package demo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetFileBaseName(key string) string {

	switch key {

	case "health insurance":
		return health_insurance
	case "bike insurance":
		return bike_insurance
	case "car insurance":
		return car_insurance
	case "life insurance":
		return life_insurance
	case "term insurance":
		return term_insurance
	case "travel insurance":
		return travel_insurance
	case "mobile insurance":
		return mobile_insurance
	case "marine insurance":
		return marine_insurance
	case "commercial vehicle insurance":
		return commercial_vehicle_insurance
	case "motor insurance":
		return motor_insurance
	case "senior citizen health insurance":
		return senior_citizen_health_insurance
	case "home insurance":
		return home_insurance
	case "crop insurance":
		return crop_insurance
	case "fire insurance":
		return fire_insurance
	case "personal accident insurance":
		return personal_accident_insurance
	case "critical illness insurance":
		return critical_illness_insurance
	case "pension plans":
		return pension_plans
	case "international travel insurance":
		return international_travel_insurance
	case "property insurance":
		return property_insurance
	case "whole life insurance":
		return whole_life_insurance
	case "liability insurance":
		return liability_insurance
	case "group health insurance":
		return group_health_insurance
	case "endowment plans":
		return endowment_plans
	case "group insurance":
		return group_health_insurance
	case "professional indemnity insurance":
		return professional_indemnity_insurance
	case "cyber insurance":
		return cyber_insurance
	case "commercial insurance":
		return commercial_insurance
	case "maternity health insurance":
		return maternity_health_insurance
	case "product liability insurance":
		return product_liability_insurance
	case "ulips":
		return ulips
	case "individual health insurance":
		return individual_health_insurance
	case "agriculture insurance":
		return agriculture_insurance
	case "domestic travel insurance":
		return domestic_travel_insurance
	case "cargo insurance":
		return cargo_insurance
	case "hull insurance":
		return hull_insurance
	case "freight insurance":
		return freight_insurance
	case "student travel insurance":
		return student_travel_insurance
	case "business interruption insurance":
		return business_interruption_insurance
	case "burglary insurance":
		return burglary_insurance
	case "directors and officers liability insurance":
		return directors_and_officers_liability_insurance
	case "child plans":
		return child_plans
	case "cyber liability insurance":
		return cyber_liability_insurance
	case "flood insurance":
		return flood_insurance
	case "livestock insurance":
		return livestock_insurance
	case "home appliance insurance":
		return home_appliance_insurance
	case "earthquake insurance":
		return earthquake_insurance
	case "electronic equipment insurance":
		return electronic_equipment_insurance
	case "home content insurance":
		return home_content_insurance
	case "marine liability insurance":
		return marine_liability_insurance
	case "data breach insurance":
		return data_breach_insurance
	case "family floater plans":
		return family_floater_plans
	case "home structure insurance":
		return home_structure_insurance
	case "farm equipment insurance":
		return farm_equipment_insurance
	case "cybercrime insurance":
		return cybercrime_insurance
	case "loans":
		return loans
	case "online banking":
		return online_banking
	case "mobile banking":
		return mobile_banking
	case "credit cards":
		return credit_cards
	case "deposits":
		return deposits
	case "cash management services":
		return cash_management_services
	case "insurance services":
		return insurance_services
	case "debit cards":
		return debit_cards
	case "letters of credit":
		return letters_of_credit
	case "investment services":
		return investment_services
	case "bill payment services":
		return bill_payment_services
	case "foreign exchange services":
		return foreign_exchange_services
	case "atm services":
		return atm_services
	case "pension services":
		return pension_services
	case "trade finance services":
		return trade_finance_services
	case "phone banking services":
		return phone_banking_services
	case "safe deposit lockers":
		return safe_deposit_lockers
	case "sms banking services":
		return sms_banking_services
	case "cheque books":
		return cheque_books
	case "mudra loan":
		return mudra_loan
	case "loan against property":
		return loan_against_property
	case "personal loans":
		return personal_loans
	case "home loans":
		return home_loans
	case "business loans":
		return business_loans
	case "car loans":
		return car_loans
	case "startup india loan":
		return startup_india_loan
	case "wedding loans":
		return wedding_loans
	case "education loans":
		return education_loans
	case "medical loans":
		return medical_loans
	case "gold loans":
		return gold_loans
	case "nri loans":
		return nri_loans
	case "travel loans":
		return travel_loans
	case "credit card loans":
		return credit_card_loans
	case "consumer durable loans":
		return consumer_durable_loans
	case "agricultural loans":
		return agricultural_loans
	case "microfinance loans":
		return microfinance_loans
	case "professional loans":
		return professional_loans
	case "rural housing loans":
		return rural_housing_loans
	case "etc":
		return etc
	case "two-wheeler loans":
		return two_wheeler_loans

	default:
		return health_insurance

	}

	return "This project is under prototyping phase hence works with limited Rates. Sorry For The inconvinience!"
}

func CustomFileReader(key string, t string) (*[][]string, error) {

	var filepath string

	if t == "trends" {
		filepath = GetFileBaseName(key)
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
					result.Difficulty.Priority = 0.10
					result.Difficulty.Value = float64(getFloatValue(field))

				} else if j == 4 {

					result.Volume.IsBeneficiary = true
					result.Volume.Priority = 0.30
					result.Volume.Value = float64(getFloatValue(field))

				} else if j == 5 {

					result.CPC.IsBeneficiary = true
					result.CPC.Priority = 0.08
					result.CPC.Value = float64(getFloatValue(field))

				} else if j == 6 {

					result.CPS.IsBeneficiary = true
					result.CPS.Priority = 0.07
					result.CPS.Value = float64(getFloatValue(field))

				} else if j == 7 {
					result.ParentKeyword = field
				} else if j == 8 {
					result.LastUpdate = (field)
				} else if j == 9 {
					result.SERP = field
				} else if j == 10 {

					result.GlobalVolume.IsBeneficiary = true
					result.GlobalVolume.Priority = 0.30
					result.GlobalVolume.Value = float64(getFloatValue(field))

				} else if j == 11 {

					result.TrafficPotential.IsBeneficiary = true
					result.TrafficPotential.Priority = 0.15
					result.TrafficPotential.Value = float64(getFloatValue(field))

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

func getFloatValue(s string) float64 {

	b2, _ := strconv.ParseFloat(s, 64)
	return b2
}

func GetCalculatedResults(keywordData []KeywordPlannerData) []RankedKeywordsResponse {

	// get the data for The vars specified
	var maxVolume float64 = 0
	var maxGlobalVolume float64 = 0
	var maxTrafficPotential float64 = 0
	var minDifficulty float64 = 999999999
	var minCPC float64 = 999999999
	var minCPS float64 = 999999999

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

		keywordData[i].Volume.Value = keywordData[i].Volume.Value * (float64(keywordData[i].Volume.Priority))
		keywordData[i].GlobalVolume.Value = keywordData[i].GlobalVolume.Value * (float64(keywordData[i].GlobalVolume.Priority))
		keywordData[i].TrafficPotential.Value = keywordData[i].TrafficPotential.Value * (float64(keywordData[i].TrafficPotential.Priority))
		keywordData[i].Difficulty.Value = (float64(keywordData[i].Difficulty.Priority)) * keywordData[i].Difficulty.Value
		keywordData[i].CPC.Value = (float64(keywordData[i].CPC.Priority)) * keywordData[i].CPC.Value
		keywordData[i].CPS.Value = (float64(keywordData[i].CPS.Priority)) * keywordData[i].CPS.Value
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
		return keywordRanks[i].Score > keywordRanks[j].Score
	})

	// return ranks and the winner of all
	for i := range keywordRanks {
		keywordRanks[i].Rank = i
	}

	return keywordRanks

}

func GetTrendsScript(keyword string) string {

	var trendsScript string = ""
	var ns string = ""

	s := `<script type="text/javascript" src="https://ssl.gstatic.com/trends_nrtr/3316_RC03/embed_loader.js"></script> <script type="text/javascript"> trends.embed.renderExploreWidget("GEO_MAP", {"comparisonItem":[{"keyword":"%s","geo":"IN","time":"today 3-m"}],"category":0,"property":""},`

	sArr := strings.Split(keyword, " ")
	for i := range sArr {
		ns = ns + sArr[i] + "%20"
	}

	trendsScript = trendsScript + fmt.Sprintf(s, keyword) + ` {"exploreQuery":"date=today%203-m&geo=IN&q=` + ns[:len(ns)-3] + `&hl=en-GB","guestPath":"https://trends.google.com:443/trends/embed/"}); </script>`
	return trendsScript
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
