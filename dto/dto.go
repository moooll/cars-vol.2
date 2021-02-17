package dto

import "time"

type SpeedDate struct {
	Speed float64 		`json:"speed"`
	Date time.Time		`json:"date"`
}

type SpeedDateReq struct {
	Speed float64 		`json:"speed"`
	Date string		`json:"date"`
}

type CarInformation struct {
	DateTime time.Time 	`json:"date_time"`
	CarNumber string	`json:"car_name"`
	Speed float64 		`json:"speed"`
}

type CreateCars struct {
	DateTime string		`json:"date_time"`
	CarNumber string	`json:"car_name"`
	Speed float64 		`json:"speed"`
}

type MinMaxSpeedReq struct {
	Date time.Time		`json:"date"`
}

type MinMaxSpeedResp struct {
	MinSpeed CarInformation 	`json:"min_speed"`
	MaxSpeed CarInformation		`json:"max_speed"`
}