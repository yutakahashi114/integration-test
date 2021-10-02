EXEC=docker-compose exec app
APP_DB_CONFIG=user=develop dbname=app sslmode=disable password=develop
TEST_DB_CONFIG=user=develop dbname=test sslmode=disable password=develop

openapigen:
	${EXEC} oapi-codegen -generate types,chi-server -package openapi -o openapi/openapi_gen.go openapi/openapi.yaml

openapigen_i:
	${EXEC} oapi-codegen -generate types,client -package openapi -o integration_test/openapi/openapi_gen.go openapi/openapi.yaml

db_up:
	${EXEC} sh -c 'cd /app/migrations && goose postgres "${APP_DB_CONFIG} host=db" up'

db_down:
	${EXEC} sh -c 'cd /app/migrations && goose postgres "${APP_DB_CONFIG} host=db" down'

test_db_up:
	${EXEC} sh -c 'cd /app/migrations && goose postgres "${TEST_DB_CONFIG} host=db" up'

test_db_down:
	${EXEC} sh -c 'cd /app/migrations && goose postgres "${TEST_DB_CONFIG} host=db" down'

TRUNCATE_DB_SQL=drop schema public cascade; create schema public;
RESET_SQL_FILE=./integration_test/reset.sql
test_i:
# テーブル全削除のSQL
	echo "${TRUNCATE_DB_SQL}" > ${RESET_SQL_FILE}
# テーブル全削除
	docker-compose exec db psql "${TEST_DB_CONFIG}" -c "${TRUNCATE_DB_SQL}"
# マイグレーション実行
	make test_db_up
# テーブル全削除のSQLにマイグレーション実行後の状態に戻すSQLを追記
	docker-compose exec db pg_dump "${TEST_DB_CONFIG}" >> ${RESET_SQL_FILE}
# テスト実行
	${EXEC} sh -c 'cd ./integration_test && go test -parallel=1 ./...'