package helper

import "cashbookTeam/models"

func MaxIdUser(UserArray []models.UserModel) int {
	var maxId = 0
	for i := 0; i < len(UserArray); i++ {
		if maxId < UserArray[i].Id {
			maxId=UserArray[i].Id
		}
	}
	return maxId+1
}
func MaxIdCard(PostArray []models.CardModel) int {
	var maxId = 0
	for i := 0; i < len(PostArray); i++ {
		if maxId < PostArray[i].Id {
			maxId=PostArray[i].Id
		}
	}
	return maxId+1
}
func MaxIdTransaction(CommentArray []models.TransactionModel) int {
	var maxId = 0
	for i := 0; i < len(CommentArray); i++ {
		if maxId < CommentArray[i].Id {
			maxId=CommentArray[i].Id
		}
	}
	return maxId+1
}