package helper

type Response struct {
	Draw            int         `json:"draw"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}

type Data struct {
	Number int    `json:"no"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

func ApiResponse(draw int, recordsFiltered int, data interface{}) Response {
	jsonResponse := Response{
		Draw:            draw,
		RecordsFiltered: recordsFiltered,
		Data:            data,
	}
	return jsonResponse
}
