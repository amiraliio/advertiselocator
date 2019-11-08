package models

//image size
const (
	OriginalSize  string = "ORIGINAL"
	MediumSize    string = "MEDIUM"
	ThumbnailSize string = "THUMBNAIL"
)

//Image types
const (
	TypeJPG  string = "jpg"
	TypeJPEG string = "jpeg"
	TypeGIF  string = "gif"
	TypePNG  string = "png"
)

//file types
const (
	TypePDF string = "pdf"
)

//video types
const (
	TypeMP4  string = "mp4"
	Type3GP  string = "3gp"
	TypeMPEG string = "mpeg"
)

//audio types
const (
	TypeMP3 string = "mp3"
)

//BaseMedia model
type BaseMedia struct {
	OriginalURL string `json:"originalURL" bson:"originalURL"`
	URL         string `json:"url" bson:"url" validate:"required,min=10,max=1000"`
	Type        string `json:"type" bson:"type"`
}

//File model
type File struct {
	BaseMedia ",inline"
}

//Image model
type Image struct {
	BaseMedia ",inline"
	Size      string `json:"size" bson:"size"`
}

//Video model
type Video struct {
	BaseMedia ",inline"
}

//AdvertiseImage model
type AdvertiseImage struct {
	Image    ",inline"
	IsMain   bool   `json:"isMain" bson:"isMain" validate:"omitempty"`
	Caption  string `json:"caption" bson:"caption" validate:"omitempty,min=10,max=255"`
	Show     bool   `json:"show" bson:"show" validate:"required"`
	Priority byte   `json:"priority" bson:"priority" validate:"required,numeric,max=3"`
}

// func ValidateMediaType(filename, mediaType string) bool{
// 	switch mediaType {
// 	case "image":
// 		if
// 	}
// }
