package competition

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
	"strconv"
)

type PublicCompetitionServer struct {
	businesslogic.ICompetitionRepository
	businesslogic.IEventRepository
	businesslogic.IEventMetaRepository
}

// GET /api/competitions
// Search competition(s). This controller is invokable without authentication
func (server PublicCompetitionServer) SearchCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	searchDTO := new(businesslogic.SearchCompetitionCriteria)
	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	} else {
		competitions, err := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{
			ID:       searchDTO.ID,
			Name:     searchDTO.Name,
			StatusID: searchDTO.StatusID,
		})
		if err != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
			return
		}
		data := make([]viewmodel.CompetitionViewModel, 0)
		for _, each := range competitions {
			data = append(data, viewmodel.CompetitionDataModelToViewModel(each, businesslogic.AccountTypeNoAuth))
		}
		output, _ := json.Marshal(data)
		w.Write(output)

	}
}

// GET /api/competition/federation
func (server PublicCompetitionServer) GetUniqueEventFederationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	if len(searchResults) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	competition := searchResults[0]

	federations, err := competition.GetEventUniqueFederations(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Federation, 0)
	for _, each := range federations {
		data = append(data, viewmodel.Federation{
			ID:           each.ID,
			Name:         each.Name,
			Abbreviation: each.Abbreviation,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/division
func (server PublicCompetitionServer) GetEventUniqueDivisionsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	if len(searchResults) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	competition := searchResults[0]

	divisions, err := competition.GetEventUniqueDivisions(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.DivisionViewModel, 0)
	for _, each := range divisions {
		data = append(data, viewmodel.DivisionViewModel{
			ID:         each.ID,
			Name:       each.Name,
			Federation: each.FederationID,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/age
func (server PublicCompetitionServer) GetEventUniqueAgesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	if len(searchResults) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	competition := searchResults[0]
	ages, err := competition.GetEventUniqueAges(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Age, 0)
	for _, each := range ages {
		data = append(data, viewmodel.Age{
			ID:       each.ID,
			Name:     each.Name,
			Division: each.DivisionID,
			Maximum:  each.AgeMaximum,
			Minimum:  each.AgeMinimum,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/proficiency
func (server PublicCompetitionServer) GetEventUniqueProficienciesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}
	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	if len(searchResults) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	competition := searchResults[0]
	proficiencies, err := competition.GetEventUniqueProficiencies(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Proficiency, 0)
	for _, each := range proficiencies {
		data = append(data, viewmodel.ProficiencyDataModelToViewModel(each))
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/style
func (server PublicCompetitionServer) GetEventUniqueStylesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	if len(searchResults) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	competition := searchResults[0]
	styles, err := competition.GetEventUniqueStyles(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Style, 0)
	for _, each := range styles {
		data = append(data, viewmodel.Style{
			ID:   each.ID,
			Name: each.Name,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}
