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
	productOperationFieldNames          = builder.RawFieldNames(&ProductOperation{})
	productOperationRows                = strings.Join(productOperationFieldNames, ",")
	productOperationRowsExpectAutoSet   = strings.Join(stringx.Remove(productOperationFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	productOperationRowsWithPlaceHolder = strings.Join(stringx.Remove(productOperationFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheProductOperationIdPrefix = "cache:productOperation:id:"
)

type (
	productOperationModel interface {
		Insert(ctx context.Context, data *ProductOperation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductOperation, error)
		Update(ctx context.Context, data *ProductOperation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductOperationModel struct {
		sqlc.CachedConn
		table string
	}

	ProductOperation struct {
		Id         int64     `db:"id"`          // 商品运营id
		ProductId  int64     `db:"product_id"`  // 商品id
		Status     int64     `db:"status"`      // 运营商品状态 0-下线 1-上线
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newProductOperationModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductOperationModel {
	return &defaultProductOperationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`product_operation`",
	}
}

func (m *defaultProductOperationModel) Insert(ctx context.Context, data *ProductOperation) (sql.Result, error) {
	productOperationIdKey := fmt.Sprintf("%s%v", cacheProductOperationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, productOperationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Status)
	}, productOperationIdKey)
	return ret, err
}

func (m *defaultProductOperationModel) FindOne(ctx context.Context, id int64) (*ProductOperation, error) {
	productOperationIdKey := fmt.Sprintf("%s%v", cacheProductOperationIdPrefix, id)
	var resp ProductOperation
	err := m.QueryRowCtx(ctx, &resp, productOperationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productOperationRows, m.table)
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

func (m *defaultProductOperationModel) Update(ctx context.Context, data *ProductOperation) error {
	productOperationIdKey := fmt.Sprintf("%s%v", cacheProductOperationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productOperationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Status, data.Id)
	}, productOperationIdKey)
	return err
}

func (m *defaultProductOperationModel) Delete(ctx context.Context, id int64) error {
	productOperationIdKey := fmt.Sprintf("%s%v", cacheProductOperationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productOperationIdKey)
	return err
}

func (m *defaultProductOperationModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProductOperationIdPrefix, primary)
}

func (m *defaultProductOperationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productOperationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductOperationModel) tableName() string {
	return m.table
}
