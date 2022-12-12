package common_db

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

//时序数据库

type InfluxClient struct {
	host     string //服务的地址
	port     string //服务的端口
	passport string //账号
	password string //密码
	dbName   string //数据库名称
	table    string //要读取的表
	c        client.Client
}

func (ic *InfluxClient) Build(host, port, passport, password string) error {
	if host == "" {
		return fmt.Errorf("host is empty")
	}
	if port == "" {
		return fmt.Errorf("port is empty")
	}
	ic.host = host
	ic.port = port
	ic.passport = passport
	ic.password = password
	return nil
}

func (ic *InfluxClient) newHttpClient() error {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://" + ic.host + ":" + ic.port,
		Username: ic.passport,
		Password: ic.password,
	})
	if err != nil {
		g.Log().Warning(context.Background(), "链接influxdb 错误", err)
	}
	ic.c = c
	return err
}

func (ic *InfluxClient) SetTable(dbName, table string) {
	ic.dbName = dbName
	ic.table = table
}

// 存储数据
// whereRecord 查询条件
// data 要存储的数据
// recordTime 数据的记录时间
func (ic *InfluxClient) SaveRecord(whereRecord map[string]string, data map[string]interface{}, recordTime int64) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{})
	if err == nil {
		intime := time.Unix(recordTime, 0)
		bp.SetDatabase(ic.dbName)
		pt, err := client.NewPoint(
			ic.table,
			whereRecord,
			data,
			intime,
		)
		if err != nil {
			g.Log().Warningf(context.Background(), "link influxdb,unexpected error actual %v \r", err)
		}
		bp.AddPoint(pt)
		err = ic.newHttpClient()
		if err == nil {
			defer ic.c.Close()
			err = ic.c.Write(bp)
			if err != nil {
				g.Log().Warningf(context.Background(), "save influxdb,unexpected error. actual %v\r", err)
			}
			return nil
		}
	}
	return err
}

//获取数据
// beginTime 数据开始时间，单位秒
// endTime 数据结束时间，单位秒
// timeSort desc 时间降序，asc 时间升序
// limitFile 限制读取文件数，=0 不限制
/*
	result := list[0].Series //结果列表
	out := make([]DeviceInfo, len(result))
	for _, v := range result { // like array_combine
		for _, fv := range v.Values {
			for i, vi := range v.Columns {
				if vi == "stringvalue" {
					out = append(out, DeviceInfo{DeviceKey: gconv.String(fv[i])})
				}
			}
		}
	}
*/
func (ic *InfluxClient) GetRecords(whereRecord map[string]string, beginTime, endTime int64, timeSort string, limitFile uint) ([]client.Result, error) {
	if timeSort == "" {
		timeSort = "desc"
	}
	err := ic.newHttpClient()
	if err == nil {
		defer ic.c.Close()
		where := ""
		gnum := 0
		for i, v := range whereRecord {
			baseWhere := fmt.Sprintf(" %s = '%s' ", i, v)
			if gnum != 0 {
				baseWhere = "and " + baseWhere
			}
			where = where + baseWhere
			gnum++
		}
		beginTime = beginTime * 1e9
		endTime = endTime * 1e9
		where = where + fmt.Sprintf(" and time >= %s and time <= %s", gconv.String(beginTime), gconv.String(endTime))
		sql := fmt.Sprintf("select * from %s where %s order by time %s ", ic.table, where, timeSort)
		if limitFile > 0 {
			sql = sql + fmt.Sprintf(" limit  %d", limitFile)
		}
		query := client.NewQuery(sql, ic.dbName, "")
		resp, err := ic.c.Query(query)
		if err == nil {
			return resp.Results, nil
		}
		return nil, err
	}
	return nil, err
}

func (ic *InfluxClient) GetRecordsPage(whereRecord map[string]string, beginTime, endTime int64, timeSort string, page, eachPage int) ([]client.Result, error) {
	if timeSort == "" {
		timeSort = "desc"
	}
	err := ic.newHttpClient()
	if err == nil {
		defer ic.c.Close()
		where := ""
		gnum := 0
		for i, v := range whereRecord {
			baseWhere := fmt.Sprintf(" %s = '%s' ", i, v)
			if gnum != 0 {
				baseWhere = "and " + baseWhere
			}
			where = where + baseWhere
			gnum++
		}
		beginTime = beginTime * 1e9
		endTime = endTime * 1e9
		where = where + fmt.Sprintf(" and time >= %s and time <= %s", gconv.String(beginTime), gconv.String(endTime))
		sql := fmt.Sprintf("select * from %s where %s order by time %s ", ic.table, where, timeSort)
		if page > 0 && eachPage > 0 {
			offset := (page - 1) * eachPage
			sql += " limit " + gconv.String(eachPage) + " offset " + gconv.String(offset)
		}
		query := client.NewQuery(sql, ic.dbName, "")
		resp, err := ic.c.Query(query)
		if err == nil {
			return resp.Results, nil
		}
		return nil, err
	}
	return nil, err
}

func (ic *InfluxClient) GetCount(whereRecord map[string]string, beginTime, endTime int64) ([]client.Result, error) {
	err := ic.newHttpClient()
	if err == nil {
		defer ic.c.Close()
		where := ""
		gnum := 0
		for i, v := range whereRecord {
			baseWhere := fmt.Sprintf(" %s = '%s' ", i, v)
			if gnum != 0 {
				baseWhere = "and " + baseWhere
			}
			where = where + baseWhere
			gnum++
		}
		beginTime = beginTime * 1e9
		endTime = endTime * 1e9
		where = where + fmt.Sprintf(" and time >= %s and time <= %s", gconv.String(beginTime), gconv.String(endTime))
		sql := fmt.Sprintf("select count(*) from %s where %s  ", ic.table, where)
		query := client.NewQuery(sql, ic.dbName, "")
		resp, err := ic.c.Query(query)
		if err == nil {
			return resp.Results, nil
		}
		return nil, err
	}
	return nil, err
}

// 删除数据
func (ic *InfluxClient) Del(whereRecord map[string]string, beginTime, endTime int64) error {
	err := ic.newHttpClient()
	if err == nil {
		defer ic.c.Close()
		beginTime = beginTime * 1e9
		endTime = endTime * 1e9
		where := ""
		g := 0
		for i, v := range whereRecord {
			baseWhere := fmt.Sprintf(" %s = '%s' ", i, v)
			if g != 0 {
				baseWhere = "and " + baseWhere
			}
			where = where + baseWhere
			g++
		}
		where = where + fmt.Sprintf(" and time >= %s and time <= %s", gconv.String(beginTime), gconv.String(endTime))
		sql := fmt.Sprintf("delete from %s where %s", ic.table, where)
		query := client.NewQuery(sql, ic.dbName, "")
		resp, err := ic.c.Query(query)
		if err == nil {
			return resp.Error()
		}
		return err
	}
	return err
}

// 通过时间获取数据  2022-04-18
func (ic *InfluxClient) GetRecordsByTime(beginTime, endTime int64, timeSort string, limitFile uint) ([]client.Result, error) {
	if timeSort == "" {
		timeSort = "desc"
	}
	err := ic.newHttpClient()
	if err == nil {
		defer ic.c.Close()
		beginTime = beginTime * 1e9
		endTime = endTime * 1e9
		where := fmt.Sprintf(" time >= %s and time <= %s", gconv.String(beginTime), gconv.String(endTime))
		sql := fmt.Sprintf("select * from %s where %s order by time %s ", ic.table, where, timeSort)
		if limitFile > 0 {
			sql = sql + fmt.Sprintf(" limit  %d", limitFile)
		}
		query := client.NewQuery(sql, ic.dbName, "")
		resp, err := ic.c.Query(query)
		if err == nil {
			return resp.Results, nil
		}
		return nil, err
	}
	return nil, err
}
