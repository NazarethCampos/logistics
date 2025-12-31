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

		w.Header().Set("Content-Type", "appliacation/json")
		json.NewEncoder(w).Encode(container)
	}
}

func CreateOrderHandler(orderStore *OrderStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var order Order
		err := json.NewDecoder(r.Body).Decode(&order)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// 데이터 검증
		if order.OrderID == "" || order.CustomerName == "" || order.Destination == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		order.Status = "주문접수"
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
