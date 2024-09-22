package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID uint) ([]Transaction, error)
	GetByUserID(userID uint) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID uint) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Find(&transactions).Error
	if err != nil{
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(userID uint) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("Campaign").Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil{
		return transactions, err
	}

	return transactions, nil
}

