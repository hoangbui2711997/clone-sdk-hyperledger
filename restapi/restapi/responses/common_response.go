package responses

//go:generate easytags $GOFILE json,xml

type ResponseBool struct {
	Data  bool `json:"data" xml:"data"`
	Error *Err `json:"error" xml:"error"`
}

type ResponseCommonSingle struct {
	Data  interface{} `json:"data" xml:"data"`
	Error *Err        `json:"error" xml:"error"`
}

type ResponseCommonSingleWithValidate struct {
	Data  interface{}  `json:"data" xml:"data"`
	Error *ValidateErr `json:"error" xml:"error"`
}

type ResponseCommonArray struct {
	Data       interface{} `json:"data" xml:"data"`
	TotalCount int64       `json:"total_count" xml:"total_count"`
	Error      *Err        `json:"error" xml:"error"`
}

type ResponseCommonArrayPtr struct {
	Data       []*interface{} `json:"data" xml:"data"`
	TotalCount int64          `json:"total_count" xml:"total_count"`
	Error      *Err           `json:"error" xml:"error"`
}
