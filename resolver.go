package resolver

import "context"

type ResolveResult struct {
	StorageKey string
	StreamURL  string
	MIMEType   string
	Size       int64
	Seekable   bool
	Streamable bool
}

type AssetType string

const (
	AssetMedia     AssetType = "media"
	AssetPoster    AssetType = "poster"
	AssetBackdrop  AssetType = "backdrop"
	AssetThumbnail AssetType = "thumbnail"
	AssetSubtitle  AssetType = "subtitle"
)

type MediaResolver interface {
	Resolve(ctx context.Context, mediaID string, assetType AssetType) (ResolveResult, error)
}
