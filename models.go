package main

import (
	"strings"
	"time"
)

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

type Shipment struct {
	ID           string    `json:"shipment_id"`
	OrderID      string    `json:"order_id"`
	CarrierName  string    `json:"carrier_name"`
	TrackingNum  string    `json:"tracking_number"`
	CurrentLoc   string    `json:"current_location"`
	EstimatedDel string    `json:"estimated_delivery"`
	CreatedAt    time.Time `json:"created_at"`
}

type ContainerStore struct {
	containers map[string]*Container
}

type OrderStore struct {
	orders map[string]*Order
}

type ShipmentStore struct {
	shipments map[string]*Shipment
}

// 컨테이너
func NewContainerStore() *ContainerStore {
	return &ContainerStore{
		containers: make(map[string]*Container),
	}
}

// 주문
func NewOrderStore() *OrderStore {
	return &OrderStore{
		orders: make(map[string]*Order),
	}
}

// 배송
func NewShipmentStore() *ShipmentStore {
	return &ShipmentStore{
		shipments: make(map[string]*Shipment),
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

func (ss *ShipmentStore) AddShipment(shipment *Shipment) {
	ss.shipments[shipment.ID] = shipment
}

func (ss *ShipmentStore) GetShipment(id string) *Shipment {
	return ss.shipments[id]
}

func (os *OrderStore) GetAllOrders() []*Order {
	orders := []*Order{}
	for _, order := range os.orders {
		orders = append(orders, order)
	}

	return orders
}

func (os *OrderStore) UpdateOrder(order_id string, order *Order) bool {
	if _, exists := os.orders[order_id]; exists {
		os.orders[order_id] = order
		return true
	}
	return false
}

func (os *OrderStore) DeleteOrder(order_id string) bool {
	if _, exists := os.orders[order_id]; exists {
		delete(os.orders, order_id)
		return true
	}
	return false
}

// 주문검색/필터링 기능
func (os *OrderStore) SearchOrders(customerName, status string) []*Order {
	var results []*Order
	for _, order := range os.orders {
		// 고객이름과 상태로 필터링
		if customerName != "" && !contains(order.CustomerName, customerName) {
			continue
		}

		// 상태로 검색
		if status != "" && order.Status != status {
			continue
		}
		results = append(results, order)
	}
	return results
}

func (ss *ShipmentStore) GetShipmentByOrder(orderID string) *Shipment {
	for _, shipment := range ss.shipments {
		if shipment.OrderID == orderID {
			return shipment
		}
	}
	return nil
}

func contains(source, substr string) bool {
	//return len(source) >= len(substr) && (source == substr || len(source) > len(substr) && (source[0:len(substr)] == substr || contains(source[1:], substr)))
	return strings.Contains(source, substr)
}
