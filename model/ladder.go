package model

type ladder struct {
	startPoint int64
	endPoint   int64
}

func InitLadder(startPoint, endPoint int64) *ladder {
	return &ladder{
		startPoint: startPoint,
		endPoint:   endPoint,
	}
}
