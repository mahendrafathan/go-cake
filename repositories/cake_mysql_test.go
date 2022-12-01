package repositories

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/mahendrafathan/go-cake/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CakeTestSuite struct {
	suite.Suite
	sqlDB   *sql.DB
	sqlMock sqlmock.Sqlmock
	sqlxDb  *sqlx.DB
}

func TestCakeRepository(t *testing.T) {
	suite.Run(t, new(CakeTestSuite))
}

func (suite *CakeTestSuite) TestFindAll() {
	testCases := []struct {
		name     string
		preTest  func(ctx context.Context)
		wantResp []models.Cake
		wantErr  bool
	}{
		{
			name: "success",
			preTest: func(ctx context.Context) {
				suite.
					sqlMock.
					ExpectQuery(regexp.QuoteMeta("SELECT c.id, c.title, c.description, c.rating, c.image FROM cakes c order by c.rating, c.title")).
					WillReturnRows(sqlmock.NewRows([]string{"c.id", "c.title", "c.description", "c.rating", "c.image"}).
						AddRow(1, "cake", "desc", 7, "image"))
			},
			wantResp: []models.Cake{
				models.Cake{
					Id:          1,
					Title:       "cake",
					Description: "desc",
					Rating:      7,
					Image:       "image",
				},
			},
			wantErr: false,
		},
		{
			name: "failed",
			preTest: func(ctx context.Context) {
				suite.
					sqlMock.
					ExpectQuery(regexp.QuoteMeta("SELECT c.id, c.title, c.description, c.rating, c.image FROM cakes c where id = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"c.id", "c.title", "c.description", "c.rating", "c.image"}).
						AddRow(1, "cake", "desc", 7, "image"))
			},
			wantResp: []models.Cake{},
			wantErr:  true,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			cakeRepo := cakeRepoMysql{
				db: suite.sqlDB,
			}
			tt.preTest(context.Background())

			res, err := cakeRepo.FindAll()
			assert.Equal(t, tt.wantResp, res)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}
func (suite *CakeTestSuite) TestFindOne() {
	testCases := []struct {
		name     string
		preTest  func(ctx context.Context)
		wantResp models.Cake
		wantErr  bool
	}{
		{
			name: "success",
			preTest: func(ctx context.Context) {
				suite.
					sqlMock.
					ExpectQuery(regexp.QuoteMeta("SELECT c.id, c.title, c.description, c.rating, c.image FROM cakes c where id = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"c.id", "c.title", "c.description", "c.rating", "c.image"}).
						AddRow(1, "cake", "desc", 7, "image"))
			},
			wantResp: models.Cake{
				Id:          1,
				Title:       "cake",
				Description: "desc",
				Rating:      7,
				Image:       "image",
			},
			wantErr: false,
		},
		{
			name: "failed",
			preTest: func(ctx context.Context) {
				suite.
					sqlMock.
					ExpectQuery(regexp.QuoteMeta("SELECT c.id, c.title, c.description, c.rating, c.image FROM cakes c where id = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"c.id", "c.title", "c.description", "c.rating", "c.image"}).
						AddRow(1, "cake", "desc", 7, "image"))
			},
			wantResp: models.Cake{
				Id:          1,
				Title:       "cake",
				Description: "desc",
				Rating:      7,
				Image:       "image",
			},
			wantErr: false,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			cakeRepo := cakeRepoMysql{
				db: suite.sqlDB,
			}
			tt.preTest(context.Background())

			res, err := cakeRepo.FindOne(1)
			assert.Equal(t, tt.wantResp, res)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}

func (suite *CakeTestSuite) SetupTest() {
	t := suite.T()
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)

	suite.sqlDB = mockDB
	suite.sqlMock = mock

	suite.sqlxDb = sqlx.NewDb(mockDB, "sqlmock")
}

func (suite *CakeTestSuite) TearDownTest() {
	suite.sqlxDb.Close()
	suite.sqlDB.Close()
}
