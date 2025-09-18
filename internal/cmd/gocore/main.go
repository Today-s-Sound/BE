package main

import (
	"context"      // 컨텍스트 관리 (요청 취소, 타임아웃 등)
	"database/sql" // 데이터베이스 연결
	"flag"         // 명령행 인수 파싱
	"fmt"          // 문자열 포맷팅
	"log"          // 기본 로깅
	"log/slog"     // 구조화된 로깅 (JSON 형태)
	"os"           // 운영체제 인터페이스
	"strings"      // 문자열 조작
	"time"         // 시간 관련 기능

	"github.com/labstack/echo/v4" // 웹 프레임워크 (Spring의 @RestController 같은)
	"go.uber.org/fx"              // 의존성 주입 (Spring의 @Autowired 같은)

	_ "github.com/go-sql-driver/mysql"

	// 우리 프로젝트의 패키지들
	"github.com/nicewook/gocore/internal/global/config"      // 설정 관리
	"github.com/nicewook/gocore/internal/global/db"          // 데이터베이스 연결
	"github.com/nicewook/gocore/internal/global/errors"      // 에러 처리
	"github.com/nicewook/gocore/internal/global/middlewares" // 미들웨어
)

func main() {
	app := fx.New( // 의존성 주입 컨테이너 생성
		fx.Provide( // "이런 것들을 만들어줘"라고 등록
			NewConfig, // 설정 객체
			NewLogger, // 로거 객체
			NewDB,     // 데이터베이스 연결
			echo.New,  // Echo 웹 서버
		),
		fx.Invoke( // "이런 것들을 실행해줘"라고 등록
			RegisterMiddlewares, // 미들웨어 등록
			RegisterRoutes,      // 라우트 등록
			StartServer,         // 서버 시작
		),
	)

	app.Run() // 모든 것을 실행!
}

func NewConfig() *config.Config {
	// 명령행에서 환경 설정 읽기 (예: go run main.go -env=prod)
	env := flag.String("env", "dev", "Environment (dev, qa, stg, prod)")
	flag.Parse() // 실제로 파싱 실행

	// 유효한 환경인지 확인
	validEnvs := map[string]bool{"dev": true, "qa": true, "stg": true, "prod": true}
	if !validEnvs[*env] {
		log.Fatalf("Invalid environment: %s. Valid environments are: dev, qa, stg, prod", *env)
	}

	// 설정 파일 로드 (config.dev.yaml, config.prod.yaml 등)
	cfg, err := config.LoadConfig(*env)
	if err != nil {
		log.Fatalf("Config load error: %v", err)
	}

	return cfg
}

func NewLogger(cfg *config.Config) *slog.Logger {
	// 설정에서 로그 레벨 읽기
	logLevel := slog.LevelInfo // 기본값
	switch strings.ToLower(cfg.App.LogLevel) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	// JSON 형태의 로그 핸들러 생성
	logHandler := slog.NewJSONHandler(
		os.Stdout, // 표준 출력으로 로그 출력
		&slog.HandlerOptions{
			Level: logLevel,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// 시간 형식을 ISO 8601로 변경
				if a.Key == slog.TimeKey {
					return slog.String(a.Key, time.Now().Format("2006-01-02T15:04:05.000Z07:00"))
				}
				return a
			},
		},
	)
	return slog.New(logHandler)
}

func NewDB(lc fx.Lifecycle, cfg *config.Config) *sql.DB {

	dbConn, err := db.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// 서버 종료 시 데이터베이스 연결도 종료하도록 등록
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return dbConn.Close() // 연결 닫기
		},
	})

	return dbConn
}

func RegisterMiddlewares(cfg *config.Config, logger *slog.Logger, e *echo.Echo) {
	// 전역 에러 핸들러 설정
	e.HTTPErrorHandler = errors.EchoErrorHandler

	// 미들웨어 등록 (CORS, 로깅, 인증 등)
	middlewares.RegisterMiddlewares(cfg, logger, e)
}

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	// 헬스체크 엔드포인트
	e.GET("/health", func(c echo.Context) error {
		return errors.OK(c, map[string]string{
			"status":  "healthy",
			"message": "Server is running",
		}, "Health check successful")
	})

	// 간단한 ping 엔드포인트
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// API v1 그룹 (모든 API가 /api/v1로 시작)
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", func(c echo.Context) error {
			return errors.OK(c, map[string]string{
				"message": "Welcome to TodaySound API v1",
			})
		})
	}
}

func StartServer(lc fx.Lifecycle, e *echo.Echo, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// 고루틴(goroutine)으로 서버 시작 (비동기)
			go func() {
				if err := e.Start(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
					log.Fatal("Shutting down the server due to:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 서버 종료 시 정리 작업
			return e.Shutdown(ctx)
		},
	})
}
