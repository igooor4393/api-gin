package models

type Program struct {
	Id        string `json:"id" binding:"required"`
	Name      string `json:"name"`
	NameEn    string `json:"nameEn"`
	IsPublic  bool   `json:"isPublic" binding:"required"`
	ProjectID string `json:"projectId" binding:"required"`
}
