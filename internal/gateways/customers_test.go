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

type CustomerRepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo     *customerGateway
	customer entities.Customer
}

func (rs *CustomerRepositorySuite) SetupSuite() {
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

	rs.repo = &customerGateway{rs.DB}
	assert.IsType(rs.T(), &customerGateway{}, rs.repo)

	rs.customer = entities.Customer{
		ID:    1,
		Name:  "John Doe",
		CPF:   "12345678901",
		Email: "test@email.com",
	}
}

func (rs *CustomerRepositorySuite) TestGetAll() {
	expectedSQL := "SELECT (.+) FROM \"customers\" WHERE (.+)"
	customers := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(customers) // avalia o resultado

	_, err := rs.repo.GetAll()  // chama o método GetAll do repository
	assert.NoError(rs.T(), err) // avalia se não houve nenhum erro na execução
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestGetAllReturnsErrorOnQueryFailure() {
	expectedSQL := "SELECT (.+) FROM \"customers\" WHERE (.+)"
	rs.mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("query error"))

	_, err := rs.repo.GetAll()
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "query error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestGetOne_shouldFound() {
	expectedSQL := "SELECT (.+) FROM \"customers\" WHERE (.+) LIMIT (.+)"
	customers := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(customers) // avalia o resultado

	_, err := rs.repo.GetOne(rs.customer)
	assert.NoError(rs.T(), err) // avalia se não houve nenhum erro na execução
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestGetOne_shouldNotFound() {
	expectedSQL := "SELECT (.+) FROM \"customers\" WHERE (.+) LIMIT (.+)"
	customers := sqlmock.NewRows([]string{"id"})
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(customers) // avalia o resultado

	_, err := rs.repo.GetOne(rs.customer) // chama o método GetOne do repository
	assert.Error(rs.T(), err)
	assert.True(rs.T(), errors.Is(err, gorm.ErrRecordNotFound))
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestCreate() {
	expectedSQL := "INSERT INTO \"customers\" (.+) VALUES (.+)"
	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	rs.mock.ExpectBegin()                                   // inicia a transação
	rs.mock.ExpectQuery(expectedSQL).WillReturnRows(addRow) // avalia o resultado
	rs.mock.ExpectCommit()                                  // commita a transação

	_, err := rs.repo.Create(rs.customer) // chama o método Create do repository
	assert.NoError(rs.T(), err)           // avalia se não houve nenhum erro na execução
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestCreateReturnsErrorOnInsertFailure() {
	expectedSQL := "INSERT INTO \"customers\" (.+) VALUES (.+)"
	rs.mock.ExpectBegin()
	rs.mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("insert error"))
	rs.mock.ExpectRollback()

	_, err := rs.repo.Create(rs.customer)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "insert error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestUpdate() {
	expectedSQL := "UPDATE \"customers\" SET .+"
	rs.mock.ExpectBegin()                                                     // inicia a transação
	rs.mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1)) // avalia o resultado
	rs.mock.ExpectCommit()                                                    // commita a transação

	_, err := rs.repo.Update(rs.customer.ID, rs.customer) // chama o método Update do repository
	assert.NoError(rs.T(), err)                           // avalia se não houve nenhum erro na execução
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestUpdateReturnsErrorOnUpdateFailure() {
	expectedSQL := "UPDATE \"customers\" SET .+"
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(expectedSQL).WillReturnError(errors.New("update error"))
	rs.mock.ExpectRollback()

	_, err := rs.repo.Update(rs.customer.ID, rs.customer)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "update error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestDelete() {
	expectedSQL := "UPDATE \"customers\" SET \"deleted_at\"=.+ WHERE \"customers\".\"id\" =.+ AND \"customers\".\"deleted_at\" IS NULL"
	rs.mock.ExpectBegin()                                                     // inicia a transação
	rs.mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1)) // avalia o resultado
	rs.mock.ExpectCommit()                                                    // commita a transação

	err := rs.repo.Delete(rs.customer.ID) // chama o método Delete do repository
	assert.NoError(rs.T(), err)           // avalia se não houve nenhum erro na execução
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *CustomerRepositorySuite) TestDeleteReturnsErrorOnDeleteFailure() {
	expectedSQL := "UPDATE \"customers\" SET \"deleted_at\"=.+ WHERE \"customers\".\"id\" =.+ AND \"customers\".\"deleted_at\" IS NULL"
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(expectedSQL).WillReturnError(errors.New("delete error"))
	rs.mock.ExpectRollback()

	err := rs.repo.Delete(rs.customer.ID)
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "delete error", err.Error())
	assert.Nil(rs.T(), rs.mock.ExpectationsWereMet())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositorySuite))
}
