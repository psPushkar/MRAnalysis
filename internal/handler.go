package demo

import (
	"github.com/gin-gonic/gin"
)

type ResponseModel struct {
	Keywords []RankedKeywordsResponse
	Winner   string
	Index    int
	Location []GoogleTrendsResponse
}

func GetAnalysisForKeyword(c *gin.Context) {

	var keywords []string
	var res ResponseModel
	//var max = 0

	err := c.BindJSON(&keywords)
	if err != nil {
		c.JSON(200, gin.H{"error": "Invalid Request"})
		return
	}

	// var keyData []KeywordPlannerResponse

	keywordData, err := GetKeywordsPlannerData(keywords)
	if err != nil {
		c.JSON(200, gin.H{"error": "Google keyword planner service is down. Soory for the inconvinience"})
		return
	}

	rankedKeywords := GetCalculatedResults(keywordData)

	// for i, keyword := range keywords {

	// 	// fetch the normalized data

	// 	nData := GetNormalizedKeywordData(keywordData)
	// 	keyData = append(keyData, nData)

	// 	// finding the max based on the values
	// 	val, err := strconv.Atoi(nData.AvgMonthlySearches)
	// 	if err != nil {
	// 		c.JSON(200, gin.H{"error": "Invalid value"})
	// 		return
	// 	}

	// 	if max < val {
	// 		max = val
	// 		res.Winner = keyword
	// 		res.Index = i
	// 	}

	// }

	trendsData, err := GetInterestByLocation(rankedKeywords[0].Keyword)
	if err != nil {
		c.JSON(200, gin.H{"error": "Google trends service is down. Soory for the inconvinience"})
		return
	}

	res.Keywords = rankedKeywords
	res.Location = trendsData
	res.Winner = rankedKeywords[0].Keyword

	c.JSON(200, gin.H{"data": res, "code": 200})

}
