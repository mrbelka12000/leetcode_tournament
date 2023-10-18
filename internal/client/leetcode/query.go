package leetcode

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
