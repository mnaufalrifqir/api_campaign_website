package campaign

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Name     string
	ShortDescription string
	Description string
	Perks string
	BackerCount int
	GoalAmount int
	CurrentAmount int
	Slug string
	UserID uint
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	gorm.Model
	FileName string
	IsPrimary bool
	CampaignID uint
}