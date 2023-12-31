package leetcode

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LeetCode struct {
	url    string
	client *http.Client
}

func New(url string) *LeetCode {
	return &LeetCode{url: url, client: http.DefaultClient}
}

func (l *LeetCode) Stats(ctx context.Context, username string) (resp LCGetProblemsSolvedResp, err error) {
	var (
		in = struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables"`
		}{
			Query: query,
			Variables: map[string]interface{}{
				"username": username,
			},
		}
		out = struct {
			Data struct {
				AllQuestionsCount []struct {
					Difficulty string `json:"difficulty"`
					Count      int    `json:"count"`
				} `json:"allQuestionsCount"`
				MatchedUser struct {
					ProblemsSolvedBeatsStats []struct {
						Difficulty string  `json:"difficulty"`
						Percentage float64 `json:"percentage"`
					} `json:"problemsSolvedBeatsStats"`
					SubmitStatsGlobal struct {
						AcSubmissionNum []struct {
							Difficulty string `json:"difficulty"`
							Count      uint64 `json:"count"`
						} `json:"acSubmissionNum"`
					} `json:"submitStatsGlobal"`
				} `json:"matchedUser"`
			} `json:"data"`
		}{}
	)

	// Marshal the request solvedReq to JSON
	jsonData, err := json.Marshal(in)
	if err != nil {
		return resp, fmt.Errorf("marshal solvedReq: %w", err)
	}

	// Make the POST request
	response, err := http.Post(l.url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return resp, fmt.Errorf("post request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return resp, fmt.Errorf("read response body from leetcode: %w", err)
	}

	if err = json.Unmarshal(body, &out); err != nil {
		return resp, fmt.Errorf("unmarshal allQuestionsCount: %w", err)
	}

	if len(out.Data.MatchedUser.SubmitStatsGlobal.AcSubmissionNum) > 3 {
		resp.Total = out.Data.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[0].Count
		resp.Easy = out.Data.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[1].Count
		resp.Medium = out.Data.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[2].Count
		resp.Hard = out.Data.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[3].Count
		return resp, nil
	}
	return resp, fmt.Errorf("no leetcode profile with this username %v", username)
}
