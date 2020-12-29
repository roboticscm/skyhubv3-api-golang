package models

type RefreshToken struct {
	Id        int64  `json:"id"`
	Token     *string `json:"token" orm:"null"`
	AccountId *int64 `json:"accountId" orm:"null"`
	CreatedAt *int64 `json:"createdAt" orm:"null"`
}
