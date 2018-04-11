package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	WORLD_HEIGHT int = 4
	WORLD_WIDTH  int = 12

	NUM_ACTIONS int = 4

	ACTION_UP    int = 0
	ACTION_DOWN  int = 1
	ACTION_LEFT  int = 2
	ACTION_RIGHT int = 3

	ε float64 = 0.1
	α float64 = 0.5
	γ float64 = 1.0
)

func main() {
	rand.Seed(time.Now().Unix())
	QLearning := QLearning{}
	QLearning.Initialize()
	QLearning.Start()
	QLearning.Print()
}

type QLearning struct {
	Q [][]float64

	xIni int
	yIni int
	xFin int
	yFin int
}

func (q *QLearning) Initialize() {
	q.xIni = 0
	q.yIni = 0
	q.xFin = 11
	q.yFin = 0

	q.Q = make([][]float64, NUM_ACTIONS)

	for i := 0; i < NUM_ACTIONS; i++ {
		q.Q[i] = make([]float64, WORLD_HEIGHT*WORLD_WIDTH)
	}

	for i := 0; i < NUM_ACTIONS; i++ {
		for j := 0; j < WORLD_HEIGHT*WORLD_WIDTH; j++ {
			q.Q[i][j] = rand.Float64()
		}
	}

	for a := 0; a < NUM_ACTIONS; a++ {
		q.Q[a][q.yFin*WORLD_HEIGHT+q.xFin] = 0
	}
}

func (q *QLearning) Start() {
	episodes := 1000
	for i := 0; i < episodes; i++ {
		//Initialize S
		Sx := q.xIni
		Sy := q.yIni

		for Sy != q.yFin || Sx != q.xFin {
			//Choose A from S using policy derived from Q
			action := q.εGreedy(Sx, Sy)

			R, _Sx, _Sy := q.Action(Sx, Sy, action)
			//Q(S, A) ← Q(S, A) + α[R + γQ(S', A') − Q(S, A)]
			QSA := q.Q[action][q.GetIdx(Sx, Sy)]
			maxAction := q.GetAction(_Sx, _Sy)
			_QSA := q.Q[maxAction][q.GetIdx(_Sx, _Sy)]
			Q := QSA + α*(R+γ*_QSA-QSA)
			q.Q[action][q.GetIdx(Sx, Sy)] = Q
			//S ← S'; A ← A'
			Sx = _Sx
			Sy = _Sy
		}
	}
}

func (q *QLearning) Action(x, y, a int) (float64, int, int) {
	_x := x
	_y := y

	switch a {
	case ACTION_UP:
		if y != WORLD_HEIGHT-1 {
			_y = y + 1
		}
	case ACTION_DOWN:
		if y != 0 {
			_y = y - 1
		}
	case ACTION_LEFT:
		if x != 0 {
			_x = x - 1
		}
	case ACTION_RIGHT:
		if x != WORLD_WIDTH-1 {
			_x = x + 1
		}
	}

	if _y == q.yFin && _x == q.xFin {
		return 0.0, q.xFin, q.yFin
	} else if _y == 0 && _x >= 1 && _x < WORLD_WIDTH-1 {
		return -100.0, 0, 0
	}

	return -1.0, _x, _y
}

func (q *QLearning) εGreedy(x, y int) int {
	action := q.GetAction(x, y)

	if rand.Float64() < 1-ε {
		return action
	}
	return rand.Intn(NUM_ACTIONS)
}

func (q *QLearning) GetAction(x, y int) int {
	Idx := q.GetIdx(x, y)
	max := q.Q[0][Idx]
	action := 0
	for i := 0; i < NUM_ACTIONS; i++ {
		if max < q.Q[i][Idx] {
			max = q.Q[i][Idx]
			action = i
		}
	}

	return action
}

func (q *QLearning) Print() {
	for i := WORLD_HEIGHT - 1; i >= 0; i-- {
		for j := 0; j < WORLD_WIDTH; j++ {
			if i != q.yFin || j != q.xFin {
				if j > 0 && j < WORLD_WIDTH-1 && i == 0 {
					fmt.Print("<-| ")
				} else {
					switch q.GetAction(j, i) {
					case ACTION_UP:
						fmt.Print("U | ")
					case ACTION_DOWN:
						fmt.Print("D | ")
					case ACTION_LEFT:
						fmt.Print("L | ")
					case ACTION_RIGHT:
						fmt.Print("R | ")
					}
				}

			} else {
				fmt.Print("G | ")
			}
		}
		fmt.Println("")
	}

	//	for i := 0; i < len(q.wind); i++ {
	//		fmt.Print(q.wind[i], "   ")
	//	}
}

func (q *QLearning) GetIdx(x, y int) int {
	return y*WORLD_WIDTH + x
}
