package models

type Brt struct {
	Id       int    `json:"id"`
	Time     string `json:"time"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Desc     string `json:"isi"`
	Path     string `json:"path"`
}

type Brtn struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Time  string `json:"time"`
}

type Dbrt struct {
	Icon  []Icon
	Logo  []Head
	Data  []Brt
	Data2 []Brtn
}

type Kbrt struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Category string `json:"category"`
}

type Kdb struct {
	Icon []Icon
	Logo []Head
	Data []Kbrt
}
