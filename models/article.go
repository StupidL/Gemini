package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

var dbName = "gemini"
var driverName = "mysql"
var dataSource = "stupidl:stupidl@/gemini?charset=utf8"
var maxIdle = 30
var maxConn = 30

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(
		dbName,
		driverName,
		dataSource,
		maxIdle,
		maxConn)
	orm.RegisterModel(new(Article))
}

type Article struct {
	Id      int64  `form:"-"`
	Title   string `form:"title"`
	Url     string `form:"url"`
	Content string `form:"content"`
}

type ArticleManager interface {
	Insert(a *Article) error
	DeleteById(id int64) error
	DeleteByTitle(t string) error
	Update(a *Article) error
	QueryById(id int64) (*Article, error)
	QueryByTitle(t string) (*Article, error)
	QueryAll() ([]*Article, error)
}

type ArticleMySQLManager struct {
	db *sql.DB
}

func (m *ArticleMySQLManager) Insert(a *Article) error {

}

func (m *ArticleMySQLManager) DeleteById(id int64) error {

}
