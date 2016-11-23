package hello

import (
    "fmt"
    "net/http"
    "time"
)

func init() {
    http.HandleFunc("/svg/month", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {


text := r.URL.Query().Get("text")
width := r.URL.Query().Get("width")
if width == "" {
   width = "300"
}

height := r.URL.Query().Get("height")
if height  == "" {
   height = "20"
}
dy := r.URL.Query().Get("dy")

if dy  == "" {
   dy = "-2.6em"
}

fill := r.URL.Query().Get("fill")

if fill  == "" {
   fill = "000000"
}


format := r.URL.Query().Get("format")

if format  == "" {
   format = "January 2006"
}



if text == "" {
   text = time.Now().Format(format)
} else {
   text = text + " " + time.Now().Format(format)
}
    w.Header().Set("Content-Type","image/svg+xml")
    


   output := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"  viewBox="0 0 %s %s">
  

    <text text-anchor="left"  style="font-family: Arial, Helvetica, Verdana" font-weight="bold" fill="#%s" dy='%s' font-size="ANY SIZE">%s</text>
</svg>`, width, height, fill, dy, text)

  //  w.Header().Set("Content-Length", fmt.Sprintf("%d",len(output)))
    fmt.Fprint(w, output)
}
