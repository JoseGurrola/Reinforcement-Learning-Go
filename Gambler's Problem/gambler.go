package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
	"strconv"
)

const (
	GOAL = 100
	ph   = 0.55   //Probabilidad de que el lanzamiento sea cara
	γ    = 1 - ph //Factor de olvido
	θ    = 0.0000000001
)

func main() {

	stateValue := make([]float64, GOAL+1) //Arreglo estado-valor
	stateValue[GOAL] = 1

	Policys := make([]int, GOAL+1) //Arreglo de politicas

	//ALGORITHM VALUE-ITERATION

	for delta := 1.0; delta > θ; {
		//∆ ← 0
		delta = 0
		for s := 1; s < GOAL; s++ {
			//v ← V (s)
			v := stateValue[s]
			//V (s) ← maxa E(s',r) p(s', r|s, a) [r + γV (s')]
			stateValue[s], _ = MaxP(s, stateValue)
			//∆ ← max(∆, |v − V (s)|)
			delta = math.Max(delta, math.Abs(v-stateValue[s]))
		}
	}
	//π(s) = argmaxa E(s',r) p(s', r | s, a)[r + γV(s')]
	for s := 0; s < GOAL; s++ {
		_, Policys[s] = MaxP(s, stateValue)
	}

	//fmt.Println(stateValue)
	//fmt.Println()
	//fmt.Print(Policys)

	GraphicSV(stateValue)
	GraphicP(Policys)
	fmt.Println("Graficas generadas!.")
}

func MaxP(s int, sv []float64) (float64, int) {

	n := int(math.Min(float64(s), float64(GOAL-s)))

	max := ph*sv[s] + γ*sv[s]
	argMax := 0
	for action := 1; action <= n; action++ {
		aux := ph*sv[s+action] + γ*sv[s-action]
		if aux > max {
			max = aux
			argMax = action
		}
	}

	return max, argMax
}

func GraphicSV(array []float64) {
	const S = 500
	const P = 32
	dc := gg.NewContext(S, (S+100)/2)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	//create points
	points := make([]gg.Point, len(array))
	for i := 0; i < len(array); i++ {
		x := float64(i)
		y := array[i]
		points[i] = gg.Point{x, y}
	}

	dc.Translate(P, P)
	dc.Scale(S-P*2, S-P)

	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		y := float64(i) / 20
		dc.MoveTo(x, 0)
		dc.LineTo(x, 0.5)
		dc.MoveTo(0, y)
		dc.LineTo(1, y)
	}

	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(1)
	dc.Stroke()

	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, 0)
	dc.LineTo(0, 0.5)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	dc.Stroke()

	// draw points
	dc.SetRGBA(0, 0, 1, 0.5)
	for _, p := range points {
		dc.DrawCircle(p.X/100.0, p.Y/2, 3.0/S)
		dc.Fill()
	}

	dc.SetLineWidth(2)

	for i := 1; i < len(points); i++ {
		dc.SetRGB(0, 0, 100)
		dc.Stroke()
		dc.DrawLine(points[i].X/100.0, points[i].Y/2, points[i-1].X/100.0, points[i-1].Y/2)
	}

	// draw text
	dc.Identity()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Capital", 260, 290, 0.5, 0.5)
	dc.DrawString("ValueFunction Found", 200, 20)

	for i := 10; i >= 0; i-- {
		dc.DrawString(strconv.FormatFloat(float64(i)/10, 'f', 1, 64), 7, 270-(float64(i)*23))
	}

	for i := 0; i <= 100; i += 10 {
		dc.DrawString(strconv.Itoa(i), float64(i)*4.3+30, 280)
	}
	dc.SavePNG("valueFunctionFound.png")
}

func GraphicP(array []int) {
	const S = 500
	const P = 32
	dc := gg.NewContext(S, (S+100)/2)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	//create points
	points := make([]gg.Point, len(array))
	for i := 0; i < len(array); i++ {
		x := float64(i)
		y := float64(array[i])
		points[i] = gg.Point{x, y}
	}

	dc.Translate(P, P)
	dc.Scale(S-P*2, S-P)

	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		y := float64(i) / 20
		dc.MoveTo(x, 0)
		dc.LineTo(x, 0.5)
		dc.MoveTo(0, y)
		dc.LineTo(1, y)
	}

	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(1)
	dc.Stroke()

	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, 0)
	dc.LineTo(0, 0.5)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	dc.Stroke()

	// draw points
	dc.SetRGBA(0, 0, 1, 0.5)
	for _, p := range points {
		dc.DrawCircle(p.X/100.0, p.Y/100.0, 3.0/S)
		dc.Fill()
	}

	dc.SetLineWidth(2)
	for i := 0; i < len(points)-1; i++ {
		dc.SetRGB(0, 0, 100)
		dc.Stroke()
		dc.DrawLine(points[i].X/100.0, points[i].Y/100.0, points[i+1].X/100.0, points[i+1].Y/100.0)
	}

	// draw text
	dc.Identity()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Capital", 260, 290, 0.5, 0.5)
	dc.DrawString("Final policy", 200, 20)

	for i := 50; i >= 0; i -= 5 {
		dc.DrawString(strconv.Itoa(i), 7, 270-(float64(i)*4.6))
	}

	for i := 0; i <= 100; i += 10 {
		dc.DrawString(strconv.Itoa(i), float64(i)*4.3+30, 280)
	}
	dc.SavePNG("FinalPolicy.png")
}
