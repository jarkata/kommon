package kate

import (
	"time"

	"github.com/jarkata/kommon/logger"
)

const (
	yyyyMMddHHmmss   = "20060102150405"
	IsoDateTimeStyle = "2006-01-02 15:04:05"
	BasicDateStyle   = "20060102"
	IsoDateStyle     = "2006-01-02'T'15:04:05"
)

/*
*
可解析格式为:yyyyMMdd的日期
*
*/
func ParseBasicDate(date string) *time.Time {
	t, err := time.Parse(BasicDateStyle, date)
	if err != nil {
		logger.Error("ParseBasicDate Error:", err)
		return nil
	}
	return &t
}

/*
*
可解析日期为yyyy-MM-dd'T'HH:mm:ss
*/
func ParseIsoDateTime2(datetime string) *time.Time {
	t, err := time.Parse(IsoDateStyle, datetime)
	if err != nil {
		logger.Error("ParseIsoDateTime2 Error:", err)
		return nil
	}
	return &t
}

/*
*
可解析日期为yyyy-MM-dd HH:mm:ss
*/
func ParseIsoDateTime(datetime string) *time.Time {
	t, err := time.Parse(IsoDateTimeStyle, datetime)
	if err != nil {
		logger.Error("ParseIsoDateTime Error:", err)
		return nil
	}
	return &t
}

/*
*
return date style : yyyyMMdd
*
*/
func FormatBasicDate(date time.Time) string {
	return date.Format(BasicDateStyle)
}

/** 格式化日期 **/
func FormatDate(date time.Time, layout string) string {
	return date.Format(layout)
}

/*
*
return date style: yyyy-MM-dd HH:mm:ss
*
*/
func FormatIsoDateTime(datetime time.Time) string {
	return datetime.Format(IsoDateTimeStyle)
}
