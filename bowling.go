package bowling

import "fmt"

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {

	if len(game) != 10 {
		return 0, fmt.Errorf("Less than 10 frame found")
	} else {

		score := 0

		for i := 0; i < len(game); i++ {
			score = score + frameScore(game, i)
		}
		return score, nil
	}
}

func frameScore(game []Frame, position int) (int, error) {
	if game[position].firstThrow >= 0 && game[position].secondThrow >= 0 {

		currentFrame := game[position].firstThrow + game[position].secondThrow

		if currentFrame <= 10 {

			if currentFrame == 10 && position+1 < len(game) {

				if game[position].firstThrow == 10 {
					// Strike Score
					currentFrame += game[position+1].firstThrow + game[position+1].secondThrow
				} else {
					// Spare Score
					currentFrame += game[position+1].firstThrow
				}
			}
			return currentFrame
		} else {
			return 0
		}
	} else {
		return 0
	}
}
