package models

type CreatePostRequest struct {
	Items []*Program `json:"items" binding:"required"`
}
