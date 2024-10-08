package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID uint) ([]Campaign, error)
	FindByID(ID uint) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID uint) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID uint) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID uint) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("id = ?", ID).Preload("User").Preload("CampaignImages").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error){
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r  *repository) MarkAllImagesAsNonPrimary(campaignID uint) (bool, error){
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error

	if err != nil{
		return false, err
	}

	return true, nil
}
