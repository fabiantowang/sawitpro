package repository

import (
	"context"
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	userInput1 := AddUserInput{
		Phone:    "12345678901234",
		Fullname: "ABC",
		Salt:     "ABC",
		Password: "ABC",
	}
	userInput2 := AddUserInput{
		Phone:    "12345",
		Fullname: "ABC",
		Salt:     "ABC",
		Password: "ABC",
	}
	mockUuid := uuid.New()
	gomock.InOrder(
		mockRepo.EXPECT().AddUser(context.TODO(), userInput1).Return(AddUserOutput{}, fmt.Errorf("%v", "insert error")).Times(1),
		mockRepo.EXPECT().AddUser(context.TODO(), userInput2).Return(AddUserOutput{Id: mockUuid}, nil).Times(1),
	)

	_, err := mockRepo.AddUser(context.TODO(), userInput1)
	assert.NotNil(t, err)

	_, err = mockRepo.AddUser(context.TODO(), userInput2)
	assert.Nil(t, err)
}

func TestRepository_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	userOutput := GetUserByIdOutput{
		Phone:    "12345678901234",
		Fullname: "ABC",
	}
	mockUuid := uuid.New()
	gomock.InOrder(
		mockRepo.EXPECT().GetUserById(context.TODO(), mockUuid.String()).Return(userOutput, nil).Times(1),
	)

	queryResult, err := mockRepo.GetUserById(context.TODO(), mockUuid.String())
	assert.Nil(t, err)
	assert.Equal(t, queryResult, userOutput)
}

func TestRepository_GetUserByPhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	mockUuid := uuid.New()
	userOutput := GetUserByPhoneOutput{
		Id:       mockUuid,
		Salt:     "ZOTP/quFNP4zkphXRTOBBQ",
		Password: "CiU9WfO7x7LwfVXs1af0SDxCy0eDIIWZ13hOb7+kNwo",
	}
	userPhone := "+621234567"
	gomock.InOrder(
		mockRepo.EXPECT().GetUserByPhone(context.TODO(), userPhone).Return(userOutput, nil).Times(1),
	)

	queryResult, err := mockRepo.GetUserByPhone(context.TODO(), userPhone)
	assert.Nil(t, err)
	assert.Equal(t, queryResult, userOutput)
}

func TestRepository_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	mockUuid := uuid.New()
	userInput := UpdateUserInput{
		Id:       mockUuid,
		Phone:    "+62123456789",
		Fullname: "John Smith",
	}
	gomock.InOrder(
		mockRepo.EXPECT().UpdateUser(context.TODO(), userInput).Return(nil).Times(1),
	)

	err := mockRepo.UpdateUser(context.TODO(), userInput)
	assert.Nil(t, err)
}

func TestRepository_IncrementSuccessfulLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepositoryInterface(ctrl)
	gomock.InOrder(
		mockRepo.EXPECT().IncrementSuccessfulLogin(context.TODO(), "123").Return(nil).Times(1),
	)

	err := mockRepo.IncrementSuccessfulLogin(context.TODO(), "123")
	assert.Nil(t, err)
}
