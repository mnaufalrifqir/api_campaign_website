package transaction

import (
	"api-campaign/user"
)

type GetCampaignTransactionsInput struct {
	ID uint `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct{
	Amount int `json:"amount" binding:"required"`
	CampaignID uint `json:"campaign_id" binding:"required"`
	User user.User
}