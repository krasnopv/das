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

package businesslogic

import (
	"errors"
	"log"
	"time"
)

const (
	// RoleApplicationStatusApproved marks the role application as "approved"
	RoleApplicationStatusApproved = 1
	// RoleApplicationStatusDenied marks the role application as "denied"
	RoleApplicationStatusDenied = 2
	// RoleApplicationStatusPending marks the role application as "pending"
	RoleApplicationStatusPending = 3
)

// SearchRoleApplicationCriteria specifies the search criteria for role application
type SearchRoleApplicationCriteria struct {
	ID             int
	AccountID      int
	AppliedRoleID  int
	StatusID       int
	ApprovalUserID int
	Responded      bool
}

// RoleApplication is an application for restricted roles, including adjudicator, scrutineer, and organizer.
// Non-restrictive roles such as emcee and deck captain can be approved by competition organizers
type RoleApplication struct {
	ID               int
	AccountID        int
	Account          Account
	AppliedRoleID    int
	Description      string
	StatusID         int
	ApprovalUserID   *int
	DateTimeApproved time.Time
	CreateUserID     int
	DateTimeCreated  time.Time
	UpdateUserID     int
	DateTimeUpdated  time.Time
}

// IRoleApplicationRepository specifies the interface that a Role Application Repository should implement
type IRoleApplicationRepository interface {
	CreateApplication(application *RoleApplication) error
	SearchApplication(criteria SearchRoleApplicationCriteria) ([]RoleApplication, error)
	UpdateApplication(application RoleApplication) error
}

// RoleProvisionService is a service that handles Role Application and provision
type RoleProvisionService struct {
	accountRepo                   IAccountRepository
	roleApplicationRepo           IRoleApplicationRepository
	roleRepo                      IAccountRoleRepository
	organizerProvisionRepo        IOrganizerProvisionRepository
	organizerProvisionHistoryRepo IOrganizerProvisionHistoryRepository
}

// NewRoleProvisionService create a service that serves Role Provision
func NewRoleProvisionService(
	accountRepo IAccountRepository,
	roleApplicationRepo IRoleApplicationRepository,
	roleRepo IAccountRoleRepository,
	organizerProvisionRepo IOrganizerProvisionRepository,
	organizerProvisionHistoryRepo IOrganizerProvisionHistoryRepository) *RoleProvisionService {
	service := RoleProvisionService{
		accountRepo:                   accountRepo,
		roleApplicationRepo:           roleApplicationRepo,
		roleRepo:                      roleRepo,
		organizerProvisionRepo:        organizerProvisionRepo,
		organizerProvisionHistoryRepo: organizerProvisionHistoryRepo,
	}
	return &service
}

// CreateRoleApplication check the validity of the role application and create it if it's valid
func (service RoleProvisionService) CreateRoleApplication(currentUser Account, application *RoleApplication) error {
	// check if current user has the role
	if currentUser.HasRole(application.AppliedRoleID) {
		return errors.New("current user already has the applied role")
	}

	// check if has a pending application
	searchResults, err := service.roleApplicationRepo.SearchApplication(SearchRoleApplicationCriteria{
		AccountID:     currentUser.ID,
		AppliedRoleID: application.AppliedRoleID,
		StatusID:      RoleApplicationStatusPending,
	})
	if err != nil {
		return err
	}
	if len(searchResults) != 0 {
		return errors.New("previous application has not been responded")
	}

	// check what role that user is applying for
	if application.AppliedRoleID == AccountTypeAthlete {
		return errors.New("athlete role should be granted when the account was created")
	}
	if application.AppliedRoleID > AccountTypeEmcee {
		return errors.New("invalid role")
	}

	return service.roleApplicationRepo.CreateApplication(application)
}

func (service RoleProvisionService) respondRoleApplication(currentUser Account, application *RoleApplication, action int) error {
	application.StatusID = action
	application.ApprovalUserID = &currentUser.ID
	application.DateTimeApproved = time.Now()
	if updateErr := service.roleApplicationRepo.UpdateApplication(*application); updateErr != nil {
		return updateErr
	}
	if action == RoleApplicationStatusApproved {
		role := AccountRole{
			AccountID:       application.AccountID,
			AccountTypeID:   application.AppliedRoleID,
			CreateUserID:    currentUser.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    currentUser.ID,
			DateTimeUpdated: time.Now(),
		}
		return service.roleRepo.CreateAccountRole(&role)
	}
	return nil
}

// UpdateApplication attempts to approve the Role application based on the privilege of current user.
// If current user is admin, any application can be approved
// If current user is organizer, only emcee and deck-captain can be approved
// If current user is other roles, current user will be prohibited from performing such action
func (service RoleProvisionService) UpdateApplication(currentUser Account, application *RoleApplication, action int) error {
	// check if action is valid
	if !(action == RoleApplicationStatusApproved || action == RoleApplicationStatusDenied) {
		return errors.New("invalid response to role application")
	}
	// check if application is pending
	if application.StatusID == RoleApplicationStatusApproved || application.StatusID == RoleApplicationStatusDenied {
		return errors.New("role application is already responded")
	}
	// Only an Admin or Organizer user ca update user's role application
	if !(currentUser.HasRole(AccountTypeOrganizer) || currentUser.HasRole(AccountTypeAdministrator)) {
		return errors.New("unauthorized")
	}
	// should not allow users to provision themselves other than Admin
	if currentUser.ID == application.AccountID && !currentUser.HasRole(AccountTypeAdministrator) {
		return errors.New("not authorized to provision your own role application")
	}
	switch application.AppliedRoleID {
	case AccountTypeAthlete:
		return nil // Athlete role does not need to be provisioned
	case AccountTypeAdjudicator:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeScrutineer:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeOrganizer:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeDeckCaptain:
		if !(currentUser.HasRole(AccountTypeAdministrator) || currentUser.HasRole(AccountTypeOrganizer)) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeEmcee:
		if !(currentUser.HasRole(AccountTypeAdministrator) || currentUser.HasRole(AccountTypeOrganizer)) {
			return errors.New("not authorized to approve user's role application")
		}
	default:
		return errors.New("invalid role application")
	}
	roleProvisionErr := service.respondRoleApplication(currentUser, application, action)
	if roleProvisionErr != nil {
		return roleProvisionErr
	}

	if application.AppliedRoleID == AccountTypeOrganizer {
		roleSearch, roleSearchErr := service.roleRepo.SearchAccountRole(SearchAccountRoleCriteria{
			AccountID:     application.AccountID,
			AccountTypeID: AccountTypeOrganizer,
		})
		if roleSearchErr != nil || len(roleSearch) != 1 {
			if roleSearchErr != nil {
				log.Println(roleSearchErr)
			}
			return errors.New("cannot find Organizer role of this account")
		}
		role := roleSearch[0]

		// create organizer provision
		orgProvision, orgProvisionHist := initializeOrganizerProvision(role.ID, currentUser.ID)
		if orgProvErr := service.organizerProvisionRepo.CreateOrganizerProvision(&orgProvision); orgProvErr != nil {
			return orgProvErr
		}
		if orgProvHistErr := service.organizerProvisionHistoryRepo.CreateOrganizerProvisionHistory(&orgProvisionHist); orgProvHistErr != nil {
			return orgProvHistErr
		}
	}
	return nil
}

// SearchRoleApplication searches the available role application based on current user's privilege
func (service RoleProvisionService) SearchRoleApplication(currentUser Account, criteria SearchRoleApplicationCriteria) ([]RoleApplication, error) {
	return service.roleApplicationRepo.SearchApplication(criteria)
}

// OrganizerProvision provision organizer competition slots for creating and hosting competitions
type OrganizerProvision struct {
	ID              int
	AccountID       int
	OrganizerID     int
	Organizer       Account
	Available       int
	Hosted          int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type UpdateOrganizerProvision struct {
	OrganizerID   int
	Amount        int
	Note          string
	CurrentUserID int
}

// OrganizerProvisionServices provides functions that allows provisioning Organizer's Competition, including updating
// and querying Organizer's Competition Provision.
type OrganizerProvisionService struct {
	accountRepo                   IAccountRepository
	organizerProvisionRepo        IOrganizerProvisionRepository
	organizerProvisionHistoryRepo IOrganizerProvisionHistoryRepository
}

func NewOrganizerProvisionService(accountRepo IAccountRepository, organizerProvisionRepo IOrganizerProvisionRepository, organizerProvisionHistoryRepo IOrganizerProvisionHistoryRepository) OrganizerProvisionService {
	return OrganizerProvisionService{
		accountRepo:                   accountRepo,
		organizerProvisionRepo:        organizerProvisionRepo,
		organizerProvisionHistoryRepo: organizerProvisionHistoryRepo,
	}
}

func (service OrganizerProvisionService) SearchOrganizerProvision(criteria SearchOrganizerProvisionCriteria) ([]OrganizerProvision, error) {
	return service.organizerProvisionRepo.SearchOrganizerProvision(criteria)
}

func (service OrganizerProvisionService) UpdateOrganizerCompetitionProvision(update UpdateOrganizerProvision) error {
	provisions, searchErr := service.organizerProvisionRepo.SearchOrganizerProvision(SearchOrganizerProvisionCriteria{OrganizerID: update.OrganizerID})
	if searchErr != nil {
		return searchErr
	}
	if len(provisions) != 1 {
		return errors.New("cannot find organizer's competition provision information")
	}

	provision := provisions[0]
	provision.Available = provision.Available + update.Amount
	provision.UpdateUserID = update.CurrentUserID
	provision.DateTimeUpdated = time.Now()

	history := OrganizerProvisionHistoryEntry{
		OrganizerID:     update.OrganizerID,
		Amount:          update.Amount,
		Note:            update.Note,
		CreateUserID:    update.CurrentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    update.CurrentUserID,
		DateTimeUpdated: time.Now(),
	}

	updateErr := service.organizerProvisionRepo.UpdateOrganizerProvision(provision)
	if updateErr != nil {
		return updateErr
	}
	createErr := service.organizerProvisionHistoryRepo.CreateOrganizerProvisionHistory(&history)
	if createErr != nil {
		return createErr
	}
	return nil
}

// SearchOrganizerProvisionCriteria specifies the search criteria of Organizer's provision information
type SearchOrganizerProvisionCriteria struct {
	ID           int    `schema:"id"`
	OrganizerID  int    `schema:"organizerID"`  // organizer's account ID, not type-account id
	OrganizerUID string `schema:"organizerUID"` // Organizer's UID,
}

// IOrganizerProvisionRepository specifies the interface that a repository should implement for Organizer Provision
type IOrganizerProvisionRepository interface {
	CreateOrganizerProvision(provision *OrganizerProvision) error
	UpdateOrganizerProvision(provision OrganizerProvision) error
	DeleteOrganizerProvision(provision OrganizerProvision) error
	SearchOrganizerProvision(criteria SearchOrganizerProvisionCriteria) ([]OrganizerProvision, error)
}

func (provision OrganizerProvision) updateForCreateCompetition(competition Competition) OrganizerProvision {
	newProvision := provision
	newProvision.Available = provision.Available - 1
	newProvision.Hosted = provision.Hosted + 1
	newProvision.UpdateUserID = competition.CreateUserID
	newProvision.DateTimeUpdated = time.Now()
	return newProvision
}

func initializeOrganizerProvision(accountRoleID, currentUserID int) (OrganizerProvision, OrganizerProvisionHistoryEntry) {
	provision := OrganizerProvision{
		OrganizerID:     accountRoleID,
		Available:       0,
		CreateUserID:    currentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUserID,
		DateTimeUpdated: time.Now(),
	}
	history := OrganizerProvisionHistoryEntry{
		OrganizerID:     accountRoleID,
		Amount:          0,
		Note:            "initialize organizer organizer",
		CreateUserID:    currentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUserID,
		DateTimeUpdated: time.Now(),
	}
	return provision, history
}

func updateOrganizerProvision(provision OrganizerProvision, history OrganizerProvisionHistoryEntry,
	organizerRepository IOrganizerProvisionRepository, historyRepository IOrganizerProvisionHistoryRepository) {
	historyRepository.CreateOrganizerProvisionHistory(&history)
	organizerRepository.UpdateOrganizerProvision(provision)
}
