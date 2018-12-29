/*
 *  Copyright 2018 The CovenantSQL Authors.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"github.com/CovenantSQL/CovenantSQL/proto"
)

// PermStat defines the permissions status structure.
type PermStat struct {
	Permission UserPermission
	Status     Status
}

// UserState defines the user state structure.
type UserState struct {
	State map[proto.AccountAddress]*PermStat
}

// NewUserState returns new user state instance.
func NewUserState() *UserState {
	return &UserState{
		State: make(map[proto.AccountAddress]*PermStat),
	}
}

// UpdatePermission sets account permission.
func (us *UserState) UpdatePermission(user proto.AccountAddress, perm UserPermission) {
	if state, ok := us.State[user]; ok {
		state.Permission = perm
	} else {
		us.State[user] = &PermStat{
			Permission: perm,
		}
	}
}

// AddPermission updates account permission.
func (us *UserState) AddPermission(user proto.AccountAddress, perm UserPermission) {
	us.UpdatePermission(user, perm)
}

// UpdateStatus update account status.
func (us *UserState) UpdateStatus(user proto.AccountAddress, stat Status) {
	if state, ok := us.State[user]; ok {
		state.Status = stat
	} else {
		us.State[user] = &PermStat{
			Status: stat,
		}
	}
}

// AddStatus is an alias for UpdateStatus.
func (us *UserState) AddStatus(user proto.AccountAddress, stat Status) {
	us.UpdateStatus(user, stat)
}

// GetPermission returns the account permission.
func (us *UserState) GetPermission(user proto.AccountAddress) (up UserPermission, ok bool) {
	permstat, ok := us.State[user]
	up = permstat.Permission
	return
}

// GetStatus returns the account status.
func (us *UserState) GetStatus(user proto.AccountAddress) (stat Status, ok bool) {
	permstat, ok := us.State[user]
	stat = permstat.Status
	return
}