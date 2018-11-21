// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package provision

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	DAS_ORGANIZER_PROVISION                  = "DAS.ORGANIZER_PROVISION"
	DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID = "ORGANIZER_ID"
	DAS_ORGANIZER_PROVISION_COL_HOSTED       = "HOSTED"
	DAS_ORGANIZER_PROVISION_COL_AVAILABLE    = "AVAILABLE"
)

type PostgresOrganizerProvisionRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresOrganizerProvisionRepository) CreateOrganizerProvision(provision *businesslogic.OrganizerProvision) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_ORGANIZER_PROVISION).
		Columns(
			DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
			DAS_ORGANIZER_PROVISION_COL_HOSTED,
			DAS_ORGANIZER_PROVISION_COL_AVAILABLE,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated,
		).Values(provision.OrganizerID, provision.Hosted, provision.Available, provision.CreateUserID, provision.DateTimeCreated, provision.UpdateUserID, provision.DateTimeUpdated)
	_, err := stmt.RunWith(repo.Database).Exec()
	if err != nil {
		log.Printf("[error] initializing organizer organizer: %s\n", err.Error())
		return err
	}

	//CreateOrganizerProvisionHistoryEntry(accountID, 0, "initial organizer", accountID)
	if err != nil {
		log.Printf("[error] initializing organizer organizer history: %s\n", err.Error())
		return err
	}
	return err
}

// UpdateOrganizerProvision update the provision summary of an organizer. It does not update the provision history
// record of the organizer.
func (repo PostgresOrganizerProvisionRepository) UpdateOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").
		Table(DAS_ORGANIZER_PROVISION).
		Set(DAS_ORGANIZER_PROVISION_COL_AVAILABLE, provision.Available).
		Set(DAS_ORGANIZER_PROVISION_COL_HOSTED, provision.Hosted).
		Set(common.ColumnDateTimeUpdated, provision.DateTimeUpdated).
		Where(squirrel.Eq{DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID: provision.OrganizerID})
	_, err := stmt.RunWith(repo.Database).Exec()
	return err
}

// SearchOrganizerProvision get the provision information of an organizer user
func (repo PostgresOrganizerProvisionRepository) SearchOrganizerProvision(
	criteria businesslogic.SearchOrganizerProvisionCriteria) ([]businesslogic.OrganizerProvision, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}

	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
		DAS_ORGANIZER_PROVISION_COL_HOSTED,
		DAS_ORGANIZER_PROVISION_COL_AVAILABLE,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).
		From(DAS_ORGANIZER_PROVISION)
	if criteria.OrganizerID > 0. {
		stmt = stmt.Where(squirrel.Eq{DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID: criteria.OrganizerID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()

	provisions := make([]businesslogic.OrganizerProvision, 0)
	for rows.Next() {
		each := businesslogic.OrganizerProvision{}
		rows.Scan(
			&each.ID,
			&each.OrganizerID,
			&each.Hosted,
			&each.Available,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		selectAccount := repo.SqlBuilder.Select(
			fmt.Sprintf("%s, %s, %s",
				common.ColumnUID,
				"FIRST_NAME",
				"LAST_NAME")).From("DAS.ACCOUNT").Where(squirrel.Eq{common.ColumnPrimaryKey: each.ID})
		selectAccount.RunWith(repo.Database).QueryRow().Scan(&each.Organizer.UID, &each.Organizer.FirstName, &each.Organizer.LastName)
		provisions = append(provisions, each)
	}
	return provisions, err
}

func (repo PostgresOrganizerProvisionRepository) DeleteOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionRepository is not specified")
	}
	return errors.New("deleting organizer provision history is prohibited")
}
