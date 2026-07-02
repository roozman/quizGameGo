package entities

type questions struct {
	ID              uint
	Text            string
	PossibleAnswers []PossibleAnswer
	CorrectAnswerID uint
	Difficulty      QuestionDifficulty
	Category        Category
}

type PossibleAnswer struct {
	ID      uint
	content string
	choice  PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	if p >= PossibleAnswera && p <= PossibleAnswerd {
		return true
	}
	return false
}

const (
	PossibleAnswera = iota + 1
	PossibleAnswerb
	PossibleAnswerc
	PossibleAnswerd
)

type QuestionDifficulty uint8

const (
	QuesttionDifficultyEasy QuestionDifficulty = iota + 1
	QuesttionDifficultyMedium
	QuestionDifficultyHard
)

func (d QuestionDifficulty) IsValid() bool {
	if d >= QuesttionDifficultyEasy && d <= QuestionDifficultyHard {
		return true
	}
	return false
}
