package model

// StorageObject 存储对象（S3/R2、WebDAV 共用文件索引）。
type StorageObject struct {
	ID         string `json:"id" gorm:"primaryKey"`
	ProviderID string `json:"providerId" gorm:"index"`
	Bucket     string `json:"bucket"`
	ObjectKey  string `json:"objectKey" gorm:"uniqueIndex"`
	PublicURL  string `json:"publicUrl"`
	MimeType   string `json:"mimeType"`
	Bytes      int64  `json:"bytes"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	SHA256     string `json:"sha256"`
	Direct     bool   `json:"direct"`
	CreatedBy  string `json:"createdBy" gorm:"index"`
	CreatedAt  string `json:"createdAt"`
	DeletedAt  string `json:"deletedAt"`
}

// AdminStorageObject 不含存储凭据，仅供后台查看同步记录。
type AdminStorageObject struct {
	StorageObject
	Username        string `json:"username" gorm:"->"`
	UserDisplayName string `json:"userDisplayName" gorm:"->"`
	AssetID         string `json:"assetId" gorm:"->"`
	ProviderName    string `json:"providerName" gorm:"-"`
	URL             string `json:"url" gorm:"-"`
}

type StorageObjectList struct {
	Items []AdminStorageObject `json:"items"`
	Total int64                `json:"total"`
}
