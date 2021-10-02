package main

import (
	"bufio"
	"context"
	"fmt"
	"integration-test/openapi"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dsn string

func TestMain(m *testing.M) {
	// ポートやDBの接続先を変えてサーバーを起動する
	os.Chdir("/app")
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = append(os.Environ(), "PORT=81", "POSTGRES_DB_NAME="+os.Getenv("POSTGRES_TEST_DB_NAME"))
	stdout, _ := cmd.StdoutPipe()
	fatalIf(cmd.Start())
	go func() {
		// 必要ならサーバーのログを出力
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}()
	os.Chdir("/app/integration_test")

	// DBに接続
	connectDB()

	// サーバーの起動が完了するまで待機
	// ヘルスチェックのエンドポイントがあればそれを叩いた方が良さそう
	time.Sleep(5 * time.Second)

	fmt.Println("============ test start ============")

	// テスト実行
	exitVal := m.Run()

	// サーバーを終了する
	cmd.Process.Kill()
	cmd.Wait()

	os.Exit(exitVal)
}

func connectDB() {
	dsn = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_TEST_DB_NAME"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fatalIf(err)
}

func resetDB() {
	// 原因は不明だがテーブルの作成に失敗することがあるので、最大5回リトライする
	for i := 0; i < 5; i++ {
		out, err := exec.Command("psql", dsn, "-f", "./reset.sql").Output()
		fatalIf(err)
		// 成功したときは4行以上出力される
		if len(strings.Split(string(out), "\n")) >= 4 {
			return
		}
		fmt.Println("retry reset db")
		time.Sleep(1 * time.Second)
	}
	log.Fatalln("failed to reset db")
}

func newClient() *openapi.ClientWithResponses {
	// 認証が必要であればトークンをセットする
	token := ""
	c, err := openapi.NewClientWithResponses("http://localhost:81", openapi.WithRequestEditorFn(
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Authorization", "Bearer "+token)
			return nil
		},
	))
	fatalIf(err)
	return c
}

func assertStatusCode(t *testing.T, expect int, actual interface{ StatusCode() int }) bool {
	if expect != actual.StatusCode() {
		assert.Fail(t, "assertStatusCode mismatch", fmt.Sprintf("status code got: %d, want:%d", actual.StatusCode(), expect))
		return false
	}
	return true
}

func assertEqual(t *testing.T, expect, actual interface{}, opts ...cmp.Option) bool {
	if diff := cmp.Diff(actual, expect, opts...); diff != "" {
		assert.Fail(t, "assertEqual mismatch", fmt.Sprintf("differs: (-got +want)\n%s", diff))
		return false
	}
	return true
}

func fatalIf(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
