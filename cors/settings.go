package cors

var DefaultSettings = Settings{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE"},
	AllowedHeaders: []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"Accept",
		"Origin",
		"Cache-Control",
		"X-Requested-With",
	},
	ExposedHeaders: []string{"Content-Range"},
}

type Settings struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	ExposedHeaders []string
}
