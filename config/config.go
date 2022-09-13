package config

import (
	"fmt"
	"os"
	"strconv"
)

// AppConfig is the set of all possible configuration options
// that our application will ever need. When starting up, all
// of these need to have some sort of a value. If anyone of them
// is missing, there's no point in continuing startup.
//
// This does not mean that, in this case, a database connection
// would _succeed_, this only means that the rest of the application
// has enough information to attempt connecting to the database.
//
// Whether that finishes or not (bad password, bad port, bad host)
// is another question.
type AppConfig struct {
	DBhost string
	DBport int
	DBuser string
	DBpass string
	DBname string
}

// Parse reads the environment variables, and does some really basic
// validation as well.
func Parse() (AppConfig, error) {
	host, err := getString("DB_HOST")
	if err != nil {
		return AppConfig{}, err
	}

	user, err := getString("DB_USER")
	if err != nil {
		return AppConfig{}, err
	}

	pass, err := getString("DB_PASS")
	if err != nil {
		return AppConfig{}, err
	}

	name, err := getString("DB_NAME")
	if err != nil {
		return AppConfig{}, err
	}

	port, err := getPort("DB_PORT")
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		DBhost: host,
		DBport: port,
		DBuser: user,
		DBpass: pass,
		DBname: name,
	}, nil
}

func getString(varName string) (string, error) {
	vn := os.Getenv(varName)
	if vn == "" {
		return "", fmt.Errorf("missing value for env var: %s", varName)
	}

	return vn, nil
}

func getPort(varName string) (int, error) {
	port := os.Getenv(varName)
	if port == "" {
		return 0, fmt.Errorf("missing value for env var: %s", varName)
	}

	// convert the string into an integer (Atoi stands for "ASCII to integer")
	p, err := strconv.Atoi(port)
	if err != nil {
		return 0, fmt.Errorf("could not convert value %s for env var %s to integer", port, varName)
	}

	if p < 1025 || p > 65535 {
		return 0, fmt.Errorf("port is out of bounds, expected 1025 - 65535, got %d", p)
	}

	return p, nil
}
