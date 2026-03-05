package service

import (
	"context"
	events "cookdie/events/internal/models"
	"cookdie/events/internal/producer"
	sql "cookdie/menu/sql/db/query_gen"
	"time"
)

type EventService struct {
	Queue producer.Queue
}

type eventService struct {
}

type EventServiceConfig struct {
	store sql.Queries
	Queue producer.Queue
}

func NewEventService(cfg *EventServiceConfig) *EventService {
	return &EventService{
		Queue: cfg.Queue,
	}
}

func (s *EventService) PostReviewRestaurant(review *events.Review) (*events.Review, error) {
	event := &events.Event{
		Type:       events.EventTypeReview,
		UserId:     review.UserId,
		BusinessId: review.BusinessId,
		Date:       "",
		Payload:    review,
	}
	err := s.Queue.Publish(context.Background(), *event)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *EventService) CheckinRestaurant(businessId string, userId string, date time.Time) (*events.Checkin, error) {
	today := time.Now()
	event := &events.Event{
		Type:       events.EventTypeCheckin,
		UserId:     userId,
		BusinessId: businessId,
		Date:       today.String(),
	}
	err := s.Queue.Publish(context.Background(), *event)
	if err != nil {
		return nil, err
	}
	return &events.Checkin{
		BusinessId: businessId,
		UserId:     userId,
		Date:       today.String(),
	}, nil
}
