package hello

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type Document struct {
	ID   int    `json:"id"`
	Hash string `json:"hash"`
	Text string `json:"text`
}

var (
	testText     = "test_data_text"
	r            = sha256.Sum256([]byte(testText))
	testDocument = Document{ID: 1, Hash: hex.EncodeToString(r[:]), Text: testText}
)

func getTestTransaction() (*sql.Tx, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("HOST_NAME"), os.Getenv("PORT"), os.Getenv("TEST_DB_NAME"))
	var err error
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS document")
	if err != nil {
		return nil, err
	}

	createDocumentTableCmd := `CREATE TABLE IF NOT EXISTS document(
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		hash varchar(255) NOT NULL,
		text text NOT NULL,
		UNIQUE KEY (hash))`
	if _, err := tx.Exec(createDocumentTableCmd); err != nil {
		return nil, err
	}

	return tx, nil
}

func TestCreateDocument(t *testing.T) {
	tx, err := getTestTransaction()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	ins, err := tx.Prepare("INSERT INTO document(hash, text) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(testDocument.Hash, testDocument.Text)
	document := Document{}
	getDocumentCmd := "SELECT `id`, `hash`, `text` FROM document WHERE `id` = ?"
	_ = tx.QueryRowContext(ctx, getDocumentCmd, testDocument.ID).Scan(
		&document.ID,
		&document.Hash,
		&document.Text,
	)
	if document != testDocument {
		t.Error("\n実際： ", document, "\n理想： ", testDocument)
	}
	if err := tx.Rollback(); err != nil {
		t.Error(err)
	}
}
