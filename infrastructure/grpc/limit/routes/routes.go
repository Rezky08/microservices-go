// Package routes contains all routes of the application
package routes

import (
	context "context"
	"github.com/Rezky08/microservices-go/pb/limit"

	// swaggerFiles for documentation
	_ "github.com/Rezky08/microservices-go/docs"
)

type Routes struct {
	limit.UnimplementedLimitServiceServer
}

func (r Routes) CreateLimit(ctx context.Context, create *limit.LimitCreate) (*limit.LimitCreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r Routes) UpdateLimit(ctx context.Context, update *limit.LimitUpdate) (*limit.LimitUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r Routes) DetailLimit(ctx context.Context, detail *limit.LimitDetail) (*limit.LimitDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r Routes) DeleteLimit(ctx context.Context, limitDelete *limit.LimitDelete) (*limit.LimitUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r Routes) mustEmbedUnimplementedLimitServiceServer() {
	//TODO implement me
	panic("implement me")
}
