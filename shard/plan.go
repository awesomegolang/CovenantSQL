/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package shard

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/CovenantSQL/go-sqlite3-encrypt"
	"github.com/CovenantSQL/sqlparser"
)

// Plan represents the execution strategy for a given query.
// For now it's a simple wrapper around the real instructions.
// An instruction (aka Primitive) is typically a tree where
// each node does its part by combining the results of the
// sub-nodes.
type Plan struct {
	// Original is the original query.
	Original string
	// Instructions contains the instructions needed to
	// fulfil the query.
	Instructions Primitive
}

// Primitive is the interface that needs to be satisfied by
// all primitives of a plan.
type Primitive interface {
	ExecContext(ctx context.Context) (result driver.Result, err error)
}

// BasePrimitive is the primitive just execute the origin query with origin args on rawConn
type BasePrimitive struct {
	// query is the original query.
	query string
	// args is the original args.
	args []driver.NamedValue
	// rawConn is the raw sqlite3 conn
	rawConn *sqlite3.SQLiteConn
	// rawDB is the userspace sql conn
	rawDB *sql.DB
}

func (dp *BasePrimitive) ExecContext(ctx context.Context) (result driver.Result, err error) {
	return dp.rawConn.ExecContext(ctx, dp.query, dp.args)
}

// BuildFromStmt builds a plan based on the AST provided.
func BuildFromStmt(query string, args []driver.NamedValue, stmt sqlparser.Statement, c *ShardingConn) (*Plan, error) {
	var err error
	plan := &Plan{}
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		return nil, errors.New("unsupported construct: select")
	case *sqlparser.Insert:
		plan.Instructions, err = buildInsertPlan(query, stmt, args, c)
	case *sqlparser.Update:
		return nil, errors.New("unsupported construct: update")
	case *sqlparser.Delete:
		return nil, errors.New("unsupported construct: delete")
	case *sqlparser.Union:
		return nil, errors.New("unsupported construct: union")
	case *sqlparser.Set:
		return nil, errors.New("unsupported construct: set")
	case *sqlparser.Show:
		return nil, errors.New("unsupported construct: show")
	case *sqlparser.DDL:
		return nil, errors.New("unsupported construct: ddl")
	case *sqlparser.DBDDL:
		return nil, errors.New("unsupported construct: ddl on database")
	default:
		panic(fmt.Sprintf("BUG: unexpected statement type: %T", stmt))
	}
	if err != nil {
		return nil, err
	}
	return plan, nil
}
