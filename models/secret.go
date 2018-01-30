package models

// User is info
type User struct {
	NickName  string `json:"nickName"`
	AvatarURL string `json:"avatarUrl"`
}

// Secret is
type Secret struct {
	Key  string `json:"key"`
	UKey string `json:"ukey"`
	User `json:"user"`
}
