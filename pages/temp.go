package pages

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/siddiq24/golang-weekly/models"
	"github.com/siddiq24/golang-weekly/utils"
)

func (db *Db) Fetch() []models.Product {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("failed to connect database postgres")
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), `select * from products`)
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		fmt.Println("error collectrows\n>> ", err)
	}
	return products
}


var CacheFile string = filepath.Join(os.TempDir(), "janjiw.json")

func (db *Db) Caching() {
	wak, er := strconv.Atoi(os.Getenv("TIME"))
	if er != nil{
		wak = 15
	}

	info, err := os.Stat(CacheFile)
	if errors.Is(err, os.ErrNotExist) {
		// Cache belum ada
		db.Products = db.Fetch()
		data, _ := json.MarshalIndent(db.Products, "", "  ")
		os.WriteFile(CacheFile, data, 0644)
		return
	}

	if err != nil {
		utils.Alert(fmt.Sprintf("Error saat cek cache: %v", err))
		return
	}

	age := time.Since(info.ModTime())
	if age >= time.Duration(wak)*time.Second {
		// if age > 15 dtik
		db.Products = db.Fetch()
		data, _ := json.MarshalIndent(db.Products, "", "  ")
		os.WriteFile(CacheFile, data, 0644)
	} else {
		// cache ada dan kurang dari 15 dtik
		body, err := os.ReadFile(CacheFile)
		if err != nil {
			utils.Alert(fmt.Sprintf("Gagal baca cache: %v", err))
			return
		}
		if err := json.Unmarshal(body, &db.Products); err != nil {
			utils.Alert(fmt.Sprintf("Gagal unmarshal cache: %v", err))
		}
	}
}

func (db *Db) ClearCache() {
	err := os.Remove(CacheFile)
	if err != nil {
		utils.Alert("\nGagal menghapus file")
	} else {
		utils.Alert("\nBerhasil menghapus cache")
	}
}
