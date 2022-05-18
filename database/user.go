package database

func GetShortenedsByAuthor(authorid string) ([]Shortened, error) {
	var shorteneds []Shortened
	if result := DB.Where("author = ?", authorid).Find(&shorteneds); result.Error != nil {
		return nil, result.Error
	}
	return shorteneds, nil
}
