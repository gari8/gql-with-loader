package entity

import (
	"fmt"
	"io"
	"strconv"
)

type Member struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Status  Status `json:"status"`
	GroupID string `json:"groupId"`
}

type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
	StatusRetired  Status = "RETIRED"
)

var AllStatus = []Status{
	StatusActive,
	StatusInactive,
	StatusRetired,
}

func (e Status) Atoi() int {
	for i, v := range AllStatus {
		if e == v {
			return i
		}
	}
	return -1
}

func (e Status) IsValid() bool {
	switch e {
	case StatusActive, StatusInactive, StatusRetired:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
