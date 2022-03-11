package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Keshav-Agrawal/mongoapi/datasource"

	"github.com/Keshav-Agrawal/mongoapi/model"
	"github.com/gorilla/mux"
)

type HomeworkSVC interface {
	GetMyAllTask(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	MarkAsDone(w http.ResponseWriter, r *http.Request)
	DeleteATask(w http.ResponseWriter, r *http.Request)
	DeleteAllTask(w http.ResponseWriter, r *http.Request)
}
type homeworkService struct {
	ds datasource.IDataSource
}

func NewHomeWorkService(ds datasource.IDataSource) HomeworkSVC {
	return &homeworkService{
		ds: ds,
	}
}

func (h homeworkService) GetMyAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	allTask, err := h.ds.GetAllTask()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(allTask)
}

func (h homeworkService) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")

	var work model.Homework
	_ = json.NewDecoder(r.Body).Decode(&work)
	err := h.ds.InsertOneTask(work)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(work)

}

func (h homeworkService) MarkAsDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)
	h.ds.UpdateOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func (h homeworkService) DeleteATask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	params := mux.Vars(r)
	h.ds.DeleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func (h homeworkService) DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")

	count := h.ds.DeleteAllTask()
	json.NewEncoder(w).Encode(count)
}
