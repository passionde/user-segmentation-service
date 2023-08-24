package service

import "github.com/passionde/user-segmentation-service/internal/repo"

type SegmentService struct {
	segmentRepo repo.Segment
}

func NewSegmentService(segmentRepo repo.Segment) *SegmentService {
	return &SegmentService{segmentRepo: segmentRepo}
}

// todo: реализовать методы сервиса
