package database_test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/m7mdkamal/webwatcher/model"

	"github.com/m7mdkamal/webwatcher/database"
)

func TestInitSQLiteDatabase(t *testing.T) {
	_, err := database.InitSQLiteDatabase("")
	ok(t, err)
}

func TestCreateTask(t *testing.T) {
	db, err := NewTestDB()
	defer db.Close()
	ok(t, err)
	taskID, err := db.CreateTask(&model.Task{
		Name: "Reddit#java",
	})
	assert(t, taskID == -1, "Task ID should equal -1")
	taskID, err = db.CreateTask(&model.Task{
		Name:     "Reddit#java",
		Interval: 32,
	})
	ok(t, err)
}

func NewTestDB() (*database.SQLiteDatabase, error) {
	return database.InitSQLiteDatabase(path.Join(os.TempDir(), "testdb.sqlite3"))
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
