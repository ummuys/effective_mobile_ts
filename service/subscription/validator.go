package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func ymtoymd(date string) (string, error) {
	parsed, err := time.Parse("01-2006", date)
	if err != nil {
		return "", err
	}
	return time.Date(parsed.Year(), parsed.Month(), 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02"), nil
}

func ymdtoym(date time.Time) string {
	return date.Format("01-2006")
}

func validServiceName(name string) error {
	if name == "" {
		return fmt.Errorf("bad service_name")
	}
	return nil
}

func validUserID(id string) error {
	if id == "" {
		return fmt.Errorf("bad user_id")
	}

	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("bad user_id")
	}

	return nil
}

func validPrice(price int) error {
	if price < 0 {
		return fmt.Errorf("bad price")
	}

	return nil
}
