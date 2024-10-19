package form

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SignUpRes struct {
	MemberId int64 `json:"member_id"`
}

type UpdateRequest struct {
}
