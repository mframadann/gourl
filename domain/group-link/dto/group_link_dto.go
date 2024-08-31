package dto

type CreateGroupPayload struct {
	GroupName string `json:"name" validate:"required"`
}

type UpdateGroupPayload struct {
	ID uint `json:"group_id"`
	CreateGroupPayload
}

type DeleteGroupPalyload struct {
	ID uint `json:"group_id"`
}
