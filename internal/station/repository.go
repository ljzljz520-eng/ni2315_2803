package station

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Catalog struct {
	StationName string
	MonthLabel  string
	Plants      []Plant
	Stories     []VolunteerStory
	Activities  []Activity
	Articles    []Article
	Categories  []string
}

type MemoryRepository struct {
	catalog       Catalog
	mu            sync.Mutex
	registrations []Registration
}

func NewMemoryRepository(catalog Catalog) *MemoryRepository {
	return &MemoryRepository{catalog: catalog}
}

func (repository *MemoryRepository) Catalog() Catalog {
	return cloneCatalog(repository.catalog)
}

func (repository *MemoryRepository) ActivityExists(activityID string) bool {
	for _, activity := range repository.catalog.Activities {
		if activity.ID == activityID {
			return true
		}
	}
	return false
}

func (repository *MemoryRepository) SaveRegistration(request RegisterRequest) Registration {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	registration := Registration{
		ID:         fmt.Sprintf("registration-%03d", len(repository.registrations)+1),
		ActivityID: request.ActivityID,
		Name:       request.Name,
		Contact:    request.Contact,
		Status:     "confirmed",
	}
	repository.registrations = append(repository.registrations, registration)
	return registration
}

func (repository *MemoryRepository) Registrations() []Registration {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	return append([]Registration(nil), repository.registrations...)
}

func cloneCatalog(catalog Catalog) Catalog {
	copyCatalog := catalog
	copyCatalog.Plants = append([]Plant(nil), catalog.Plants...)
	copyCatalog.Stories = append([]VolunteerStory(nil), catalog.Stories...)
	copyCatalog.Activities = append([]Activity(nil), catalog.Activities...)
	copyCatalog.Categories = append([]string(nil), catalog.Categories...)
	copyCatalog.Articles = make([]Article, len(catalog.Articles))
	for index, article := range catalog.Articles {
		copyCatalog.Articles[index] = article
		copyCatalog.Articles[index].Tags = append([]string(nil), article.Tags...)
	}
	return copyCatalog
}

func availableTags(articles []Article) []string {
	seen := make(map[string]struct{})
	for _, article := range articles {
		for _, tag := range article.Tags {
			seen[tag] = struct{}{}
		}
	}
	tags := make([]string, 0, len(seen))
	for tag := range seen {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	return tags
}

func filterArticles(articles []Article, query string, requestedTags []string) []Article {
	normalizedQuery := strings.ToLower(strings.TrimSpace(query))
	normalizedTags := normalizeTags(requestedTags)
	filtered := make([]Article, 0, len(articles))
	for _, article := range articles {
		if normalizedQuery != "" && !articleMatchesQuery(article, normalizedQuery) {
			continue
		}
		if !articleHasAllTags(article, normalizedTags) {
			continue
		}
		article.Tags = append([]string(nil), article.Tags...)
		filtered = append(filtered, article)
	}
	return filtered
}

func articleMatchesQuery(article Article, query string) bool {
	values := []string{article.Title, article.Summary, article.Category, strings.Join(article.Tags, " ")}
	return strings.Contains(strings.ToLower(strings.Join(values, " ")), query)
}

func articleHasAllTags(article Article, requested []string) bool {
	if len(requested) == 0 {
		return true
	}
	articleTags := make(map[string]struct{}, len(article.Tags))
	for _, tag := range article.Tags {
		articleTags[strings.ToLower(tag)] = struct{}{}
	}
	for _, requestedTag := range requested {
		if _, ok := articleTags[requestedTag]; !ok {
			return false
		}
	}
	return true
}

func normalizeTags(tags []string) []string {
	normalized := make([]string, 0, len(tags))
	seen := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		value := strings.ToLower(strings.TrimSpace(tag))
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		normalized = append(normalized, value)
	}
	return normalized
}
