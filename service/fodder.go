package service

import (
	"Waste2Feed/constants"
	"Waste2Feed/dto"
	"Waste2Feed/entity"
	"Waste2Feed/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	FodderService interface {
		Order(context.Context, string, string, dto.FodderRequest) (dto.FodderResponse, error)
		Payment(context.Context, string, dto.FodderPaymentRequest) (dto.FodderResponse, error)
		GetPendingPayments(context.Context, string) ([]dto.FodderResponse, error)
		GetPendingDetails(context.Context, string, string) (dto.FodderResponse, error)
	}

	fodderService struct {
		fodderRepo repository.FodderRepository
	}
)

func NewFodderService(fr repository.FodderRepository) FodderService {
	return &fodderService{
		fodderRepo: fr,
	}
}

func (s *fodderService) Payment(ctx context.Context, userID string, req dto.FodderPaymentRequest) (dto.FodderResponse, error) {
	fodder, err := s.fodderRepo.GetByID(req.ID)
	if err != nil {
		return dto.FodderResponse{}, err
	}

	if fodder.FarmerID != userID {
		return dto.FodderResponse{}, dto.ErrFodderNotOwnedByRequester
	}

	if fodder.Paid {
		return dto.FodderResponse{}, dto.ErrFodderAlreadyPaid
	}

	fodder.Paid = true

	fodder, err = s.fodderRepo.Update(fodder)
	if err != nil {
		return dto.FodderResponse{}, err
	}

	resp := dto.FodderResponse{
		ID:          fodder.ID.String(),
		FarmerID:    fodder.FarmerID,
		Date:        fodder.Date,
		Time:        fodder.Time,
		Address:     fodder.Address,
		Coordinates: fodder.Coordinates,
		Amount:      fodder.Amount,
		Status:      fodder.Status,
		Type:        fodder.Type,
		Paid:        fodder.Paid,
	}

	if fodder.CourierID != nil {
		resp.CourierID = fodder.CourierID.String()
	} else {
		resp.CourierID = ""
	}

	return resp, nil
}

func (s *fodderService) Order(ctx context.Context, farmerID string, _type string, req dto.FodderRequest) (dto.FodderResponse, error) {
	if _type != constants.ENUM_FODDER_TYPE_CHICKEN && _type != constants.ENUM_FODDER_TYPE_FISH {
		return dto.FodderResponse{}, dto.ErrInvalidFodderType
	}

	tDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return dto.FodderResponse{}, err
	}

	tTime, err := time.Parse("15:04", req.Time)
	if err != nil {
		return dto.FodderResponse{}, err
	}

	_date := tDate.Format("2006-01-02")
	_time := tTime.Format("15:04")

	fodder := entity.Fodder{
		ID:          uuid.New(),
		FarmerID:    farmerID,
		Date:        _date,
		Time:        _time,
		Address:     req.Address,
		Coordinates: req.Coordinates,
		Amount:      req.Amount,
		Status:      constants.ENUM_PICKUP_STATUS_PROGRESS,
		Type:        _type,
	}

	fodder, err = s.fodderRepo.Create(fodder)
	if err != nil {
		return dto.FodderResponse{}, err
	}

	return dto.FodderResponse{
		ID:          fodder.ID.String(),
		FarmerID:    fodder.FarmerID,
		Date:        fodder.Date,
		Time:        fodder.Time,
		Address:     fodder.Address,
		Coordinates: fodder.Coordinates,
		Amount:      fodder.Amount,
		Status:      fodder.Status,
		Type:        fodder.Type,
	}, nil
}

func (s *fodderService) GetPendingPayments(ctx context.Context, farmerID string) ([]dto.FodderResponse, error) {
	fodders, err := s.fodderRepo.GetAllUnpaid(farmerID)
	if err != nil {
		return nil, err
	}

	var responses []dto.FodderResponse
	for _, fodder := range fodders {
		tmp := dto.FodderResponse{
			ID:          fodder.ID.String(),
			FarmerID:    fodder.FarmerID,
			Date:        fodder.Date,
			Time:        fodder.Time,
			Address:     fodder.Address,
			Coordinates: fodder.Coordinates,
			Amount:      fodder.Amount,
			Status:      fodder.Status,
			Type:        fodder.Type,
			Paid:        fodder.Paid,
		}

		if fodder.CourierID != nil {
			tmp.CourierID = fodder.CourierID.String()
		} else {
			tmp.CourierID = ""
		}

		responses = append(responses, tmp)
	}

	return responses, nil
}

func (s *fodderService) GetPendingDetails(ctx context.Context, userID, id string) (dto.FodderResponse, error) {
	fodder, err := s.fodderRepo.GetByID(id)
	if err != nil {
		return dto.FodderResponse{}, err
	}

	if fodder.FarmerID != userID {
		return dto.FodderResponse{}, dto.ErrFodderNotOwnedByRequester
	}

	resp := dto.FodderResponse{
		ID:          fodder.ID.String(),
		FarmerID:    fodder.FarmerID,
		Date:        fodder.Date,
		Time:        fodder.Time,
		Address:     fodder.Address,
		Coordinates: fodder.Coordinates,
		Amount:      fodder.Amount,
		Status:      fodder.Status,
		Type:        fodder.Type,
		Paid:        fodder.Paid,
	}

	if fodder.CourierID != nil {
		resp.CourierID = fodder.CourierID.String()
	} else {
		resp.CourierID = ""
	}

	return resp, nil
}
