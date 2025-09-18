package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nicewook/gocore/internal/global/config"
)

// NewDBConnection 새로운 데이터베이스 연결을 생성합니다
// cfg: 데이터베이스 설정 정보가 담긴 Config 구조체
// 반환값: *sql.DB (데이터베이스 연결 객체), error (에러)
func NewDBConnection(cfg *config.Config) (*sql.DB, error) {
	// DSN(Data Source Name) 생성 - 데이터베이스 연결 문자열
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
	)

	// mysql 데이터베이스 연결 생성
	// sql.Open은 실제 연결을 생성하지 않고 연결 가능한 객체만 반환
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("데이터베이스 연결 생성 실패: %w", err)
	}

	// 연결 풀 설정 (성능 최적화)
	db.SetMaxOpenConns(10)                 // 동시에 열 수 있는 최대 연결 수: 10개
	db.SetMaxIdleConns(5)                  // 유휴 상태로 유지할 연결 수: 5개
	db.SetConnMaxLifetime(5 * time.Minute) // 연결의 최대 수명: 5분 (5분 후 재연결)

	if err := createUserTable(db); err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	if err := createSubscriptionTable(db); err != nil {
		return nil, fmt.Errorf("failed to create subscriptions table: %w", err)
	}

	if err := createSummaryTable(db); err != nil {
		return nil, fmt.Errorf("failed to create summary table: %w", err)
	}
	// 실제 데이터베이스 연결 테스트
	// Ping()을 통해 데이터베이스가 실제로 연결 가능한지 확인
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("데이터베이스 연결 테스트 실패: %w", err)
	}

	// 연결 성공 시 데이터베이스 객체 반환
	return db, nil
}

func createUserTable(db *sql.DB) error {

	const query = `
        CREATE TABLE IF NOT EXISTS user (
            id BIGINT NOT NULL AUTO_INCREMENT,
            name VARCHAR(30) NOT NULL,
            created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
            updated_at DATETIME(6) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
            PRIMARY KEY (id)
        )
	`
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to create User table: %w", err)
	}
	return nil
}

func createSubscriptionTable(db *sql.DB) error {

	const query = `
        CREATE TABLE IF NOT EXISTS subscription (
            id BIGINT NOT NULL AUTO_INCREMENT,
            user_id BIGINT NOT NULL,
            url TEXT NOT NULL,
            alias VARCHAR(30),
            keyword_filter VARCHAR(100),
            is_urgent BOOLEAN NOT NULL DEFAULT FALSE,
            created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
            updated_at DATETIME(6) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
            PRIMARY KEY (id),
            FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
        )
	`
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to create products table: %w", err)
	}
	return nil
}

func createSummaryTable(db *sql.DB) error {

	const query = `
        CREATE TABLE IF NOT EXISTS summary (
            id BIGINT NOT NULL AUTO_INCREMENT,
            subscription_id BIGINT NOT NULL,
            hash TEXT NOT NULL,
            summary TEXT NOT NULL,
            is_read BOOLEAN NOT NULL DEFAULT FALSE,
            created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
            updated_at DATETIME(6) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
            PRIMARY KEY (id),
            FOREIGN KEY (subscription_id) REFERENCES subscription(id) ON DELETE CASCADE
        )
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to create orders table: %w", err)
	}
	return nil
}
