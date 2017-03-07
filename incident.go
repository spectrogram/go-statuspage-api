package statuspage

import (
	"fmt"
	"time"
)

type IncidentUpdate struct {
	Body               *string    `json:"body,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	DisplayAt          *time.Time `json:"display_at,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncidentID         *string    `json:"incident_id,omitempty"`
	Status             *string    `json:"status,omitempty"`
	TwitterUpdatedAt   *time.Time `json:"twitter_updated_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	WantsTwitterUpdate *bool      `json:"wants_twitter-update,omitempty"`
}

type Incident struct {
	Backfilled                    *bool             `json:"backfilled,omitempty"`
	Components                    *[]Component      `json:"components,omitempty"`
	CreatedAt                     *time.Time        `json:"created_at,omitempty"`
	ID                            *string           `json:"id,omitempty"`
	Impact                        *string           `json:"impact,omitempty"`
	ImpactOverride                *string           `json:"impact_override,omitempty"`
	IncidentUpdates               []*IncidentUpdate `json:"incident_updates,omitempty"`
	MonitoringAt                  *time.Time        `json:"monitoring_at,omitempty"`
	Name                          *string           `json:"name,omitempty"`
	PageID                        *string           `json:"page_id,omitempty"`
	PostmortemBody                *string           `json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt   *time.Time        `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored             *bool             `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers *bool             `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter     *bool             `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt         *time.Time        `json:"postmorem_published_at,omitempty"`
	ResolvedAt                    *time.Time        `json:"resolved_at,omitempty"`
	ScheduledAutoInProgress       *bool             `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledAutoCompleted        *bool             `json:"scheduled_auto_completed,omitempty"`
	ScheduledFor                  *time.Time        `json:"scheduled_for"`
	ScheduledRemindPrior          *bool             `json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt           *time.Time        `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil                *time.Time        `json:"scheduled_until,omitempty"`
	Shortlink                     *string           `json:shortlink,omitempty"`
	Status                        *string           `json:status,omitempty"`
	UpdatedAt                     *time.Time        `json:updated_at,omitempty"`
}

type IncidentResponse struct {
	Offset *int       `json:"offset,omitempty"`
	Limit  *int       `json:"limit,omitempty"`
	Total  *int       `json:"total,omitempty"`
	Data   []Incident `json:"data,omitempty"`
}

type NewIncident struct {
	Name               string
	Status             string
	Message            string
	WantsTwitterUpdate bool
	ImpactOverride     string
	ComponentIDs       []string
}

func (i *NewIncident) String() string {
	return encodeParams(map[string]interface{}{
		"incident[name]":                 i.Name,
		"incident[status]":               i.Status,
		"incident[wants_twitter_update]": i.WantsTwitterUpdate,
		"incident[message]":              i.Message,
		"incident[impact_override]":      i.ImpactOverride,
		"incident[component_ids][]":      i.ComponentIDs,
	})
}

// TODO: Paging
func (c *Client) doGetIncidents(path string) ([]Incident, error) {
	resp := &IncidentResponse{}
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) GetAllIncidents() ([]Incident, error) {
	return c.doGetIncidents("incidents.json")
}

func (c *Client) GetOpenIncidents() ([]Incident, error) {
	return c.doGetIncidents("incidents/unresolved.json")
}

func (c *Client) GetScheduledIncidents() ([]Incident, error) {
	return c.doGetIncidents("incidents/scheduled.json")
}

func (c *Client) CreateIncident(name, status, message, component string) ([]Incident, error) {
	switch status {
	case "investigating", "identified", "monitoring", "resolved":
		break
	default:
		return nil, fmt.Errorf("create error: status not (investigating|identified|monitoring|resolved), got %s", status)
	}
	cp, err := c.GetComponentByName(component)
	if err != nil {
		return nil, err
	}
	i := &NewIncident{
		Name:               name,
		Status:             status,
		Message:            message,
		WantsTwitterUpdate: false,
		ImpactOverride:     "none",
		ComponentIDs:       []string{*cp.ID},
	}
	resp := &IncidentResponse{}
	err = c.doPost("incidents.json", i, resp)
	if err != nil {
		return nil, err
	}
	fmt.Println(i)
	return nil, nil
}
