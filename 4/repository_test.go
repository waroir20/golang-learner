package db_test

import (
	"context"
	"github.com/rs/zerolog/log"
	"golang-learner/4"
	"golang-learner/4/libs"
	"golang-learner/4/models"
	"os"
	"testing"
)

const (
	username = "testuser"
	password = "testpass"
	dbName   = "testdb"
	dbSchema = "testschema"
)

var (
	testContainer database.PostgresContainer
	con           *database.Connection
	target        db.Repository
)

func TestMain(t *testing.M) {
	err := setup()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error during start up")
		os.Exit(1)
		return
	}
	code := t.Run()
	shutdown()
	os.Exit(code)
}

func setup() error {
	var err error
	testContainer, err = database.NewTestContainer(context.Background(), dbName, dbSchema, username, password)
	if err != nil {
		return err
	}
	con = database.New(testContainer.Config())
	if err = con.Open(); err != nil {
		_ = testContainer.Shutdown(context.Background()) //Terminate container on error
		return err
	}
	if err = models.Migrate(con.Active()); err != nil {
		shutdown() //Close existing connections and terminate container on error
		return err
	}
	target = db.New(con)
	return nil
}

func shutdown() {
	_ = con.Disconnect()
	_ = testContainer.Shutdown(context.Background())
}

func TestRepository_All(t *testing.T) {
	t.Cleanup(cleanup) //Used to wipe DB in between tests

	//TODO implement me
	// Cases to test:
	//  Empty array and NO error are returned when tables empty
	//  Empty array and NO error are returned when tables not empty but all records marked deleted
	//  When pageSize <= 0 no limitation placed on array size of records returned
	//  When pageSize > 0 pagination should be used to limit number of records returned
	panic("implement me")
}

func TestRepository_Create(t *testing.T) {
	t.Cleanup(cleanup) //Used to wipe DB in between tests

	//TODO implement me
	// Cases to test:
	//  Error returned when User struct has no FirstName specified (i.e. FirstName is an empty string) as FirstName is marked required
	//    (User and Address models already have a lifecycle hook to run struct validation on create or update)
	//  When create returns without error, returned user object has the auto-generated id value set (gorm does this by default, but you will need to test for it)
	//  Use GetById method to verify record exists in DB with newly generated ID
	panic("implement me")
}

func TestRepository_DeleteById(t *testing.T) {
	t.Cleanup(cleanup) //Used to wipe DB in between tests

	//TODO implement me
	// Cases to test:
	//  User record marked deleted when record exists with matching ID
	//     (gorm by default uses soft deletes setting the DeletedAt field to a non-nil value)
	//  Use GetById method to verify record not returned by ID
	panic("implement me")
}

func TestRepository_GetById(t *testing.T) {
	t.Cleanup(cleanup) //Used to wipe DB in between tests

	//TODO implement me
	// Cases to test:
	//  Error is returned when record does not exist with matching ID
	//  Error is returned when record with matching ID is marked deleted
	//  User and embedded Addresses are returned when record with matching ID exists and is NOT marked deleted
	panic("implement me")
}

func TestRepository_Update(t *testing.T) {
	t.Cleanup(cleanup) //Used to wipe DB in between tests

	//TODO implement me
	// Cases to test:
	//  User object and its embedded Addresses are updated when record exists
	//  Deleted User object and its embedded Addresses are updated when record exists, AND record is marked NOT deleted
	//  Error returned when User struct has no FirstName specified (i.e. FirstName is an empty string) as FirstName is marked required
	//    (User and Address models already have a lifecycle hook to run struct validation on create or update)
	panic("implement me")
}

func cleanup() {
	tx1 := con.Active().Unscoped().Where("1=1").Delete(&models.User{})
	if tx1.Error != nil {
		log.Error().Stack().Err(tx1.Error).Msg("Error clearing the data from user table")
	}
	tx2 := con.Active().Unscoped().Where("1=1").Delete(&models.Address{})
	if tx2.Error != nil {
		log.Error().Stack().Err(tx2.Error).Msg("Error clearing the data from address table")
	}
}
