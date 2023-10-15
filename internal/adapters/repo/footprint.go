package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

type footprint struct {
	db *sql.DB
}

func newFootprint(db *sql.DB) *footprint {
	return &footprint{db: db}
}

func (f *footprint) Create(ctx context.Context, obj *models.Footprint) (int64, error) {
	query := `
		INSERT INTO footprint (usr_id, easy, medium, hard, total)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	var id int64
	err := f.db.QueryRowContext(ctx, query, obj.UsrID, obj.Easy, obj.Medium, obj.Hard, obj.Total).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create footprint: %w", err)
	}

	return id, nil
}

func (f *footprint) Update(ctx context.Context, obj *models.Footprint, id int64) error {
	query := `
		UPDATE footprint
		SET usr_id=$1, easy=$2, medium=$3, hard=$4, total=$5
		WHERE id=$6`

	_, err := f.db.ExecContext(ctx, query, obj.UsrID, obj.Easy, obj.Medium, obj.Hard, obj.Total, id)
	if err != nil {
		return fmt.Errorf("update footprint: %w", err)
	}

	return nil
}

func (f *footprint) Get(ctx context.Context, pars *models.FootprintGetPars) (*models.Footprint, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT f.id, f.usr_id, f.easy, f.medium, f.hard, f.total, f.active `
	queryFrom := ` FROM footprint f`
	queryWhere := ` WHERE 1 = 1`

	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND f.id = $` + strconv.Itoa(len(filterValues))
	}

	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND f.usr_id = $` + strconv.Itoa(len(filterValues))
	}

	fp := &models.Footprint{}

	row := f.db.QueryRowContext(ctx, querySelect+queryFrom+queryWhere, filterValues...)

	err := row.Scan(
		&fp.ID,
		&fp.UsrID,
		&fp.Easy,
		&fp.Medium,
		&fp.Hard,
		&fp.Total,
		&fp.Active,
	)
	if err != nil {
		return nil, fmt.Errorf("footprint row scan: %w", err)
	}

	err = row.Err()
	if err != nil {
		return nil, fmt.Errorf("footprint row: %w", err)
	}

	return fp, nil
}

func (f *footprint) List(ctx context.Context, pars *models.FootprintListPars) ([]*models.Footprint, int64, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT f.id, f.usr_id, f.easy, f.medium, f.hard, f.total, f.active`
	queryFrom := ` FROM footprint f`
	queryWhere := ` WHERE 1 = 1`
	queryOrderBy := ` order by f.id desc`
	queryOffset := ``
	queryLimit := ``

	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND f.id = $` + strconv.Itoa(len(filterValues))
	}

	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND f.usr_id = $` + strconv.Itoa(len(filterValues))
	}

	var tCount int64
	if pars.Limit > 0 {
		err := f.db.QueryRowContext(ctx, `
		select count(*) `+queryFrom+queryWhere, filterValues...).Scan(&tCount)
		if err != nil {
			return nil, 0, fmt.Errorf("get footprints count: %w", err)
		}
		if pars.OnlyCount {
			return nil, tCount, nil
		}
		queryOffset = ` offset ` + strconv.FormatInt(pars.Offset, 10)
		queryLimit = ` limit ` + strconv.FormatInt(pars.Limit, 10)
	}

	rows, err := f.db.Query(querySelect+queryFrom+queryWhere+queryOrderBy+queryOffset+queryLimit, filterValues...)
	if err != nil {
		return nil, 0, fmt.Errorf("query footprints: %w", err)
	}
	defer rows.Close()

	var fps []*models.Footprint

	for rows.Next() {
		fp := &models.Footprint{}
		err := rows.Scan(
			&fp.ID,
			&fp.UsrID,
			&fp.Easy,
			&fp.Medium,
			&fp.Hard,
			&fp.Total,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("rows scan: %w", err)
		}

		fps = append(fps, fp)
	}

	err = rows.Err()
	if err != nil {
		return nil, 0, fmt.Errorf("rows err: %w", err)
	}

	return fps, tCount, nil

}
