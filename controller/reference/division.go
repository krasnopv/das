// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type DivisionServer struct {
	referencebll.IDivisionRepository
}

func (server DivisionServer) SearchDivisionHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchDivisionCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if divisions, err := server.IDivisionRepository.SearchDivision(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	} else {
		data := make([]viewmodel.DivisionViewModel, 0)
		for _, each := range divisions {
			view := viewmodel.DivisionViewModel{
				ID:         each.ID,
				Name:       each.Name,
				Federation: each.FederationID,
			}
			data = append(data, view)
		}
		output, _ := json.Marshal(data)
		w.Write(output)
	}

}

func (server DivisionServer) CreateDivisionHandler(w http.ResponseWriter, r *http.Request) {}
func (server DivisionServer) UpdateDivisionHandler(w http.ResponseWriter, r *http.Request) {}
func (server DivisionServer) DeleteDivisionHandler(w http.ResponseWriter, r *http.Request) {}
