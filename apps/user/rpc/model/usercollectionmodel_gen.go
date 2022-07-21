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
	userCollectionFieldNames          = builder.RawFieldNames(&UserCollection{})
	userCollectionRows                = strings.Join(userCollectionFieldNames, ",")
	userCollectionRowsExpectAutoSet   = strings.Join(stringx.Remove(userCollectionFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userCollectionRowsWithPlaceHolder = strings.Join(stringx.Remove(userCollectionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserCollectionIdPrefix           = "cache:userCollection:id:"
	cacheUserCollectionUidProductIdPrefix = "cache:userCollection:uid:productId:"
)

type (
	userCollectionModel interface {
		Insert(ctx context.Context, data *UserCollection) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserCollection, error)
		FindOneByUidProductId(ctx context.Context, uid int64, productId int64) (*UserCollection, error)
		Update(ctx context.Context, data *UserCollection) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserCollectionModel struct {
		sqlc.CachedConn
		table string
	}

	UserCollection struct {
		Id         int64     `db:"id"`          // 收藏Id
		Uid        int64     `db:"uid"`         // 用户id
		ProductId  int64     `db:"product_id"`  // 商品id
		IsDelete   int64     `db:"is_delete"`   // 是否删除
		CreateTime time.Time `db:"create_time"` // 数据创建时间[禁止在代码中赋值]
		UpdateTime time.Time `db:"update_time"` // 数据更新时间[禁止在代码中赋值]
	}
)

func newUserCollectionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserCollectionModel {
	return &defaultUserCollectionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_collection`",
	}
}

func (m *defaultUserCollectionModel) Insert(ctx context.Context, data *UserCollection) (sql.Result, error) {
	userCollectionIdKey := fmt.Sprintf("%s%v", cacheUserCollectionIdPrefix, data.Id)
	userCollectionUidProductIdKey := fmt.Sprintf("%s%v:%v", cacheUserCollectionUidProductIdPrefix, data.Uid, data.ProductId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userCollectionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.ProductId, data.IsDelete)
	}, userCollectionIdKey, userCollectionUidProductIdKey)
	return ret, err
}

func (m *defaultUserCollectionModel) FindOne(ctx context.Context, id int64) (*UserCollection, error) {
	userCollectionIdKey := fmt.Sprintf("%s%v", cacheUserCollectionIdPrefix, id)
	var resp UserCollection
	err := m.QueryRowCtx(ctx, &resp, userCollectionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userCollectionRows, m.table)
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

func (m *defaultUserCollectionModel) FindOneByUidProductId(ctx context.Context, uid int64, productId int64) (*UserCollection, error) {
	userCollectionUidProductIdKey := fmt.Sprintf("%s%v:%v", cacheUserCollectionUidProductIdPrefix, uid, productId)
	var resp UserCollection
	err := m.QueryRowIndexCtx(ctx, &resp, userCollectionUidProductIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `uid` = ? and `product_id` = ? limit 1", userCollectionRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uid, productId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserCollectionModel) Update(ctx context.Context, data *UserCollection) error {
	userCollectionIdKey := fmt.Sprintf("%s%v", cacheUserCollectionIdPrefix, data.Id)
	userCollectionUidProductIdKey := fmt.Sprintf("%s%v:%v", cacheUserCollectionUidProductIdPrefix, data.Uid, data.ProductId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userCollectionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.ProductId, data.IsDelete, data.Id)
	}, userCollectionIdKey, userCollectionUidProductIdKey)
	return err
}

func (m *defaultUserCollectionModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userCollectionIdKey := fmt.Sprintf("%s%v", cacheUserCollectionIdPrefix, id)
	userCollectionUidProductIdKey := fmt.Sprintf("%s%v:%v", cacheUserCollectionUidProductIdPrefix, data.Uid, data.ProductId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userCollectionIdKey, userCollectionUidProductIdKey)
	return err
}

func (m *defaultUserCollectionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserCollectionIdPrefix, primary)
}

func (m *defaultUserCollectionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userCollectionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserCollectionModel) tableName() string {
	return m.table
}
