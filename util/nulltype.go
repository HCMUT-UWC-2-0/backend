package util

import (
	"database/sql"
	"time"
)

func MapInt32ToSqlNullInt32(value int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: value,
		Valid: value != 0,
	}

}

func MapStringToNullString(data string) sql.NullString {

	var result = sql.NullString{
		String: data,
		Valid:  false,
	}
	if data != "" {
		result.Valid = true
	}
	return result
}

func MapStringToNullTime(data string) (sql.NullTime, error) {

	var parseData, err = time.Parse(LayoutDate, data)
	if err != nil {
		return sql.NullTime{Time: time.Now(), Valid: false}, err
	}
	var result = sql.NullTime{
		Time:  parseData,
		Valid: false,
	}
	if data != "" {
		result.Valid = true
	}
	return result, err

}


func MapStringToTime(data string) (time.Time, error) {

	var parseData, err = time.Parse(LayoutDate, data)
	if err != nil {
		return  time.Now(), err
	}
	return parseData, err

}