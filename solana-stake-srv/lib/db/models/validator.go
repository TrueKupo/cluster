package models

import (
	"time"
)

type Validator struct {
	Id                           int64
	Network                      string    `json:"network" db:"network"`
	Account                      string    `json:"account" db:"account"`
	Name                         string    `json:"name" db:"name"`
	WwwUrl                       string    `json:"www_url" db:"www_url"`
	Details                      string    `json:"details" db:"details"`
	AvatarUrl                    string    `json:"avatar_url" db:"avatar_url"`
	CreatedAt                    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at" db:"updated_at"`
	TotalScore                   int64     `json:"total_score" db:"total_score"`
	RootDistanceScore            int64     `json:"root_distance_score" db:"root_distance_score"`
	VoteDistanceScore            int64     `json:"vote_distance_score" db:"vote_distance_score"`
	SkippedSlotScore             int64     `json:"skipped_slot_score" db:"skipped_slot_score"`
	SoftwareVersion              string    `json:"software_version" db:"software_version"`
	SoftwareVersionScore         int64     `json:"software_version_score" db:"software_version_score"`
	StakeConcentrationScore      int64     `json:"stake_concentration_score" db:"stake_concentration_score"`
	DataCenterConcentrationScore int64     `json:"data_center_concentration_score" db:"data_center_concentration_score"`
	PublishedInformationScore    int64     `json:"published_information_score" db:"published_information_score"`
	SecurityReportScore          int64     `json:"security_report_score" db:"security_report_score"`
	ActiveStake                  int64     `json:"active_stake" db:"active_stake"`
	Commission                   int64     `json:"commission" db:"commission"`
	Delinquent                   bool      `json:"delinquent" db:"delinquent"`
	DataCenterKey                string    `json:"data_center_key" db:"data_center_key"`
	DataCenterHost               string    `json:"data_center_host" db:"data_center_host"`
	AutonomousSystemNumber       int64     `json:"autonomous_system_number" db:"autonomous_system_number"`
	VoteAccount                  string    `json:"vote_account" db:"vote_account"`
	EpochCredits                 int64     `json:"epoch_credits" db:"epoch_credits"`
	SkippedSlots                 int64     `json:"skipped_slots" db:"skipped_slots"`
	SkippedSlotPercent           string    `json:"skipped_slot_percent" db:"skipped_slot_percent"`
	PingTime                     string    `json:"ping_time" db:"ping_time"`
	Url                          string    `json:"url" db:"url"`
}
