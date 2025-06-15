package service

import (
	"context"

	"Waste2Feed/constants"
	"Waste2Feed/dto"
	"Waste2Feed/entity"
	"Waste2Feed/repository"
	"Waste2Feed/utils"

	"gorm.io/gorm"
)

type (
	UserService interface {
		Register(context.Context, string, dto.UserAuthRequest) (dto.UserResponse, error)
		NewCourier(context.Context, dto.UserAuthRequest) (dto.UserResponse, error)
		DeleteCourier(context.Context, string) error
		EditCourier(context.Context, string, dto.UserUpdateRequest) error
		GetAllByRoleID(context.Context, string) ([]entity.User, error)
		GetCourierByID(context.Context, string) (entity.User, error)
		Update(context.Context, string, dto.UserUpdateRequest) error
		Login(context.Context, dto.UserAuthRequest) (entity.User, error)
		Me(context.Context, string) (dto.UserResponse, error)
		Redeem(context.Context, string, dto.UserRedeemRequest) error
	}

	userService struct {
		userRepo repository.UserRepository
		roleRepo repository.RoleRepository
	}
)

func NewUserService(ur repository.UserRepository, rr repository.RoleRepository) UserService {
	return &userService{
		userRepo: ur,
		roleRepo: rr,
	}
}

func (s *userService) Register(ctx context.Context, _role string, req dto.UserAuthRequest) (dto.UserResponse, error) {
	var user entity.User

	if _role == "" {
		return dto.UserResponse{}, dto.ErrMissingRole
	}

	if _role != constants.ENUM_ROLE_FARMER && _role != constants.ENUM_ROLE_USER {
		return dto.UserResponse{}, dto.ErrInvalidRoleName
	}

	if req.Email != "" {
		_, err := s.userRepo.GetByEmail(req.Email)
		if err == nil && err != gorm.ErrRecordNotFound {
			return dto.UserResponse{}, dto.ErrEmailAlreadyExists
		}

		user = entity.User{
			Email: req.Email,
		}
	} else if req.PhoneNumber != "" {
		_, err := s.userRepo.GetByPhoneNumber(req.PhoneNumber)
		if err == nil && err != gorm.ErrRecordNotFound {
			return dto.UserResponse{}, dto.ErrPhoneNumberAlreadyExists
		}

		user = entity.User{
			PhoneNumber: req.PhoneNumber,
		}
	} else {
		return dto.UserResponse{}, dto.ErrEmailOrPhoneNumberMissing
	}

	role, err := s.roleRepo.GetByName(_role)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user.Password = req.Password
	user.RoleID = role.ID.String()

	userReg, err := s.userRepo.Create(user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:          userReg.ID.String(),
		Email:       userReg.Email,
		PhoneNumber: userReg.PhoneNumber,
		Role:        role.Name,
	}, nil
}

func (s *userService) GetAllByRoleID(ctx context.Context, rolename string) ([]entity.User, error) {
	role, err := s.roleRepo.GetByName(rolename)
	if err != nil {
		return nil, err
	}

	users, err := s.userRepo.GetAllByRoleID(role.ID.String())
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) EditCourier(ctx context.Context, userID string, req dto.UserUpdateRequest) error {
	var data entity.User

	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return err
	}

	role, err := s.roleRepo.GetbyId(user.RoleID)
	if err != nil {
		return err
	}

	if role.Name != constants.ENUM_ROLE_COURIER {
		return dto.ErrUserNotCourier
	}

	var newPassword string
	if req.NewPassword != "" {
		newPassword, err = utils.HashPassword(req.NewPassword)
		if err != nil {
			return err
		}
	}

	if req.UserName != "" {
		_, err := s.userRepo.GetByUsername(req.UserName)
		if err == nil && err != gorm.ErrRecordNotFound {
			return dto.ErrUsernameAlreadyExists
		}
		data.UserName = req.UserName
	}

	if req.Gender != "" && (req.Gender != "Pria" && req.Gender != "Wanita") {
		return dto.ErrInvalidGender
	}

	data = entity.User{
		ID:       user.ID,
		Name:     req.Name,
		UserName: req.UserName,
		Address:  req.Address,
		Password: newPassword,
		Gender:   req.Gender,
	}

	_, err = s.userRepo.Update(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) NewCourier(ctx context.Context, req dto.UserAuthRequest) (dto.UserResponse, error) {
	var user entity.User

	if req.Email != "" {
		_, err := s.userRepo.GetByEmail(req.Email)
		if err == nil && err != gorm.ErrRecordNotFound {
			return dto.UserResponse{}, dto.ErrEmailAlreadyExists
		}

		user = entity.User{
			Email: req.Email,
		}
	} else if req.PhoneNumber != "" {
		_, err := s.userRepo.GetByPhoneNumber(req.PhoneNumber)
		if err == nil && err != gorm.ErrRecordNotFound {
			return dto.UserResponse{}, dto.ErrPhoneNumberAlreadyExists
		}

		user = entity.User{
			PhoneNumber: req.PhoneNumber,
		}
	} else {
		return dto.UserResponse{}, dto.ErrEmailOrPhoneNumberMissing
	}

	role, err := s.roleRepo.GetByName(constants.ENUM_ROLE_COURIER)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user.Password = req.Password
	user.RoleID = role.ID.String()

	userReg, err := s.userRepo.Create(user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:          userReg.ID.String(),
		Email:       userReg.Email,
		PhoneNumber: userReg.PhoneNumber,
		Role:        constants.ENUM_ROLE_COURIER,
	}, nil
}

func (s *userService) DeleteCourier(ctx context.Context, userID string) error {
	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return err
	}

	role, err := s.roleRepo.GetbyId(user.RoleID)
	if err != nil {
		return err
	}

	if role.Name != constants.ENUM_ROLE_COURIER {
		return dto.ErrUserNotCourier
	}

	err = s.userRepo.Delete(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Update(ctx context.Context, userID string, req dto.UserUpdateRequest) error {
	var data entity.User

	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return err
	}

	var newPassword string
	if req.OldPassword != "" && req.NewPassword != "" {
		hashedPassword, err := utils.HashPassword(req.OldPassword)
		if err != nil {
			return err
		}

		if hashedPassword != user.Password {
			return dto.ErrCredentialsNotMatched
		}

		newPassword, err = utils.HashPassword(req.NewPassword)
		if err != nil {
			return err
		}
	}

	if req.UserName != "" {
		_, err := s.userRepo.GetByUsername(req.UserName)
		if err == nil && err != gorm.ErrRecordNotFound {
			return dto.ErrUsernameAlreadyExists
		}
		data.UserName = req.UserName
	}

	if req.Gender != "" && (req.Gender != "Pria" && req.Gender != "Wanita") {
		return dto.ErrInvalidGender
	}

	data = entity.User{
		ID:       user.ID,
		Name:     req.Name,
		UserName: req.UserName,
		Address:  req.Address,
		Password: newPassword,
		Gender:   req.Gender,
	}

	_, err = s.userRepo.Update(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(ctx context.Context, req dto.UserAuthRequest) (entity.User, error) {
	var user entity.User
	var err error

	if req.Email != "" {
		user, err = s.userRepo.GetByEmail(req.Email)
		if err != nil {
			return entity.User{}, dto.ErrCredentialsNotMatched
		}

		checkPassword, err := utils.CheckPassword(user.Password, []byte(req.Password))
		if err != nil || !checkPassword {
			return entity.User{}, dto.ErrCredentialsNotMatched
		}
	} else if req.PhoneNumber != "" {
		user, err = s.userRepo.GetByPhoneNumber(req.PhoneNumber)
		if err != nil {
			return entity.User{}, dto.ErrCredentialsNotMatched
		}

		checkPassword, err := utils.CheckPassword(user.Password, []byte(req.Password))
		if err != nil || !checkPassword {
			return entity.User{}, dto.ErrCredentialsNotMatched
		}
	} else {
		return entity.User{}, dto.ErrEmailOrPhoneNumberMissing
	}

	role, err := s.roleRepo.GetbyId(user.RoleID)
	if err != nil {
		return entity.User{}, err
	}
	user.Role = &role

	return user, nil
}

func (s *userService) Me(ctx context.Context, userID string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	role, err := s.roleRepo.GetbyId(user.RoleID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:          user.ID.String(),
		UserName:    user.UserName,
		Name:        user.Name,
		Address:     user.Address,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        role.Name,
		Verified:    user.Verified,
		Gender:      user.Gender,
		Points:      *user.Points,
	}, nil
}

func (s *userService) Redeem(ctx context.Context, userID string, req dto.UserRedeemRequest) error {
	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return err
	}

	if *user.Points <= 0 {
		return dto.ErrNoPointsAvailable
	}

	if *req.Nominal > *user.Points {
		return dto.ErrInvalidPointsAmount
	}

	new := *user.Points - *req.Nominal
	_user := entity.User{
		ID:     user.ID,
		Points: &new,
	}
	_, err = s.userRepo.Update(_user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetCourierByID(ctx context.Context, userID string) (entity.User, error) {
	user, err := s.userRepo.GetById(userID)
	if err != nil {
		return entity.User{}, err
	}

	role, err := s.roleRepo.GetbyId(user.RoleID)
	if err != nil {
		return entity.User{}, err
	}

	if role.Name != constants.ENUM_ROLE_COURIER {
		return entity.User{}, dto.ErrUserNotCourier
	}

	user.Role = &role

	return user, nil
}
