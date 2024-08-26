// internal/storage/note_repository.go

package storage

import (
    "log"

    "github.com/marktsarkov/notesapp/internal/models"
)

// NoteRepository предоставляет методы для работы с заметками
type NoteRepository struct {
    db *DB
}

// NewNoteRepository создает новый репозиторий для заметок
func NewNoteRepository(db *DB) *NoteRepository {
    return &NoteRepository{db: db}
}

// Create создает новую заметку в базе данных
func (r *NoteRepository) Create(note *models.Note) error {
    query := `
        INSERT INTO notes (title, content, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `

    err := r.db.Conn.QueryRow(query, note.Title, note.Content, note.UserID).
        Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)

    if err != nil {
        return err
    }

    log.Printf("Создана заметка: %v", note)
    return nil
}

func (r *NoteRepository) GetAll(userID int) ([]*models.Note, error) {
    var notes []*models.Note
    query := `
        SELECT id, title, content, user_id, created_at, updated_at
        FROM notes
        WHERE user_id = $1
        ORDER BY created_at DESC
    `

    rows, err := r.db.Conn.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var note models.Note
        if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.UserID, &note.CreatedAt, &note.UpdatedAt); err != nil {
            return nil, err
        }
        notes = append(notes, &note)
    }

    return notes, nil
}
