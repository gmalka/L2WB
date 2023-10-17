package handler

import (
	"L10WB3/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type createEventReq struct {
	Username string `json:"username"`
	EventId  int    `json:"event_id"`
	Name     string `json:"event_name"`
	Date     string `json:"date"`
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	er := createEventReq{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.errorLog.Printf("Can't Read from body: %v\n", err)
		http.Error(w, "503 Server error", http.StatusServiceUnavailable)
		return
	}

	err = json.Unmarshal(b, &er)
	if err != nil {
		h.errorLog.Printf("Can't Unmarshal from body: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", er.Date)
	if err != nil {
		h.errorLog.Printf("Can't Parse date: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	event := model.Event{
		ID: er.EventId,
		Name: er.Name,
		Date: date,
	}
	err = h.e.AddEvent(er.Username, event)
	if err != nil {
		h.errorLog.Printf("Can't Add event: %v\n", err)
		http.Error(w, fmt.Sprintf("400 %s", err), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Success created event %s for user %s", event.Name, er.Username)))
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	er := createEventReq{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.errorLog.Printf("Can't Read from body: %v\n", err)
		http.Error(w, "503 Server error", http.StatusServiceUnavailable)
		return
	}

	err = json.Unmarshal(b, &er)
	if err != nil {
		h.errorLog.Printf("Can't Unmarshal from body: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", er.Date)
	if err != nil {
		h.errorLog.Printf("Can't Parse date: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	event := model.Event{
		ID: er.EventId,
		Name: er.Name,
		Date: date,
	}
	err = h.e.UpdateEvent(er.Username, event)
	if err != nil {
		h.errorLog.Printf("Can't Update event: %v\n", err)
		http.Error(w, fmt.Sprintf("400 %s", err), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Success update event %s for user %s", event.Name, er.Username)))
}

type deleteeEventReq struct {
	Username string `json:"username"`
	Id       int    `json:"event_id"`
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	er := deleteeEventReq{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.errorLog.Printf("Can't Read from body: %v\n", err)
		http.Error(w, "500 Server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &er)
	if err != nil {
		h.errorLog.Printf("Can't Unmarshal from body: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	err = h.e.DeleteEvent(er.Username, er.Id)
	if err != nil {
		h.errorLog.Printf("Can't Delete event: %v\n", err)
		http.Error(w, fmt.Sprintf("400 %s", err), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Success delete event by id %d for user %s", er.Id, er.Username)))
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	query := r.URL.Query()
	username := query.Get("username")
	date, err := time.Parse("2006-01-02", query.Get("date"))
	if err != nil {
		h.errorLog.Printf("Can't Parse date: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	events, err := h.e.FindDayEvents(username, date)
	if err != nil {
		h.errorLog.Printf("Can't get events for day: %v\n", err)
		http.Error(w, fmt.Sprintf("503 %s", err), http.StatusServiceUnavailable)
		return
	}

	b, err := json.Marshal(events)
	if err != nil {
		h.errorLog.Printf("Can't Unmarshal from body: %v\n", err)
		http.Error(w, "503 Some server error", http.StatusServiceUnavailable)
		return
	}

	w.Write(b)
}

func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	query := r.URL.Query()
	username := query.Get("username")
	date, err := time.Parse("2006-01-02", query.Get("date"))
	if err != nil {
		h.errorLog.Printf("Can't Parse date: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	events, err := h.e.FindWeekEvents(username, date)
	if err != nil {
		h.errorLog.Printf("Can't get events for day: %v\n", err)
		http.Error(w, fmt.Sprintf("503 %s", err), http.StatusServiceUnavailable)
		return
	}

	b, err := json.Marshal(events)
	if err != nil {
		h.errorLog.Printf("Can't Unmarshal from body: %v\n", err)
		http.Error(w, "503 Some server error", http.StatusServiceUnavailable)
		return
	}

	w.Write(b)
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	query := r.URL.Query()
	username := query.Get("username")
	date, err := time.Parse("2006-01-02", query.Get("date"))
	if err != nil {
		h.errorLog.Printf("Can't Parse date: %v\n", err)
		http.Error(w, "400 Incorrect input data", http.StatusBadRequest)
		return
	}

	events, err := h.e.FindMonthEvents(username, date)
	if err != nil {
		h.errorLog.Printf("Can't get events for day: %v\n", err)
		http.Error(w, fmt.Sprintf("503 %s", err), http.StatusServiceUnavailable)
		return
	}

	b, err := json.Marshal(events)
	if err != nil {
		h.errorLog.Printf("Can't Unmarshal from body: %v\n", err)
		http.Error(w, "503 Some server error", http.StatusServiceUnavailable)
		return
	}

	w.Write(b)
}