package models

import "time"

type Ticket struct {
	ID        int       `json:"id" db:"id"`                 // 티켓 ID
	Name      string    `json:"name" db:"name"`             // 이벤트 이름
	Status    int       `json:"status" db:"status"`         // 티켓 상태
	StartDate time.Time `json:"start_date" db:"start_date"` // 이벤트 시작 날짜
	EndDate   time.Time `json:"end_date" db:"end_date"`     // 이벤트 종료 날짜
	Capacity  int       `json:"capacity" db:"capacity"`     // 인원수 제한
}
