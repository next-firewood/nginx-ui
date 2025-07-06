package database

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type GormModel struct {
	Id        uint64     `gorm:"primarykey" json:"id"`
	Uuid      string     `gorm:"uniqueIndex;size:50;comment:uuid" json:"uuid"`
	CreatedAt SQLiteTime `gorm:"type:datetime(0);autoCreateTime" json:"createdAt"`
	UpdatedAt SQLiteTime `gorm:"type:datetime(0);autoUpdateTime" json:"updatedAt"`
}

func ScreenTime(db *gorm.DB, timeAt []string, timeField string) *gorm.DB {
	for i, t := range timeAt {
		switch i {
		case 0:
			db = db.Where(timeField+" >= ?", t)
		case 1:
			db = db.Where(timeField+" <= ?", t)
		}
	}

	return db
}

const (
	dataLen = 2
)

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*dataLen)
	j := false
	num := len(s)

	for i := 0; i < num; i++ {
		d := s[i]

		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}

		if d != '_' {
			j = true
		}

		data = append(data, d)
	}

	return strings.ToLower(string(data[:]))
}

func OrderByGBK(s string) string {
	if s == "" {
		return ""
	}

	split := strings.Split(s, " ")

	split[0] = fmt.Sprintf("convert(%s USING gbk)", split[0])

	return strings.Join(split, " ")
}

type OrderParam struct {
	Field  string
	IsGBK  bool
	Prefix string
}

type Order struct {
	Params []OrderParam
	Def    string
}

const (
	_sortFieldsLen = 2
	_ascSort       = "asc"
	_descSort      = "desc"
)

func (o Order) OrderBySingleField(field string) string {
	if field == "" {
		return o.Def
	}

	split := strings.Split(field, " ")
	// 检查长度<=2且第二个元素是否为asc或者desc
	if len(split) == _sortFieldsLen && !(strings.ToLower(split[1]) == _ascSort ||
		strings.ToLower(split[1]) == _descSort) || len(split) > _sortFieldsLen {
		return o.Def
	}
	// 轮询可匹配的排序字段
	for _, p := range o.Params {
		if p.Field == split[0] {
			if p.IsGBK {
				return OrderByGBK(SnakeString(p.Prefix + field))
			} else {
				return SnakeString(p.Prefix + field)
			}
		}
	}

	return o.Def
}

type SQLiteTime time.Time

func (t *SQLiteTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		tt, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = SQLiteTime(tt)
	case time.Time:
		*t = SQLiteTime(v)
	default:
		return fmt.Errorf("cannot convert %T to time.Time", v)
	}
	return nil
}

func (t SQLiteTime) Value() (driver.Value, error) {
	return time.Time(t).Format("2006-01-02 15:04:05"), nil
}

func (t SQLiteTime) Time() time.Time {
	return time.Time(t)
}
