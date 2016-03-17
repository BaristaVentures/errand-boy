package tracker

import (
	"strconv"

	"github.com/Sirupsen/logrus"
)

// ActivityPayload represents a Pivotal Tracker webhook request's body.
type ActivityPayload struct {
	Highlight        string    `json:"highlight"`
	PrimaryResources Resources `json:"primary_resources"`
	Actor            *Actor    `json:"performed_by"`
	Project          *Project  `json:"project"`
}

// Resource is a Pivotal Tracker resource.
type Resource struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	URL  string `json:"url"`
	ID   int    `json:"id"`
}

// Resources is a *Resource array.
type Resources []*Resource

// Actor represents the activity's Actor
type Actor struct {
	Name string `json:"name"`
}

// Project is the project where the activity happened.
type Project struct {
	ID int `json:"id"`
}

// GetContext implements logging.Logger for ActivityPayload
func (act *ActivityPayload) GetContext() logrus.Fields {
	fields := logrus.Fields{
		"highlight": act.Highlight,
		"resources": act.PrimaryResources.GetContext(),
		"actor":     act.Actor.GetContext(),
		"project":   act.Project.GetContext(),
	}
	return fields
}

// GetContext implements logging.Logger for Resource
func (res *Resource) GetContext() logrus.Fields {
	fields := logrus.Fields{
		"name": res.Name,
		"kind": res.Kind,
		"url":  res.URL,
		"id":   res.ID,
	}
	return fields
}

// GetContext implements logging.Logger for Resource
func (res *Resources) GetContext() logrus.Fields {
	fields := logrus.Fields{}
	for i, r := range *res {
		fields[strconv.Itoa(i)] = r.GetContext()
	}
	return fields
}

// GetContext implements logging.Logger for Actor
func (actor *Actor) GetContext() logrus.Fields {
	return logrus.Fields{"name": actor.Name}
}

// GetContext implements logging.Logger for Project
func (proj *Project) GetContext() logrus.Fields {
	return logrus.Fields{"id": proj.ID}
}
