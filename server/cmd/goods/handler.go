package main

import (
	goods "NihiStore/server/shared/kitex_gen/goods"
	"context"
)

// GoodsServiceImpl implements the last service interface defined in the IDL.
type GoodsServiceImpl struct{}

// CreateGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) CreateGoods(ctx context.Context, req *goods.CreateGoodsRequest) (resp *goods.CreateGoodsResponse, err error) {
	// TODO: Your code here...

	return
}
