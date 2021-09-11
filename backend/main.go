//menjalankan aksi yang ada pada controller
//sumber belajar 1. https://kodingin.com/golang-koneksi-database-mysql/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"encoding/json"
	"NewsAPISPE/backend/controller"
	"NewsAPISPE/backend/models"
	"NewsAPISPE/backend/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	handleRequest()

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//tukang panggil fungsi
func handleRequest() {

	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "application/json; charset=UTF-8"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/bisnis", getBisnis).Methods("GET")
	router.HandleFunc("/detailbisnis", getSelectedBisnis).Methods("GET")

	router.HandleFunc("/health", getHealth).Methods("GET")
	router.HandleFunc("/detailhealth", getSelectedHealth).Methods("GET")

	router.HandleFunc("/entertainment", getEntertainment).Methods("GET")
	router.HandleFunc("/detailentertainment", getSelectedEntertainment).Methods("GET")

	fmt.Println("Terhubung ke 8090")
	http.ListenAndServe(":8090", handlers.CORS(headers, methods, origins)(router))
}

//fungsi middleware cek isian dan method GET untuk ambil data
func getBisnis(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		articles, err := controller.GetBisnis(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, articles, http.StatusOK)
		return
	}

	http.Error(w, "Tidak diijinkan", http.StatusNotFound)
	return
}

func getSelectedBisnis(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var art models.Article

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		art.Id, _ = strconv.Atoi(id)

		articles, err := controller.GetSelectedBisnis(ctx, art)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, articles, http.StatusOK)
		return

	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		articles, err := controller.GetHealth(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, articles, http.StatusOK)
		return
	}

	http.Error(w, "Tidak diijinkan", http.StatusNotFound)
	return
}

func getSelectedHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var art models.Article

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		art.Id, _ = strconv.Atoi(id)

		articles, err := controller.GetSelectedHealth(ctx, art)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, articles, http.StatusOK)
		return

	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func getEntertainment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		articles, err := controller.GetEntertainment(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, articles, http.StatusOK)
		return
	}

	http.Error(w, "Tidak diijinkan", http.StatusNotFound)
	return
}

func getSelectedEntertainment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var art models.Article

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		art.Id, _ = strconv.Atoi(id)

		articles, err := controller.GetSelectedEntertainment(ctx, art)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, articles, http.StatusOK)
		return

	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
