package model

type Shoplist struct {
}

type Order struct {
	OutTradeNo  string
	BuyerId     int64
	BuyerAliId  string
	SellerId    int64
	SellerAliId string
	TradeNo     string
	TradeStatus string
	Subject     string
	GmtCreate   string
	GmtPayment  string
}
