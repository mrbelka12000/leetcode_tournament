package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/lib/pq"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type tournament struct {
	db *sql.DB
}

func newTournament(db *sql.DB) *tournament {
	return &tournament{db: db}
}

func (t *tournament) Create(ctx context.Context, obj *models.TournamentCU) (int64, error) {
	var id int64

	err := t.db.QueryRowContext(ctx, `
		INSERT INTO tournament
		(usr_id, start_time, end_time, goal, cost, prize_pool, status_id)
		VALUES 
		($1,$2,$3,$4,$5,$6,$7)
		RETURNING id
		`,
		*obj.UsrID, *obj.StartTime, *obj.EndTime, *obj.Goal, *obj.Cost, *obj.PrizePool, *obj.StatusID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("query row context: %w", err)
	}

	return id, nil
}

func (t *tournament) Update(ctx context.Context, obj *models.TournamentCU, id int64) error {
	updateValues := []interface{}{id}
	queryUpdate := ` UPDATE tournament t`
	querySet := ` SET id = id`

	queryWhere := ` WHERE id = $1`
	if obj.UsrID != nil {
		updateValues = append(updateValues, *obj.UsrID)
		querySet += ` , usr_id = $` + strconv.Itoa(len(updateValues))
	}
	if obj.StartTime != nil {
		updateValues = append(updateValues, *obj.StartTime)
		querySet += ` , start_time = $` + strconv.Itoa(len(updateValues))
	}
	if obj.EndTime != nil {
		updateValues = append(updateValues, *obj.EndTime)
		querySet += ` , end_time = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Goal != nil {
		updateValues = append(updateValues, *obj.Goal)
		querySet += ` , goal = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Cost != nil {
		updateValues = append(updateValues, *obj.Cost)
		querySet += ` , cost = $` + strconv.Itoa(len(updateValues))
	}
	if obj.PrizePool != nil {
		updateValues = append(updateValues, *obj.PrizePool)
		querySet += ` , prize_pool = $` + strconv.Itoa(len(updateValues))
	}
	if obj.StatusID != nil {
		updateValues = append(updateValues, *obj.StatusID)
		querySet += ` , status_id = $` + strconv.Itoa(len(updateValues))
	}

	_, err := t.db.ExecContext(ctx, queryUpdate+querySet+queryWhere, updateValues...)
	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

func (t *tournament) Get(ctx context.Context, pars *models.TournamentGetPars) (*models.Tournament, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT t.id, t.usr_id, t.start_time, t.end_time, t.goal, t.cost, t.prize_pool, t.status_id
`

	queryFrom := ` FROM tournament t`
	queryWhere := ` WHERE 1 = 1`
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND t.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND t.usr_id = $` + strconv.Itoa(len(filterValues))
	}

	tr := &models.Tournament{}
	row := t.db.QueryRowContext(ctx, querySelect+queryFrom+queryWhere, filterValues...)
	err := row.Scan(
		&tr.ID,
		&tr.UsrID,
		&tr.StartTime,
		&tr.EndTime,
		&tr.Goal,
		&tr.Cost,
		&tr.PrizePool,
		&tr.StatusID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("row scan: %w", err)
	}

	err = row.Err()
	if err != nil {
		return nil, fmt.Errorf("row err: %w", err)
	}

	return tr, nil
}

func (t *tournament) List(ctx context.Context, pars *models.TournamentListPars) ([]*models.Tournament, int64, error) {
	var err error

	var filterValues []interface{}
	querySelect := `
	SELECT t.id, t.usr_id, t.start_time, t.end_time, t.goal, t.cost, t.prize_pool, t.status_id
`
	queryFrom := ` FROM tournament t`
	queryWhere := ` WHERE 1 = 1`
	queryOffset := ``
	queryLimit := ``
	queryOrderBy := " order by t.id desc"

	// filter
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND t.usr_id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND t.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.IDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.IDs))
		queryWhere += fmt.Sprintf(` AND t.id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.UsrIDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.UsrIDs))
		queryWhere += fmt.Sprintf(` AND t.usr_id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.StatusIDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.StatusIDs))
		queryWhere += fmt.Sprintf(` AND t.status_id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}

	var tCount int64
	if pars.Limit > 0 {
		err := t.db.QueryRowContext(ctx, `select count(*)`+queryFrom+queryWhere, filterValues...).Scan(&tCount)
		if err != nil {
			return nil, 0, err
		}
		if pars.OnlyCount {
			return nil, tCount, nil
		}
		queryOffset = ` offset ` + strconv.FormatInt(pars.Offset, 10)
		queryLimit = ` limit ` + strconv.FormatInt(pars.Limit, 10)
	}

	rows, err := t.db.QueryContext(ctx, querySelect+queryFrom+queryWhere+queryOrderBy+queryOffset+queryLimit, filterValues...)
	if err != nil {
		return nil, 0, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var tournaments []*models.Tournament
	for rows.Next() {
		tr := &models.Tournament{}
		err := rows.Scan(
			&tr.ID,
			&tr.UsrID,
			&tr.StartTime,
			&tr.EndTime,
			&tr.Goal,
			&tr.Cost,
			&tr.PrizePool,
			&tr.StatusID,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("rows scan: %w", err)
		}

		tournaments = append(tournaments, tr)
	}

	return tournaments, tCount, nil
}
