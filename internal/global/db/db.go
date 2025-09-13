package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL 드라이버 (사용하지 않지만 import 필요)

	"github.com/nicewook/gocore/internal/global/config"
)

// NewDBConnection 새로운 데이터베이스 연결을 생성합니다
// cfg: 데이터베이스 설정 정보가 담긴 Config 구조체
// 반환값: *sql.DB (데이터베이스 연결 객체), error (에러)
func NewDBConnection(cfg *config.Config) (*sql.DB, error) {
	// DSN(Data Source Name) 생성 - 데이터베이스 연결 문자열
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,     // 데이터베이스 호스트 주소
		cfg.DB.Port,     // 데이터베이스 포트 번호
		cfg.DB.User,     // 데이터베이스 사용자명
		cfg.DB.Password, // 데이터베이스 비밀번호
		cfg.DB.DBName,   // 데이터베이스 이름
		cfg.DB.SSLMode,  // SSL 연결 모드 (disable, require 등)
	)

	// PostgreSQL 데이터베이스 연결 생성
	// sql.Open은 실제 연결을 생성하지 않고 연결 가능한 객체만 반환
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("데이터베이스 연결 생성 실패: %w", err)
	}

	// 연결 풀 설정 (성능 최적화)
	db.SetMaxOpenConns(10)                 // 동시에 열 수 있는 최대 연결 수: 10개
	db.SetMaxIdleConns(5)                  // 유휴 상태로 유지할 연결 수: 5개
	db.SetConnMaxLifetime(5 * time.Minute) // 연결의 최대 수명: 5분 (5분 후 재연결)

	// 실제 데이터베이스 연결 테스트
	// Ping()을 통해 데이터베이스가 실제로 연결 가능한지 확인
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("데이터베이스 연결 테스트 실패: %w", err)
	}

	// 연결 성공 시 데이터베이스 객체 반환
	return db, nil
}
