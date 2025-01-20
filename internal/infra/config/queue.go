package config

var (
	QUEUE_PROCESS_VIDEO             = GetEnv("QUEUE_PROCESS_VIDEO", "process-video-queue")
	DEAD_LETTER_QUEUE_PROCESS_VIDEO = GetEnv("DEAD_LETTER_QUEUE_PROCESS_VIDEO", "process-video-dead-letter-queue")
)
