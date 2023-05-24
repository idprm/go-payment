package repository

import (
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/gorm"
)

type ApplicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

type IApplicationRepository interface {
	GetAll() (*[]entity.Application, error)
	GetById(int) (*entity.Application, error)
	GetByCode(string) (*entity.Application, error)
	GetByUrl(string) (*entity.Application, error)
	GetByUrlCallback(string) (*entity.Application, error)
	CountByUrlCallback(string) (int64, error)
	Save(*entity.Application) (*entity.Application, error)
	Update(*entity.Application) (*entity.Application, error)
	Delete(int) error
}

func (r *ApplicationRepository) GetAll() (*[]entity.Application, error) {
	var applications []entity.Application
	err := r.db.Order("id desc").Find(&applications).Error
	if err != nil {
		return nil, err
	}
	return &applications, err
}

func (r *ApplicationRepository) GetById(id int) (*entity.Application, error) {
	var application entity.Application
	err := r.db.Where("id = ?", id).Take(&application).Error
	if err != nil {
		return nil, err
	}
	return &application, err
}

func (r *ApplicationRepository) GetByCode(code string) (*entity.Application, error) {
	var application entity.Application
	err := r.db.Where("code = ?", code).Take(&application).Error
	if err != nil {
		return nil, err
	}
	return &application, err
}

func (r *ApplicationRepository) GetByUrl(url string) (*entity.Application, error) {
	var application entity.Application
	err := r.db.Where("url = ?", url).Take(&application).Error
	if err != nil {
		return nil, err
	}
	return &application, err
}

func (r *ApplicationRepository) GetByUrlCallback(urlCallback string) (*entity.Application, error) {
	var application entity.Application
	err := r.db.Where("url_callback = ?", urlCallback).Take(&application).Error
	if err != nil {
		return nil, err
	}
	return &application, err
}

func (r *ApplicationRepository) CountByUrlCallback(urlCallback string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Application{}).Where("url_callback = ?", urlCallback).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ApplicationRepository) Save(application *entity.Application) (*entity.Application, error) {
	err := r.db.Create(&application).Error
	if err != nil {
		return nil, err
	}
	return application, nil
}

func (r *ApplicationRepository) Update(application *entity.Application) (*entity.Application, error) {
	err := r.db.Save(&application).Error
	if err != nil {
		return nil, err
	}
	return application, nil
}

func (r *ApplicationRepository) Delete(id int) error {
	var application entity.Application
	err := r.db.Where("id = ?", id).Delete(&application).Error
	if err != nil {
		return err
	}
	return nil
}
