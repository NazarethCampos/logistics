package main

import "time"

type Container struct {
	ID               string    `json:"id"`
	Location         string    `json:"status"`
	Status           string    `json:"last_update"`
	LastUpdated      time.Time `json:"last_updated"`
	EstimatedArrival string    `json:"estimated_arrival"`
}

type ContainerStore struct {
	containers map[string]*Container
}

func NewContainerStore() *ContainerStore {
	return &ContainerStore{
		containers: make(map[string]*Container),
	}
}

func (cs *ContainerStore) AddContainer(container *Container) {
	cs.containers[container.ID] = container
}

func (cs *ContainerStore) GetContainer(id string) *Container {
	return cs.containers[id]
}
