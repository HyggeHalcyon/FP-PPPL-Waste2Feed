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
	PickupService interface {
		Create(context.Context, string, dto.PickupCreateRequest) (dto.PickupResponse, error)
		GetPaginated(context.Context, string, string, dto.PaginationQuery) (dto.PickupPaginationResponse, error)
	}

	pickupService struct {
		pickupRepo repository.PickupRepository
	}
)

func NewPickupService(pr repository.PickupRepository) PickupService {
	return &pickupService{
		pickupRepo: pr,
	}
}

func (s *pickupService) Create(ctx context.Context, userID string, req dto.PickupCreateRequest) (dto.PickupResponse, error) {
	tDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return dto.PickupResponse{}, err
	}

	tTime, err := time.Parse("15:04", req.Time)
	if err != nil {
		return dto.PickupResponse{}, err
	}

	_date := tDate.Format("2006-01-02")
	_time := tTime.Format("15:04")

	pickup := entity.Pickup{
		ID:          uuid.New(),
		UserID:      userID,
		Date:        _date,
		Time:        _time,
		Address:     req.Address,
		Coordinates: req.Coordinates,
		Amount:      req.Amount,
		Status:      constants.ENUM_PICKUP_STATUS_PROGRESS,
	}

	pickup, err = s.pickupRepo.Create(pickup)
	if err != nil {
		return dto.PickupResponse{}, err
	}

	return dto.PickupResponse{
		ID:          pickup.ID.String(),
		UserID:      pickup.UserID,
		Date:        pickup.Date,
		Time:        pickup.Time,
		Address:     pickup.Address,
		Coordinates: pickup.Coordinates,
		Amount:      pickup.Amount,
		Status:      pickup.Status,
	}, nil
}

func (s *pickupService) GetPaginated(ctx context.Context, userID, role string, req dto.PaginationQuery) (dto.PickupPaginationResponse, error) {
	var limit int
	var page int

	limit = req.PerPage
	if limit <= 0 {
		limit = constants.ENUM_PAGINATION_LIMIT
	}

	page = req.Page
	if page <= 0 {
		page = constants.ENUM_PAGINATION_PAGE
	}

	var datas []entity.Pickup
	var err error
	var maxPage, count int64

	if role == constants.ENUM_ROLE_USER {
		datas, maxPage, count, err = s.pickupRepo.GetPaginationByUser(userID, limit, page)
		if err != nil {
			return dto.PickupPaginationResponse{}, err
		}
	} else if role == constants.ENUM_ROLE_COURIER {
		datas, maxPage, count, err = s.pickupRepo.GetPaginationByCourier(userID, limit, page)
		if err != nil {
			return dto.PickupPaginationResponse{}, err
		}
	} else {
		return dto.PickupPaginationResponse{}, dto.ErrRoleNotAllowed
	}

	var result []dto.PickupResponse
	for _, data := range datas {
		tmp := dto.PickupResponse{
			ID:          data.ID.String(),
			UserID:      data.UserID,
			Date:        data.Date,
			Time:        data.Time,
			Address:     data.Address,
			Coordinates: data.Coordinates,
			Amount:      data.Amount,
			Status:      data.Status,
		}
		if data.CourierID != nil {
			tmp.CourierID = data.CourierID.String()
		} else {
			tmp.CourierID = ""
		}
		result = append(result, tmp)

	}

	return dto.PickupPaginationResponse{
		Data: result,
		PaginationMetadata: dto.PaginationMetadata{
			Page:    page,
			PerPage: limit,
			MaxPage: maxPage,
			Count:   count,
		},
	}, nil
}
