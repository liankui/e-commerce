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
	productFieldNames          = builder.RawFieldNames(&Product{})
	productRows                = strings.Join(productFieldNames, ",")
	productRowsExpectAutoSet   = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	productRowsWithPlaceHolder = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheProductIdPrefix = "cache:product:id:"
)

type (
	productModel interface {
		Insert(ctx context.Context, data *Product) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Product, error)
		Update(ctx context.Context, data *Product) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductModel struct {
		sqlc.CachedConn
		table string
	}

	Product struct {
		Id         int64          `db:"id"`          // 商品id
		Cateid     int64          `db:"cateid"`      // 类别Id
		Name       string         `db:"name"`        // 商品名称
		Subtitle   string         `db:"subtitle"`    // 商品副标题
		Images     sql.NullString `db:"images"`      // 图片地址,json格式,扩展用
		Detail     sql.NullString `db:"detail"`      // 商品详情
		Price      float64        `db:"price"`       // 价格,单位-元保留两位小数
		Stock      int64          `db:"stock"`       // 库存数量
		Status     int64          `db:"status"`      // 商品状态.1-在售 2-下架 3-删除
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"` // 更新时间
	}
)

func newProductModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductModel {
	return &defaultProductModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`product`",
	}
}

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, productRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Cateid, data.Name, data.Subtitle, data.Images, data.Detail, data.Price, data.Stock, data.Status)
	}, productIdKey)
	return ret, err
}

func (m *defaultProductModel) FindOne(ctx context.Context, id int64) (*Product, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, id)
	var resp Product
	err := m.QueryRowCtx(ctx, &resp, productIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
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

func (m *defaultProductModel) Update(ctx context.Context, data *Product) error {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Cateid, data.Name, data.Subtitle, data.Images, data.Detail, data.Price, data.Stock, data.Status, data.Id)
	}, productIdKey)
	return err
}

func (m *defaultProductModel) Delete(ctx context.Context, id int64) error {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productIdKey)
	return err
}

func (m *defaultProductModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProductIdPrefix, primary)
}

func (m *defaultProductModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductModel) tableName() string {
	return m.table
}
