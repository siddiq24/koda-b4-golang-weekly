package pages

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/siddiq24/golang-weekly/models"
	"github.com/siddiq24/golang-weekly/utils"
)

func (db *Db) Fetch() []models.Product {
	resp, err := http.Get("https://raw.githubusercontent.com/siddiq24/all-products-janji-jiwa/refs/heads/main/all-products.json")
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}
	defer resp.Body.Close()

	var data []models.Product
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("err:", err)
		return nil
	}
	return data
}


var CacheFile string = filepath.Join(os.TempDir(), "janjiw.json")

func (db *Db) Caching() {
	wak, er := strconv.Atoi(os.Getenv("TIME"))
	if er != nil{
		utils.Alert("Gagal ambil data dari ENV")
		return
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
