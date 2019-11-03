package main

import (
	"context"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
)

type handler struct {
	repository
}

func (s *handler) DownscaleCap(ctx context.Context, req *cap.CapDownscale) (*cap.DownscaleResponse, error) {
	println("DownscaleCap : ", req.AccountID, " ", req.Value)
	if err := s.repository.Create(req); err != nil {
		print(err)
	}

	res := cap.DownscaleResponse{
		Accepted: true,
	}
	//println(res.Accepted)
	return &res, nil
}

// GetConsignments -
func (s *handler) GetDownscales(ctx context.Context, req *cap.GetRequest, res *cap.DownscaleResponse) error {
	println("GetDownscales")
	downscales, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	res.Downscales = downscales
	return nil
}
