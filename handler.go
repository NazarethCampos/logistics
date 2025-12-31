package main

import (
	"encoding/json"
	"net/http"
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
