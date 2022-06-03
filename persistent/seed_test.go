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

	err = Seed(repo)
	suite.NoError(err)

	suite.repo = repo
}

func (suite *ShoppingSeedTestSuite) TestGetOrder() {
	order, err := suite.repo.GetOrder(1)
	suite.NoError(err)

	suite.Equal("Mirror", order.Customer.Name)
	suite.Equal("犀牛盾", order.Store.Title)

	suite.Len(order.Items, 1)
	suite.Equal("犀牛盾 適用Airpods Pro/第2代/第1代 防摔保護套(含扣環)/玩具總動員-倒吊三眼怪", order.Items[0].Product.Title)
	suite.Equal(float64(690), order.Items[0].Product.Price)
	suite.Equal(609, order.Items[0].Product.Likes)
	suite.Equal("白", order.Items[0].ProductStyle.Style)

	suite.Equal("現折 $100", order.Items[0].Coupon.Title)
	suite.Equal(float64(100), order.Items[0].Coupon.Discount)
	suite.Equal(float64(590), order.Items[0].Price())
}

func (suite *ShoppingSeedTestSuite) TearDownTest() {
	os.Remove("shopping.db")
}

func TestShoppingSeedTestSuite(t *testing.T) {
	suite.Run(t, new(ShoppingSeedTestSuite))
}
