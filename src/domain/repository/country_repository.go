package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

type ICountryRepository interface {
	GetAll() (*[]entity.Country, error)
	GetByLocale(string) (*entity.Country, error)
	CountByLocale(string) (int64, error)
	Save(*entity.Country) (*entity.Country, error)
	Update(*entity.Country) (*entity.Country, error)
	Delete(int) error
}

func (r *CountryRepository) GetAll() (*[]entity.Country, error) {
	var countrys []entity.Country
	err := r.db.Order("id asc").Preload("Gateway.Channel").Find(&countrys).Error
	if err != nil {
		return nil, err
	}
	return &countrys, err
}

func (r *CountryRepository) GetByLocale(locale string) (*entity.Country, error) {
	var country entity.Country
	err := r.db.Where("locale = ?", locale).Preload("Gateway.Channel").Take(&country).Error
	if err != nil {
		return nil, err
	}
	return &country, err
}

func (r *CountryRepository) CountByLocale(locale string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Country{}).Where("locale = ?", locale).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *CountryRepository) Save(country *entity.Country) (*entity.Country, error) {
	err := r.db.Create(&country).Error
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (r *CountryRepository) Update(country *entity.Country) (*entity.Country, error) {
	err := r.db.Save(&country).Error
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (r *CountryRepository) Delete(id int) error {
	var country entity.Country
	err := r.db.Where("id = ?", id).Delete(&country).Error
	if err != nil {
		return err
	}
	return nil
}
