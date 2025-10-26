package collection

import (
	"api-requester/context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

/*
*	Return array of all collections.
 */
func GetAllCollection(ctx *context.AppContext) ([]*Collection, error) {
	rows, err := ctx.DB.Query("SELECT id, name, description, created_at, updated_at FROM collection")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var collections []*Collection

	for rows.Next() {
		var collec Collection
		err := rows.Scan(
			&collec.ID,
			&collec.Name,
			&collec.Description,
			&collec.Created_at,
			&collec.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		collections = append(collections, &collec)
	}

	return collections, nil
}

/*
*	Create and add collection to database. Created_at and Updated_at automatically set to
*	local date time. Returns pointer to created collection.
 */
func AddCollection(ctx *context.AppContext, name string, description *string) (*Collection, error) {
	// prepared statment, secure against sql injection
	stmt, err := ctx.DB.Prepare("INSERT INTO collection (name, description) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// go doesnt allow strings to be null, so I must use sql nullString
	var desc sql.NullString
	if description != nil {
		desc = sql.NullString{String: *description, Valid: true}
	} else {
		desc = sql.NullString{Valid: false}
	}

	result, err := stmt.Exec(name, desc.String)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Collection{
		ID:          int(lastId),
		Name:        name,
		Created_at:  time.Now().String(),
		Updated_at:  time.Now().String(),
		Description: desc,
		Requests:    nil,
	}, nil
}

/*
*	Update and saves collection in db. This DOES NOT uninformed values, only replaces new ones.
*	ID, Created_at and Updated_at will be ignored because this function is not mean
*	to update then manually. Updated_at is automatic updated to local date time.
 */
func UpdateCollection(ctx *context.AppContext, collection_id int, collection *Collection) error {
	queryClauses := []string{}
	args := []interface{}{}

	if collection.Name != "" {
		queryClauses = append(queryClauses, "name = ?")
		args = append(args, collection.Name)
	}

	if collection.Description.Valid {
		queryClauses = append(queryClauses, "description = ?")
		args = append(args, collection.Description.String)
	}

	if len(queryClauses) == 0 {
		return fmt.Errorf("nothing to update")
	}

	queryClauses = append(queryClauses, "updated_at = ?")
	args = append(args, time.Now().Format(time.DateTime))

	query := fmt.Sprintf("UPDATE collection SET %s WHERE id = ?;", strings.Join(queryClauses, ", "))
	args = append(args, collection_id)
	_, err := ctx.DB.Exec(query, args...)
	return err
}

/*
*	Delete collection with matching id from database.
 */
func DeleteCollectionById(ctx *context.AppContext, collection_id int) error {
	stmt, err := ctx.DB.Prepare("DELETE FROM collection WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(collection_id)
	return err
}

/*
*	Return collection with matching id or ErrNoRows if not found.
 */
func SearchCollectionById(ctx *context.AppContext, collection_id int) (*Collection, error) {
	row := ctx.DB.QueryRow("SELECT * FROM collection WHERE id = ?;", collection_id)
	var collection Collection
	err := row.Scan(
		&collection.ID,
		&collection.Name,
		&collection.Created_at,
		&collection.Updated_at,
		&collection.Description)

	if err != nil {
		return nil, err
	}

	return &collection, nil
}

/*
*	Return collection with containing name or ErrNoRows if not found.
 */
func SearchCollectionContainingName(collections []*Collection, collection_name string) []*Collection {
	var cols []*Collection
	for _, c := range collections {
		if strings.Contains(c.Name, collection_name) {
			cols = append(cols, c)
		}
	}
	return cols
}
