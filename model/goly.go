package model

func GetAllGolies() ([]Url, error) {
	var urlies []Url

	tx := db.Find(&urlies)

	if tx.Error != nil {
		return []Url{}, tx.Error
	}

	return urlies, nil
}

func GetGoly(id uint64) (Url, error) {
	var urlies Url

	tx := db.Where("id = ?", id).First(&urlies)

	if tx.Error != nil {
		return Url{}, tx.Error
	}

	return urlies, nil
}

func CreateRewrite(goly Url) error {
	tx := db.Create(&goly)
	return tx.Error
}

func UpdateRewrite(goly Url) error {
	tx := db.Save(&goly)
	return tx.Error
}

func DeleteUrl(id uint64) error {
	tx := db.Unscoped().Delete(&Url{}, id)
	return tx.Error
}

func FindByGolyUrl(url string) (Url, error) {
	var urlies Url
	tx := db.Where("urltext = ?", url).First(&urlies)
	return urlies, tx.Error
}
