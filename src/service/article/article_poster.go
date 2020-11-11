package article

type ArticlePoster struct {
	PosterName string
	*Article
}

func NewArticlePoster(posterName string, article *Article) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article:    article,
	}
}

func GetPosterFlag() string {
	return "poster"
}
