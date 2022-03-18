package models

import (
	"github.com/jmoiron/sqlx"
)

type Station struct {
	LineCd      int    `db:"line_cd"`
	LineName    string `db:"line_name"`
	StationCd   int    `db:"station_cd"`
	StationGCd  int    `db:"station_g_cd"`
	StationName string `db:"station_name"`
	Address     string `db:"address"`
}

func StationsByStationCD(db *sqlx.DB, stationCD int) ([]*Station, error) {
	const sqlstr = `select l.line_cd, l.line_name, s.station_cd, s.station_g_cd, s.station_name, s.address ` +
		`from stations s inner join station_lines l on s.line_cd = l.line_cd ` +
		`where s.station_cd = ? and s.e_status = 0`

	rows, err := db.Queryx(sqlstr, stationCD)
	if err != nil {
		return nil, err
	}
	var res []*Station
	for rows.Next() {
		s := Station{}
		if err := rows.StructScan(&s); err != nil {
			return nil, err
		}
		res = append(res, &s)
	}
	return res, nil
}

func StationsByStationName(db *sqlx.DB, stationName string) ([]*Station, error) {
	const sqlstr = `select l.line_cd, l.line_name, s.station_cd, s.station_g_cd, s.station_name, s.address ` +
		`from stations s inner join station_lines l on s.line_cd = l.line_cd ` +
		`where s.station_name = ? and s.e_status = 0`

	rows, err := db.Queryx(sqlstr, stationName)
	if err != nil {
		return nil, err
	}
	var res []*Station
	for rows.Next() {
		s := Station{}
		if err := rows.StructScan(&s); err != nil {
			return nil, err
		}
		res = append(res, &s)
	}
	return res, nil
}
