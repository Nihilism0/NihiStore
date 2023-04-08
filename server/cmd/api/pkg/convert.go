package pkg

import (
	hbase "NihiStore/server/cmd/api/biz/model/base"
	kbase "NihiStore/server/shared/kitex_gen/base"
)

func ConvertGoodsInformation(in *hbase.Goods) *kbase.Goods {
	return &kbase.Goods{
		Name:        in.Name,
		Description: in.Description,
		Cost:        in.Cost,
	}
}
