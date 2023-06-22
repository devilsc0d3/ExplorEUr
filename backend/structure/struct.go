package structure

import "exploreur/backend/register"

type DataHub struct {
	Role                 string
	Database             []Posts
	Category             []string
	IsConnected          bool
	ReportPostContent    string
	ReportCommentContent string
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
var CatId int

func ManageData(content []string, postId []int, message []string, postIdComment []int, userId []int, userIdComment []int, postIdLike []int) []Posts {
	var database []Posts
	countLike := CountLike(postIdLike)
	for i := 0; i < len(content); i++ {
		var temp Posts
		temp.Content = content[i]
		temp.Id = postId[i]
		temp.UserId = userId[i]
		temp.NicknamePost, _ = register.GetNicknameByID(userId[i])

		for idPost, nbrLike := range countLike {
			if idPost == postId[i] {
				temp.CountLike = nbrLike
			}
		}
		database = append(database, temp)
	}

	for i := 0; i < len(message); i++ {
		var temp2 Comment
		for j := 0; j < len(database); j++ {
			if postIdComment[i] == database[j].Id {
				temp2.Message = message[i]
				temp2.PostId = postIdComment[i]
				temp2.NicknameComment, _ = register.GetNicknameByID(userIdComment[i])
				database[j].Comments = append(database[j].Comments, temp2)
				break
			}
		}
	}

	return database
}

func CountLike(postId []int) map[int]int {
	counts := make(map[int]int)

	for _, idPost := range postId {
		counts[idPost]++
	}

	return counts
}

func InitRole(token string) {
	_, connectedRole, err := register.DecodeJWTToken(token)
	if err != nil {
		panic("decode token error")
	}
	switch connectedRole {
	case "":
		DataHub1.Role = ""
		break
	case "user":
		DataHub1.Role = "user"
		break
	case "moderator":
		DataHub1.Role = "moderator"
		break
	case "admin":
		DataHub1.Role = "admin"
		break
	}
}
