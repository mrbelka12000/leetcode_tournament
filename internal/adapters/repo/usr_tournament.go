package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/lib/pq"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

type usrTournament struct {
	db *sql.DB
}

func newUsrTournament(db *sql.DB) *usrTournament {
	return &usrTournament{db: db}
}

func (u usrTournament) Create(ctx context.Context, obj *models.UsrTournamentCU) (int64, error) {
	var id int64

	err := u.db.QueryRowContext(ctx, `
		INSERT INTO usr_tournament
		(usr_id, tournament_id, active, winner)
		VALUES 
		($1,$2,$3,$4)
		RETURNING id
		`,
		*obj.UsrID, *obj.TournamentID, *obj.Active, *obj.Winner).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("query row context: %w", err)
	}

	return id, nil
}

func (u usrTournament) Update(ctx context.Context, obj *models.UsrTournamentCU, id int64) error {
	updateValues := []interface{}{id}
	queryUpdate := ` UPDATE usr_tournament u`
	querySet := ` SET id = id`

	queryWhere := ` WHERE id = $1`
	if obj.UsrID != nil {
		updateValues = append(updateValues, *obj.UsrID)
		querySet += ` , usr_id = $` + strconv.Itoa(len(updateValues))
	}
	if obj.TournamentID != nil {
		updateValues = append(updateValues, *obj.TournamentID)
		querySet += ` , tournament_id = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Active != nil {
		updateValues = append(updateValues, *obj.Active)
		querySet += ` , active = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Winner != nil {
		updateValues = append(updateValues, *obj.Winner)
		querySet += ` , winner = $` + strconv.Itoa(len(updateValues))
	}

	_, err := u.db.ExecContext(ctx, queryUpdate+querySet+queryWhere, updateValues...)
	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

func (u usrTournament) Get(ctx context.Context, pars *models.UsrTournamentGetPars) (*models.UsrTournament, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT u.id, u.usr_id, u.tournament_id, u.active, u.winner
`

	queryFrom := ` FROM usr_tournament u`
	queryWhere := ` WHERE 1 = 1`
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND u.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND u.usr_id = $` + strconv.Itoa(len(filterValues))
	}

	ue := &models.UsrTournament{}
	row := u.db.QueryRowContext(ctx, querySelect+queryFrom+queryWhere, filterValues...)
	err := row.Scan(
		&ue.ID,
		&ue.UsrID,
		&ue.TournamentID,
		&ue.Active,
		&ue.Winner,
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

	return ue, nil
}

func (u usrTournament) List(ctx context.Context, pars *models.UsrTournamentListPars) ([]*models.UsrTournament, int64, error) {
	var err error

	var filterValues []interface{}
	querySelect := `
	SELECT u.id, u.usr_id, u.tournament_id, u.active, u.winner
`
	queryFrom := ` FROM usr_tournament u`
	queryWhere := ` WHERE 1 = 1`
	queryOffset := ``
	queryLimit := ``
	queryOrderBy := " order by u.id desc"

	// filter
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND u.usr_id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND u.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.IDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.IDs))
		queryWhere += fmt.Sprintf(` AND u.id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.UsrIDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.UsrIDs))
		queryWhere += fmt.Sprintf(` AND u.usr_id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.Winner != nil {
		filterValues = append(filterValues, *pars.Winner)
		queryWhere += ` AND u.winner = $` + strconv.Itoa(len(filterValues))
	}
	if pars.Active != nil {
		filterValues = append(filterValues, *pars.Active)
		queryWhere += ` AND u.active = $` + strconv.Itoa(len(filterValues))
	}
	var tCount int64
	if pars.Limit > 0 {
		err := u.db.QueryRowContext(ctx, `select count(*)`+queryFrom+queryWhere, filterValues...).Scan(&tCount)
		if err != nil {
			return nil, 0, err
		}
		if pars.OnlyCount {
			return nil, tCount, nil
		}
		queryOffset = ` offset ` + strconv.FormatInt(pars.Offset, 10)
		queryLimit = ` limit ` + strconv.FormatInt(pars.Limit, 10)
	}

	rows, err := u.db.QueryContext(ctx, querySelect+queryFrom+queryWhere+queryOrderBy+queryOffset+queryLimit, filterValues...)
	if err != nil {
		return nil, 0, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var usrTournament []*models.UsrTournament
	for rows.Next() {
		ev := &models.UsrTournament{}
		err := rows.Scan(
			&ev.ID,
			&ev.UsrID,
			&ev.TournamentID,
			&ev.Active,
			&ev.Winner,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("rows scan: %w", err)
		}

		usrTournament = append(usrTournament, ev)
	}

	return usrTournament, tCount, nil
}
