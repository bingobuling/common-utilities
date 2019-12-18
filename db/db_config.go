//author xinbing
//time 2018/9/4 17:55
package db

import "github.com/pkg/errors"

type DBConfig struct {
	DBAddr            string
	AutoCreateTables  []interface{} //自动创建的表，不设置则不创建表
	MaxIdleConns      int           //数据库连接池设置——最大空闲数，不设置则为10
	MaxOpenConns      int           //数据库连接池设置——最大打开的连接数，不设置则为100
	LogMode           bool          //是否打印gorm的日志
	DetectionInterval int           //心跳检测间隔，单位为s，默认30s,小于0则不检测
}

func (p *DBConfig) check() error {
	if p.DBAddr == "" {
		return errors.New("empty sql addr")
	}
	if p.MaxIdleConns <= 0 {
		p.MaxIdleConns = 10
	}
	if p.MaxOpenConns <= 0 {
		p.MaxOpenConns = 100
	}
	if p.DetectionInterval == 0 {
		p.DetectionInterval = 30
	}
	return nil
}
