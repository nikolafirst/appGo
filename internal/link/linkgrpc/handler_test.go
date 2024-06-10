package linkgrpc

import (
 "context"
 "errors"
 "testing"
 "time"

 "github.com/golang/mock/gomock"
 "github.com/stretchr/testify/assert"

 "appGo/internal/database/mocks"
 "appGo/pkg/pb"
)

func TestHandler_GetLinkByUserID(t *testing.T) {
 ctrl := gomock.NewController(t)
 defer ctrl.Finish()

 mockRepository := mocks.NewMocklinksRepository(ctrl)
 mockPublisher := mocks.NewMockamqpPublisher(ctrl)

 handler := New(mockRepository, time.Second, mockPublisher)

 // Создаем фэйковый контекст
 ctx := context.Background()

 // Создаем фэйковые данные
 userID := "user123"
 fakeLinks := []database.Link{
  {
   ID:        "link1",
   Title:     "Link 1",
   URL:       "https://example.com/link1",
   Images:    []string{"image1.jpg"},
   Tags:      []string{"tag1", "tag2"},
   UserID:    userID,
   CreatedAt: time.Now(),
   UpdatedAt: time.Now(),
  },
 }

 // Устанавливаем ожидания вызова метода FindByUserID у mockRepository
 mockRepository.EXPECT().FindByUserID(gomock.Any(), userID).Return(fakeLinks, nil)

 // Вызываем тестируемый метод
 resp, err := handler.GetLinkByUserID(ctx, &pb.GetLinksByUserId{UserId: userID})

 // Проверяем результаты
 assert.NoError(t, err)
 assert.NotNil(t, resp)
 assert.Equal(t, len(fakeLinks), len(resp.Links))
}

func TestHandler_CreateLink(t *testing.T) {
 ctrl := gomock.NewController(t)
 defer ctrl.Finish()

 mockRepository := mocks.NewMocklinksRepository(ctrl)
 mockPublisher := mocks.NewMockamqpPublisher(ctrl)

 handler := New(mockRepository, time.Second, mockPublisher)

 // Создаем фэйковый контекст
 ctx := context.Background()

 // Создаем фэйковые данные
 objectID := "object123"
 request := &pb.CreateLinkRequest{
  Id:     objectID,
  Url:    "https://example.com",
  Title:  "Example Link",
  Tags:   []string{"tag1", "tag2"},
  Images: []string{"image1.jpg"},
  UserId: "user123",
 }

 // Устанавливаем ожидания вызова метода Create у mockRepository
 mockRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

 // Устанавливаем ожидания вызова метода Publish у mockPublisher
 mockPublisher.EXPECT().Publish(gomock.Any(), gomock.Any(), false, false, gomock.Any()).Return(nil)

 // Вызываем тестируемый метод
 resp, err := handler.CreateLink(ctx, request)

 // Проверяем результаты
 assert.NoError(t, err)
 assert.NotNil(t, resp)
}

func TestHandler_GetLink(t *testing.T) {
 ctrl := gomock.NewController(t)
 defer ctrl.Finish()

 mockRepository := mocks.NewMocklinksRepository(ctrl)
 mockPublisher := mocks.NewMockamqpPublisher(ctrl)

 handler := New(mockRepository, time.Second, mockPublisher)

 // Создаем фэйковый контекст
 ctx := context.Background()

 // Создаем фэйковые данные
 objectID := "object123"
 fakeLink := database.Link{
  ID:        objectID,
  Title:     "Example Link",
  URL:       "https://example.com",
  Tags:      []string{"tag1", "tag2"},
  Images:    []string{"image1.jpg"},
  UserID:    "user123",
  CreatedAt: time.Now(),
  UpdatedAt: time.Now(),
 }

 // Устанавливаем ожидания вызова метода FindByID у mockRepository
 mockRepository.EXPECT().FindByID(gomock.Any(), objectID).Return(fakeLink, nil)

 // Вызываем тестируемый метод
 resp, err := handler.GetLink(ctx, &pb.GetLinkRequest{Id: objectID})

 // Проверяем результаты
 assert.NoError(t, err)
 assert.NotNil(t, resp)
 assert.Equal(t, objectID, resp.Id)
}

func TestHandler_UpdateLink(t *testing.T) {
 ctrl := gomock.NewController(t)
 defer ctrl.Finish()

 mockRepository := mocks.NewMocklinksRepository(ctrl)
 mockPublisher := mocks.NewMockamqpPublisher(ctrl)

 handler := New(mockRepository, time.Second, mockPublisher)

 // Создаем фэйковый контекст
 ctx := context.Background()

 // Создаем фэйковые данные
 objectID := "object123"
 request := &pb.UpdateLinkRequest{
  Id:     objectID,
  Url:    "https://example.com",
  Title:  "Updated Link",
  Tags:   []string{"tag1", "tag2"},
  Images: []string{"image1.jpg"},
  UserI
D: "user123",
 }

 // Устанавливаем ожидания вызова метода Update у mockRepository
 mockRepository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

 // Вызываем тестируемый метод
 resp, err := handler.UpdateLink(ctx, request)

 // Проверяем результаты
 assert.NoError(t, err)
 assert.NotNil(t, resp)
}