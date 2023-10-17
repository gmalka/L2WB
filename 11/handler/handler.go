package handler

import (
	"L10WB3/model"
	"log"
	"net/http"
	"time"
)

type EventStorage interface {
	AddEvent(username string, e model.Event) error
	UpdateEvent(username string, e model.Event) error
	DeleteEvent(username string, event_id int) error
	FindDayEvents(username string, date time.Time) ([]model.Event, error)
	FindWeekEvents(username string, date time.Time) ([]model.Event, error)
	FindMonthEvents(username string, date time.Time) ([]model.Event, error)
}

func NewHandler(e EventStorage, errorLog *log.Logger, infoLog *log.Logger) Handler {
	return Handler{
		e: e,

		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

type Handler struct {
	e EventStorage

	errorLog *log.Logger
	infoLog  *log.Logger
}

func (h *Handler) InitHandler() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("/create_event", h.Logging(http.HandlerFunc(h.CreateEvent)))
	r.Handle("/update_event", h.Logging(http.HandlerFunc(h.UpdateEvent)))
	r.Handle("/delete_event", h.Logging(http.HandlerFunc(h.DeleteEvent)))
	r.Handle("/events_for_day", h.Logging(http.HandlerFunc(h.EventsForDay)))
	r.Handle("/events_for_week", h.Logging(http.HandlerFunc(h.EventsForWeek)))
	r.Handle("/events_for_month", h.Logging(http.HandlerFunc(h.EventsForMonth)))

	return r
}

func (h *Handler) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedWriter := &customWriter{
			ResponseWriter: w,
			code: http.StatusOK,
		}
		next.ServeHTTP(wrappedWriter, r)
		h.infoLog.Printf("%s %s %s Status %d", r.Method, r.RequestURI, time.Since(start), wrappedWriter.code)
	})
}

type customWriter struct {
	http.ResponseWriter
	code int
}

func (w *customWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}