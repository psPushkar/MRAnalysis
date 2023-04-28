package demo

import (
	"html/template"
	"net/http"
	"strings"

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
	key := c.PostForm("keywords")
	key = strings.TrimSpace(key)
	key = strings.Trim(key, ",")
	keywords = strings.Split(key, ",")
	// c.JSON(200, gin.H{"data": keywords})

	for i := range keywords {
		keywords[i] = strings.ToLower(strings.TrimSpace(keywords[i]))
	}

	// err := c.BindJSON(&keywords)
	// if err != nil {
	// 	c.JSON(200, gin.H{"error": "Invalid Request"})
	// 	return
	// }

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

	// trendsData, err := GetInterestByLocation(rankedKeywords[0].Keyword)
	// if err != nil {
	// 	c.JSON(200, gin.H{"error": "Google trends service is down. Soory for the inconvinience"})
	// 	return
	// }

	res.Keywords = rankedKeywords
	//res.Location = trendsData
	res.Winner = rankedKeywords[0].Keyword

	// c.JSON(200, gin.H{
	// 	"winner":                    res.Winner,
	// 	"locations for The keyword": res.Location,
	// })

	// trendsData := GetFileBaseName(rankedKeywords[0].Keyword)

	trendsData := GetTrendsScript(rankedKeywords[0].Keyword)

	// template.Must(template.New("").Parse(trendsData)).Execute(os.Stdout, d)

	// var trendsConfig strings.Builder
	// for _, keyword := range rankedKeywords {
	// 	// Modify the widget type and configuration parameters as per your requirements
	// 	trendsConfig.WriteString(fmt.Sprintf(`trends.embed.renderExploreWidget("GEO_MAP", {"comparisonItem":[{"keyword":"%s","geo":"IN","time":"today 3-m"}],"category":0,"property":""}, {"exploreQuery":"date=today%%203-m&geo=IN&q=%s&hl=en","guestPath":"https://trends.google.com:443/trends/embed/"});\n`, keyword.Keyword, url.QueryEscape(keyword.Keyword)))
	// }

	c.HTML(http.StatusOK, "data.html", gin.H{
		"winner": res.Winner,
		"trend":  template.JS(trendsData),
	})

}

func GetKeywords(ctx *gin.Context) {

	content := []string{
		"health insurance ,",
		"bike insurance ,",
		"car insurance ,",
		"life insurance ,",
		"term insurance ,",
		"travel insurance ,",
		"mobile insurance ,",
		"marine insurance ,",
		"commercial vehicle insurance ,",
		"motor insurance ,",
		"senior citizen health insurance ,",
		"home insurance ,",
		"crop insurance ,",
		"fire insurance ,",
		"personal accident insurance ,",
		"critical illness insurance ,",
		"pension plans ,",
		"international travel insurance ,",
		"property insurance ,",
		"whole life insurance ,",
		"liability insurance ,",
		"group health insurance ,",
		"endowment plans ,",
		"group insurance ,",
		"professional indemnity insurance ,",
		"cyber insurance ,",
		"commercial insurance ,",
		"maternity health insurance ,",
		"product liability insurance ,",
		"ulips ,",
		"individual health insurance ,",
		"agriculture insurance ,",
		"domestic travel insurance ,",
		"cargo insurance ,",
		"hull insurance ,",
		"freight insurance ,",
		"student travel insurance ,",
		"business interruption insurance ,",
		"burglary insurance ,",
		"directors and officers liability insurance ,",
		"child plans ,",
		"cyber liability insurance ,",
		"flood insurance ,",
		"livestock insurance ,",
		"home appliance insurance ,",
		"earthquake insurance ,",
		"electronic equipment insurance ,",
		"home content insurance ,",
		"marine liability insurance ,",
		"data breach insurance ,",
		"family floater plans ,",
		"home structure insurance ,",
		"farm equipment insurance ,",
		"cybercrime insurance ,",
		"loans ,",
		"online banking ,",
		"mobile banking ,",
		"credit cards ,",
		"deposits ,",
		"cash management services ,",
		"insurance services ,",
		"debit cards ,",
		"letters of credit ,",
		"investment services ,",
		"bill payment services ,",
		"foreign exchange services ,",
		"atm services ,",
		"pension services ,",
		"trade finance services ,",
		"phone banking services ,",
		"safe deposit lockers ,",
		"sms banking services ,",
		"cheque books ,",
		"mudra loan ,",
		"loan against property ,",
		"personal loans ,",
		"home loans ,",
		"business loans ,",
		"car loans ,",
		"startup india loan ,",
		"wedding loans ,",
		"education loans ,",
		"medical loans ,",
		"gold loans ,",
		"nri loans ,",
		"travel loans ,",
		"credit card loans ,",
		"consumer durable loans ,",
		"agricultural loans ,",
		"microfinance loans ,",
		"professional loans ,",
		"rural housing loans ,",
		"etc ,",
		"two-wheeler loans ,",
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"contentKeywordsList": content,
	})
}
