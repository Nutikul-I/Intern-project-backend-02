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

	tsql := model.SQL_GET_EMPLOYEES

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

	createdAt := time.Now()

	query := model.SQL_CREATE_EMPLOYEES

	log.Infof("Query: %s", query)
	log.Infof("Body: %+v", body)

	_, err = conn.ExecContext(ctx, query,
		body.Prefix,
		body.FirstName,
		body.LastName,
		body.NickName,
		body.PositionID,
		body.Email,
		body.BankName,
		body.BankBranch,
		body.AccountNumber,
		body.PayDate,
		body.WithholdingTax,
		body.SocialSecurity,
		body.SocialSecurityID,
		body.OtRate,
		body.LeaveRightsYear,
		body.LeaveRightsSick,
		body.LeaveRightsPersonal,
		body.Color,
		body.Password,
		body.RoleID,
		true,
		createdAt,
		createdAt,
		createdAt,
		nil,
		body.SeatRate,
		body.PaymentChannel,
		body.AccountType,
	)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return model.UpdateResponse{400, "Error creating employee"}, err
	}

	return model.UpdateResponse{200, "Employee created successfully"}, nil

	// if err != nil {
	// 	var createdAt = time.Now()
	// 	var updateAt = time.Now()
	// 	if body.SocialSecurity == "" {
	// 		body.SocialSecurity = "0" // หรือใช้ค่า 0 หรือ false ตามที่ต้องการ
	// 	}
	// 	if body.SocialSecurity == "0" {
	// 		body.SocialSecurityID = ""
	// 	}
	// 	_, err := conn.ExecContext(ctx, model.SQL_CREATE_EMPLOYEES,
	// 		body.Prefix,
	// 		body.FirstName,
	// 		body.LastName,
	// 		body.Nickname,
	// 		body.PositionID,
	// 		body.Email,
	// 		body.BankName,
	// 		body.BankBranch,
	// 		body.AccountNumber,
	// 		body.PayDate,
	// 		body.WithholdingTax,
	// 		body.SocialSecurity,
	// 		body.SocialSecurityID,
	// 		body.OtRate,
	// 		body.LeaveRightsYear,
	// 		body.LeaveRightsSick,
	// 		body.LeaveRightsPersonal,
	// 		body.Color,
	// 		body.Password,
	// 		body.RoleID,
	// 		createdAt,
	// 		updateAt,
	// 		body.SeatRate,
	// 		body.PaymentChannel,
	// 		body.AccountType,
	// 	)
	// 	if err != nil {
	// 		log.Errorf("Error executing query: %v", err)
	// 		return model.UpdateResponse{}, err
	// 	}

	// 	return model.UpdateResponse{StatusCode: 200, Message: "Customer created successfully"}, nil
	// } else {
	// 	// Customer already exists
	// 	return model.UpdateResponse{StatusCode: 400, Message: "Customer already exists"}, nil
	// }
}
