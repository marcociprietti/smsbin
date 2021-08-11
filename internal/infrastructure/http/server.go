package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cipma/smsbin/app"
	"github.com/cipma/smsbin/app/command"
	"github.com/cipma/smsbin/app/query"
	"github.com/cipma/smsbin/app/sms"
	"github.com/cipma/smsbin/infrastructure/config"
	"github.com/cipma/smsbin/infrastructure/dto"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type Server struct {
	application app.Application
}

func RunHttpServer(application app.Application) {
	server := Server{application: application}

	router := httprouter.New()
	router.POST("/sms", server.SaveSms)
	router.GET("/sms/:id", server.FindSms)
	router.GET("/sms", server.FindAllSms)

	configuration := config.GetConfig()
	host := configuration.GetString("server.host")
	port := configuration.GetInt("server.port")

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router)
	if err != nil {
		panic(err)
	}
}

func (s Server) SaveSms(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	inputData := dto.CreateSms{}
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		writeError(w, InvalidJsonError{}, http.StatusBadRequest)
		return
	}

	cmd := command.SaveSmsCommand{
		Uuid: uuid.NewString(),
		Text: inputData.Text,
		From: inputData.From,
		To:   inputData.To,
	}

	err = s.application.Commands.SaveSmsHandler.Handle(r.Context(), cmd)
	if err != nil {
		if errors.As(err, &sms.InvalidSmsError{}) {
			writeError(w, err, http.StatusBadRequest)
		} else {
			writeError(w, err, http.StatusInternalServerError)
		}

		return
	}

	writeJsonResponse(w, dto.CreateSmsResponse{Id: cmd.Uuid}, http.StatusCreated)
}

func (s Server) FindSms(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data, err := s.application.Queries.FindSmsHandler.Handle(r.Context(), params.ByName("id"))
	if err != nil {
		if errors.As(err, &query.SmsNotFoundError{}) {
			writeError(w, err, http.StatusNotFound)
		} else {
			writeError(w, err, http.StatusInternalServerError)
		}

		return
	}

	writeJsonResponse(w, data, http.StatusOK)
}

func (s Server) FindAllSms(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	lastUuid := r.URL.Query().Get("lastUuid")
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	data, err := s.application.Queries.AllSmsHandler.Handle(r.Context(), size, lastUuid)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, data, http.StatusOK)
}
