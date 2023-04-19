package models

type KworkResponseModel struct {
	ID          int    `json:"id"`
	Lang        string `json:"lang"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	URL         string `json:"url"`
	Files       []struct {
		Name string `json:"name"`
		Url  string `json: "url"`
	} `json:"files"`
	IsActive           bool   `json:"isActive"`
	IsHigherPrice      bool   `json:"isHigherPrice"`
	IsSymbolRu         bool   `json:"isSymbolRu"`
	CategoryMinPrice   int    `json:"categoryMinPrice"`
	PriceLimit         string `json:"priceLimit"`
	PossiblePriceLimit int    `json:"possiblePriceLimit"`
	DateView           any    `json:"dateView"`
	DateExpire         string `json:"dateExpire"`
	DateCreate         string `json:"dateCreate"`
	DateExpireText     string `json:"dateExpireText"`
	DateCreateText     string `json:"dateCreateText"`
	KworkCount         string `json:"kworkCount"`
	ProjectReviewType  any    `json:"projectReviewType"`
	GetReviewCanChange bool   `json:"getReviewCanChange"`
	TimeLeft           string `json:"timeLeft"`
	Turnover           int    `json:"turnover"`
	IsUserWant         bool   `json:"isUserWant"`
	UserID             int    `json:"userId"`
	UserName           string `json:"userName"`
	UserAvatar         string `json:"userAvatar"`
	UserAvatarSrcSet   string `json:"userAvatarSrcSet"`
	UserBackground     string `json:"userBackground"`
	UserIsOnline       bool   `json:"userIsOnline"`
	UserBadges         []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		PayerLevel   any    `json:"payer_level"`
		IsSuperPayer bool   `json:"is_super_payer"`
		ImageURL     string `json:"image_url"`
	} `json:"userBadges"`
	UserActiveWants        int    `json:"userActiveWants"`
	UserWants              string `json:"userWants"`
	UserIsOtherActiveWants bool   `json:"userIsOtherActiveWants"`
	UserWantsHiredPercent  string `json:"userWantsHiredPercent"`
	UserAlreadyWork        any    `json:"userAlreadyWork"`
	CurrentUserReviewType  any    `json:"currentUserReviewType"`
	CategoryName           string `json:"categoryName"`
	ParentCategoryName     string `json:"parentCategoryName"`
	CompetitionLevel       any    `json:"competitionLevel"`
}
