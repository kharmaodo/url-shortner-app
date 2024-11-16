package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strings"
	"url-shortner-app/internal/db"
	"url-shortner-app/internal/url"
)

func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		originalUrl := r.FormValue("url")
		if originalUrl == "" {
			http.Error(w, "Missing original_url", http.StatusBadRequest)
			return
		}
		if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://") {
			originalUrl = "https://" + originalUrl
		}

		_, shorURL := url.Shorten(originalUrl)

		if err := db.StoreURL(lite, shorURL, originalUrl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]string{
			"ShortURL": shorURL,
		}
		t, err := template.ParseFiles("internal/views/shorten.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Proxy(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:]
		if shortUrl == "" {
			http.Error(w, "Missing short_url", http.StatusBadRequest)
			return
		}
		originalUrl, err := db.GetOriginalURL(lite, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)
	}
}
