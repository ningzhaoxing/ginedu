package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnError(errors.New("xxxx"))

	_, _ = db.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", 2, 3)

	//if err != nil {
	//	t.Fatalf("错误%v", err)
	//	return
	//}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
