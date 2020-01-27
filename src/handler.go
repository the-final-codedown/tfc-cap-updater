package main

import (
	"context"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto"
	"log"
)

type handler struct {
	repository
}

func (h *handler) DownscaleCap(ctx context.Context, req *cap.CapDownscale) (*cap.DownscaleResponse, error) {
	log.Println("DownscaleCap : ", req.AccountID, " ", req.Value)

	if err := h.repository.Downscale(req); err != nil {
		log.Println(err)
		return &cap.DownscaleResponse{
			Accepted:  false,
			Downscale: req,
		}, err
	}

	return &cap.DownscaleResponse{
		Accepted:  true,
		Downscale: req,
	}, nil
}

func (h *handler) UpscaleCap(ctx context.Context, req *cap.CapDownscale) (*cap.DownscaleResponse, error) {
	log.Println("UpscaleCap : ", req.AccountID, " ", req.Value)
	if err := h.repository.Upscale(req); err != nil {
		log.Println(err)
	}

	res := cap.DownscaleResponse{
		Accepted: true,
	}
	//println(res.Accepted)
	return &res, nil
}

// GetConsignments -
func (h *handler) GetDownscales(ctx context.Context, req *cap.GetRequest, res *cap.DownscaleResponse) error {
	println("GetDownscales")
	downscales, err := h.repository.GetAll()
	if err != nil {
		return err
	}
	res.Downscales = downscales
	return nil
}
