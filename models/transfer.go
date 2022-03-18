package models

import "github.com/jmoiron/sqlx"

type Transfer struct {
	StationCd           int    `db:"station_cd"`
	LineCd              int    `db:"line_cd"`
	LineName            string `db:"line_name"`
	StationName         string `db:"station_name"`
	StationGCd          int    `db:"station_g_cd"`
	Address             string `db:"address"`
	TransferLineCd      int    `db:"transfer_line_cd"`
	TransferLineName    string `db:"transfer_line_name"`
	TransferStationCd   int    `db:"transfer_station_cd"`
	TransferStationName string `db:"transfer_station_name"`
	TransferAddress     string `db:"transfer_address"`
}

func TransfersByStationCD(db *sqlx.DB, stationCD int) ([]*Transfer, error) {
	const sqlstr = `select s.station_cd, ls.line_cd, ls.line_name, s.station_name, s.station_g_cd, s.address, ` +
		`coalesce(lt.line_cd, 0) as transfer_line_cd, coalesce(lt.line_name, '') as transfer_line_name, ` +
		`coalesce(t.station_cd, 0) as transfer_station_cd, coalesce(t.station_name, '') as transfer_station_name, ` +
		`coalesce(t.address, '') as transfer_address ` +
		`from stations s left outer join stations t on s.station_g_cd = t.station_g_cd and s.station_cd <> t.station_cd ` +
		`left outer join station_lines ls on s.line_cd = ls.line_cd ` +
		`left outer join station_lines lt on t.line_cd = lt.line_cd ` +
		`where s.station_cd = ?`

	rows, err := db.Queryx(sqlstr, stationCD)
	if err != nil {
		return nil, err
	}
	var res []*Transfer
	for rows.Next() {
		t := Transfer{}
		if err := rows.StructScan(&t); err != nil {
			return nil, err
		}
		res = append(res, &t)
	}
	return res, nil
}
