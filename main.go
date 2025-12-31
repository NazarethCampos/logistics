package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 컨테이너 저장소 생성
	store := NewContainerStore()

	// 샘플데이터 추가ontainer{
	store.AddContainer(&Container{
		ID:               "CONT001",
		Location:         "서울 강남구",
		Status:           "운송중",
		LastUpdated:      time.Now(),
		EstimatedArrival: "2025-01-02",
	})

	store.AddContainer(&Container{
		ID:               "CONT002",
		Location:         "부산 해운대구",
		Status:           "보관중",
		LastUpdated:      time.Now(),
		EstimatedArrival: "2025-01-05",
	})

	// API 핸들러 등록
	http.HandleFunc("/api/containers", GetContainerHandler(store))

	// 서버 시작
	fmt.Println("서버시작: http://localhost8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
