package expect

type RequestData struct {
	Vars      map[string]string
	Body      interface{}
	DecodeErr string
}

func SendVars(vars map[string]string) RequestData {
	return RequestData{Vars: vars}
}

func SendBody(body interface{}) RequestData {
	return RequestData{Body: body}
}

func SendVarsAndBody(vars map[string]string, body interface{}) RequestData {
	return RequestData{
		Vars: vars,
		Body: body,
	}
}
