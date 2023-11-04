package repo

import (
	"context"
	"database/sql"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	eventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/event"
	scoreservice "github.com/mrbelka12000/leetcode_tournament/internal/service/score"
	sessionservice "github.com/mrbelka12000/leetcode_tournament/internal/service/session"
	tournamentservice "github.com/mrbelka12000/leetcode_tournament/internal/service/tournament"
	txservice "github.com/mrbelka12000/leetcode_tournament/internal/service/tx"
	usrservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usr"
	usreventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usrevent"
	usrtournamentservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usrtournament"
)

type conSt interface {
	ExecContext(ctx context.Context, sql string, arguments ...any) (sql.Result, error)
	QueryContext(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, sql string, args ...any) *sql.Row
}

type Repo struct {
	Usr           usrservice.Repo
	Score         scoreservice.Repo
	Event         eventservice.Repo
	Tournament    tournamentservice.Repo
	UsrEvent      usreventservice.Repo
	UsrTournament usrtournamentservice.Repo
	Session       sessionservice.Repo
	Tx            txservice.Repo
}

func New(db *sql.DB) *Repo {
	return &Repo{
		Usr:           newUsr(db),
		Score:         newScore(db),
		Event:         newEvent(db),
		Tournament:    newTournament(db),
		UsrEvent:      newUsrEvent(db),
		UsrTournament: newUsrTournament(db),
		Session:       newSession(db),
		Tx:            newTx(db),
	}
}

func Exec(ctx context.Context, db *sql.DB, query string, args ...any) (sql.Result, error) {
	return coalesceConn(ctx, db).ExecContext(ctx, query, args...)
}

func Query(ctx context.Context, db *sql.DB, query string, args ...any) (*sql.Rows, error) {
	return coalesceConn(ctx, db).QueryContext(ctx, query, args...)
}

func QueryRow(ctx context.Context, db *sql.DB, query string, args ...any) *sql.Row {
	return coalesceConn(ctx, db).QueryRowContext(ctx, query, args...)
}

func getContextTransactionContainer(ctx context.Context) *txContainerSt {
	contextV := ctx.Value(consts.TransactionCtxKey)
	if contextV == nil {
		return nil
	}

	switch tx := contextV.(type) {
	case *txContainerSt:
		return tx
	default:
		return nil
	}
}

func getContextTransaction(ctx context.Context) *sql.Tx {
	container := getContextTransactionContainer(ctx)
	if container != nil {
		return container.tx
	}

	return nil
}

func coalesceConn(ctx context.Context, db *sql.DB) conSt {
	if tx := getContextTransaction(ctx); tx != nil {
		return tx
	}
	return db
}
