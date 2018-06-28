// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/partnership/request"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

var partnershipRequestStatusServer = request.PartnershipRequestStatusServer{
	database.PartnershipRequestStatusRepository,
}

var PartnershipRequestStatusController = util.DasController{
	Name:         "PartnershipRequestStatusController",
	Description:  "Search partnership request status in DAS",
	Method:       http.MethodGet,
	Endpoint:     "/api/partnership/request/status",
	Handler:      partnershipRequestStatusServer.GetPartnershipRequestStatusHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
