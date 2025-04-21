package config

const (
	MongoURI     = "mongodb://localhost:27017"
	DatabaseName = "auth_service"
	SessionCol   = "sessions"
	RabbitMQURL  = "amqp://guest:guest@localhost:5672/"
	SessionQueue = "session_events"
)
