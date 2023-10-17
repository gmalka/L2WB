package memorystorage

import (
	"L10WB3/model"
	"fmt"
	"time"
)

type MemoryStorage struct {
	storage map[string]model.User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		storage: make(map[string]model.User),
	}
}

func (m *MemoryStorage) AddEvent(username string, e model.Event) error {
	if v, ok := m.storage[username]; ok {
		if _, ok := v.Events[e.ID]; !ok {
			v.Events[e.ID] = e
			m.storage[username] = v
		} else {
			return fmt.Errorf("event %s with id %d already exists", e.Name, e.ID)
		}
	} else {
		m.storage[username] = model.User{
			Name:   username,
			Events: make(map[int]model.Event, 10),
		}

		m.storage[username].Events[e.ID] = e
	}

	return nil
}

func (m *MemoryStorage) UpdateEvent(username string, e model.Event) error {
	if v, ok := m.storage[username]; ok {
		if _, ok := v.Events[e.ID]; !ok {
			return fmt.Errorf("event %s with id %d doesn't exists", e.Name, e.ID)
		} else {
			v.Events[e.ID] = e
			m.storage[username] = v
		}
	} else {
		return fmt.Errorf("user %s doesn't exists", username)
	}

	return nil
}

func (m *MemoryStorage) DeleteEvent(username string, event_id int) error {
	if v, ok := m.storage[username]; ok {
		if _, ok := v.Events[event_id]; !ok {
			return fmt.Errorf("event %d doesn't exists", event_id)
		} else {
			delete(v.Events, event_id)
		}
	} else {
		return fmt.Errorf("user %s doesn't exists", username)
	}

	return nil
}

func (m *MemoryStorage) FindDayEvents(username string, date time.Time) ([]model.Event, error) {
	result := make([]model.Event, 0, 10)

	if v, ok := m.storage[username]; ok {
		for _, v := range v.Events {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() && v.Date.Day() == date.Day() {
				result = append(result, v)
			}
		}
	} else {
		return nil, fmt.Errorf("user %s doesn't exists", username)
	}

	return result, nil
}

func (m *MemoryStorage) FindWeekEvents(username string, date time.Time) ([]model.Event, error) {
	result := make([]model.Event, 0, 10)

	weekStart := date.AddDate(0, 0, -int(date.Weekday()))
	weekEnd := weekStart.AddDate(0, 0, 6)

	if v, ok := m.storage[username]; ok {
		for _, v := range v.Events {
			if v.Date.After(weekStart) && v.Date.Before(weekEnd) {
				result = append(result, v)
			}
		}
	} else {
		return nil, fmt.Errorf("user %s doesn't exists", username)
	}

	return result, nil
}

func (m *MemoryStorage) FindMonthEvents(username string, date time.Time) ([]model.Event, error) {
	result := make([]model.Event, 0, 10)

	if v, ok := m.storage[username]; ok {
		for _, v := range v.Events {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
				result = append(result, v)
			}
		}
	} else {
		return nil, fmt.Errorf("user %s doesn't exists", username)
	}

	return result, nil
}
