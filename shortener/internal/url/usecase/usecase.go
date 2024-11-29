package usecase

type Resposiotry interface {
}

type URLUsecase struct {
	repo Resposiotry
}

func NewURLUsecase(repo Resposiotry) *URLUsecase {
	return &URLUsecase{
		repo: repo,
	}
}
