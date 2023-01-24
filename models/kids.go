package models

type Jkids struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Img  string `json:"img"`
	Path string `json:"path"`
}

type Jkd struct {
	Data []Jkids
}

type Ikids struct {
	Id    int    `json:"id"`
	Img   string `json:"img"`
	Menu  string `json:"menu"`
	Title string `json:"judul"`
	Desc  string `json:"penjelasan"`
	Path  string `json:"path"`
}

type Dkds struct {
	Data []Ikids
}
