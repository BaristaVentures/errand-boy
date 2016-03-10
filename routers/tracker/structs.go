package tracker

// ActivityPayload represents a Pivotal Tracker webhook request's body.
type ActivityPayload struct {
	Highlight        string      `json:"highlight"`
	PrimaryResources []*Resource `json:"primary_resources"`
	Actor            *Actor      `json:"performed_by"`
}

// Resource is a Pivotal Tracker resource.
type Resource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	ID   int    `json:"id"`
}

// Actor represents the activity's Actor
type Actor struct {
	Name string `json:"name"`
}
