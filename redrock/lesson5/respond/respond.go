package respond

type Respond struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func (r Respond) Error() string {
	return r.Info
}

type Finalrespond struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	Data   string `json:"data"`
}

var (
	Ok        = Respond{Status: "10000", Info: "successful"}
	WrongName = Respond{Status: "40001", Info: "Wrong username"}
	WrongPwd  = Respond{Status: "40002", Info: "wrong password"}
)

func InternalError(err error) Respond {
	return Respond{
		Status: "500",
		Info:   err.Error(),
	}
}
