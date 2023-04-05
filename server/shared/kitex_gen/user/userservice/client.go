// Code generated by Kitex v0.5.1. DO NOT EDIT.

package userservice

import (
	user "NihiStore/server/shared/kitex_gen/user"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	CreateFavorites(ctx context.Context, req *user.CreateFavoritesRequest, callOptions ...callopt.Option) (r *user.CreateFavoritesResponse, err error)
	WatchFavorites(ctx context.Context, req *user.WatchFavoritesRequset, callOptions ...callopt.Option) (r *user.WatchFavoritesResponse, err error)
	WatchGoodsInFavorites(ctx context.Context, req *user.WatchGoodsInFavoritesRequest, callOptions ...callopt.Option) (r *user.WatchGoodsInFavoritesResponse, err error)
	DeleteFavorites(ctx context.Context, req *user.DeleteFavoritesRequest, callOptions ...callopt.Option) (r *user.DeleteFavoritesResponse, err error)
	CollectGoods(ctx context.Context, req *user.CollectGoodsRequest, callOptions ...callopt.Option) (r *user.CollectGoodsResponse, err error)
	AddToCart(ctx context.Context, req *user.AddToCartRequest, callOptions ...callopt.Option) (r *user.AddToCartResponse, err error)
	AddAmountCart(ctx context.Context, req *user.AddAmountCartRequest, callOptions ...callopt.Option) (r *user.AddAmountCartResponse, err error)
	DeleteAmountCart(ctx context.Context, req *user.DeleteAmountCartRequest, callOptions ...callopt.Option) (r *user.DeleteAmountCartResponse, err error)
	RemoveOutCart(ctx context.Context, req *user.RemoveOutCartRequest, callOptions ...callopt.Option) (r *user.RemoveOutCartResponse, err error)
	WatchCart(ctx context.Context, req *user.AddToCartRequest, callOptions ...callopt.Option) (r *user.WatchCartResponse, err error)
	CleanCart(ctx context.Context, req *user.CleanCartRequest, callOptions ...callopt.Option) (r *user.CleanCartResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kUserServiceClient) Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kUserServiceClient) CreateFavorites(ctx context.Context, req *user.CreateFavoritesRequest, callOptions ...callopt.Option) (r *user.CreateFavoritesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFavorites(ctx, req)
}

func (p *kUserServiceClient) WatchFavorites(ctx context.Context, req *user.WatchFavoritesRequset, callOptions ...callopt.Option) (r *user.WatchFavoritesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.WatchFavorites(ctx, req)
}

func (p *kUserServiceClient) WatchGoodsInFavorites(ctx context.Context, req *user.WatchGoodsInFavoritesRequest, callOptions ...callopt.Option) (r *user.WatchGoodsInFavoritesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.WatchGoodsInFavorites(ctx, req)
}

func (p *kUserServiceClient) DeleteFavorites(ctx context.Context, req *user.DeleteFavoritesRequest, callOptions ...callopt.Option) (r *user.DeleteFavoritesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFavorites(ctx, req)
}

func (p *kUserServiceClient) CollectGoods(ctx context.Context, req *user.CollectGoodsRequest, callOptions ...callopt.Option) (r *user.CollectGoodsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CollectGoods(ctx, req)
}

func (p *kUserServiceClient) AddToCart(ctx context.Context, req *user.AddToCartRequest, callOptions ...callopt.Option) (r *user.AddToCartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddToCart(ctx, req)
}

func (p *kUserServiceClient) AddAmountCart(ctx context.Context, req *user.AddAmountCartRequest, callOptions ...callopt.Option) (r *user.AddAmountCartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddAmountCart(ctx, req)
}

func (p *kUserServiceClient) DeleteAmountCart(ctx context.Context, req *user.DeleteAmountCartRequest, callOptions ...callopt.Option) (r *user.DeleteAmountCartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteAmountCart(ctx, req)
}

func (p *kUserServiceClient) RemoveOutCart(ctx context.Context, req *user.RemoveOutCartRequest, callOptions ...callopt.Option) (r *user.RemoveOutCartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveOutCart(ctx, req)
}

func (p *kUserServiceClient) WatchCart(ctx context.Context, req *user.AddToCartRequest, callOptions ...callopt.Option) (r *user.WatchCartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.WatchCart(ctx, req)
}

func (p *kUserServiceClient) CleanCart(ctx context.Context, req *user.CleanCartRequest, callOptions ...callopt.Option) (r *user.CleanCartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CleanCart(ctx, req)
}
