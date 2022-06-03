package persistent

import (
	"os"
	"testing"

	"github.com/mirror520/rs/model/shopping"
	"github.com/mirror520/rs/persistent/db"

	"github.com/stretchr/testify/suite"
)

type ShoppingSeedTestSuite struct {
	suite.Suite
	repo shopping.Repository
}

func (suite *ShoppingSeedTestSuite) SetupSuite() {
	repo, err := db.NewShoppingRepository()
	suite.NoError(err)

	suite.repo = repo
}

func (suite *ShoppingSeedTestSuite) TestShoppingSeed() {
	err := Seed(suite.repo)
	suite.NoError(err)
}

func (suite *ShoppingSeedTestSuite) TearDownTest() {
	os.Remove("shopping.db")
}

func TestShoppingSeedTestSuite(t *testing.T) {
	suite.Run(t, new(ShoppingSeedTestSuite))
}
