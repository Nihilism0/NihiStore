package mysql

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/model"
)

type MysqlFavoGenerator struct{}

func (*MysqlFavoGenerator) SelectFavoFromUserIdAndFavoName(UserId int64, FavoritesName string) model.Favorites {
	favorites := model.Favorites{}
	config.DB.Where("user_id = ? AND name = ?", UserId, FavoritesName).First(&favorites)
	return favorites
}

func (*MysqlFavoGenerator) SelectFavoInUserId(UserID int64) []model.Favorites {
	var favoriteses []model.Favorites
	config.DB.Where("id = ?", UserID).Find(&favoriteses)
	return favoriteses
}

func (*MysqlFavoGenerator) SelectCollectionsByUserAndFavo(UserId, FavoritesId int64) []model.Collection {
	var collections []model.Collection
	config.DB.Where("user_id = ? AND favorites_id = ?", UserId, FavoritesId).Find(&collections)
	return collections
}

func (*MysqlFavoGenerator) SelectFavoByIdAndUserId(FavoritesId, UserId int64) model.Favorites {
	var favorites model.Favorites
	config.DB.Where("id = ? AND user_id = ?", FavoritesId, UserId).First(&favorites)
	return favorites
}

func (*MysqlFavoGenerator) DeleteFavo(favo *model.Favorites) {
	config.DB.Unscoped().Delete(&favo)
}

func (*MysqlFavoGenerator) SelectCollectionByAllId(FavoritesId, UserId, GoodsId int64) model.Collection {
	var collection model.Collection
	config.DB.Where("favorites_id = ? AND user_id = ? AND goods_id = ?", FavoritesId, UserId, GoodsId).First(&collection)
	return collection
}
