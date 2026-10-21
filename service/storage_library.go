package service

import (
	"path"
	"strings"

	"github.com/tigerowo/infinite-canvas/model"
	"github.com/tigerowo/infinite-canvas/repository"
)

func ListAdminStorageObjects(q model.Query) (model.StorageObjectList, error) {
	settings, err := repository.GetSettings()
	if err != nil {
		return model.StorageObjectList{}, err
	}
	providers := normalizePrivateStorageSetting(settings.Private.Storage).Providers
	ids := make([]string, 0, len(providers))
	names := make(map[string]string, len(providers))
	for _, provider := range providers {
		ids = append(ids, provider.ID)
		names[provider.ID] = provider.Name
	}
	result, err := repository.ListStorageObjects(q, ids)
	for i := range result.Items {
		item := &result.Items[i]
		item.ProviderName = names[item.ProviderID]
		item.URL = storageLibraryURL(item.ID)
	}
	return result, err
}

func PublishStorageObject(id string) (model.Asset, error) {
	object, err := repository.GetStorageObject(id)
	if err != nil {
		return model.Asset{}, err
	}
	settings, err := repository.GetSettings()
	if err != nil {
		return model.Asset{}, err
	}
	global := false
	for _, provider := range normalizePrivateStorageSetting(settings.Private.Storage).Providers {
		if provider.ID == object.ProviderID {
			global = true
			break
		}
	}
	if !global || object.Direct || object.DeletedAt != "" {
		return model.Asset{}, safeMessageError{message: "只能发布系统对象存储中已同步的媒体"}
	}
	kind := model.AssetType(strings.SplitN(object.MimeType, "/", 2)[0])
	if kind != model.AssetTypeImage && kind != model.AssetTypeVideo && kind != model.AssetTypeAudio {
		return model.Asset{}, safeMessageError{message: "仅支持发布图片、视频和音频"}
	}
	url := storageLibraryURL(object.ID)
	item := model.Asset{
		ID: "storage-" + object.ID, StorageObjectID: object.ID,
		Title: path.Base(object.ObjectKey), Type: kind, URL: url,
		Tags: []string{}, CreatedAt: now(), UpdatedAt: now(),
	}
	if kind == model.AssetTypeImage {
		item.CoverURL = url
	}
	return repository.PublishStorageAsset(item)
}

func storageLibraryURL(id string) string {
	return "/api/files/" + id + "/content"
}
