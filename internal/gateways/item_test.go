package gateways

import (
	"database/sql"
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type ItemRepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo *itemGateway
	item entities.Item
}

func (rs *ItemRepositorySuite) SetupSuite() {
	var (
		err error
	)

	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)

	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       rs.conn,
	})

	rs.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)

	rs.repo = &itemGateway{rs.DB}
	assert.IsType(rs.T(), &itemGateway{}, rs.repo)

	rs.item = entities.Item{
		ID:   1,
		Name: "Burger",
	}
}

func (rs *ItemRepositorySuite) TestGetAll() {
	expectedSQL := "SELECT (.+) FROM \"items\" WHERE (.+)"
	items := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(items) // evaluate the result

	_, err := rs.repo.GetAll(rs.item) // call the GetAll method of the repository
	assert.NoError(rs.T(), err)       // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestGetAllReturnsErrorOnQueryFailure() {
	expectedSQL := "SELECT (.+) FROM \"items\" WHERE (.+)"
	rs.mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("query error"))

	_, err := rs.repo.GetAll(rs.item)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "query error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestGetOne_shouldFound() {
	expectedSQL := "SELECT (.+) FROM \"items\" WHERE (.+) LIMIT (.+)"
	items := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(items) // evaluate the result

	_, err := rs.repo.GetOne(rs.item)
	assert.NoError(rs.T(), err) // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestGetOne_shouldNotFound() {
	expectedSQL := "SELECT (.+) FROM \"items\" WHERE (.+) LIMIT (.+)"
	items := sqlmock.NewRows([]string{"id"})
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(items) // evaluate the result

	_, err := rs.repo.GetOne(rs.item) // call the GetOne method of the repository
	assert.Error(rs.T(), err)
	assert.True(rs.T(), errors.Is(err, gorm.ErrRecordNotFound))
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestCreate() {
	expectedSQL := "INSERT INTO \"items\" (.+) VALUES (.+)"
	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectBegin()                                   // start the transaction
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(addRow) // evaluate the result
	rs.mock.ExpectCommit()                                  // commit the transaction

	_, err := rs.repo.Create(rs.item) // call the Create method of the repository
	assert.NoError(rs.T(), err)       // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestCreateReturnsErrorOnInsertFailure() {
	expectedSQL := "INSERT INTO \"items\" (.+) VALUES (.+)"
	rs.mock.ExpectBegin()
	rs.mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("insert error"))
	rs.mock.ExpectRollback()

	_, err := rs.repo.Create(rs.item)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "insert error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestUpdate() {
	expectedSQL := "UPDATE \"items\" SET .+"
	rs.mock.ExpectBegin()                                                     // start the transaction
	rs.mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1)) // evaluate the result
	rs.mock.ExpectCommit()                                                    // commit the transaction

	_, err := rs.repo.Update(rs.item.ID, rs.item) // call the Update method of the repository
	assert.NoError(rs.T(), err)                   // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestUpdateReturnsErrorOnUpdateFailure() {
	expectedSQL := "UPDATE \"items\" SET .+"
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(expectedSQL).WillReturnError(errors.New("update error"))
	rs.mock.ExpectRollback()

	_, err := rs.repo.Update(rs.item.ID, rs.item)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "update error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestDelete() {
	expectedSQL := "UPDATE \"items\" SET \"deleted_at\"=.+ WHERE \"items\".\"id\" =.+ AND \"items\".\"deleted_at\" IS NULL"
	rs.mock.ExpectBegin()                                                     // start the transaction
	rs.mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1)) // evaluate the result
	rs.mock.ExpectCommit()                                                    // commit the transaction

	err := rs.repo.Delete(rs.item.ID) // call the Delete method of the repository
	assert.NoError(rs.T(), err)       // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *ItemRepositorySuite) TestDeleteReturnsErrorOnDeleteFailure() {
	expectedSQL := "UPDATE \"items\" SET \"deleted_at\"=.+ WHERE \"items\".\"id\" =.+ AND \"items\".\"deleted_at\" IS NULL"
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(expectedSQL).WillReturnError(errors.New("delete error"))
	rs.mock.ExpectRollback()

	err := rs.repo.Delete(rs.item.ID)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "delete error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func TestItemSuite(t *testing.T) {
	suite.Run(t, new(ItemRepositorySuite))
}
