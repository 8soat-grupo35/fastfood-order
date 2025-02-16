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

type OrderRepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo  *orderGateway
	order entities.Order
}

func (rs *OrderRepositorySuite) SetupSuite() {
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

	rs.repo = &orderGateway{rs.DB}
	assert.IsType(rs.T(), &orderGateway{}, rs.repo)

	rs.order = entities.Order{
		ID: 1,
	}
}

func (rs *OrderRepositorySuite) TestGetAll() {
	expectedSQL := "SELECT (.+) FROM \"orders\" WHERE (.+) ORDER BY CASE status WHEN 'PRONTO' THEN 1 WHEN 'EM_PREPARACAO' THEN 2 WHEN 'RECEBIDO' THEN 3 ELSE 4 END,created_at ASC"
	orders := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(orders) // evaluate the result

	expectedOrderItemsSQL := "SELECT (.+) FROM \"order_items\" WHERE \"order_items\".\"order_id\" = (.+)"
	orderItems := sqlmock.NewRows([]string{"order_id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedOrderItemsSQL).WithArgs(rs.order.ID).WillReturnRows(orderItems) // evaluate the result

	_, err := rs.repo.GetAll()  // call the GetAll method of the repository
	assert.NoError(rs.T(), err) // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *OrderRepositorySuite) TestGetAllReturnsErrorOnQueryFailure() {
	expectedSQL := "SELECT (.+) FROM \"orders\" WHERE (.+) ORDER BY CASE status WHEN 'PRONTO' THEN 1 WHEN 'EM_PREPARACAO' THEN 2 WHEN 'RECEBIDO' THEN 3 ELSE 4 END,created_at ASC"
	rs.mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("query error"))

	_, err := rs.repo.GetAll()
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "query error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *OrderRepositorySuite) TestGetById_shouldFound() {
	expectedOrderSQL := "SELECT (.+) FROM \"orders\" WHERE (.+) ORDER BY \"orders\".\"id\" LIMIT (.+)"
	orders := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedOrderSQL).WillReturnRows(orders) // evaluate the result

	expectedOrderItemsSQL := "SELECT (.+) FROM \"order_items\" WHERE \"order_items\".\"order_id\" = (.+)"
	orderItems := sqlmock.NewRows([]string{"order_id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedOrderItemsSQL).WithArgs(rs.order.ID).WillReturnRows(orderItems) // evaluate the result

	_, err := rs.repo.GetById(rs.order.ID)
	assert.NoError(rs.T(), err) // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *OrderRepositorySuite) TestGetById_shouldNotFound() {
	expectedSQL := "SELECT (.+) FROM \"orders\" WHERE (.+) LIMIT (.+)"
	orders := sqlmock.NewRows([]string{"id"})
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(orders) // evaluate the result

	_, err := rs.repo.GetById(rs.order.ID) // call the GetById method of the repository
	assert.Error(rs.T(), err)
	assert.True(rs.T(), errors.Is(err, gorm.ErrRecordNotFound))
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *OrderRepositorySuite) TestCreate() {
	expectedSQL := "INSERT INTO \"orders\" (.+) VALUES (.+)"
	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectBegin()                                   // start the transaction
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(addRow) // evaluate the result
	rs.mock.ExpectCommit()                                  // commit the transaction

	_, err := rs.repo.Create(rs.order) // call the Create method of the repository
	assert.NoError(rs.T(), err)        // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *OrderRepositorySuite) TestCreateReturnsErrorOnInsertFailure() {
	expectedSQL := "INSERT INTO \"orders\" (.+) VALUES (.+)"
	rs.mock.ExpectBegin()
	rs.mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("insert error"))
	rs.mock.ExpectRollback()

	_, err := rs.repo.Create(rs.order)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "insert error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *OrderRepositorySuite) TestUpdate() {
	expectedSQL := "UPDATE \"orders\" SET .+"
	rs.mock.ExpectBegin()                                                     // start the transaction
	rs.mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1)) // evaluate the result
	rs.mock.ExpectCommit()                                                    // commit the transaction

	_, err := rs.repo.Update(rs.order.ID, rs.order) // call the Update method of the repository
	assert.NoError(rs.T(), err)                     // evaluate if there was no error in execution
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *OrderRepositorySuite) TestUpdateReturnsErrorOnUpdateFailure() {
	expectedSQL := "UPDATE \"orders\" SET .+"
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(expectedSQL).WillReturnError(errors.New("update error"))
	rs.mock.ExpectRollback()

	_, err := rs.repo.Update(rs.order.ID, rs.order)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "update error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositorySuite))
}
