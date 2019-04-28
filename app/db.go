package app

import (
	"errors"
	"sync"

	_ "github.com/go-sql-driver/mysql" // this is required to drive the the database we connect to
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

var db *xorm.Engine
var once sync.Once

func NewDB(connectionString string) *xorm.Engine {

	once.Do(func() {
		engine, err := xorm.NewEngine("mysql", connectionString)

		if err != nil {
			log.Panic(err)
		}

		engine.SetColumnMapper(core.GonicMapper{})

		if err = engine.Ping(); err != nil {
			log.Panic(err)
		}

		// Use an adapter to the logrus Standard logger
		logger := &logrusAdapter{Logger: log.StandardLogger()}
		engine.SetLogger(logger)
		engine.ShowSQL(true)

		results, err := engine.QueryString("SELECT VERSION() as v")
		log.Infof("Connected to SQL server:%s", results[0]["v"])

		results, err = engine.QueryString("show tables")
		log.Infof("Tables are ", results)

		results, err = engine.QueryString("select * from pet")
		log.Infof("Table contents of table pet: ", results)

		db = engine
	})

	return db

}

//GetDB returns the already initialized DB otherwise error will be returned.
func GetDB() (*xorm.Engine, error) {
	if db == nil {
		return nil, errors.New("DB -- expected DB instance to be in initialized state but found instance as nil")
	}

	return db, nil
}

//RunInTx runs the specified function f within a transaction and rolling back changes during failures
func RunInTx(f func(s *xorm.Session) (interface{}, error)) (interface{}, error) {
	//Assumption - db instance is already initialized via the construction of repo.

	if db == nil {
		log.Error("Expected db ptr is not initialized - should have been done via repository constructors")
	}

	session := db.NewSession()
	defer session.Close()

	returnVal, err := f(session)

	if err != nil {
		session.Rollback()
		return nil, err
	}

	err = session.Commit()

	if err != nil {
		return nil, err
	}

	return returnVal, nil

}

type logrusAdapter struct {
	*log.Logger
	showSQL bool
}

func (logger *logrusAdapter) Level() core.LogLevel {
	// Convert between logrus and xorm log levels
	logint := -1 * (int(logger.Logger.Level) - 5)

	return core.LogLevel(logint)
}

func (logger *logrusAdapter) SetLevel(logLevel core.LogLevel) {
	logint := -1 * (logLevel - 5)
	logger.Logger.SetLevel(log.Level(logint))
}

func (logger *logrusAdapter) ShowSQL(show ...bool) {
	if len(show) > 0 {
		logger.showSQL = show[0]
	} else {
		logger.showSQL = true
	}
}

func (logger *logrusAdapter) IsShowSQL() bool {
	return logger.showSQL
}
