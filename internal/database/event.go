package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Event struct {
	bun.BaseModel  `bun:"table:events"`
	Id             string     `bun:"id,pk"`
	TenantId       string     `bun:"tenant_id"`
	StreamId       string     `bun:"stream_id"`
	Data           []byte     `bun:"data"`
	IdempotencyKey *string    `bun:"idempotency_key"`
	CustomHeaders  []byte     `bun:"custom_headers"`
	ScheduledAt    *time.Time `bun:"scheduled_at"`
	CreatedAt      time.Time  `bun:"created_at"`
}
