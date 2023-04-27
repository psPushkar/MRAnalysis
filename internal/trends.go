package demo

type GoogleTrendsResponse struct {
	Location string
	Value    string
}

func GetInterestByLocation(keyword string) ([]GoogleTrendsResponse, error) {

	var data []GoogleTrendsResponse

	fileData, err := CustomFileReader(keyword, "trends")
	if err != nil {
		return data, err
	}

	data = ExtractTrendsData(*fileData)
	return data, nil

}
