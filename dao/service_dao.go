package dao

import "time"

type ServiceEntity struct {
  ID          uint32
  Name        string
  Description string
  Role        string
  Destination string
  NoticeType  string
  CreatedAt   time.Time
}

func (ServiceEntity) TableName() string {
  return "tbl_service"
}

type ServiceDAO interface {
  create(service *ServiceEntity) error

}
