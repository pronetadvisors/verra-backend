package models

import "time"

type Email struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Email      string    `json:"email"`
	TimeViewed uint      `json:"time_viewed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (e *Email) CreateEmail() (*Email, error) {
	if err := DB.Create(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Email) GetEmails() ([]Email, error) {
	var emails []Email

	if err := DB.Find(&emails).Error; err != nil {
		return nil, err
	}

	return emails, nil
}

func (e *Email) GetEmailByID(id string) (*Email, error) {
	if err := DB.Where("id = ?", id).First(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}
func GetEmailByEmail(email string) (*Email, error) {
	var e Email

	if err := DB.Where("email = ?", email).First(&e).Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Email) UpdateEmail() (*Email, error) {
	if err := DB.Save(&e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Email) DeleteEmail() error {
	if err := DB.Delete(&e).Error; err != nil {
		return err
	}

	return nil
}
