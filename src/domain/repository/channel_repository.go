package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{
		db: db,
	}
}

type IChannelRepository interface {
	GetAll(int) (*[]entity.Channel, error)
	GetBySlug(int, string) (*entity.Channel, error)
	CountBySlug(int, string) (int64, error)
	Save(*entity.Channel) (*entity.Channel, error)
	Update(*entity.Channel) (*entity.Channel, error)
	Delete(int) error
}

func (r *ChannelRepository) GetAll(gateId int) (*[]entity.Channel, error) {
	var channels []entity.Channel
	err := r.db.Where("gateway_id = ?", gateId).Preload("Gateway").Order("id asc").Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return &channels, err
}

func (r *ChannelRepository) GetBySlug(gateId int, slug string) (*entity.Channel, error) {
	var channel entity.Channel
	err := r.db.Where("gateway_id = ?", gateId).Where("slug = ?", slug).Preload("Gateway").Take(&channel).Error
	if err != nil {
		return nil, err
	}
	return &channel, err
}

func (r *ChannelRepository) CountBySlug(gateId int, slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Channel{}).Where("gateway_id = ?", gateId).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ChannelRepository) Save(channel *entity.Channel) (*entity.Channel, error) {
	err := r.db.Create(&channel).Error
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (r *ChannelRepository) Update(channel *entity.Channel) (*entity.Channel, error) {
	err := r.db.Save(&channel).Error
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (r *ChannelRepository) Delete(id int) error {
	var channel entity.Channel
	err := r.db.Where("id = ?", id).Delete(&channel).Error
	if err != nil {
		return err
	}
	return nil
}
