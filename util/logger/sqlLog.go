// Package logger provides request logging via middleware. See _examples/http_request/request-logger
package logger

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
	"unicode"
)

type SqlLoggerMiddleware struct {
	Zap *zap.Logger
}

// Println format & print log
func (l *SqlLoggerMiddleware) Println(values []interface{}) {
	l.Zap.Info( "gorm" , createLog(values).toZapFields()...)
}

// Print passes arguments to Println
func (l *SqlLoggerMiddleware) Print(values ...interface{}) {
	l.Println(values)
}

type sqlLog struct {
	occurredAt time.Time
	source     string
	duration   time.Duration
	sql        string
	values     []string
	other      []string
}

func createLog(values []interface{}) *sqlLog {
	ret := &sqlLog{}
	ret.occurredAt = gorm.NowFunc()

	if len(values) > 1 {
		var level = values[0]
		ret.source = fmt.Sprint(values[1])

		if level == "sql" {
			ret.duration = values[2].(time.Duration)
			ret.sql = values[3].(string)
			ret.values = getFormattedValues(values)
		} else {
			ret.other = append(ret.other, fmt.Sprint(values[2:]))
		}
	}
	return ret
}

//已zap的格式 性能更加
func (l *sqlLog) toZapFields() []zapcore.Field {
	return []zapcore.Field{
		zap.String("occurredAt", l.occurredAt.Format("2006-01-02 15:04:05")),
		//zap.String("source", l.source),
		zap.Int64("duration", l.duration.Milliseconds()),
		zap.String("sql", l.sql),
		zap.Strings("values", l.values),
		zap.Strings("other", l.other),
	}
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// do sql params
func getFormattedValues(values []interface{}) []string {
	rawValues := values[4].([]interface{})
	formattedValues := make([]string, 0, len(rawValues))
	for _, value := range rawValues {
		switch v := value.(type) {
		case time.Time:
			formattedValues = append(formattedValues, fmt.Sprint(v))
		case []byte:
			if str := string(v); isPrintable(str) {
				formattedValues = append(formattedValues, fmt.Sprint(str))
			} else {
				formattedValues = append(formattedValues, "<binary>")
			}
		default:
			str := "NULL"
			if v != nil {
				str = fmt.Sprint(v)
			}
			formattedValues = append(formattedValues, str)
		}
	}
	return formattedValues
}
