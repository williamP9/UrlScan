package util

type Domain struct {
	Url        string
	Title      string
	Path       map[string]string
	Urls       map[string]string
	Statuscode string
	Urlnum     uint
}

func (d *Domain) GetTile() map[string]string {
	title := make(map[string]string)
	title["https://172.16.0.1/#/link"] = "345"
	title["https://app.tuta.com/mail/OAqTSjL-0R-9/OCW2LhE--3-9?mail"] = "123"
	return title
}
