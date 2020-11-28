package db

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/log"
)

var instance *gorm.DB

// parsePostgreSQLHostPort fork https://github.com/gogs/gogs/blob/main/internal/db/db.go#L31
func parsePostgreSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "5432"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

// parseMSSQLHostPort fork https://github.com/gogs/gogs/blob/main/internal/db/db.go#L43
func parseMSSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "1433"
	switch {
	case strings.Contains(info, ":"):
		host = strings.Split(info, ":")[0]
		port = strings.Split(info, ":")[1]
	case strings.Contains(info, ","):
		host = strings.Split(info, ",")[0]
		port = strings.TrimSpace(strings.Split(info, ",")[1])
	case len(info) > 0:
		host = info
	}

	return host, port
}

// DSN return parsed DSN
func DSN() (dsn string, err error) {
	concate := "?"
	if strings.Contains(conf.Database.Name, concate) {
		concate = "&"
	}

	switch conf.Database.Type {
	case Mysql, Tidb:
		var format string
		if conf.Database.Host[0] == '/' { // Looks like a unix socket
			format = "%s:%s@unix(%s)/%s%scharset=utf8mb4&parseTime=true"
		} else {
			format = "%s:%s@tcp(%s)/%s%scharset=utf8mb4&parseTime=true"
		}

		dsn = fmt.Sprintf(format, conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Name, concate)
	case Postgres:
		host, port := parsePostgreSQLHostPort(conf.Database.Host)
		if host[0] == '/' { // looks like a unix socket
			dsn = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
				url.QueryEscape(conf.Database.User), url.QueryEscape(conf.Database.Password), port, conf.Database.Name, concate, conf.Database.SSLMode, host)
		} else {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
				url.QueryEscape(conf.Database.User), url.QueryEscape(conf.Database.Password), host, port, conf.Database.Name, concate, conf.Database.SSLMode)
		}
	case Mssql:
		host, port := parseMSSQLHostPort(conf.Database.Host)
		dsn = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;",
			host, port, conf.Database.Name, conf.Database.User, conf.Database.Password)
	case Sqlite3:
		dsn = "file:" + conf.Database.Path + "?cache=shared&mode=rwc"
	default:
		return "", fmt.Errorf("unrecognized dialect: %s", conf.Database.Type)
	}

	return dsn, nil
}

func open(cfg *gorm.Config) (*gorm.DB, error) {
	dsn, err := DSN()
	if err != nil {
		return nil, fmt.Errorf("parse DSN: %w", err)
	}

	var dialector gorm.Dialector
	switch conf.Database.Type {
	case Mysql, Tidb:
		dialector = mysql.Open(dsn)
	case Postgres:
		dialector = postgres.Open(dsn)
	case Mssql:
		dialector = sqlserver.Open(dsn)
	case Sqlite3:
		dialector = sqlite.Open(dsn)
	default:
		log.Fatal("unreachable")
	}

	return gorm.Open(dialector, cfg)
}

// Init for db
func Init() {
	logger.Default = logger.New(
		new(Logger),
		logger.Config{
			SlowThreshold: time.Second,    // 慢 SQL 阈值
			LogLevel:      logger.Info,    // Log level
			Colorful:      conf.App.Debug, // 禁用彩色打印
		},
	)

	db, err := open(&gorm.Config{
		PrepareStmt: true, // 开启预编译 SQL
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC().Truncate(time.Microsecond)
		},
	})

	if err != nil {
		log.Fatal(fmt.Errorf("open database: %w", err).Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(fmt.Errorf("get underlying *sql.DB: %w", err).Error())
	}

	sqlDB.SetMaxOpenConns(conf.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(conf.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Minute)

	switch conf.Database.Type {
	case "postgres":
		conf.UsePostgreSQL = true
	case "mysql":
		conf.UseMySQL = true
		db = db.Set("gorm:table_options", "ENGINE=InnoDB").Session(&gorm.Session{})
	case "sqlite3":
		conf.UseSQLite3 = true
	case "mssql":
		conf.UseMSSQL = true
	default:
		log.Fatal("unreachable")
	}

	var tables []string
	_ = db.Raw("SHOW TABLES").Scan(&tables)
	log.Debug("tbales = %+v", tables)
	instance = db
}

// Close sql db connect
func Close() {
	sqlDB, err := instance.DB()
	if err != nil {
		log.Error("get *sql.DB error: %s", err)
	} else {
		_ = sqlDB.Close()
	}

	instance = nil
}
