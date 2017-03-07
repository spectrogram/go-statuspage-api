package statuspage

import (
	"time"
)

type IncidentUpdate struct {
	Body               *string
	CreatedAt          *time.Time
	DisplayAt          *time.Time
	ID                 *string
	IncidentID         *string
	Status             *string
	TwitterUpdatedAt   *time.Time
	UpdatedAt          *time.Time
	WantsTwitterUpdate *bool
}

type Incident struct {
	Backfilled                    *bool
	CreatedAt                     *time.Time
	ID                            *string
	Impact                        *string
	ImpactOverride                *string
	IncidentUpdates               []*IncidentUpdate
	MonitoringAt                  *time.Time
	Name                          *string
	PageID                        *string
	PostmortemBody                *string
	PostmortemBodyLastUpdatedAt   *time.Time
	PostmortemIgnored             *bool
	PostmortemNotifiedSubscribers *bool
	PostmortemNotifiedTwitter     *bool
	PostmortemPublishedAt         *time.Time
	ResolvedAt                    *time.Time
	ScheduledAutoInProgress       *bool
	ScheduledAutoCompleted        *bool
	ScheduledFor                  *time.Time
	ScheduledRemindPrior          *bool
	ScheduledRemindedAt           *time.Time
	ScheduledUntil                *time.Time
	Shortlink                     *string
	Status                        *string
	UpdatedAt                     *time.Time
}
