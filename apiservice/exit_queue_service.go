package apiservice

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExitQueueService struct {
	api.UnimplementedExitQueueServiceServer
}

func (s *ExitQueueService) GetExitQueue(ctx context.Context, req *api.GetExitQueueRequest) (*api.GetExitQueueResponse, error) {
	skip := req.GetSkip()
	first := req.GetFirst()
	if first <= 0 {
		first = 20
	}
	if first > 100 {
		first = 100
	}

	query := db.DB().WithContext(ctx).Model(&model.CandidateExitQueue{})
	if statuses := req.GetStatuses(); len(statuses) > 0 {
		for _, st := range statuses {
			switch st {
			case "requested", "scheduled", "confirmed":
			default:
				return nil, status.Errorf(codes.InvalidArgument, "invalid status filter %q: must be requested, scheduled, or confirmed", st)
			}
		}
		query = query.Where("status IN ?", statuses)
	}
	if id := req.GetCandidateIdentity(); id != "" {
		query = query.Where("candidate_identity = ?", id)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count exit queue entries")
	}

	var rows []model.CandidateExitQueue
	if err := query.Order("id DESC").Offset(int(skip)).Limit(int(first)).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query exit queue")
	}

	entries := make([]*api.ExitQueueEntry, 0, len(rows))
	for _, r := range rows {
		entries = append(entries, &api.ExitQueueEntry{
			CandidateName:     r.CandidateName,
			CandidateIdentity: r.CandidateIdentity,
			Status:            r.Status,
			RequestHeight:     r.RequestHeight,
			RequestHash:       r.RequestHash,
			ScheduleHeight:    r.ScheduleHeight,
			ScheduleHash:      r.ScheduleHash,
			ConfirmHeight:     r.ConfirmHeight,
			ConfirmHash:       r.ConfirmHash,
			ScheduledAt:       r.ScheduledAt,
		})
	}

	return &api.GetExitQueueResponse{
		Exits: entries,
		Count: total,
	}, nil
}
