package models

type GetToken struct {
	IsGenerate bool
	Token      string
}

func (r *GetToken) CheckTokenParams(params map[string][]string) bool {
	return len(params) == 1
}
