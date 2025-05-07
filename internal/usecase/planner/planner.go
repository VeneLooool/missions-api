package planner

import (
	"math"

	"github.com/VeneLooool/missions-api/internal/model"
)

const (
	coverageWidth = 30.0      // ширина охвата (в метрах)
	overlap       = 0.2       // процент перекрытия
	stepY         = 10.0      // шаг по вертикали (1 м для точности фильтрации)
	earthRadius   = 6378137.0 // радиус Земли для EPSG:3857
)

func lonLatToWebMercator(lat, lon float64) (x, y float64) {
	const originShift = math.Pi * earthRadius
	x = lon * originShift / 180.0
	y = math.Log(math.Tan((90+lat)*math.Pi/360.0)) / (math.Pi / 180.0)
	y = y * originShift / 180.0
	return
}

func webMercatorToLonLat(x, y float64) (lat, lon float64) {
	const originShift = math.Pi * earthRadius
	lon = (x / originShift) * 180.0
	lat = (y / originShift) * 180.0
	lat = 180 / math.Pi * (2*math.Atan(math.Exp(lat*math.Pi/180.0)) - math.Pi/2)
	return
}

func pointInPolygon(x, y float64, polygon [][2]float64) bool {
	n := len(polygon)
	inside := false
	j := n - 1
	for i := 0; i < n; i++ {
		xi, yi := polygon[i][0], polygon[i][1]
		xj, yj := polygon[j][0], polygon[j][1]
		if ((yi > y) != (yj > y)) &&
			(x < (xj-xi)*(y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
		j = i
	}
	return inside
}

func GenerateLawnMowerPath(coords model.Coordinates) model.Coordinates {
	if len(coords) == 0 {
		return nil
	}

	// Переводим полигон в метры
	var poly [][2]float64
	for _, c := range coords {
		x, y := lonLatToWebMercator(float64(c.Latitude), float64(c.Longitude))
		poly = append(poly, [2]float64{x, y})
	}

	// Bounding box
	minX, maxX := poly[0][0], poly[0][0]
	minY, maxY := poly[0][1], poly[0][1]
	for _, p := range poly {
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
		if p[1] < minY {
			minY = p[1]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	}

	stepX := coverageWidth * (1 - overlap)
	var result model.Coordinates
	up := true

	for x := minX; x <= maxX; x += stepX {
		var inPoint, outPoint *model.Coordinate

		var yStart, yEnd, yStep float64
		if up {
			yStart, yEnd, yStep = minY, maxY, stepY
		} else {
			yStart, yEnd, yStep = maxY, minY, -stepY
		}
		up = !up

		for y := yStart; (yStep > 0 && y <= yEnd) || (yStep < 0 && y >= yEnd); y += yStep {
			if pointInPolygon(x, y, poly) {
				lat, lon := webMercatorToLonLat(x, y)
				pt := model.Coordinate{
					Latitude:  float32(lat),
					Longitude: float32(lon),
				}
				if inPoint == nil {
					inPoint = &pt
				}
				outPoint = &pt
			}
		}

		if inPoint != nil && outPoint != nil {
			// Добавляем маршрутный отрезок
			result = append(result, *inPoint, *outPoint)
		}
	}

	return result
}
