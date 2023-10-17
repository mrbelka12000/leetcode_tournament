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

type event struct {
	db *sql.DB
}

func newEvent(db *sql.DB) *event {
	return &event{db: db}
}

func (e *event) Create(ctx context.Context, obj *models.EventCU) (int64, error) {
	var id int64

	err := e.db.QueryRowContext(ctx, `
		INSERT INTO event
		(usr_id, start_time, end_time, goal, condition, status_id)
		VALUES 
		($1,$2,$3,$4,$5,$6)
		RETURNING id
		`,
		*obj.UsrID, *obj.StartTime, *obj.EndTime, *obj.Goal, *obj.Condition, *obj.StatusID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("query row context: %w", err)
	}

	return id, nil
}

func (e *event) Update(ctx context.Context, obj *models.EventCU, id int64) error {
	updateValues := []interface{}{id}
	queryUpdate := ` UPDATE event e`
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
	if obj.Condition != nil {
		updateValues = append(updateValues, *obj.Condition)
		querySet += ` , condition = $` + strconv.Itoa(len(updateValues))
	}
	if obj.StatusID != nil {
		updateValues = append(updateValues, *obj.StatusID)
		querySet += ` , status_id = $` + strconv.Itoa(len(updateValues))
	}

	_, err := e.db.ExecContext(ctx, queryUpdate+querySet+queryWhere, updateValues...)
	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

func (e *event) Get(ctx context.Context, pars *models.EventGetPars) (*models.Event, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT e.id, e.usr_id, e.start_time, e.end_time, e.goal, e.condition, e.status_id
`

	queryFrom := ` FROM event e`
	queryWhere := ` WHERE 1 = 1`
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND e.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND e.usr_id = $` + strconv.Itoa(len(filterValues))
	}

	ev := &models.Event{}
	row := e.db.QueryRowContext(ctx, querySelect+queryFrom+queryWhere, filterValues...)
	err := row.Scan(
		&ev.ID,
		&ev.UsrID,
		&ev.StartTime,
		&ev.EndTime,
		&ev.Goal,
		&ev.Condition,
		&ev.StatusID,
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

	return ev, nil
}

func (e *event) List(ctx context.Context, pars *models.EventListPars) ([]*models.Event, int64, error) {
	var err error

	var filterValues []interface{}
	querySelect := `
	SELECT e.id, e.usr_id, e.start_time, e.end_time, e.goal, e.condition, e.status_id
`
	queryFrom := ` FROM event e`
	queryWhere := ` WHERE 1 = 1`
	queryOffset := ``
	queryLimit := ``
	queryOrderBy := " order by e.id desc"

	// filter
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND e.usr_id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND e.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.IDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.IDs))
		queryWhere += fmt.Sprintf(` AND e.id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.UsrIDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.UsrIDs))
		queryWhere += fmt.Sprintf(` AND e.usr_id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.StatusIDs != nil {
		filterValues = append(filterValues, pq.Array(*pars.StatusIDs))
		queryWhere += fmt.Sprintf(` AND e.status_id in (select * from unnest($%v :: bigint[]))`, strconv.Itoa(len(filterValues)))
	}
	if pars.Condition != nil {
		filterValues = append(filterValues, *pars.Condition)
		queryWhere += ` AND e.condition = $` + strconv.Itoa(len(filterValues))
	}

	var tCount int64
	if pars.Limit > 0 {
		err := e.db.QueryRowContext(ctx, `select count(*)`+queryFrom+queryWhere, filterValues...).Scan(&tCount)
		if err != nil {
			return nil, 0, err
		}
		if pars.OnlyCount {
			return nil, tCount, nil
		}
		queryOffset = ` offset ` + strconv.FormatInt(pars.Offset, 10)
		queryLimit = ` limit ` + strconv.FormatInt(pars.Limit, 10)
	}

	rows, err := e.db.QueryContext(ctx, querySelect+queryFrom+queryWhere+queryOrderBy+queryOffset+queryLimit, filterValues...)
	if err != nil {
		return nil, 0, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var events []*models.Event
	for rows.Next() {
		ev := &models.Event{}
		err := rows.Scan(
			&ev.ID,
			&ev.UsrID,
			&ev.StartTime,
			&ev.EndTime,
			&ev.Goal,
			&ev.Condition,
			&ev.StatusID,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("rows scan: %w", err)
		}

		events = append(events, ev)
	}

	return events, tCount, nil
}
