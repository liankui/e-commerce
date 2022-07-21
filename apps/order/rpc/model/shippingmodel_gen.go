// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	shippingFieldNames          = builder.RawFieldNames(&Shipping{})
	shippingRows                = strings.Join(shippingFieldNames, ",")
	shippingRowsExpectAutoSet   = strings.Join(stringx.Remove(shippingFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	shippingRowsWithPlaceHolder = strings.Join(stringx.Remove(shippingFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheShippingIdPrefix = "cache:shipping:id:"
)

type (
	shippingModel interface {
		Insert(ctx context.Context, data *Shipping) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Shipping, error)
		Update(ctx context.Context, data *Shipping) error
		Delete(ctx context.Context, id int64) error
	}

	defaultShippingModel struct {
		sqlc.CachedConn
		table string
	}

	Shipping struct {
		Id               int64     `db:"id"`                // 收货信息表id
		Orderid          string    `db:"orderid"`           // 订单id
		Userid           int64     `db:"userid"`            // 用户id
		ReceiverName     string    `db:"receiver_name"`     // 收货姓名
		ReceiverPhone    string    `db:"receiver_phone"`    // 收货固定电话
		ReceiverMobile   string    `db:"receiver_mobile"`   // 收货移动电话
		ReceiverProvince string    `db:"receiver_province"` // 省份
		ReceiverCity     string    `db:"receiver_city"`     // 城市
		ReceiverDistrict string    `db:"receiver_district"` // 区/县
		ReceiverAddress  string    `db:"receiver_address"`  // 详细地址
		CreateTime       time.Time `db:"create_time"`       // 创建时间
		UpdateTime       time.Time `db:"update_time"`       // 更新时间
	}
)

func newShippingModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultShippingModel {
	return &defaultShippingModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`shipping`",
	}
}

func (m *defaultShippingModel) Insert(ctx context.Context, data *Shipping) (sql.Result, error) {
	shippingIdKey := fmt.Sprintf("%s%v", cacheShippingIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, shippingRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.ReceiverName, data.ReceiverPhone, data.ReceiverMobile, data.ReceiverProvince, data.ReceiverCity, data.ReceiverDistrict, data.ReceiverAddress)
	}, shippingIdKey)
	return ret, err
}

func (m *defaultShippingModel) FindOne(ctx context.Context, id int64) (*Shipping, error) {
	shippingIdKey := fmt.Sprintf("%s%v", cacheShippingIdPrefix, id)
	var resp Shipping
	err := m.QueryRowCtx(ctx, &resp, shippingIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shippingRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShippingModel) Update(ctx context.Context, data *Shipping) error {
	shippingIdKey := fmt.Sprintf("%s%v", cacheShippingIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, shippingRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.ReceiverName, data.ReceiverPhone, data.ReceiverMobile, data.ReceiverProvince, data.ReceiverCity, data.ReceiverDistrict, data.ReceiverAddress, data.Id)
	}, shippingIdKey)
	return err
}

func (m *defaultShippingModel) Delete(ctx context.Context, id int64) error {
	shippingIdKey := fmt.Sprintf("%s%v", cacheShippingIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, shippingIdKey)
	return err
}

func (m *defaultShippingModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheShippingIdPrefix, primary)
}

func (m *defaultShippingModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shippingRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShippingModel) tableName() string {
	return m.table
}
