package tools

import "strings"

func ParseOutTradeNo(otn string) (buyerId, goodsId string) {
	parts := strings.Split(otn, "-")
	return parts[0], parts[1]
}
