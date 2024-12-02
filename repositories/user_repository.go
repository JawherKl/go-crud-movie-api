package repositories

import (
    "example.com/go-crud-api/models"
    "gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *models.User) (*models.User, error)
    GetUserByID(id uint) (*models.User, error)
    GetUserByEmail(email string) (*models.User, error)
    UpdateUser(user *models.User) (*models.User, error)
    DeleteUser(id uint) error
}

type GormUserRepository struct {
    DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
    return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) CreateUser(user *models.User) (*models.User, error) {
    result := r.DB.Create(user)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

func (r *GormUserRepository) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    result := r.DB.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := r.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *GormUserRepository) UpdateUser(user *models.User) (*models.User, error) {
    result := r.DB.Save(user)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

func (r *GormUserRepository) DeleteUser(id uint) error {
    var user models.User
    result := r.DB.Delete(&user, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}
