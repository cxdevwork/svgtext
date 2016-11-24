package hello

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func init() {
	http.HandleFunc("/svg/month", mthHandler)
	http.HandleFunc("/svg/mth", mthHandler)
	http.HandleFunc("/svg/date", dateHandler)
}

func mthHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")
	size := r.URL.Query().Get("size")

	if size == "" {
		size = "30"
	}

	y, err := strconv.ParseFloat(size, 64)
	if err != nil {
		y = 20
	} else {
		y = y / 1.5
	}
	width := r.URL.Query().Get("width")
	if width == "" {
		width = "600"
	}

	x, err := strconv.Atoi(width)
	if err != nil {
		x = 150
	} else {
		x = x / 2
	}

	height := r.URL.Query().Get("height")
	if height == "" {
		height = "20"
	}

	dx := r.URL.Query().Get("dx")

	if dx == "" {
		dx = "0"
	}

	dy := r.URL.Query().Get("dy")

	if dy == "" {
		dy = fmt.Sprintf("%f", y)
	}

	fill := r.URL.Query().Get("fill")

	if fill == "" {
		fill = "000000"
	}

	format := r.URL.Query().Get("format")

	if format == "" {
		format = "January 2006"
	}

	if text == "" {
		text = time.Now().Format(format)
	} else {
		text = text + " " + time.Now().Format(format)
	}
	w.Header().Set("Content-Type", "image/svg+xml")

	output := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"  viewBox="0 0 %s %s" >
    <text dominant-baseline="top" text-anchor="middle" x='%s' y='%s' style="font-family: Arial, Helvetica, Verdana" font-weight="bold" fill="#%s" dx='%s' dy='%s' font-size="%s">%s</text>
</svg>`, width, height, fmt.Sprintf("%d", x), fmt.Sprintf("%d", y), fill, dx, dy, size, text)

	//  w.Header().Set("Content-Length", fmt.Sprintf("%f",len(output)))
	fmt.Fprint(w, output)
}

func dateHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")
	size := r.URL.Query().Get("size")

	if size == "" {
		size = "20"
	}

	y, err := strconv.ParseFloat(size, 64)
	if err != nil {
		y = 20
	} else {
		y = y / 1.15
	}
	width := r.URL.Query().Get("width")
	if width == "" {
		width = "100"
	}

	x, err := strconv.Atoi(width)
	if err != nil {
		x = 100
	} else {

	}

	height := r.URL.Query().Get("height")
	if height == "" {
		height = "20"
	}

	dx := r.URL.Query().Get("dx")

	if dx == "" {
		dx = "0"
	}

	dy := r.URL.Query().Get("dy")

	if dy == "" {
		dy = fmt.Sprintf("%f", y)
	}

	fill := r.URL.Query().Get("fill")

	if fill == "" {
		fill = "000000"
	}

	format := r.URL.Query().Get("format")

	if format == "" {
		format = "2006-01-02"
	}

	if text == "" {
		text = time.Now().Format(format)
	} else {
		text = text + " " + time.Now().Format(format)
	}
	w.Header().Set("Content-Type", "image/svg+xml")

	output := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"  viewBox="0 0 %s %s" >
    <text   text-anchor="end" x='%s' y='%s' style="font-family: Arial, Helvetica, Verdana" font-weight="bold" fill="#%s" dx='%s' dy='%s' font-size="%s">%s</text>
</svg>`, width, height, fmt.Sprintf("%d", x), fmt.Sprintf("%d", y), fill, dx, dy, size, text)

	//  w.Header().Set("Content-Length", fmt.Sprintf("%f",len(output)))
	fmt.Fprint(w, output)
}
