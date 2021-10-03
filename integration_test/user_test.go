package main

import (
	"context"
	"integration-test/openapi"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func Test_User(t *testing.T) {
	resetDB()
	client := newClient()
	ctx := context.Background()
	// データ作成
	{
		err := db.Table("users").Create(
			[]map[string]interface{}{
				{"name": "name1", "email": "email1"},
			},
		).Error
		assert.NoError(t, err)
	}
	// ユーザー作成
	{
		res, err := client.CreateUserWithResponse(ctx, openapi.CreateUserJSONRequestBody{Name: "name2", Email: "email2"})
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
		assertEqual(t, &openapi.ID{Id: 2}, res.JSON200)
	}
	// ユーザー作成
	{
		res, err := client.CreateUserWithResponse(ctx, openapi.CreateUserJSONRequestBody{Name: "name3", Email: "email3"})
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
		assertEqual(t, &openapi.ID{Id: 3}, res.JSON200)
	}
	// ユーザー一覧取得
	{
		res, err := client.FindUsersWithResponse(ctx)
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
		want := &[]openapi.User{
			{Id: 1, Name: "name1", Email: "email1"},
			{Id: 2, Name: "name2", Email: "email2"},
			{Id: 3, Name: "name3", Email: "email3"},
		}
		assertEqual(t, want, res.JSON200, cmpopts.IgnoreFields(openapi.User{}, "CreatedAt"))
	}
	// ユーザー取得
	{
		res, err := client.GetUserByIDWithResponse(ctx, 1)
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
		want := &openapi.User{Id: 1, Name: "name1", Email: "email1"}
		assertEqual(t, want, res.JSON200, cmpopts.IgnoreFields(openapi.User{}, "CreatedAt"))
	}
	// ユーザー更新
	{
		res, err := client.UpdateUserWithResponse(ctx, 1, openapi.UpdateUserJSONRequestBody{Name: "name1_update", Email: "email1_update"})
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
	}
	// ユーザー取得
	{
		res, err := client.GetUserByIDWithResponse(ctx, 1)
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
		want := &openapi.User{Id: 1, Name: "name1_update", Email: "email1_update"}
		assertEqual(t, want, res.JSON200, cmpopts.IgnoreFields(openapi.User{}, "CreatedAt"))
	}
	// ユーザー削除
	{
		res, err := client.DeleteUserWithResponse(ctx, 2)
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
	}
	// ユーザー一覧取得
	{
		res, err := client.FindUsersWithResponse(ctx)
		assert.NoError(t, err)
		assertStatusCode(t, 200, res)
		want := &[]openapi.User{
			{Id: 1, Name: "name1_update", Email: "email1_update"},
			{Id: 3, Name: "name3", Email: "email3"},
		}
		assertEqual(t, want, res.JSON200, cmpopts.IgnoreFields(openapi.User{}, "CreatedAt"))
	}
}
