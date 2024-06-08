package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"taskmanager/db"
	"taskmanager/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	//decode incoming body params to task model
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//create task and send created task
	db.DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	//encode to json and send task
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task

	result := db.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Save(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	result := db.DB.First(&task, id)

	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
