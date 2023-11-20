package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.05
	angle         = math.Pi / 6
)

var sin, cos = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/draw", draw)
	log.Fatal(http.ListenAndServe("localhost:3333", nil))
}

func draw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			color := "red"
			if az < 0 || bz < 0 || cz < 0 || dz < 0 {
				color = "blue"
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				" style='fill: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos*xyscale
	sy := height/2 + (x+y)*sin*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Sin(x) - math.Sin(y)
	return r
}
