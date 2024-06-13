package test

import (
	test "selfchain/x/migration/tests"
	"selfchain/x/migration/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestShouldFailIfAddWhenSignerIsNotAdmin() {
	ctx := sdk.WrapSDKContext(suite.ctx)

	// Alice tries to add herself to the list of migrators
	_, err := suite.msgServer.AddMigrator(ctx, &types.MsgAddMigrator{
		Creator:  test.Alice,
		Migrator: test.Alice,
	})

	suite.Require().ErrorIs(err, types.ErrOnlyAdmin)
}

func (suite *IntegrationTestSuite) TestShouldSetTheNewMigrator() {
	ctx := sdk.WrapSDKContext(suite.ctx)

	_, err := suite.msgServer.AddMigrator(ctx, &types.MsgAddMigrator{
		Creator:  test.AclAdmin,
		Migrator: test.Alice,
	})

	suite.Require().Nil(err)

	_, exists := suite.app.MigrationKeeper.GetMigrator(suite.ctx, test.Alice)
	suite.Require().True(exists)
}
