package utils

import "gorm.io/gorm"

func HandleError(tx *gorm.DB) error {
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
