package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id "`
	FullName      string             `json:"full_name" validate:"required,min=2,max=20"`
	Phone         string             `json:"Phone" validate:"required,min=2,max=20"`
	Email         string             `json:"email" validate:"email , required"`
	Password      string             `json:"password" validate:"required,min=6"`
	User_type     string             `json:"user_type" validate:"required, eq=ADMIN|eq=USER`
	Token         string             `json:"token_id" `
	Refresh_token string             `json:"refresh_token"`
}
