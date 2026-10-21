package service

import (
	"strings"
	"time"

	"github.com/tigerowo/infinite-canvas/model"
	"github.com/tigerowo/infinite-canvas/repository"
)

func ListAssets(q model.Query, featured bool) (model.AssetList, error) {
	items, total, err := repository.ListAssets(q, featured)
	if err != nil {
		return model.AssetList{}, err
	}
	tags, categories, err := repository.ListAssetFilters(q)
	if err != nil {
		return model.AssetList{}, err
	}
	return model.AssetList{Items: items, Tags: tags, Categories: categories, Total: int(total)}, nil
}

func SaveAsset(item model.Asset) (model.Asset, error) {
	now := time.Now().Format(time.RFC3339)
	if item.Type == "" {
		item.Type = model.AssetTypeText
	}
	item.Title = strings.TrimSpace(item.Title)
	if item.Title == "" {
		return item, safeMessageError{message: "请输入素材标题"}
	}
	switch item.Type {
	case model.AssetTypeText:
		if strings.TrimSpace(item.Content) == "" {
			return item, safeMessageError{message: "请输入文本内容"}
		}
	case model.AssetTypeImage, model.AssetTypeVideo, model.AssetTypeAudio:
		if strings.TrimSpace(item.URL) == "" {
			return item, safeMessageError{message: "请输入素材链接"}
		}
	default:
		return item, safeMessageError{message: "不支持的素材类型"}
	}
	if item.ID == "" {
		item.ID = newID("asset")
		item.CreatedAt = now
	}
	item.UpdatedAt = now
	if item.CoverURL == "" {
		item.CoverURL = assetCoverURL(item)
	}
	return repository.SaveAsset(item)
}

func DeleteAsset(id string) error {
	return repository.DeleteAsset(id)
}

func assetCoverURL(item model.Asset) string {
	if item.CoverURL != "" {
		return item.CoverURL
	}
	if item.Type == model.AssetTypeImage {
		return item.URL
	}
	return ""
}
