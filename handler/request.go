package handler

import "reflect"

type CreateOpeningReq struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   bool   `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}

type UpdateOpeningReq struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningReq) IsEmpty() bool {
	return reflect.DeepEqual(r, UpdateOpeningReq{})
}
