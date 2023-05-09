package repository

import "gorm.io/gorm"

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{
		db: db,
	}
}

type IChannelRepository interface {
	GetAll()
	Get()
	Save()
	Update()
}

func (r *ChannelRepository) GetAll() {

}

func (r *ChannelRepository) Get() {

}

func (r *ChannelRepository) Save() {

}

func (r *ChannelRepository) Update() {

}
