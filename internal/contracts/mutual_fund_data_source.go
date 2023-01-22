package contracts

import "aya-money-go/internal/entities"

type MutualFundDataSource interface {
	GetMutualFundData(productId string) (*entities.MutualFund, error)
}
