package config

// ENV ...
type ENV struct {
	// Common
	Env string

	// Port
	Port struct {
		App string
	}

	// Database
	Database struct {
		URI  string
		Name string
		Auth struct {
			Mechanism string
			Source    string
			Username  string
			Password  string
		}
	}

	// Redis
	Redis struct {
		URI      string
		Password string
	}

	// Auth
	Auth struct {
		SecretKey string
	}
}
