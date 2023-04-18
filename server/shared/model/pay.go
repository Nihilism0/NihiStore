package model

type Shoplist struct {
}

type Order struct {
	GoodsId     string
	OutTradeNo  string
	BuyerId     string
	BuyerAliId  string
	SellerId    int64
	SellerAliId string
	TradeNo     string
	TradeStatus string
	Subject     string
	GmtCreate   string
	GmtPayment  string
}
