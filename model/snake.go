package model

type snake struct {
	startPoint int64
	endPoint   int64
}

func InitSnake(startPoint, endPoint int64) *snake {
	return &snake{
		startPoint: startPoint,
		endPoint:   endPoint,
	}
}
