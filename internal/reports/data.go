package reports

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// GetReportById ищем репорт по id
func (r *Repo) GetReportById(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.Id, &report.Title, &report.Description, &report.CreatedAt, &report.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// CreateNewReport добавляет новый репорт в бд
func (r *Repo) CreateNewReport(title, description string, createdAt time.Time) error {
	_, err := r.db.Exec(`INSERT INTO reports (title, description, created_at) VALUES ($1, $2, $3)`, title, description, createdAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateReportById обновление репорта в бд
func (r *Repo) UpdateReportById(id int, title, description string, updatedAt time.Time) error {
	_, err := r.db.Exec(`UPDATE reports SET (title, description, updated_at) = ($2, $3, $4) WHERE id = $1`, id, title, description, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteReportById удаление репорта из бд
func (r *Repo) DeleteReportById(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
