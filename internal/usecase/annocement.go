package usecase

import (
	"context"

	"github.com/kyh0703/go-project-layout/internal/domain"
)

type announcementUsecase struct{}

func (u announcementUsecase) CreateAnnouncement(
	ctx context.Context,
	chatID string, userID int, content string,
) (*domain.Announcement, error) {
	// if !validateContent(content) {
	// 	return nil, errors.New("invalid content")
	// }

	// writer, err := u.memberRepo.GetMember(ctx, userID)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := u.checkPermission(writer); err != nil {
	// 	return nil, err
	// }

	// anno, err := u.announcementRepo.CreateAnnouncement(ctx, chatID, writer, content)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (u announcementUsecase) GetAnnouncement(ctx context.Context, chatID string) (*domain.Announcement, error) {
	return &domain.Announcement{
		ChatID: chatID,
	}, nil
}
