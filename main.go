package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// CORS 미들웨어
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// 컨테이너 저장소 생성
	containerStore := NewContainerStore()

	// 샘플데이터 추가 Container{
	containerStore.AddContainer(&Container{
		ID:               "CONT001",
		Location:         "서울 강남구",
		Status:           "운송중",
		LastUpdated:      time.Now(),
		EstimatedArrival: "2025-01-02",
	})

	containerStore.AddContainer(&Container{
		ID:               "CONT002",
		Location:         "부산 해운대구",
		Status:           "보관중",
		LastUpdated:      time.Now(),
		EstimatedArrival: "2025-01-05",
	})

	// 주문 저장소 생성
	orderStore := NewOrderStore()

	// 컨테이너 API 핸들러 등록
	http.HandleFunc("/api/containers", GetContainerHandler(containerStore))

	// 주문 API 핸들러 등록
	// http.HandleFunc("/api/orders", CreateOrderHandler(orderStore))
	// http.HandleFunc("/api/orders", GetOrdersHandler(orderStore))
	// http.HandleFunc("/api/orders/all", GetAllOrdersHandler(orderStore))

	// 주문API 핸들러 통합 등록
	// 주문 API
	http.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			CreateOrderHandler(orderStore)(w, r)
		case http.MethodGet:
			orderID := r.URL.Query().Get("order_id")
			customerName := r.URL.Query().Get("customer_name")
			status := r.URL.Query().Get("status")

			if orderID != "" {
				GetOrdersHandler(orderStore)(w, r)
			} else if customerName != "" || status != "" {
				SearchOrdersHandler(orderStore)(w, r)
			} else {
				GetAllOrdersHandler(orderStore)(w, r)
			}
		case http.MethodPut:
			UpdateOrderHandler(orderStore)(w, r)
		case http.MethodDelete:
			DeleteOrderHandler(orderStore)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 서버 시작
	fmt.Println("서버시작: http://localhost8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux)))
}
