package dao

import "time"

type ServiceEntity struct {
  ID          uint32  `gorm:"primary_key"`
  Name        string
  Description string
  Role        string
  Destination string
  NoticeType
  CreatedAt   time.Time
}
