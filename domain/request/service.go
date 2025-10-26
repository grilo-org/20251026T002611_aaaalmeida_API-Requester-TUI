package request

import (
	"api-requester/context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

/**
* 	Create and insert request into database. Status_code, headers and body are optional and can
*	be pass as nil. Created_at and updated_at are automatically set as local date time.
*	Returns pointer to created request.
 */
func AddRequest(ctx *context.AppContext, req *Request) (*Request, error) {
	stmt, err := ctx.DB.Prepare(`
		INSERT INTO request
		(name, url, method_id, collection_id, status_code, headers, body, body_type)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	headers, err := json.Marshal(req.Headers)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(
		req.Name,
		req.Url,
		req.Method_id,
		req.Collection_id,
		req.Expected_Status_code,
		headers,
		req.Body,
		req.BodyType)

	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	req.ID = int(lastId)
	req.Created_at = time.Now().String()
	req.Updated_at = time.Now().String()

	return req, nil
}

// /*
// *	Return array of all requests.
//  */
// func GetAllRequest(ctx *context.AppContext) ([]Request, error) {
// 	rows, err := ctx.DB.Query(`SELECT id, name, url, method_id, collection_id,
// 	 status_code, headers, body, created_at, updated_at FROM request;`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var requests []Request
// 	for rows.Next() {
// 		var req Request
// 		err := rows.Scan(
// 			&req.ID,
// 			&req.Name,
// 			&req.Url,
// 			&req.Method_id,
// 			&req.Collection_id,
// 			&req.Status_code,
// 			&req.Headers,
// 			&req.Body,
// 			&req.Created_at,
// 			&req.Updated_at,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		requests = append(requests, req)
// 	}
// 	return requests, nil
// }

// /*
// *	Return array of all requests with matching method_id.
//  */
// func SearchRequestByMethodId(ctx *context.AppContext, method_id int) ([]Request, error) {
// 	rows, err := ctx.DB.Query(`SELECT id, name, url, method_id, collection_id,
// 	 status_code, headers, body, created_at, updated_at FROM request;`, method_id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var requests []Request
// 	for rows.Next() {
// 		var req Request
// 		err := rows.Scan(
// 			&req.ID,
// 			&req.Name,
// 			&req.Url,
// 			&req.Method_id,
// 			&req.Collection_id,
// 			&req.Status_code,
// 			&req.Headers,
// 			&req.Body,
// 			&req.Created_at,
// 			&req.Updated_at,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		requests = append(requests, req)
// 	}
// 	return requests, nil
// }

// /*
// *	Return request with matching id or ErrNoRows if not found.
//  */
// func SearchRequestById(ctx *context.AppContext, request_id int) (*Request, error) {
// 	row := ctx.DB.QueryRow(`SELECT id, name, url, method_id, collection_id,
// 	 status_code, headers, body, created_at, updated_at FROM request;`, request_id)

// 	var request Request
// 	err := row.Scan(
// 		&request.ID,
// 		&request.Name,
// 		&request.Url,
// 		&request.Method_id,
// 		&request.Collection_id,
// 		&request.Status_code,
// 		&request.Headers,
// 		&request.Body,
// 		&request.Created_at,
// 		&request.Updated_at,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &request, nil
// }

/*
Return array of all requests with matching collection_id.
*/
func SearchRequestByCollectionId(ctx *context.AppContext, collection_id int) ([]*Request, error) {
	rows, err := ctx.DB.Query(`SELECT id, name, url, method_id, collection_id,
	 status_code, headers, body, body_type, created_at, updated_at FROM request WHERE collection_id = ?;`, collection_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var requests []*Request
	for rows.Next() {
		var req Request
		var headers sql.NullString
		err := rows.Scan(
			&req.ID,
			&req.Name,
			&req.Url,
			&req.Method_id,
			&req.Collection_id,
			&req.Expected_Status_code,
			&headers,
			&req.Body,
			&req.BodyType,
			&req.Created_at,
			&req.Updated_at,
		)

		if err != nil {
			return nil, err
		}

		if headers.Valid {
			// unmarshal json into string: any map
			var temp map[string]any
			if err := json.Unmarshal([]byte(headers.String), &temp); err != nil {
				return nil, err
			}

			// transform temp map into a string: string map
			target := make(map[string]string, len(temp))
			for k, v := range temp {
				target[k] = fmt.Sprint(v)
			}

			req.Headers = target
		} else {
			req.Headers = make(map[string]string) // empty
		}
		requests = append(requests, &req)
	}
	return requests, nil
}

// // TODO: remover referencia do ponteiro da lista de requests durante execução
/*
*	Delete request with matching id from database.
 */
func DeleteRequestById(ctx *context.AppContext, request_id int) error {
	stmt, err := ctx.DB.Prepare("DELETE FROM request WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(request_id)
	return err
}

/*
*	Update and saves request in db. This DOES NOT override uninformed values, only replaces new ones.
*	ID, Collection_ID, Created_at and Updated_at will be ignored because this function is not mean
*	to update then manually. Updated_at is automatic updated to local date time.
 */
func UpdateRequest(ctx *context.AppContext, request_id int, req *Request) error {
	queryClauses := []string{}
	args := []interface{}{}

	if req.Name != "" {
		queryClauses = append(queryClauses, "name = ?")
		args = append(args, req.Name)
	}

	if req.Url != "" {
		queryClauses = append(queryClauses, "url = ?")
		args = append(args, req.Url)
	}

	if req.Method_id != 0 {
		queryClauses = append(queryClauses, "method_id = ?")
		args = append(args, req.Method_id)
	}

	if req.Expected_Status_code != nil {
		queryClauses = append(queryClauses, "status_code = ?")
		args = append(args, &req.Expected_Status_code)
	}

	if len(req.Headers) != 0 {
		queryClauses = append(queryClauses, "headers = ?")
		headers, err := json.Marshal(req.Headers)
		if err != nil {
			return err
		}
		args = append(args, headers)
	}

	if len(req.Body) != 0 {
		queryClauses = append(queryClauses, "body = ?")
		args = append(args, string(req.Body))
		queryClauses = append(queryClauses, "body_type = ?")
		args = append(args, req.BodyType)
	} else {
		queryClauses = append(queryClauses, "body_type = ?")
		args = append(args, 0) // BodyTypeNull
	}

	if len(queryClauses) == 0 {
		// TODO: arrumar esse erro
		return fmt.Errorf("nothing to update")
	}

	queryClauses = append(queryClauses, "updated_at = ?")
	args = append(args, time.Now().Format(time.DateTime))

	query := fmt.Sprintf("UPDATE request SET %s WHERE id = ?;", strings.Join(queryClauses, ", "))
	args = append(args, request_id)
	_, err := ctx.DB.Exec(query, args...)
	return err
}

/*
*	Call HTTP Request using Headers (if valid). Returns body stringified.
 */
func CallRequest(req *Request) (string, error) {
	body := strings.NewReader(string(req.Body))

	var method string
	switch req.Method_id {
	case 2:
		method = http.MethodPost
	case 3:
		method = http.MethodPut
	case 4:
		method = http.MethodDelete
	case 5:
		method = http.MethodPatch
	case 6:
		method = http.MethodHead
	case 7:
		method = http.MethodTrace
	case 8:
		method = http.MethodOptions
	default:
		method = http.MethodGet
	}

	httpRequest, err := http.NewRequest(method, req.Url, body)
	if err != nil {
		return "", err
	}

	for key, value := range req.Headers {
		httpRequest.Header.Set(key, value)
	}

	// If req.Headers dont have content type and BodyType is different than null,
	// automatically add it based on req.BodyType
	if _, ok := req.Headers["Content-Type"]; !ok && req.BodyType != BodyTypeNull {
		httpRequest.Header.Set("Content-Type", req.BodyType.String())
	}

	client := &http.Client{}
	response, err := client.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(responseBody), nil
}

// /*
//   - Call HTTP Request with matching id. Returns body stringified.
//     */
func CallRequestById(ctx *context.AppContext, request_id int) (string, error) {
	row := ctx.DB.QueryRow(`SELECT id, name, url, method_id, collection_id,
	 status_code, headers, body, body_type, created_at, updated_at FROM request WHERE id = ?;`, request_id)

	var request Request
	err := row.Scan(
		&request.ID,
		&request.Url,
		&request.Name,
		&request.Method_id,
		&request.Collection_id,
		&request.Expected_Status_code,
		&request.Headers,
		&request.Body,
		&request.BodyType,
		&request.Created_at,
		&request.Updated_at)

	if err != nil {
		return "", err
	}

	return CallRequest(&request)
}
