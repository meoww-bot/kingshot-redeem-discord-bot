package main

import (
	"time"
)

const (
	BASE_URL       = "https://kingshot-giftcode.centurygame.com/api"
	SECRET         = "mN4!pQs6JrYwV9"
	MAX_RETRIES    = 20
	RETRY_DELAY    = 8 * time.Second
	DELAY_DURATION = 2 * time.Second
)
