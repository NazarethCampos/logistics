package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func GetContainerHandler(store *ContainerStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containerID := r.URL.Query().Get("id")

		if containerID == "" {
			http.Error(w, "Container ID is required", http.StatusBadRequest)
			return
		}

		container := store.GetContainer(containerID)

		if container == nil {
			http.Error(w, "Container Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(container)
	}
}

// func CreateOrderHandler(orderStore *OrderStore) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodPost {
// 			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 			return
// 		}

// 		var order Order
// 		err := json.NewDecoder(r.Body).Decode(&order)

// 		if err != nil {
// 			http.Error(w, "Invalid request body", http.StatusBadRequest)
// 			return
// 		}

// 		// 데이터 검증
// 		if order.OrderID == "" || order.CustomerName == "" || order.Destination == "" {
// 			http.Error(w, "Missing required fields", http.StatusBadRequest)
// 			return
// 		}

// 		order.Status = "주문접수"
// 		order.CreatedAt = time.Now()

// 		orderStore.AddOrder(&order)

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(order)
// 	}
// }

func CreateOrderHandler(orderStore *OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var order Order
		err := json.NewDecoder(r.Body).Decode(&order)

		if err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		// 데이터 검증
		if order.OrderID == "" || order.CustomerName == "" || order.Destination == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		order.Status = "주문 접수"
		order.CreatedAt = time.Now()

		orderStore.AddOrder(&order)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(order)
	}
}

func GetOrdersHandler(orderStore *OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Query().Get("order_id")

		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}

		order := orderStore.GetOrder(orderID)

		if order == nil {
			http.Error(w, "Order Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	}
}

func GetAllOrdersHandler(orderStore *OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders := orderStore.GetAllOrders()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}

func UpdateOrderHandler(orderStore *OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		orderID := r.URL.Query().Get("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}

		var order Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		order.OrderID = orderID
		success := orderStore.UpdateOrder(orderID, &order)
		if !success {
			http.Error(w, "Order Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	}
}

func DeleteOrderHandler(orderStore *OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		orderID := r.URL.Query().Get("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}

		success := orderStore.DeleteOrder(orderID)
		if !success {
			http.Error(w, "Order Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Order deleted successfully"}`))
	}
}

// 검색필터링
func SearchOrdersHandler(orderStore *OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerName := r.URL.Query().Get("customer_name")
		status := r.URL.Query().Get("status")

		orders := orderStore.SearchOrders(customerName, status)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}

// 배송정보 생성
func CreateShipmentHandler(shipmentStore *ShipmentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var shipment Shipment
		err := json.NewDecoder(r.Body).Decode(&shipment)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		// 데이터 검증
		if shipment.ID == "" || shipment.OrderID == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		shipment.CreatedAt = time.Now()
		shipmentStore.AddShipment(&shipment)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(shipment)
	}
}

// 배송정보 조회 핸들러
func GetShipmentHandler(shipmentStore *ShipmentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shipmentID := r.URL.Query().Get("shipment_id")
		if shipmentID == "" {
			http.Error(w, "Shipment ID is required", http.StatusBadRequest)
			return
		}

		shipment := shipmentStore.GetShipment(shipmentID)
		if shipment == nil {
			http.Error(w, "Shipment Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(shipment)
	}
}

// 주문ID로 배송정보 조회 핸들러
func GetShipmentByOrderHandler(shipmentStore *ShipmentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Query().Get("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		shipment := shipmentStore.GetShipmentByOrder(orderID)
		if shipment == nil {
			http.Error(w, "Shipment Not Found for the given Order ID", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(shipment)
	}
}
