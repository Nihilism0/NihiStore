package convert

import (
	"NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
)

type ConvertGenerator struct{}

func (*ConvertGenerator) ConvertUser(req *user.RegisterRequest, user *model.User) {
	user.Username = req.Username
	user.Password = req.Password
}

func (*ConvertGenerator) ConvertFavorites(req *user.CreateFavoritesRequest, favorites *model.Favorites) {
	favorites.Name = req.FavoritesName
	favorites.UserId = req.UserId
	favorites.Description = req.Description
}
