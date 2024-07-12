package grpc

import (
	"context"
	"github.com/dadaxiaoxiao/account/internal/domain"
	"github.com/dadaxiaoxiao/account/internal/service"
	accountv1 "github.com/dadaxiaoxiao/api-repository/api/proto/gen/account/v1"
	"github.com/ecodeclub/ekit/slice"
	"google.golang.org/grpc"
)

type AccountServiceServer struct {
	accountv1.UnimplementedAccountServiceServer
	svc service.AccountService
}

func NewAccountServiceServer(svc service.AccountService) *AccountServiceServer {
	return &AccountServiceServer{
		svc: svc,
	}
}

func (a *AccountServiceServer) Register(server *grpc.Server) {
	accountv1.RegisterAccountServiceServer(server, a)
}

func (a *AccountServiceServer) Credit(ctx context.Context, req *accountv1.CreditRequest) (*accountv1.CreditResponse, error) {
	err := a.svc.Credit(ctx, a.toDomain(req))
	return &accountv1.CreditResponse{}, err
}

func (a *AccountServiceServer) toDomain(c *accountv1.CreditRequest) domain.Credit {
	return domain.Credit{
		Biz:   c.GetBiz(),
		BizId: c.GetBizId(),
		Items: slice.Map(c.Items, func(idx int, src *accountv1.CreditItem) domain.CreditItem {
			return a.itemToDomain(src)
		}),
	}
}

func (a *AccountServiceServer) itemToDomain(c *accountv1.CreditItem) domain.CreditItem {
	return domain.CreditItem{
		Account: c.Account,
		Amt:     c.Amt,
		Uid:     c.Uid,
		// 因为这两个取值一样,可以直接转换
		AccountType: domain.AccountType(c.AccountType),
		Currency:    c.Currency,
	}
}
