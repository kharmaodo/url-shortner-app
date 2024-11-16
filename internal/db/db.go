package db

import "database/sql"

// CreateTable createTable ensures the URLs table exists
func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS urlS (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_url TEXT NULL,
		original_url TEXT NULL,
	);`
	_, err := db.Exec(query)
	return err
}

// StoreURL storeURL inserts a new short URL and the original URL into the database
func StoreURL(db *sql.DB, shortUrl string, originalUrl string) error {
	query := `INSERT INTO urlS (short_url, original_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortUrl, originalUrl)
	return err
}

// GetOriginalURL getOriginalURL fetches the original URL by the short URL
func GetOriginalURL(db *sql.DB, shortUrl string) (string, error) {
	var originalUrl string
	query := `SELECT original_url FROM urlS WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortUrl).Scan(&originalUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}
