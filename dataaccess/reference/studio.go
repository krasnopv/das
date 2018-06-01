package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_STUDIO_TABLE = "DAS.STUDIO"
)

type PostgresStudioRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresStudioRepository) SearchStudio(criteria *reference.SearchStudioCriteria) ([]reference.Studio, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf(`DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, 
		DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s`,
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ADDRESS,
			common.COL_CITY_ID,
			common.COL_WEBSITE,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_STUDIO_TABLE).OrderBy(common.PRIMARY_KEY)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.ID})
	}
	if criteria.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_CITY_ID: criteria.CityID})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Join(`DAS.CITY C ON C.ID = DAS.STUDIO.CITY_ID`).
			Join(`DAS.STATE S ON S.ID = C.STATE_ID`).Where(squirrel.Eq{`S.ID`: criteria.StateID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	studios := make([]reference.Studio, 0)
	if err != nil {
		return studios, err
	}

	for rows.Next() {
		each := reference.Studio{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Address,
			&each.CityID,
			&each.Website,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		studios = append(studios, each)
	}
	rows.Close()
	return studios, err
}

func (repo PostgresStudioRepository) CreateStudio(studio *reference.Studio) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_STUDIO_TABLE).Columns(
		common.COL_NAME,
		common.COL_ADDRESS,
		common.COL_CITY_ID,
		common.COL_WEBSITE,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		studio.Name,
		studio.Address,
		studio.CityID,
		studio.Website,
		studio.CreateUserID,
		studio.DateTimeCreated,
		studio.UpdateUserID,
		studio.DateTimeUpdated,
	).Suffix(
		fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY),
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&studio.ID)
		if err = tx.Commit(); err != nil {
			tx.Rollback()
		}
	}
	return err
}

func (repo PostgresStudioRepository) UpdateStudio(studio reference.Studio) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_STUDIO_TABLE)
	if studio.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, studio.Name).
			Set(common.COL_ADDRESS, studio.Address).
			Set(common.COL_CITY_ID, studio.CityID).
			Set(common.COL_WEBSITE, studio.Website).
			Set(common.COL_UPDATE_USER_ID, studio.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, studio.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	} else {
		return errors.New("studio is not specified")
	}
}

func (repo PostgresStudioRepository) DeleteStudio(studio reference.Studio) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_STUDIO_TABLE).Where(squirrel.Eq{common.PRIMARY_KEY: studio.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}
