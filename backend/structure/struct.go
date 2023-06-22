package structure

type DataHub struct {
	Role        string
	Database    []Posts
	Category    []string
	IsConnected bool
}

type Posts struct {
	Content      string
	Id           int
	Comments     []Comment
	UserId       int
	NicknamePost string
	CountLike    int
}

type Comment struct {
	Message         string
	PostId          int
	NicknameComment string
}

var DataHub1 DataHub
