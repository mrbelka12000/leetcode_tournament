package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type usr struct {
	db *sql.DB
}

func newUsr(db *sql.DB) *usr {
	return &usr{db: db}
}

func (u *usr) Create(ctx context.Context, obj *models.UsrCU) (int64, error) {
	var id int64

	err := u.db.QueryRowContext(ctx, `
		INSERT INTO usr
		(name, username, email, password, u_group, status_id, type_id)
		VALUES 
		($1,$2,$3,$4,$5,$6,$7)
		RETURNING id
		`,
		*obj.Name, *obj.Username, *obj.Email, *obj.Password, *obj.Group, *obj.StatusID, *obj.TypeID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("query row context: %w", err)
	}

	return id, nil
}

func (u *usr) Update(ctx context.Context, obj *models.UsrCU, id int64) error {
	updateValues := []interface{}{id}
	queryUpdate := ` UPDATE usr u`
	querySet := ` SET id = id`

	queryWhere := ` WHERE id = $1`
	if obj.Name != nil {
		updateValues = append(updateValues, *obj.Name)
		querySet += ` , name = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Password != nil {
		updateValues = append(updateValues, *obj.Password)
		querySet += ` , password = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Email != nil {
		updateValues = append(updateValues, *obj.Email)
		querySet += ` , email = $` + strconv.Itoa(len(updateValues))
	}
	if obj.Group != nil {
		updateValues = append(updateValues, *obj.Group)
		querySet += ` , u_group = $` + strconv.Itoa(len(updateValues))
	}

	if obj.Username != nil {
		updateValues = append(updateValues, *obj.Username)
		querySet += ` , username = $` + strconv.Itoa(len(updateValues))
	}
	if obj.StatusID != nil {
		updateValues = append(updateValues, *obj.StatusID)
		querySet += ` , status_id = $` + strconv.Itoa(len(updateValues))
	}
	if obj.TypeID != nil {
		updateValues = append(updateValues, *obj.TypeID)
		querySet += ` , type_id = $` + strconv.Itoa(len(updateValues))
	}

	_, err := u.db.ExecContext(ctx, queryUpdate+querySet+queryWhere, updateValues...)
	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

func (u *usr) Get(ctx context.Context, pars *models.UsrGetPars) (*models.Usr, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT u.id, u.name, u.username,u.email,u.password, u.u_group, u.status_id, u.type_id
`
	queryFrom := ` FROM usr u`
	queryWhere := ` WHERE 1 = 1`

	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND u.id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.Email != nil {
		filterValues = append(filterValues, *pars.Email)
		queryWhere += ` AND u.email = $` + strconv.Itoa(len(filterValues))
	}
	if pars.Username != nil {
		filterValues = append(filterValues, *pars.Username)
		queryWhere += ` AND u.username = $` + strconv.Itoa(len(filterValues))
	}
	if pars.UsernameEmail != nil {
		filterValues = append(filterValues, *pars.UsernameEmail)
		queryWhere += ` AND u.username = $` + strconv.Itoa(len(filterValues)) + ` OR u.email = $` + strconv.Itoa(len(filterValues))
	}
	if pars.StatusID != nil {
		filterValues = append(filterValues, *pars.StatusID)
		queryWhere += ` AND u.status_id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.TypeID != nil {
		filterValues = append(filterValues, *pars.TypeID)
		queryWhere += ` AND u.type_id = $` + strconv.Itoa(len(filterValues))
	}

	usr := &models.Usr{}
	row := u.db.QueryRowContext(ctx, querySelect+queryFrom+queryWhere, filterValues...)
	err := row.Scan(
		&usr.ID,
		&usr.Name,
		&usr.Username,
		&usr.Email,
		&usr.Password,
		&usr.Group,
		&usr.StatusID,
		&usr.TypeID,
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

	return usr, nil
}

func (u *usr) List(ctx context.Context, pars *models.UsrListPars) ([]*models.Usr, int64, error) {
	var err error

	var filterValues []interface{}
	querySelect := `
	SELECT u.id, u.name, u.username,u.email,u.password, u.u_group, u.status_id, u.type_id
`
	queryFrom := ` FROM usr u`
	queryWhere := ` WHERE 1 = 1`
	queryOffset := ``
	queryLimit := ``
	queryOrderBy := " order by u.id desc"

	// filter
	if pars.IDs != nil {
		filterValues = append(filterValues, *pars.IDs)
		queryWhere += ` AND u.id (select * from unnest($ :: bigint[]))` + strconv.Itoa(len(filterValues))
	}
	if pars.TypeID != nil {
		filterValues = append(filterValues, *pars.TypeID)
		queryWhere += ` AND u.type_id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.TypeIDs != nil {
		filterValues = append(filterValues, *pars.TypeID)
		queryWhere += ` AND u.type_id (select * from unnest($ :: smallint[]))` + strconv.Itoa(len(filterValues))
	}
	if pars.StatusID != nil {
		filterValues = append(filterValues, *pars.StatusID)
		queryWhere += ` AND u.status_id = $` + strconv.Itoa(len(filterValues))
	}
	if pars.UsernameEmail != nil {
		filterValues = append(filterValues, *pars.UsernameEmail)
		queryWhere += ` AND u.username = $` + strconv.Itoa(len(filterValues)) + ` OR u.email = $` + strconv.Itoa(len(filterValues))
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

	var users []*models.Usr
	for rows.Next() {
		usr := &models.Usr{}
		err := rows.Scan(
			&usr.ID,
			&usr.Name,
			&usr.Username,
			&usr.Email,
			&usr.Password,
			&usr.Group,
			&usr.StatusID,
			&usr.TypeID,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("rows scan: %w", err)
		}

		users = append(users, usr)
	}

	return users, tCount, nil
}
