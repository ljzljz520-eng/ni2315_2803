package station

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeIncludesFeaturedSectionsAndDarkTheme(t *testing.T) {
	service, _ := NewFixtureService(&RecordingNotifier{})
	home, err := service.Home(HomeRequest{Mode: ThemeDark})
	if err != nil {
		t.Fatal(err)
	}
	if len(home.Plants) == 0 || len(home.VolunteerStories) == 0 || len(home.Activities) == 0 {
		t.Fatalf("featured sections are incomplete: %+v", home)
	}
	if len(home.Categories) != 4 {
		t.Fatalf("got %d categories", len(home.Categories))
	}
	if home.Theme.Border != "#465044" || home.Theme.ImageTreatment == "none" {
		t.Fatalf("unexpected dark theme: %+v", home.Theme)
	}
}

func TestHomeCombinesSearchAndTags(t *testing.T) {
	service, _ := NewFixtureService(&RecordingNotifier{})
	home, err := service.Home(HomeRequest{Query: "修剪", Tags: []string{"夏季"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(home.Articles) != 1 || home.Articles[0].ID != "tomato-pruning" {
		t.Fatalf("unexpected articles: %+v", home.Articles)
	}
}

func TestActivityRegistrationWorkflow(t *testing.T) {
	service, repository := NewFixtureService(nil)
	handler := NewHandler(service)
	request := httptest.NewRequest(http.MethodPost, "/api/activities/registrations", strings.NewReader(`{"activity_id":"seedling-swap","name":"Lin","contact":"lin@example.test"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("got status %d with body %s", response.Code, response.Body.String())
	}
	var result RegistrationResult
	if err := json.Unmarshal(response.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Registration.Status != "confirmed" || result.Message == "" {
		t.Fatalf("unexpected result: %+v", result)
	}
	if len(repository.Registrations()) != 1 {
		t.Fatalf("got %d registrations", len(repository.Registrations()))
	}
}

func TestActivityRegistrationWithChannel(t *testing.T) {
	notifier := &RecordingNotifier{}
	service, _ := NewFixtureService(notifier)
	result, err := service.Register(context.Background(), RegisterRequest{ActivityID: "night-insect-watch", Name: "Mei", Contact: "mei@example.test"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Registration.ID != "registration-001" || len(notifier.Records) != 1 {
		t.Fatalf("unexpected registration result: %+v", result)
	}
}
