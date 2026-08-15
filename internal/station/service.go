package station

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrInvalidTheme    = errors.New("theme mode must be light or dark")
	ErrActivityMissing = errors.New("activity does not exist")
	ErrNameRequired    = errors.New("name is required")
	ErrContactRequired = errors.New("contact is required")
)

type Notifier interface {
	Notify(context.Context, Registration) error
}

type RecordingNotifier struct {
	mu      sync.Mutex
	Records []Registration
}

func (notifier *RecordingNotifier) Notify(_ context.Context, registration Registration) error {
	notifier.mu.Lock()
	defer notifier.mu.Unlock()
	notifier.Records = append(notifier.Records, registration)
	return nil
}

type Service struct {
	repository *MemoryRepository
	notifier   Notifier
}

func NewFixtureService(notifier Notifier) (*Service, *MemoryRepository) {
	repository := NewMemoryRepository(fixtureCatalog())
	return &Service{repository: repository, notifier: notifier}, repository
}

func (service *Service) Home(request HomeRequest) (HomePage, error) {
	mode := request.Mode
	if mode == "" {
		mode = ThemeLight
	}
	if mode != ThemeLight && mode != ThemeDark {
		return HomePage{}, ErrInvalidTheme
	}
	catalog := service.repository.Catalog()
	return HomePage{
		StationName:      catalog.StationName,
		MonthLabel:       catalog.MonthLabel,
		Plants:           catalog.Plants,
		VolunteerStories: catalog.Stories,
		Activities:       catalog.Activities,
		Categories:       catalog.Categories,
		Articles:         filterArticles(catalog.Articles, request.Query, request.Tags),
		AvailableTags:    availableTags(catalog.Articles),
		Filters: AppliedFilters{
			Query: strings.TrimSpace(request.Query),
			Tags:  append([]string(nil), request.Tags...),
		},
		Theme: themeFor(mode),
	}, nil
}

func (service *Service) Register(ctx context.Context, request RegisterRequest) (RegistrationResult, error) {
	request.ActivityID = strings.TrimSpace(request.ActivityID)
	request.Name = strings.TrimSpace(request.Name)
	request.Contact = strings.TrimSpace(request.Contact)
	if !service.repository.ActivityExists(request.ActivityID) {
		return RegistrationResult{}, ErrActivityMissing
	}
	if request.Name == "" {
		return RegistrationResult{}, ErrNameRequired
	}
	if request.Contact == "" {
		return RegistrationResult{}, ErrContactRequired
	}

	registration := service.repository.SaveRegistration(request)
	// Notification is an optional channel. When none is configured the core
	// business must still complete and return a clear success result.
	if service.notifier != nil {
		if err := service.notifier.Notify(ctx, registration); err != nil {
			return RegistrationResult{}, err
		}
	}
	return RegistrationResult{Registration: registration, Message: "活动报名成功"}, nil
}

func themeFor(mode ThemeMode) Theme {
	if mode == ThemeDark {
		return Theme{
			Mode:           ThemeDark,
			Background:     "#161914",
			Text:           "#EEF2E9",
			Border:         "#465044",
			ImageTreatment: "brightness(0.82) contrast(1.08)",
		}
	}
	return Theme{
		Mode:           ThemeLight,
		Background:     "#F7F8F3",
		Text:           "#20261E",
		Border:         "#C8D0C3",
		ImageTreatment: "none",
	}
}
