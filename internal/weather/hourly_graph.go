package weather

import (
	"fmt"
	"math"
	"strings"
)

var dotBits = [2][4]rune{
	{0x01, 0x02, 0x04, 0x40},
	{0x08, 0x10, 0x20, 0x80},
}

func interp(vals []float64, n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		pos := float64(i) / float64(n-1) * float64(len(vals)-1)
		lo := int(math.Floor(pos))
		hi := min(lo+1, len(vals)-1)
		f := pos - float64(lo)
		out[i] = vals[lo]*(1-f) + vals[hi]*f
	}
	return out
}

func labelRow(w int, pos []int, labels []string) string {
	row := []rune(strings.Repeat(" ", w))
	for i, l := range labels {
		r := []rune(l)
		start := max(0, min(pos[i]-len(r)/2, w-len(r)))
		copy(row[start:], r)
	}
	return string(row)
}

func RenderHourly(hours []HourPoint, w, h int, u Units) string {
	if len(hours) < 2 || w < 10 {
		return ""
	}

	temps := make([]float64, len(hours))
	lo, hi := math.Inf(1), math.Inf(-1)
	for i, p := range hours {
		temps[i] = u.Temp(p.TempC)
		lo, hi = math.Min(lo, temps[i]), math.Max(hi, temps[i])
	}
	span := math.Max(hi-lo, 1)
	floor := lo - span*0.5
	rng := hi - floor

	levels := h * 4
	samples := interp(temps, w*2)
	height := make([]int, len(samples))
	for i, v := range samples {
		height[i] = int(math.Round((v-floor)/rng*float64(levels-1))) + 1
	}

	step := max(1, int(math.Ceil(float64(len(hours))/float64(max(1, w/8)))))
	var pos []int
	var tempLabels, timeLabels []string
	for i := 0; i < len(hours); i += step {
		pos = append(pos, i*(w-1)/(len(hours)-1))
		tempLabels = append(tempLabels, fmt.Sprintf("%.0f°", temps[i]))
		timeLabels = append(timeLabels, hours[i].Time.Format("3 PM"))
	}

	var b strings.Builder
	b.WriteString(labelRow(w, pos, tempLabels) + "\n")
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			var cell rune
			for sc := 0; sc < 2; sc++ {
				for dy := 0; dy < 4; dy++ {
					fromBottom := (h-1-r)*4 + (3 - dy) + 1
					if fromBottom <= height[c*2+sc] {
						cell |= dotBits[sc][dy]
					}
				}
			}
			b.WriteRune(0x2800 + cell)
		}
		b.WriteByte('\n')
	}
	b.WriteString(labelRow(w, pos, timeLabels))
	return b.String()
}
