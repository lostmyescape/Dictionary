package words

import "database/sql"

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetWordById ищем слово по id
func (r *Repo) RGetWordById(id int) (*Word, error) {
	var word Word
	err := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id = $1`, id).
		Scan(&word.Id, &word.Title, &word.Translation)
	if err != nil {
		return nil, err
	}

	return &word, nil
}

// CreateNewWords добавляет новые переводы в базу данных
func (r *Repo) CreateNewWords(word, translate string) error {
	_, err := r.db.Exec(`INSERT INTO ru_en (title, translation) VALUES ($1, $2)`, word, translate)
	if err != nil {
		return err
	}

	return nil
}

// UpdateWord обновление слова в бд
func (r *Repo) UpdateWord(word, translate string) error {
	_, err := r.db.Exec(`UPDATE ru_en SET translation = $2 WHERE title = $1`, word, translate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteWordsById удаление слова из бд
func (r *Repo) DeleteWordsById(id int) error {
	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

// SearchWordsByParam поиск слов по title
func (r *Repo) SearchWordsByParam(title string) ([]Word, error) {
	var words []Word

	query := `
		SELECT id, title, translation 
		FROM ru_en
		WHERE title ILIKE '%' || $1 || '%'
		ORDER BY similarity(title, $1)
		DESC LIMIT 100
		`

	rows, err := r.db.Query(query, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var word Word
		err := rows.Scan(&word.Id, &word.Title, &word.Translation)
		if err != nil {
			return nil, err
		}

		words = append(words, word)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
