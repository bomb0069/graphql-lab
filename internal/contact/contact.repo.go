package contact

import (
	"fmt"
	"graphql-api/pkg/data"
	"graphql-api/pkg/data/models"

	_ "github.com/mattn/go-sqlite3"
)

// ContactRepo represents the repository for contact operations
type ContactRepo struct {
  DB *data.DB
}

// NewContactRepo creates a new instance of ContactRepo
func NewContactRepo() *ContactRepo {
  db := data.NewDB()
  return &ContactRepo{DB: db}
}

// Get Contacts fetches contacts from the database with support for text search, limit, and offset
func (cr *ContactRepo) GetContactsBySearchText(searchText string, limit, offset int) ([]*models.ContactModel, error) {
  var contacts []*models.ContactModel

  query := fmt.Sprintf(`
            SELECT * FROM contact
             Where name like '%%%s%%' OR first_name like '%%%s%%' OR last_name like '%%%s%%' OR email like '%%%s%%' OR phone like '%%%s%%' OR address like '%%%s%%' OR photo_path like '%%%s%%'
            LIMIT ? OFFSET ?
        `, searchText, searchText, searchText, searchText, searchText, searchText, searchText)

  rows, err := cr.DB.Query(query, limit, offset)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var contact models.ContactModel
    err := rows.Scan(
      &contact.ContactId,
      &contact.Name,
      &contact.FirstName,
      &contact.LastName,
      &contact.GenderId,
      &contact.Dob,
      &contact.Email,
      &contact.Phone,
      &contact.Address,
      &contact.PhotoPath,
      &contact.CreatedAt,
      &contact.CreatedBy,
    )
    if err != nil {
      return nil, err
    }
    contacts = append(contacts, &contact)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }

  return contacts, nil
}
