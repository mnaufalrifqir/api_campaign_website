package transaction

import (
	"api-campaign/campaign"
	"api-campaign/user"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CampaignID uint
	UserID     uint
}
