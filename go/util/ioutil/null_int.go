package ioutil

import "database/sql"

func IntToNullInt(s int) sql.NullInt64 {
	if s != 0 {
		return sql.NullInt64{
			Int64: int64(s),
			Valid: true,
		}
	}
	return sql.NullInt64{}
}

func NullIntToInt(s sql.NullInt64) int {
	if s.Valid {
		return int(s.Int64)
	}
	return 0
}
