package controller

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type EntryServer struct {
	Service businesslogic.CompetitionRegistrationService
}

// SearchCompetitionEntryHandler handles the request
//	GET /api/v1.0/competition/entries
// Public view for competitive event entry
func (server EntryServer) SearchCompetitionEntryHandler(w http.ResponseWriter, r *http.Request) {
	form := new(viewmodel.SearchEntryForm)
	if parseErr := util.ParseRequestData(r, form); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	criteria := businesslogic.SearchEntryCriteria{
		CompetitionID: form.CompetitionID,
		EventID:       form.EventID,
		FederationID:  form.FederationID,
		DivisionID:    form.DivisionID,
		ProficiencyID: form.ProficiencyID,
		StyleID:       form.StyleID,
		AthleteID:     form.AthleteID,
		PartnershipID: form.PartnershipID,
	}
	entries, err := server.Service.SearchCompetitionEntries(criteria) // TODO: the underlying query may need optimization

	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data := viewmodel.CompetitionEntriesToViewModel(entries)
	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/v1.0/event/entries
func (server EntryServer) SearchEventEntryHandler(w http.ResponseWriter, r *http.Request) {

}

// GET /api/v1.0/athlete/entries
func (server EntryServer) SearchAthleteEntryHandler(w http.ResponseWriter, r *http.Request) {

}

// GET /api/v1.0/partnership/entries
func (server EntryServer) SearchPartnershipEntryHandler(w http.ResponseWriter, r *http.Request) {

}
