package contest

import (
	"time"
	"log"
)

type Penalty struct {
	ProblemsMap map[int]int
}

type ScoreboardTime struct {
	StartTime int64
	FailPenalty int

	Rows []Row
	Problemset []ContestProblem
	Penalties map[int]Penalty
}

func (sb *ScoreboardTime) FindUserRow(user int) *Row {
	for _, row := range sb.Rows {
		if row.UserID == user {
			return &row;
		}
	}

	return nil;
}

func (sb *ScoreboardTime) AddUserRow(user int) {
	nrow := Row{
		UserID: user,
		Rankings: make([]ProblemRanking, len(sb.Problemset)),
		Total: 0,
	};

	sb.Rows = append(sb.Rows, nrow);
}

func (sb *ScoreboardTime) GetProblemIndex(problem_id int) int {
	for idx, prob := range sb.Problemset {
		if prob.Problem.Id == problem_id {
			return idx;
		}
	}

	log.Printf("[Scoreboard] ERROR: failed to get problem index for problem %d\n", problem_id);

	return -1;
}

func (sb *ScoreboardTime) AddPenaltyForProblem(user int, problem_id int) {
	oldVal := sb.GetPenaltyForProblem(user, problem_id);

	sb.Penalties[user].ProblemsMap[problem_id] = oldVal + 1;
}

func (sb *ScoreboardTime) GetPenaltyForProblem(user int, problem_id int) int {
	pen, ok := sb.Penalties[user];

	if !ok {
		sb.Penalties[user] = Penalty{ProblemsMap: make(map[int]int)};
		return 0;
	}

	val, ok := pen.ProblemsMap[problem_id];

	if !ok {
		pen.ProblemsMap[problem_id] = 0;
		return 0;
	}

	return val * sb.FailPenalty;
}

func (sb *ScoreboardTime) Init(problemset []ContestProblem, start_time int64, fail_penalty int) {
	sb.Problemset = problemset;
	sb.StartTime = start_time;
	sb.FailPenalty = fail_penalty;
	sb.Penalties = make(map[int]Penalty)
}

func (sb *ScoreboardTime) ProcessSubmissionVerdict(user int, problem_id int, verdict string, passedTests int, totalTests int) {
	row := sb.FindUserRow(user);

	if row == nil {
		sb.AddUserRow(user);
		row = sb.FindUserRow(user);
	}

	pidx := sb.GetProblemIndex(problem_id);

	ranking := row.Rankings[pidx];

	if ranking.Solved {
		return;
	}

	elapsed := int(time.Since(time.Unix(sb.StartTime / 1000, 0)).Minutes()); 

	penalty := sb.GetPenaltyForProblem(user, problem_id) + elapsed;

	if passedTests != totalTests {
		sb.AddPenaltyForProblem(user, problem_id);
	} else {
		row.Rankings[pidx].Score = sb.Problemset[pidx].Points - penalty;
		row.Rankings[pidx].Solved = true;
	}
}

func (sb *ScoreboardTime) GetRows() []Row {
	return sb.Rows;
}