package dto

type CreateLinkPayload struct {
	GroupID    *uint  `json:"group_id" `
	Title      string `json:"link_title"  validate:"required"`
	ShortedURL string `json:"shorted_url"`
	OriginURL  string `json:"origin_url" validate:"required"`
}

type GetLinkQueries struct {
	OrderByGrup bool `query:"order_by_group"`
}

type UpdateLinkPayload struct {
	ID uint `json:"link_id"`
	CreateLinkPayload
}

type DeleteLinkPalyload struct {
	ID uint `json:"link_id"`
}
