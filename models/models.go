package models

import (
	"gorm.io/gorm"
)

type Event struct {
	Username        string  `json:"username"`
	StartTime       float64 `json:"start_time"`
	EndTime         float64 `json:"end_time"`
	EventCount      float64 `json:"event_count"`
	Hostname        string  `json:"hostname"`
	SourceIP        string  `json:"source_ip"`
	SourcePort      string  `json:"source_port"`
	DestinationPort string  `json:"destination_port"`
	DestinationIP   string  `json:"destination_ip"`
	Owner           float64 `json:"owner"`
}
type Data struct {
	Events []Event `json:"events"`
}
type FailedLoginModel struct {
	gorm.Model
	ID            int  `gorm:"column:id;primaryKey;autoIncrement"`
	Username      string  `gorm:"column:username"`
	EventCount    float64 `gorm:"column:event_count"`
	StartTime     float64 `gorm:"column:start_time"`
	Owner         float64 `gorm:"column:owner"`
	SourceIP      string  `gorm:"column:source_ip"`
	EndTime       float64 `gorm:"column:end_time"`
	DestinationIP string  `gorm:"column:destination_ip"`
}

func (FailedLoginModel) TableName() string {
	return "failed_logins"
}

type AdminLoginModel struct {
	gorm.Model
	ID            int  `gorm:"column:id;primaryKey;autoIncrement"`
	Username      string  `gorm:"column:username"`
	EventCount    float64 `gorm:"column:event_count"`
	StartTime     float64 `gorm:"column:start_time"`
	Owner         float64 `gorm:"column:owner"`
	SourceIP      string  `gorm:"column:source_ip"`
	EndTime       float64 `gorm:"column:end_time"`
	DestinationIP string  `gorm:"column:destination_ip"`
}

func (AdminLoginModel) TableName() string {
	return "admin_logins"
}

type NetworkTrafficModel struct {
	gorm.Model
	ID              int `gorm:"column:id;primaryKey;autoIncrement"`
	Username        string  `gorm:"column:username"`
	EventCount      float64 `gorm:"column:event_count"`
	StartTime       float64 `gorm:"column:start_time"`
	Owner           float64 `gorm:"column:owner"`
	SourceIP        string  `gorm:"column:source_ip"`
	SourcePort      string  `gorm:"column:source_port"`
	EndTime         float64 `gorm:"column:end_time"`
	DestinationIP   string  `gorm:"column:destination_ip"`
	DestinationPort string  `gorm:"column:destination_port"`
	Weight          float64 `gorm:"column:weight"`
}

func (NetworkTrafficModel) TableName() string {
	return "network_traffic"
}

type FirewallDenyModel struct {
	gorm.Model
	ID              int `gorm:"column:id;primaryKey;autoIncrement"`
	Username        string  `gorm:"column:username"`
	EventCount      float64 `gorm:"column:event_count"`
	StartTime       float64 `gorm:"column:start_time"`
	Owner           float64 `gorm:"column:owner"`
	SourceIP        string  `gorm:"column:source_ip"`
	SourcePort      string  `gorm:"column:source_port"`
	EndTime         float64 `gorm:"column:end_time"`
	DestinationIP   string  `gorm:"column:destination_ip"`
	DestinationPort string  `gorm:"column:destination_port"`
}

func (FirewallDenyModel) TableName() string {
	return "firewall_deny"
}

type BlockedAccountsModel struct {
	gorm.Model
	ID              int `gorm:"column:id;primaryKey;autoIncrement"`
	Username        string  `gorm:"column:username"`
	Hostname        string  `gorm:"column:hostname"`
	EventCount      float64 `gorm:"column:event_count"`
	StartTime       float64 `gorm:"column:start_time"`
	Owner           float64 `gorm:"column:owner"`
	SourceIP        string  `gorm:"column:source_ip"`
	SourcePort      string  `gorm:"column:source_port"`
	EndTime         float64 `gorm:"column:end_time"`
	DestinationIP   string  `gorm:"column:destination_ip"`
	DestinationPort string  `gorm:"column:destination_port"`
	Weight          float64 `gorm:"column:weight"`
}
func (BlockedAccountsModel) TableName() string {
	return "blocked_accounts"
}
