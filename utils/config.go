package utils

import (
    "encoding/json"
    "os"
)

type MySQLConfig struct {
    DBUser     string `json:"dbUser"`
    DBPassword string `json:"dbPassword"`
    DBHost     string `json:"dbHost"`
    DBName     string `json:"dbName"`
    MaxConnection   int    `json:"maxConnection"`
    MaxIdleConnection int  `json:"maxIdleConnection"`
    ConnMaxLifetime int    `json:"connMaxLifetime"`
    ConnMaxIdleTime int    `json:"connMaxIdleTime"`
}

type Config struct {
    EnableDebug bool      `json:"enableDebug"`
    AppName   string      `json:"appName"`
    ChatId    string      `json:"chatId"`
    BotToken  string      `json:"botToken"`
    K8sToken  string      `json:"k8sToken"`
    ApiServer string      `json:"apiServer"`
    MySQL     MySQLConfig `json:"mysql"`
    PodFailureInvestigation bool `json:"pod_failure_investigate"`
}

// AppConfig represents the configuration of the application
var AppConfig Config

func LoadConfig() error {
    file, err := os.Open("config.json") // Open the config.json file
    if err != nil {
        return err
    }
    // Close the file when the function returns
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&AppConfig)
    if err != nil {
        return err
    }

    return nil
}