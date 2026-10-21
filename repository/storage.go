package repository

import (
	"github.com/tigerowo/infinite-canvas/model"
)

// ListStorageObjects 只查询系统存储提供商的已同步文件。
func ListStorageObjects(q model.Query, providerIDs []string) (model.StorageObjectList, error) {
	db, err := DB()
	if err != nil {
		return model.StorageObjectList{}, err
	}
	q.Normalize()
	tx := db.Table("storage_objects AS objects").
		Joins("LEFT JOIN users ON users.id = objects.created_by").
		Joins("LEFT JOIN assets ON assets.storage_object_id = objects.id").
		Where("objects.provider_id IN ? AND objects.deleted_at = ?", providerIDs, "")
	if q.Keyword != "" {
		like := "%" + q.Keyword + "%"
		tx = tx.Where("objects.object_key LIKE ? OR objects.created_by LIKE ? OR users.username LIKE ? OR users.display_name LIKE ?", like, like, like, like)
	}
	if q.Type != "" {
		tx = tx.Where("objects.mime_type LIKE ?", q.Type+"/%")
	}
	result := model.StorageObjectList{Items: []model.AdminStorageObject{}}
	if err := tx.Count(&result.Total).Error; err != nil {
		return result, err
	}
	err = tx.Select("objects.*, COALESCE(users.username, '') AS username, COALESCE(users.display_name, '') AS user_display_name, COALESCE(assets.id, '') AS asset_id").
		Order("objects.created_at DESC, objects.id DESC").Offset(q.Offset()).Limit(q.PageSize).Scan(&result.Items).Error
	return result, err
}

// SaveStorageObject 保存存储对象记录。
func SaveStorageObject(object model.StorageObject) (model.StorageObject, error) {
	db, err := DB()
	if err != nil {
		return model.StorageObject{}, err
	}
	return object, db.Save(&object).Error
}

// GetStorageObject 根据 ID 获取存储对象。
func GetStorageObject(id string) (model.StorageObject, error) {
	db, err := DB()
	if err != nil {
		return model.StorageObject{}, err
	}
	var object model.StorageObject
	err = db.First(&object, "id = ?", id).Error
	return object, err
}

// DeleteStorageObjectRecord 删除存储对象记录（软删除）。
func DeleteStorageObjectRecord(id string) error {
	db, err := DB()
	if err != nil {
		return err
	}
	return db.Delete(&model.StorageObject{}, "id = ?", id).Error
}
