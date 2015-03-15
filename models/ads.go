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
	ZoneField   string
	MediaFiles  []Media
}

func (ad *Ad) SetMediaFiles(mediaFiles []Media) {
	ad.MediaFiles = mediaFiles
}
