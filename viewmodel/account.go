// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"time"
)

type AccountType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AccountTypeDataModelToViewModel(dm businesslogic.AccountType) AccountType {
	return AccountType{
		ID:   dm.ID,
		Name: dm.Name,
	}
}

type CreateAccount struct {
	AccountType int       `json:"accounttype"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	FirstName   string    `json:"firstname"`
	MiddleNames string    `json:"middlenames"`
	LastName    string    `json:"lastname"`
	DateOfBirth time.Time `json:"dateofbirth"`
	Gender      int       `json:"gender"`
	Password    string    `json:"password"`
	ToSAccepted bool      `json:"tosaccepted"`
	PPAccepted  bool      `json:"ppaccepted"`
	ByGuardian  bool      `json:"byguardian"`
	Signature   string    `json:"signature"`
}

type Login struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}
