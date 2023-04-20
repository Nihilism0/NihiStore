package tools

func CreateHeadMinioPath(Id string) string {
	return "headphoto/" + Id
}

func CreateGoodsPhotoMinioPath(Id string) string {
	return "goodsphoto/" + Id
}
