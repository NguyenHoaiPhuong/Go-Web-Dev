package rabbitmq

// RabbitMQName : queue name
type RabbitMQName string

// Rabbitmq: define all of queues name
const (
	QueueScanService         RabbitMQName = "ScanService"
	QueueNotificationService RabbitMQName = "NotificationService"
	QueueEmailService        RabbitMQName = "EmailService"
	QueueProgramService      RabbitMQName = "ProgramService"
)
