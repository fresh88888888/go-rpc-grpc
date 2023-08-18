package services

import (
	context "context"
	"fmt"
	"time"
)

type UserService struct {
}

func (*UserService) GetUserScoreByServerStream(request *UserScoreRequest, strem UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for index, user := range request.Users {
		user.UserScore = score
		score++
		users = append(users, user)

		if index > 0 && (index+1)%2 == 0 {
			//fmt.Println(users)
			err := strem.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			users = (users)[0:0]
		}
		time.Sleep(time.Second * 1)
	}
	if len(users) > 0 {
		strem.Send(&UserScoreResponse{Users: users})
	}

	return nil
}

func (*UserService) mustEmbedUnimplementedUserServiceServer() {

}

func (*UserService) GetUserScore(ctx context.Context, request *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	fmt.Println(request.Users)
	for _, user := range request.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}

	return &UserScoreResponse{Users: users}, nil
}
