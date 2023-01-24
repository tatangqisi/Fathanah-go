package models

type Qrn struct {
	Id      int    `json:"id"`
	Surah   string `json:"surah"`
	Arab    string `json:"arab"`
	Latin   string `json:"latin"`
	Meaning string `json:"arti"`
}

type Dqrn struct {
	Icon []Icon
	Logo []Head
	Data []Qrn
}

type Qrns struct {
	Id      int    `json:"id"`
	Name    string `json:"surah"`
	Meaning string `json:"arti"`
}

type Dqs struct {
	Icon []Icon
	Logo []Head
	Data []Qrns
}
