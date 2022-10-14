package models

type GetToken struct {
	IsGenerate bool
	Token      string
}

func (r *GetToken) CheckParams(params map[string][]string) bool {
	return len(params) == 1
}
