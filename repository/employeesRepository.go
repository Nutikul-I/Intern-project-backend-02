package repository

import (
	"context"
	"payso-internal-api/model"
	"time"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
)

func GetEmployeesRepository(mid string, page int, row int) ([]model.EmployeesPayload, error) {
	conn := ConnectDB()
	ctx := context.Background()

	err := conn.PingContext(ctx)
	if err != nil {
		log.Errorf("Error PingContext: %v", err)
		return nil, err
	}

	offset := 0
	if page > 0 {
		offset = (page - 1) * row
	}

	tsql := model.SQL_GET_CUSTOMERS

	rows, err := conn.QueryContext(ctx, tsql, row, offset)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var employees []model.EmployeesPayload
	err = scan.Rows(&employees, rows)
	if err != nil {
		log.Errorf("Error scanning rows: %v", err)
		return employees, err
	}

	return employees, nil
}

func GetTotalEmployeesRepository(row int) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()

	var totalCount int
	err := conn.QueryRowContext(ctx, model.SQL_COUNT_EMPLOYEES).Scan(&totalCount)
	if err != nil {
		log.Errorf("Error counting employees: %v", err)
		return 0, err
	}

	totalPages := (totalCount + row - 1) / row
	return totalPages, nil
}

func CreateEmployeesRepository(body model.CreateEmployeesPayload) (model.UpdateResponse, error) {
	conn := ConnectDB()
	ctx := context.Background()

	// Check if database is alive.
	err := conn.PingContext(ctx)
	if err != nil {
		log.Errorf("Error PingContext: %v", err)
		return model.UpdateResponse{}, err
	}

	// Check if employee already exists
	rows_check, err := conn.QueryContext(ctx, model.SQL_CHECK_EMPLOYEES)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return model.UpdateResponse{}, err
	}
	defer rows_check.Close()

	var EmployeesData model.Employees
	err = scan.Row(&EmployeesData, rows_check)
	if err != nil {
		// Employee does not exist, create a new one

		_, err := conn.ExecContext(ctx, model.SQL_CREATE_EMPLOYEES,
			body.Prefix,
			body.FirstName,
			body.LastName,
			body.Nickname,
			body.Gender,
			body.StartDate,
			body.PositionID,
			body.Phone,
			body.Email,
			body.BankName,
			body.BankBranch,
			body.AccountName,
			body.AccountNumber,
			body.PayDate,
			body.WithholdingTax,
			body.SocialSecurity,
			"", // social_security_id
			0,  // ot_rate
			0,  // bonus_rate
			0,  // leave_rights_year
			0,  // leave_rights_sick
			0,  // leave_rights_personal
			"", // color
			"", // password
			body.RoleID,
			true,       // is_active
			nil,        // last_login
			time.Now(), // created_at
		)

		if err != nil {
			log.Errorf("Error executing insert: %v", err)
			return model.UpdateResponse{}, err
		}

		return model.UpdateResponse{StatusCode: 200, Message: "Employee created successfully"}, nil
	} else {
		// Employee already exists
		return model.UpdateResponse{StatusCode: 400, Message: "Employee already exists"}, nil
	}
}

// func DeleteMerchantRepository(ReqMasterMerchantID string, ReqMerchantID string) (model.UpdateResponse, error) {

// 	conn := ConnectDB()
// 	ctx := context.Background()

// 	// Check if database is alive.
// 	err := conn.PingContext(ctx)
// 	if err != nil {
// 		log.Errorf("Error PingContext: %v", err)
// 		return model.UpdateResponse{}, err
// 	}

// 	tsql_check := model.SQL_CHECK_MERCHANT
// 	rows_check, err := conn.QueryContext(ctx, tsql_check,
// 		sql.Named("MasterMerchantID", ReqMasterMerchantID),
// 		sql.Named("MerchantID", ReqMerchantID))
// 	if err != nil {
// 		log.Errorf("Error executing query: %v", err)
// 		return model.UpdateResponse{}, err
// 	}
// 	defer rows_check.Close()

// 	var MerchantData model.MasterMerchant
// 	err = scan.Row(&MerchantData, rows_check)
// 	if err != nil {
// 		return model.UpdateResponse{StatusCode: 400, Message: "deleted  merchant fail"}, nil
// 	} else {
// 		tsql := model.SQL_DELETE_MERCHANT
// 		rows, err := conn.QueryContext(ctx, tsql,
// 			sql.Named("MasterMerchantID", ReqMasterMerchantID),
// 			sql.Named("MerchantID", ReqMerchantID))
// 		if err != nil {
// 			log.Errorf("Error executing query: %v", err)
// 			return model.UpdateResponse{}, err
// 		}
// 		defer rows.Close()
// 		return model.UpdateResponse{StatusCode: 200, Message: "deleted  merchant success"}, nil
// 	}
// }
