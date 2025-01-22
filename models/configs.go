package models

type AppConfig struct {
	LogParams      LogParams      `json:"log_params"`      // Log parameters
	AppParams      AppParams      `json:"app_params"`      // Application parameters
	PostgresParams PostgresParams `json:"postgres_params"` // PostgreSQL database parameters
}

type LogParams struct {
	LogDirectory     string `json:"log_directory"`      // Directory for logs
	LogInfo          string `json:"log_info"`           // Log for informational messages
	LogError         string `json:"log_error"`          // Log for error messages
	LogWarn          string `json:"log_warn"`           // Log for warning messages
	LogDebug         string `json:"log_debug"`          // Log for debug messages
	MaxSizeMegabytes int    `json:"max_size_megabytes"` // Maximum log file size in megabytes
	MaxBackups       int    `json:"max_backups"`        // Maximum number of log backups
	MaxAge           int    `json:"max_age"`            // Maximum age of logs in days
	Compress         bool   `json:"compress"`           // Whether to compress logs
	LocalTime        bool   `json:"local_time"`         // Whether to use local time for logs
}

type AppParams struct {
	GinMode    string `json:"gin_mode"`     // Gin mode (e.g., debug or release)
	PortRun    string `json:"port_run"`     // Port on which the server will run
	ApiPortRun string `json:"api_port_run"` // Port on which the api will run
	ServerURL  string `json:"server_url"`   // Server URL
	ServerName string `json:"server_name"`  // Server name
	ApiURL     string `json:"api_url"`      // API URL
}

type PostgresParams struct {
	Host     string `json:"host"`     // Database host
	Port     string `json:"port"`     // Database port
	User     string `json:"user"`     // Database username
	Database string `json:"database"` // Database name
}
