package models

type Media struct {
	Type     string
	Width    int
	Height   int
	Delivery string
	Url      string
}

type Ad struct {
	Id          int
	Title       string
	Description string
	Priority    string
	ZoneId      string
	MediaFiles  []Media
}

func NullAd() Ad {
	return Ad{Id: 0, Title: "", Description: "", Priority: "", ZoneId: "", MediaFiles: nil}
}

func (ad *Ad) SetMediaFiles(mediaFiles []Media) {
	ad.MediaFiles = mediaFiles
}

func (ad Ad) IsNully() bool {
	return ad.Id == 0
}
