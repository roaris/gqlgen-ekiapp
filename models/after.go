package models

import "github.com/jmoiron/sqlx"

type After struct {
	LineCd              int    `db:"line_cd"`
	LineName            string `db:"line_name"`
	StationCd           int    `db:"station_cd"`
	StationName         string `db:"station_name"`
	Address             string `db:"address"`
	AfterStationCd      int    `db:"after_station_cd"`
	AfterStationName    string `db:"after_station_name"`
	AfterStationGCd     int    `db:"after_station_g_cd"`
	AfterStationAddress string `db:"after_station_address"`
}

func AftersByStationCD(db *sqlx.DB, stationCD int) ([]*After, error) {
	const sqlstr = `select l.line_cd, l.line_name, s.station_cd, s.station_name, s.address, ` +
		`coalesce(js.station_cd, 0) as after_station_cd, coalesce(js.station_name, '') as after_station_name, ` +
		`coalesce(js.station_g_cd, 0) as after_station_g_cd, coalesce(js.address, '') as after_station_address ` +
		`from stations s left outer join station_lines l on s.line_cd = l.line_cd ` +
		`left outer join station_joins j on s.line_cd = j.line_cd and s.station_cd = j.station_cd2 ` +
		`left outer join stations js on j.station_cd1 = js.station_cd ` +
		`where s.e_status = 0 ` + `and s.station_cd = ?`

	rows, err := db.Queryx(sqlstr, stationCD)
	if err != nil {
		return nil, err
	}
	var res []*After
	for rows.Next() {
		a := After{}
		if err := rows.StructScan(&a); err != nil {
			return nil, err
		}
		res = append(res, &a)
	}
	return res, nil
}
