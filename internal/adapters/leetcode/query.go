package leetcode

type LCGetProblemsSolvedReq struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

const query = `
	query userProblemsSolved($username: String!) {
		allQuestionsCount {
			difficulty
			count
		}
		matchedUser(username: $username) {
			problemsSolvedBeatsStats {
				difficulty
				percentage
			}
			submitStatsGlobal {
				acSubmissionNum {
					difficulty
					count
				}
			}
		}
	}`

type LCGetProblemsSolvedResp struct {
	Easy   int
	Medium int
	Hard   int
	Total  int
}
