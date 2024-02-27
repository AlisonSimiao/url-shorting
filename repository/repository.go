package repository

import (
	db "url-shorting/database"
	"gorm.io/gorm"
)

func (r *Repository) Super(Table string) {
	r.Table = Table
}

func (r *Repository) Update(condition string, values any, T any) {
	result := db.GetDatabase().Table(r.Table).Where(condition, values).Updates(T)
	if result.Error != nil {
		T = nil
	}
}

func (r *Repository) Delete(condition string, values any) error {
	result := db.GetDatabase().Table(r.Table).Where(condition, values).Delete(nil)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) Count(condition string, values any) (int, error) {
	var count int64
	result := db.GetDatabase().Table(r.Table).Where(condition, values).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *Repository) Paginate(condition string, values any, T any, page int, limit int) *PaginateData {
	totalEntities, err := r.Count(condition, values)

	if err != nil {
		return nil
	}

	result := db.GetDatabase().Table(r.Table).Where(condition, values).Offset((page - 1) * limit).Limit(limit).Find(T)

	if result.Error != nil {
		return nil
	}

	return &PaginateData{
		Data:          T,
		TotalEntities: totalEntities,
		CurrentPage:   page,
		TotalPages:    int(totalEntities / limit),
	}
}

func (r *Repository) PaginateWithJoin(selectQuery string, join string, condition string, values any, T any, page int, limit int) *PaginateData {

	totalEntities, err := r.Count(condition, values)

	if err != nil {
		return nil
	}

	result := db.GetDatabase().Table(r.Table).Select(selectQuery).Joins(join).Where(condition, values).Offset((page - 1) * limit).Limit(limit).Find(T)

	if result.Error != nil {
		return nil
	}

	return &PaginateData{
		Data:          T,
		TotalEntities: totalEntities,
		CurrentPage:   page,
		TotalPages:    int(totalEntities / limit),
	}
}

func Page(limit int, page int) *PageInfo {
	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	return &PageInfo{
		Limit: limit,
		Page:  page,
	}
}

func (r *Repository) Create(T any) {
	result := db.GetDatabase().Table(r.Table).Create(T)

	if result.Error != nil {
		T = nil
	}
}

func (r *Repository) FindOne(condition string, values map[string]interface{}, T any) {
	result := db.GetDatabase().Table(r.Table).Where(condition, values).First(T)

	if result.Error != nil {
		T = nil
	}
}

func (r *Repository) FindOneWithJoin(selectQuery string, join string, condition string, values any, T any) {

	result := db.GetDatabase().Table(r.Table).Select(selectQuery).Joins(join).Where(condition, values).Limit(1).First(T)

	if result.Error != nil {
		T = nil
	}
}

func (r *Repository) Raw() *gorm.DB {
	return db.GetDatabase().Table(r.Table)
}

func (r *Repository) FindAll(condition string, values any, T any) {
	if condition == "" {
		db.GetDatabase().Table(r.Table).Find(&T)
		return
	}

	result := db.GetDatabase().Table(r.Table).Find(&T, condition, values)
	if result.Error != nil {
		T = nil
	}
}
