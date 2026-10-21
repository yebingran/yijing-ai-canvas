package service

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/tigerowo/infinite-canvas/config"
	"github.com/tigerowo/infinite-canvas/model"
	"github.com/tigerowo/infinite-canvas/repository"
)

func TestStorageLibraryLifecycle(t *testing.T) {
	config.Cfg.StorageDriver = "sqlite"
	config.Cfg.DatabaseDSN = filepath.Join(t.TempDir(), "library.db")
	settings := model.Settings{}
	settings.Private.Storage.Providers = []model.StorageProvider{{ID: "system-oss", Name: "新加坡 OSS", Type: "s3"}}
	if _, err := repository.SaveSettings(settings, now()); err != nil {
		t.Fatal(err)
	}
	user := model.User{ID: "creator", Username: "artist", DisplayName: "素材创作者"}
	if _, err := repository.SaveUser(user); err != nil {
		t.Fatal(err)
	}
	for _, object := range []model.StorageObject{
		{ID: "picture", ProviderID: "system-oss", ObjectKey: "creator/picture.png", MimeType: "image/png", CreatedBy: user.ID, CreatedAt: now()},
		{ID: "movie", ProviderID: "system-oss", ObjectKey: "creator/movie.mp4", MimeType: "video/mp4", CreatedBy: user.ID, CreatedAt: now()},
		{ID: "private", ProviderID: "user-oss", ObjectKey: "creator/private.png", MimeType: "image/png", CreatedBy: user.ID},
		{ID: "document", ProviderID: "system-oss", ObjectKey: "doc.pdf", MimeType: "application/pdf"},
		{ID: "deleted", ProviderID: "system-oss", ObjectKey: "deleted.png", MimeType: "image/png", DeletedAt: now()},
	} {
		if _, err := repository.SaveStorageObject(object); err != nil {
			t.Fatal(err)
		}
	}
	list, err := ListAdminStorageObjects(model.Query{Keyword: "artist", PageSize: 1})
	if err != nil || list.Total != 2 || len(list.Items) != 1 || list.Items[0].UserDisplayName != "素材创作者" || list.Items[0].ProviderName != "新加坡 OSS" {
		t.Fatalf("user search and pagination: %#v, %v", list, err)
	}
	list, err = ListAdminStorageObjects(model.Query{Type: "image"})
	if err != nil || list.Total != 1 || list.Items[0].ID != "picture" {
		t.Fatalf("system provider and media filter: %#v, %v", list, err)
	}
	public, err := ListAssets(model.Query{}, false)
	if err != nil || public.Total != 0 {
		t.Fatalf("unpublished uploads must stay out of public library: %#v, %v", public, err)
	}
	for _, id := range []string{"private", "document", "deleted"} {
		if _, err := PublishStorageObject(id); err == nil {
			t.Fatalf("publishing %s must be rejected", id)
		}
	}
	asset, err := PublishStorageObject("picture")
	if err != nil || asset.URL != "/api/files/picture/content" || asset.Featured {
		t.Fatalf("publish: %#v, %v", asset, err)
	}
	asset.Title, asset.Featured = "首页精选", true
	asset.StorageObjectID = "" // API 编辑不会提交内部来源字段。
	asset, err = SaveAsset(asset)
	if err != nil || asset.StorageObjectID != "picture" {
		t.Fatalf("editing must preserve source: %#v, %v", asset, err)
	}
	again, err := PublishStorageObject("picture")
	if err != nil || again.ID != asset.ID || again.Title != "首页精选" || !again.Featured {
		t.Fatalf("repeat publishing must preserve edits: %#v, %v", again, err)
	}
	if _, err := PublishStorageObject("movie"); err != nil {
		t.Fatal(err)
	}
	featured, err := ListAssets(model.Query{}, true)
	if err != nil || featured.Total != 1 || featured.Items[0].ID != asset.ID {
		t.Fatalf("homepage selection: %#v, %v", featured, err)
	}
	public, err = ListAssets(model.Query{}, false)
	if err != nil || public.Total != 2 {
		t.Fatalf("public library: %#v, %v", public, err)
	}
	list, err = ListAdminStorageObjects(model.Query{Type: "image"})
	if err != nil || list.Items[0].AssetID != asset.ID {
		t.Fatalf("published status: %#v, %v", list, err)
	}
	if err := DeleteStorageObject(context.Background(), "picture", nil); err == nil {
		t.Fatal("published files must survive automatic storage cleanup")
	}
	asset.Featured = false
	if _, err := SaveAsset(asset); err != nil {
		t.Fatal(err)
	}
	featured, err = ListAssets(model.Query{}, true)
	if err != nil || featured.Total != 0 {
		t.Fatalf("unfeature: %#v, %v", featured, err)
	}
	if err := DeleteAsset(asset.ID); err != nil {
		t.Fatal(err)
	}
	if published, err := repository.StorageObjectPublished("picture"); err != nil || published {
		t.Fatalf("removing a public asset must release storage protection: %v, %v", published, err)
	}
}
