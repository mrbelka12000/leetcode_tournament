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
	Easy   uint64
	Medium uint64
	Hard   uint64
	Total  uint64
}
