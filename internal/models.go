package demo

type IntBeneficiary struct {
	Priority      int
	IsBeneficiary bool
	Value         int64
}

type RankedKeywordsResponse struct {
	Keyword string
	Score   float64
	Rank    int
}

type FloatBeneficiary struct {
	Priority      float64
	IsBeneficiary bool
	Value         float64
}

type KeywordPlannerData struct {
	Index            string
	Keyword          string
	Country          string
	ParentKeyword    string
	SERP             string
	Difficulty       FloatBeneficiary
	Volume           FloatBeneficiary
	CPC              FloatBeneficiary
	CPS              FloatBeneficiary
	LastUpdate       string
	GlobalVolume     FloatBeneficiary
	TrafficPotential FloatBeneficiary
}

type KeywordDataSetModel struct {
	Index            int
	Keyword          string
	Country          string
	Difficulty       int
	Volume           int64
	CPC              float64
	CPS              float64
	ParentKeyword    string
	LastUpdate       string
	SERP             string
	GlobalVolume     int64
	TrafficPotential int64
}

type KeywordPlannerResponse struct {
	Keyword            string
	Currency           string
	AvgMonthlySearches string
	ThreeMonthChange   string
	Yoy                string
	Competition        string
	CompetitionIndexed string
	TopPageBidLow      string
	TopPageBidHigh     string
}
