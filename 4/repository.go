package db

import (
	"context"
	"golang-learner/4/libs"
	"golang-learner/4/models"
	"gorm.io/gorm"
)

type repository struct {
	con *gorm.DB
}

type Repository interface {
	// All : Returns all Users limited by pagination details
	All(ctx context.Context, pageNumber, pageSize int) ([]*models.User, error)
	// Create - Creates User
	Create(ctx context.Context, user *models.User) (*models.User, error)
	// DeleteById : Deletes User by User.ID
	DeleteById(ctx context.Context, id uint) error
	// GetById : Returns User by User.ID
	GetById(ctx context.Context, id uint) (*models.User, error)
	// Update : Updates User
	Update(ctx context.Context, user *models.User) (*models.User, error)
}

func New(con *database.Connection) Repository {
	return &repository{con: con.Active()}
}

func (r repository) All(ctx context.Context, pageNumber, pageSize int) ([]*models.User, error) {
	//TODO implement me
	// Requirements:
	//  All User and embedded Addresses objects are returned when records are not marked deleted
	//     (gorm already adds a filter to remove soft-deleted records)
	//  If no User records exists or are not deleted an empty array is returned
	//  Error returned only if gorm returns non-nil error
	//  When pageSize <= 0 pagination should be skipped
	//  When pageSize > 0 pagination should be used to limit number of records returned
	panic("implement me")
}

func (r repository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	//TODO implement me
	// Requirements:
	//  User object and its embedded Addresses are saved (User and Address models already have a lifecycle hook to run struct validation on create or update)
	//  Error returned if gorm returns non-nil error
	//  Returned user object has the auto-generated id value set (gorm does this by default, but you will need to test for it)
	//  Feel free to add your own additional requirements for a record to be created
	panic("implement me")
}

func (r repository) DeleteById(ctx context.Context, id uint) error {
	//TODO implement me
	// Requirements:
	//  User object and its embedded Addresses are marked deleted when record exists with matching user.ID
	//     (gorm by default uses soft deletes setting the DeletedAt field to a non-nil value)
	//  When record does not exist by ID, NO error is returned
	//  Error returned if gorm returns non-nil error
	//  Feel free to add your own additional requirements for a record to be deleted
	panic("implement me")
}

func (r repository) GetById(ctx context.Context, id uint) (*models.User, error) {
	//TODO implement me
	// Requirements:
	//  User object where user.ID matches and its embedded Addresses are returned when record exists and is NOT marked deleted
	//     (gorm already adds a filter to remove soft-deleted records)
	//  When User record does not exist by id error is returned
	//  Otherwise, error returned only if gorm returns non-nil error
	panic("implement me")
}

func (r repository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	//TODO implement me
	// Requirements:
	//  User object and its embedded Addresses are updated when record exists with matching user.ID
	//    You can optionally add logic to perform an UPSERT (create if not exists)
	//  Error returned if gorm returns non-nil error
	//  Feel free to add your own additional requirements for a record to be updated
	panic("implement me")
}
