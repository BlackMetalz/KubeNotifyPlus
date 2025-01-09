package utils

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
	"time"
)

var DB *sql.DB

// InitMySQL initializes the MySQL connection
func InitMySQL() error {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		AppConfig.MySQL.DBUser,     // Database user
		AppConfig.MySQL.DBPassword, // Database password
		AppConfig.MySQL.DBHost,     // Database host
		AppConfig.MySQL.DBName,     // Database name
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }

    // Set the maximum number of open connections to the database
    DB.SetMaxOpenConns(AppConfig.MySQL.MaxConnection)

    // Set the maximum number of idle connections in the pool
    DB.SetMaxIdleConns(AppConfig.MySQL.MaxIdleConnection)

    // Set the maximum lifetime of a connection
    DB.SetConnMaxLifetime(time.Duration(AppConfig.MySQL.ConnMaxLifetime) * time.Second)

    // Set the maximum idle time of a connection
    DB.SetConnMaxIdleTime(time.Duration(AppConfig.MySQL.ConnMaxIdleTime) * time.Second)


    // Verify the connection
    err = DB.Ping()
    if err != nil {
        return err
    }

    log.Println("MySQL connection established")
    return nil
}


// GetTelegramID retrieves the telegram_id from the users table based on userID
func GetTelegramID(userID int) (string, error) {
    var telegramID string
    query := "SELECT telegram_id FROM users WHERE id = ?"
    err := DB.QueryRow(query, userID).Scan(&telegramID)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", fmt.Errorf("no user found with id %d", userID)
        }
        return "", err
    }
    return telegramID, nil
}

// GetServiceMappings retrieves the service mappings and calls GetTelegramID to get the telegramID
func GetServiceMappings(serviceName string) (string, error) {
    var userID int
    query := "SELECT user_id FROM service_mappings WHERE service_name = ?"
    err := DB.QueryRow(query, serviceName).Scan(&userID)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", fmt.Errorf("no service mapping found for service %s", serviceName)
        }
        return "", err
    }

    telegramID, err := GetTelegramID(userID)
    if err != nil {
        return "", err
    }

    return telegramID, nil
}