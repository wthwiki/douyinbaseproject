// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q        = new(Query)
	Comment  *comment
	Message  *message
	User     *user
	UserInfo *userInfo
	Video    *video
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Comment = &Q.Comment
	Message = &Q.Message
	User = &Q.User
	UserInfo = &Q.UserInfo
	Video = &Q.Video
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:       db,
		Comment:  newComment(db, opts...),
		Message:  newMessage(db, opts...),
		User:     newUser(db, opts...),
		UserInfo: newUserInfo(db, opts...),
		Video:    newVideo(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Comment  comment
	Message  message
	User     user
	UserInfo userInfo
	Video    video
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:       db,
		Comment:  q.Comment.clone(db),
		Message:  q.Message.clone(db),
		User:     q.User.clone(db),
		UserInfo: q.UserInfo.clone(db),
		Video:    q.Video.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:       db,
		Comment:  q.Comment.replaceDB(db),
		Message:  q.Message.replaceDB(db),
		User:     q.User.replaceDB(db),
		UserInfo: q.UserInfo.replaceDB(db),
		Video:    q.Video.replaceDB(db),
	}
}

type queryCtx struct {
	Comment  ICommentDo
	Message  IMessageDo
	User     IUserDo
	UserInfo IUserInfoDo
	Video    IVideoDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Comment:  q.Comment.WithContext(ctx),
		Message:  q.Message.WithContext(ctx),
		User:     q.User.WithContext(ctx),
		UserInfo: q.UserInfo.WithContext(ctx),
		Video:    q.Video.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}