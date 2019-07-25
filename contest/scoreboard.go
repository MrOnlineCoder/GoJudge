package contest

type ProblemRanking struct {
	Score int `json:"score"`
	Solved bool	`json:"solved"`
}

type Row struct {
	UserID int `json:"user_id"`
	Rankings []ProblemRanking `json:"rankings"`
	Total int `json:"total"`
}

type Scoreboard interface {
	ProcessSubmissionVerdict(user int, problem_id int, verdict string, passedTests int, totalTests int)
	GetRows() []Row
}

var scoreboard ScoreboardTime;

func GetScoreboard() *ScoreboardTime {
	return &scoreboard;
}