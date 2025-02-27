package types

import "time"

type DeviceEvent struct {
	Id       int64          `json:"id"`
	DeviceId string         `json:"device_id" xorm:"index"`
	Name     string         `json:"name"`
	Label    string         `json:"label"`
	Output   map[string]any `json:"output" xorm:"json"`
	Created  time.Time      `json:"created" xorm:"created"`
}
