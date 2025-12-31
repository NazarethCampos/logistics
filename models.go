package main

import "time"

type Container struct {
	ID               string    `json:"id"`
	Location         string    `json:"status"`
	Status           string    `json:"last_update"`
	LastUpdated      time.Time `json:"last_updated"`
	EstimatedArrival string    `json:"estimated_arrival"`
}

type Order struct {
	OrderID      string    `json:"order_id"`
	ContainerID  string    `json:"container_id"`
	CustomerName string    `json:"customer_name"`
	Destination  string    `json:"destination"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type ContainerStore struct {
	containers map[string]*Container
}

type OrderStore struct {
	orders map[string]*Order
}

func NewContainerStore() *ContainerStore {
	return &ContainerStore{
		containers: make(map[string]*Container),
	}
}

func NewOrderStore() *OrderStore {
	return &OrderStore{
		orders: make(map[string]*Order),
	}
}

func (cs *ContainerStore) AddContainer(container *Container) {
	cs.containers[container.ID] = container
}

func (cs *ContainerStore) GetContainer(id string) *Container {
	return cs.containers[id]
}

func (os *OrderStore) AddOrder(order *Order) {
	os.orders[order.OrderID] = order
}

func (os *OrderStore) GetOrder(id string) *Order {
	return os.orders[id]
}

func (os *OrderStore) GetAllOrders() []*Order {
	orders := []*Order{}
	for _, order := range os.orders {
		orders = append(orders, order)
	}

	return orders
}
