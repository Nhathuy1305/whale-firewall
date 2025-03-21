// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addContainerStmt, err = db.PrepareContext(ctx, addContainer); err != nil {
		return nil, fmt.Errorf("error preparing query AddContainer: %w", err)
	}
	if q.addContainerAddrStmt, err = db.PrepareContext(ctx, addContainerAddr); err != nil {
		return nil, fmt.Errorf("error preparing query AddContainerAddr: %w", err)
	}
	if q.addContainerAliasStmt, err = db.PrepareContext(ctx, addContainerAlias); err != nil {
		return nil, fmt.Errorf("error preparing query AddContainerAlias: %w", err)
	}
	if q.addEstContainerStmt, err = db.PrepareContext(ctx, addEstContainer); err != nil {
		return nil, fmt.Errorf("error preparing query AddEstContainer: %w", err)
	}
	if q.addWaitingContainerRuleStmt, err = db.PrepareContext(ctx, addWaitingContainerRule); err != nil {
		return nil, fmt.Errorf("error preparing query AddWaitingContainerRule: %w", err)
	}
	if q.containerExistsStmt, err = db.PrepareContext(ctx, containerExists); err != nil {
		return nil, fmt.Errorf("error preparing query ContainerExists: %w", err)
	}
	if q.deleteContainerStmt, err = db.PrepareContext(ctx, deleteContainer); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteContainer: %w", err)
	}
	if q.deleteContainerAddrsStmt, err = db.PrepareContext(ctx, deleteContainerAddrs); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteContainerAddrs: %w", err)
	}
	if q.deleteContainerAliasesStmt, err = db.PrepareContext(ctx, deleteContainerAliases); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteContainerAliases: %w", err)
	}
	if q.deleteEstContainersStmt, err = db.PrepareContext(ctx, deleteEstContainers); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEstContainers: %w", err)
	}
	if q.deleteWaitingContainerRulesStmt, err = db.PrepareContext(ctx, deleteWaitingContainerRules); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteWaitingContainerRules: %w", err)
	}
	if q.getContainerAddrsStmt, err = db.PrepareContext(ctx, getContainerAddrs); err != nil {
		return nil, fmt.Errorf("error preparing query GetContainerAddrs: %w", err)
	}
	if q.getContainerIDStmt, err = db.PrepareContext(ctx, getContainerID); err != nil {
		return nil, fmt.Errorf("error preparing query GetContainerID: %w", err)
	}
	if q.getContainerIDAndNameFromAliasStmt, err = db.PrepareContext(ctx, getContainerIDAndNameFromAlias); err != nil {
		return nil, fmt.Errorf("error preparing query GetContainerIDAndNameFromAlias: %w", err)
	}
	if q.getContainerNameStmt, err = db.PrepareContext(ctx, getContainerName); err != nil {
		return nil, fmt.Errorf("error preparing query GetContainerName: %w", err)
	}
	if q.getContainersStmt, err = db.PrepareContext(ctx, getContainers); err != nil {
		return nil, fmt.Errorf("error preparing query GetContainers: %w", err)
	}
	if q.getEstContainersStmt, err = db.PrepareContext(ctx, getEstContainers); err != nil {
		return nil, fmt.Errorf("error preparing query GetEstContainers: %w", err)
	}
	if q.getWaitingContainerRulesStmt, err = db.PrepareContext(ctx, getWaitingContainerRules); err != nil {
		return nil, fmt.Errorf("error preparing query GetWaitingContainerRules: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addContainerStmt != nil {
		if cerr := q.addContainerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addContainerStmt: %w", cerr)
		}
	}
	if q.addContainerAddrStmt != nil {
		if cerr := q.addContainerAddrStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addContainerAddrStmt: %w", cerr)
		}
	}
	if q.addContainerAliasStmt != nil {
		if cerr := q.addContainerAliasStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addContainerAliasStmt: %w", cerr)
		}
	}
	if q.addEstContainerStmt != nil {
		if cerr := q.addEstContainerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addEstContainerStmt: %w", cerr)
		}
	}
	if q.addWaitingContainerRuleStmt != nil {
		if cerr := q.addWaitingContainerRuleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addWaitingContainerRuleStmt: %w", cerr)
		}
	}
	if q.containerExistsStmt != nil {
		if cerr := q.containerExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing containerExistsStmt: %w", cerr)
		}
	}
	if q.deleteContainerStmt != nil {
		if cerr := q.deleteContainerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteContainerStmt: %w", cerr)
		}
	}
	if q.deleteContainerAddrsStmt != nil {
		if cerr := q.deleteContainerAddrsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteContainerAddrsStmt: %w", cerr)
		}
	}
	if q.deleteContainerAliasesStmt != nil {
		if cerr := q.deleteContainerAliasesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteContainerAliasesStmt: %w", cerr)
		}
	}
	if q.deleteEstContainersStmt != nil {
		if cerr := q.deleteEstContainersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEstContainersStmt: %w", cerr)
		}
	}
	if q.deleteWaitingContainerRulesStmt != nil {
		if cerr := q.deleteWaitingContainerRulesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteWaitingContainerRulesStmt: %w", cerr)
		}
	}
	if q.getContainerAddrsStmt != nil {
		if cerr := q.getContainerAddrsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getContainerAddrsStmt: %w", cerr)
		}
	}
	if q.getContainerIDStmt != nil {
		if cerr := q.getContainerIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getContainerIDStmt: %w", cerr)
		}
	}
	if q.getContainerIDAndNameFromAliasStmt != nil {
		if cerr := q.getContainerIDAndNameFromAliasStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getContainerIDAndNameFromAliasStmt: %w", cerr)
		}
	}
	if q.getContainerNameStmt != nil {
		if cerr := q.getContainerNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getContainerNameStmt: %w", cerr)
		}
	}
	if q.getContainersStmt != nil {
		if cerr := q.getContainersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getContainersStmt: %w", cerr)
		}
	}
	if q.getEstContainersStmt != nil {
		if cerr := q.getEstContainersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEstContainersStmt: %w", cerr)
		}
	}
	if q.getWaitingContainerRulesStmt != nil {
		if cerr := q.getWaitingContainerRulesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getWaitingContainerRulesStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                 DBTX
	tx                                 *sql.Tx
	addContainerStmt                   *sql.Stmt
	addContainerAddrStmt               *sql.Stmt
	addContainerAliasStmt              *sql.Stmt
	addEstContainerStmt                *sql.Stmt
	addWaitingContainerRuleStmt        *sql.Stmt
	containerExistsStmt                *sql.Stmt
	deleteContainerStmt                *sql.Stmt
	deleteContainerAddrsStmt           *sql.Stmt
	deleteContainerAliasesStmt         *sql.Stmt
	deleteEstContainersStmt            *sql.Stmt
	deleteWaitingContainerRulesStmt    *sql.Stmt
	getContainerAddrsStmt              *sql.Stmt
	getContainerIDStmt                 *sql.Stmt
	getContainerIDAndNameFromAliasStmt *sql.Stmt
	getContainerNameStmt               *sql.Stmt
	getContainersStmt                  *sql.Stmt
	getEstContainersStmt               *sql.Stmt
	getWaitingContainerRulesStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		addContainerStmt:                   q.addContainerStmt,
		addContainerAddrStmt:               q.addContainerAddrStmt,
		addContainerAliasStmt:              q.addContainerAliasStmt,
		addEstContainerStmt:                q.addEstContainerStmt,
		addWaitingContainerRuleStmt:        q.addWaitingContainerRuleStmt,
		containerExistsStmt:                q.containerExistsStmt,
		deleteContainerStmt:                q.deleteContainerStmt,
		deleteContainerAddrsStmt:           q.deleteContainerAddrsStmt,
		deleteContainerAliasesStmt:         q.deleteContainerAliasesStmt,
		deleteEstContainersStmt:            q.deleteEstContainersStmt,
		deleteWaitingContainerRulesStmt:    q.deleteWaitingContainerRulesStmt,
		getContainerAddrsStmt:              q.getContainerAddrsStmt,
		getContainerIDStmt:                 q.getContainerIDStmt,
		getContainerIDAndNameFromAliasStmt: q.getContainerIDAndNameFromAliasStmt,
		getContainerNameStmt:               q.getContainerNameStmt,
		getContainersStmt:                  q.getContainersStmt,
		getEstContainersStmt:               q.getEstContainersStmt,
		getWaitingContainerRulesStmt:       q.getWaitingContainerRulesStmt,
	}
}
