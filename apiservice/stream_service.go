package apiservice

import (
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
)

// StreamService is the service to handle stream related requests
type StreamService struct {
	api.UnimplementedStreamServiceServer
}

// Supply returns the supply info, including total supply, circulating supply and exact circulating supply
// the API is streaming, so the client can get the supply info in real time
func (s *StreamService) Supply(req *api.SupplyRequest, res api.StreamService_SupplyServer) error {
	height := req.GetStartHeight()
	var totalCirculatingSupply, exactCirculatingSupply string
	for height <= req.GetEndHeight() {
		totalSupply, err := common.GetTotalSupply(height)
		if err != nil {
			return err
		}
		if req.GetIncludeCirculating() {
			totalCirculatingSupply, err = common.GetTotalCirculatingSupply(height, totalSupply)
			if err != nil {
				return err
			}
		}
		if req.GetIncludeExactCirculating() {
			exactCirculatingSupply, err = common.GetExactCirculatingSupply(height, totalSupply)
			if err != nil {
				return err
			}
		}
		if err := res.Send(&api.SupplyResponse{
			Height:                 height,
			TotalSupply:            totalSupply,
			CirculatingSupply:      totalCirculatingSupply,
			ExactCirculatingSupply: exactCirculatingSupply,
		}); err != nil {
			return err
		}
		height++
	}
	return nil
}
