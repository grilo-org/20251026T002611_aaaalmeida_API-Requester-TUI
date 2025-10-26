package method

import "api-requester/context"

/*
*	Return method with matching id or ErrNoRows if not found.
 */
func SearchMethodById(ctx *context.AppContext, method_id int) (*Method, error) {
	row := ctx.DB.QueryRow("SELECT * FROM method WHERE id = ?;", method_id)
	var method Method
	err := row.Scan(&method.ID, &method.Name)
	if err != nil {
		return nil, err
	}
	return &method, nil
}

/*
*	Return array of methods.
 */
func GetAllMethod(ctx *context.AppContext) ([]Method, error) {
	rows, err := ctx.DB.Query("SELECT * FROM method;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var methods []Method
	for rows.Next() {
		var m Method
		err := rows.Scan(&m.ID, &m.Name)
		if err != nil {
			return nil, err
		}
		methods = append(methods, m)
	}

	return methods, nil
}
