package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahServiceService) CreateFeedConfig(
	ctx context.Context,
	req *pb.CreateFeedConfigRequest,
) (*pb.CreateFeedConfigResponse, error) {
	id, err := s.y.CreateFeedConfig(ctx, converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateFeedConfigResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateFeedConfig(
	ctx context.Context,
	req *pb.UpdateFeedConfigRequest,
) (*pb.UpdateFeedConfigResponse, error) {
	err := s.y.UpdateFeedConfig(ctx, converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFeedConfigResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListFeedConfigs(
	ctx context.Context,
	req *pb.ListFeedConfigsRequest,
) (*pb.ListFeedConfigsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	feeds, total, err := s.y.ListFeeds(ctx,
		model.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		converter.ToBizInternalIDList(req.GetIdFilter()),
		converter.ToBizInternalIDList(req.GetAuthorIdFilter()),
		converter.ToBizFeedConfigSourceList(req.GetSourceFilter()),
		converter.ToBizFeedConfigStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedConfigsResponse{
		Paging:          &librarian.PagingResponse{TotalSize: int64(total)},
		FeedsWithConfig: converter.ToPBFeedWithConfigList(feeds),
	}, nil
}
func (s *LibrarianSephirahServiceService) ListFeedItems(
	ctx context.Context,
	req *pb.ListFeedItemsRequest,
) (*pb.ListFeedItemsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	items, total, err := s.y.ListFeedItems(ctx,
		model.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		converter.ToBizInternalIDList(req.GetFeedIdFilter()),
		converter.ToBizInternalIDList(req.GetAuthorIdFilter()),
		req.GetPublishPlatformFilter(),
		converter.ToBizTimeRange(req.GetPublishTimeRange()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedItemsResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Items:  converter.ToPBItemIDWithFeedIDList(items),
	}, nil
}
func (s *LibrarianSephirahServiceService) GroupFeedItems(
	ctx context.Context,
	req *pb.GroupFeedItemsRequest,
) (*pb.GroupFeedItemsResponse, error) {
	itemMap, err := s.y.GroupFeedItems(ctx,
		converter.ToBizGroupFeedItemsBy(req.GetGroupBy()),
		converter.ToBizInternalIDList(req.GetFeedIdFilter()),
		converter.ToBizInternalIDList(req.GetAuthorIdFilter()),
		req.GetPublishPlatformFilter(),
		converter.ToBizTimeRange(req.GetPublishTimeRange()),
		int(req.GetGroupSize()),
	)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.GroupFeedItemsResponse_FeedItemsGroup, 0, len(itemMap))
	for timeRange, items := range itemMap {
		res = append(res, &pb.GroupFeedItemsResponse_FeedItemsGroup{
			TimeRange: converter.ToPBTimeRange(&timeRange),
			Items:     converter.ToPBItemIDWithFeedIDList(items),
		})
	}
	return &pb.GroupFeedItemsResponse{Groups: res}, nil
}
func (s *LibrarianSephirahServiceService) GetFeedItem(
	ctx context.Context,
	req *pb.GetFeedItemRequest,
) (*pb.GetFeedItemResponse, error) {
	item, err := s.y.GetFeedItem(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetFeedItemResponse{
		Item: converter.ToPBFeedItem(item),
	}, nil
}

func (s *LibrarianSephirahServiceService) GetBatchFeedItems(
	ctx context.Context,
	req *pb.GetBatchFeedItemsRequest,
) (*pb.GetBatchFeedItemsResponse, error) {
	items, err := s.y.GetFeedItems(ctx, converter.ToBizInternalIDList(req.GetIds()))
	if err != nil {
		return nil, err
	}
	return &pb.GetBatchFeedItemsResponse{
		Items: converter.ToPBFeedItemList(items),
	}, nil
}
