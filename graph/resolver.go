package graph

import (
	"context"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/roaris/gqlgen-ekiapp/graph/model"
	"github.com/roaris/gqlgen-ekiapp/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var db *sqlx.DB

func init() {
	conn, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/db_dev?parseTime=True")
	if err != nil {
		panic(err)
	}
	db = conn
}

func (r *Resolver) getStationsByName(ctx context.Context, name *string) ([]*model.Station, error) {
	stations, err := models.StationsByStationName(db, *name)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.Station, 0, len(stations))
	for _, v := range stations {
		resp = append(resp, &model.Station{
			StationCd:       v.StationCd,
			StationName:     v.StationName,
			LineName:        &v.LineName,
			Address:         &v.Address,
			BeforeStation:   beforeStation(v.StationCd),
			AfterStation:    afterStation(v.StationCd),
			TransferStation: transferStations(v.StationCd),
		})
	}
	return resp, nil
}

func (r *Resolver) getStationByCD(ctx context.Context, stationCd *int) (*model.Station, error) {
	stations, err := models.StationsByStationCD(db, *stationCd)
	if err != nil {
		return nil, err
	}
	if len(stations) == 0 {
		return nil, errors.New("not found")
	}
	first := stations[0]

	return &model.Station{
		StationCd:       first.StationCd,
		StationName:     first.StationName,
		LineName:        &first.LineName,
		Address:         &first.Address,
		BeforeStation:   beforeStation(first.StationCd),
		AfterStation:    afterStation(first.StationCd),
		TransferStation: transferStations(first.StationCd),
	}, nil
}

func transferStations(stationCd int) []*model.Station {
	records, err := models.TransfersByStationCD(db, stationCd)
	if err != nil {
		return nil
	}

	resp := make([]*model.Station, 0, len(records))
	for _, v := range records {
		if v.TransferStationName == "" {
			continue
		}
		resp = append(resp, &model.Station{
			StationCd:   v.TransferStationCd,
			StationName: v.TransferStationName,
			LineName:    &v.TransferLineName,
			Address:     &v.TransferAddress,
		})
	}
	return resp
}

func beforeStation(stationCd int) *model.Station {
	records, err := models.BeforesByStationCD(db, stationCd)
	if err != nil {
		return nil
	}

	if len(records) == 0 || records[0].BeforeStationName == "" {
		return nil
	}

	return &model.Station{
		StationCd:   records[0].BeforeStationCd,
		StationName: records[0].BeforeStationName,
		LineName:    &records[0].LineName,
		Address:     &records[0].BeforeStationAddress,
	}
}

func afterStation(stationCd int) *model.Station {
	records, err := models.AftersByStationCD(db, stationCd)
	if err != nil {
		return nil
	}

	if len(records) == 0 || records[0].AfterStationName == "" {
		return nil
	}

	return &model.Station{
		StationCd:   records[0].AfterStationCd,
		StationName: records[0].AfterStationName,
		LineName:    &records[0].LineName,
		Address:     &records[0].AfterStationAddress,
	}
}
