package station

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if recover() != nil {
			writeError(writer, http.StatusInternalServerError, "request_failed")
		}
	}()

	switch {
	case request.Method == http.MethodGet && request.URL.Path == "/":
		handler.home(writer, request)
	case request.Method == http.MethodGet && request.URL.Path == "/api/home":
		handler.home(writer, request)
	case request.Method == http.MethodPost && request.URL.Path == "/api/activities/registrations":
		handler.register(writer, request)
	default:
		writeError(writer, http.StatusNotFound, "not_found")
	}
}

func (handler *Handler) home(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	home, err := handler.service.Home(HomeRequest{
		Query: query.Get("q"),
		Tags:  query["tag"],
		Mode:  ThemeMode(query.Get("mode")),
	})
	if err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	writeResponse(writer, http.StatusOK, home)
}

func (handler *Handler) register(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	decoder := json.NewDecoder(http.MaxBytesReader(writer, request.Body, 1<<20))
	decoder.DisallowUnknownFields()
	var input RegisterRequest
	if err := decoder.Decode(&input); err != nil {
		writeError(writer, http.StatusBadRequest, "invalid_request")
		return
	}
	result, err := handler.service.Register(request.Context(), input)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, ErrActivityMissing) || errors.Is(err, ErrNameRequired) || errors.Is(err, ErrContactRequired) {
			status = http.StatusUnprocessableEntity
		}
		writeError(writer, status, err.Error())
		return
	}
	writeResponse(writer, http.StatusCreated, result)
}

func writeResponse(writer http.ResponseWriter, status int, value any) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(value)
}

func writeError(writer http.ResponseWriter, status int, message string) {
	writeResponse(writer, status, map[string]string{"error": message})
}
