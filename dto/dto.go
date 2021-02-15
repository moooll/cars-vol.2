package dto

import "time"

type SpeedDateReq struct {
	Speed float64 `json:"speed"`
	Date time.Time `json:"date"`
}

type SpeedDateResp struct {
	Date time.Time `json:"date"`
	Time time.Time	`json:"time"`
	CarNumber string	`json:"car_name"`
	Speed float64 `json:"speed"`
}

type MinMaxSpeedReq struct {
	Date time.Time	`json:"date"`
}

type MinMaxSpeedResp struct {
	MinSpeed SpeedDateResp `json:"min_speed"`
	MaxSpeed SpeedDateResp	`json:"max_speed"`
}