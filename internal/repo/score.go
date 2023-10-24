package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type score struct {
	db *sql.DB
}

func newScore(db *sql.DB) *score {
	return &score{db: db}
}

func (s *score) Create(ctx context.Context, obj models.ScoreCU) (int64, error) {
	var id int64

	err := QueryRowContext(ctx, s.db, `
		INSERT INTO score
		(usr_id, current, active)
		VALUES 
		($1,$2,$3)
		RETURNING id
		`,
		*obj.UsrID, *obj.Current, *obj.Active).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("query row context: %w", err)
	}

	return id, nil
}

func (s *score) Update(ctx context.Context, obj models.ScoreCU, id int64) error {
	updateValues := []interface{}{id}
	queryUpdate := ` UPDATE score s`
	querySet := ` SET id = id`

	queryWhere := ` WHERE id = $1`
	if obj.UsrID != nil {
		updateValues = append(updateValues, *obj.UsrID)
		querySet += ` , usr_id = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Current != nil {
		updateValues = append(updateValues, *obj.Current)
		querySet += ` , current = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Footprint != nil {
		updateValues = append(updateValues, *obj.Footprint)
		querySet += ` , footprint = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Active != nil {
		updateValues = append(updateValues, *obj.Active)
		querySet += ` , active = $` + strconv.Itoa(len(updateValues))
	}

	_, err := ExecContext(ctx, s.db, queryUpdate+querySet+queryWhere, updateValues...)
	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

func (s *score) Get(ctx context.Context, pars models.ScoreGetPars) (models.Score, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT s.id, s.usr_id, s.current,s.footprint,s.active
`

	queryFrom := ` FROM score s`
	queryWhere := ` WHERE 1 = 1`
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND s.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND s.usr_id = $` + strconv.Itoa(len(filterValues))
	}

	var sc models.Score
	row := QueryRowContext(ctx, s.db, querySelect+queryFrom+queryWhere, filterValues...)
	err := row.Scan(
		&sc.ID,
		&sc.UsrID,
		&sc.Current,
		&sc.Footprint,
		&sc.Active,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Score{}, nil
		}
		return models.Score{}, fmt.Errorf("row scan: %w", err)
	}

	err = row.Err()
	if err != nil {
		return models.Score{}, fmt.Errorf("row err: %w", err)
	}

	return sc, nil
}

func (s *score) List(ctx context.Context, pars models.ScoreListPars) ([]models.Score, int64, error) {
	var err error

	var filterValues []interface{}
	querySelect := `
	SELECT s.id,s.usr_id, s.current, s.footprint, s.active
`
	queryFrom := ` FROM score s`
	queryWhere := ` WHERE 1 = 1`
	queryOffset := ``
	queryLimit := ``
	queryOrderBy := " order by s.id desc"

	// filter
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND s.usr_id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND s.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.IDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.IDs))
		queryWhere += fmt.Sprintf(` AND s.id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.UsrIDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.UsrIDs))
		queryWhere += fmt.Sprintf(` AND s.usr_id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}

	var tCount int64
	if pars.Limit > 0 {
		err := QueryRowContext(ctx, s.db, `select count(*)`+queryFrom+queryWhere, filterValues...).Scan(&tCount)
		if err != nil {
			return nil, 0, err
		}
		if pars.OnlyCount {
			return nil, tCount, nil
		}
		queryOffset = ` offset ` + strconv.FormatInt(pars.Offset, 10)
		queryLimit = ` limit ` + strconv.FormatInt(pars.Limit, 10)
	}

	rows, err := QueryContext(ctx, s.db, querySelect+queryFrom+queryWhere+queryOrderBy+queryOffset+queryLimit, filterValues...)
	if err != nil {
		return nil, 0, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var scores []models.Score
	for rows.Next() {
		var sc models.Score
		err := rows.Scan(
			&sc.ID,
			&sc.UsrID,
			&sc.Current,
			&sc.Footprint,
			&sc.Active,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("rows scan: %w", err)
		}

		scores = append(scores, sc)
	}

	return scores, tCount, nil
}
