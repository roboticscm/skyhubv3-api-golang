package authentication

type CustomClaim struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"username"`
	FullName string `json:"fullName"`
}

type AccountInfo struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
