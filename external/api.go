package external

import "github.com/p-jirayusakul/go-flat-arch-template/pkg/config"

type ExternalAPI interface {
	GetPosts() (string, error)
	CreatePost(p CreatePostParams) (string, error)
}

type APIs struct {
	cfg *config.Config
}

func New(cfg *config.Config) *APIs {
	return &APIs{
		cfg: cfg,
	}
}

var _ ExternalAPI = (*APIs)(nil)
