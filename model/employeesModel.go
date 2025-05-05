package model

import "time"

type (
	EmployeesPayload struct {
		ID        int     `json:"id"`
		FirstName *string `json:"firstName"`
		LastName  *string `json:"lastName"`
		Email     *string `json:"email"`
	}

	EmployeesPagination struct {
		TotalPages    int                `json:"totalPages"`
		EmployeesList []EmployeesPayload `json:"employeesList"`
	}

	CreateEmployeesPayload struct {
		Prefix              string  `json:"prefix"`
		FirstName           string  `json:"firstName"`
		LastName            string  `json:"lastName"`
		NickName            string  `json:"nickName"`
		PositionID          uint    `json:"positionId"`
		Email               string  `json:"email"`
		BankName            string  `json:"bankName"`
		BankBranch          string  `json:"bankBranch"`
		AccountNumber       string  `json:"accountNumber"`
		PayDate             string  `json:"payDate"`
		WithholdingTax      float64 `json:"withholdingTax"`
		SocialSecurity      bool    `json:"socialSecurity"`
		SocialSecurityID    string  `json:"socialSecurityId"`
		OtRate              float64 `json:"otRate"`
		LeaveRightsYear     int     `json:"leaveRightsYear"`
		LeaveRightsSick     int     `json:"leaveRightsSick"`
		LeaveRightsPersonal int     `json:"leaveRightsersonal"`
		Color               string  `json:"color"`
		Password            string  `json:"password"`
		RoleID              uint    `json:"roleId"`
		SeatRate            float64 `json:"seatRate"`
		PaymentChannel      string  `json:"paymentChannel"`
		AccountType         string  `json:"accountType"`
	}

	UpdateEmployeesPayload struct {
		ID               int     `json:"id"`
		Prefix           string  `json:"prefix"`
		FirstName        string  `json:"firstName"`
		LastName         string  `json:"lastName"`
		Nickname         string  `json:"nickname"`
		Gender           string  `json:"gender"`
		StartDate        string  `json:"startDate"`
		PositionID       int     `json:"positionId"`
		Email            string  `json:"email"`
		BankName         string  `json:"bankName"`
		BankBranch       string  `json:"bankBranch"`
		AccountName      string  `json:"accountName"`
		AccountNumber    string  `json:"accountNumber"`
		PayDate          string  `json:"payDate"`
		WithholdingTax   float64 `json:"withholdingTax"`
		SocialSecurity   bool    `json:"socialSecurity"`
		SocialSecurityID string  `json:"socialSecurityId"`
		RoleID           int     `json:"roleId"`
		UpdatedBy        int     `json:"updatedBy"`
	}

	Employees struct {
		ID                  uint       `json:"id"`
		Prefix              string     `json:"prefix"`
		FirstName           string     `json:"first_name"`
		LastName            string     `json:"last_name"`
		Nickname            string     `json:"nickname"`
		PositionID          uint       `json:"position_id"`
		Email               string     `json:"email"`
		BankName            string     `json:"bank_name"`
		BankBranch          string     `json:"bank_branch"`
		AccountNumber       string     `json:"account_number"`
		PayDate             string     `json:"pay_date"`
		WithholdingTax      string     `json:"withholding_tax"`
		SocialSecurity      string     `json:"social_security"`
		SocialSecurityID    string     `json:"social_security_id"`
		OtRate              string     `json:"ot_rate"`
		LeaveRightsYear     int        `json:"leave_rights_year"`
		LeaveRightsSick     int        `json:"leave_rights_sick"`
		LeaveRightsPersonal int        `json:"leave_rights_personal"`
		Color               string     `json:"color"`
		Password            string     `json:"password"`
		RoleID              uint       `json:"role_id"`
		IsActive            bool       `json:"is_active"`
		LastLogin           *time.Time `json:"last_login,omitempty"`
		CreatedAt           time.Time  `json:"created_at"`
		UpdatedAt           time.Time  `json:"updated_at"`
		DeletedAt           *time.Time `json:"deleted_at,omitempty"`
		SeatRate            string     `json:"seat_rate"`
		PaymentChannel      string     `json:"payment_channel"`
		AccountType         string     `json:"account_type"`
	}
)

var SQL_GET_EMPLOYEES = `
SELECT
    id,
    prefix,
    first_name,
    last_name,
    nickname,
    position_id,
    email,
    bank_name,
    bank_branch,
    account_number,
    pay_date,
    withholding_tax,
    social_security,
    social_security_id,
    ot_rate,
    leave_rights_year,
    leave_rights_sick,
    leave_rights_personal,
    color,
    password,
    role_id,
    is_active,
    last_login,
    created_at,
    updated_at,
    deleted_at,
    seat_rate,
    payment_channel,
    account_type
FROM
    WS_employees
WHERE
    deleted_at IS NULL AND role_id = ? And id = ?;
`

var SQL_GET_EMPLOYEES_BY_ID = `
SELECT 
    id,
    COALESCE(prefix, '') AS prefix,
    COALESCE(first_name, '') AS first_name,
    COALESCE(last_name, '') AS last_name,
    COALESCE(nickname, '') AS nickname,
    COALESCE(gender, '') AS gender,
    start_date,
    position_id,
    COALESCE(phone, '') AS phone,
    COALESCE(email, '') AS email,
    COALESCE(bank_name, '') AS bank_name,
    COALESCE(bank_branch, '') AS bank_branch,
    COALESCE(account_name, '') AS account_name,
    COALESCE(account_number, '') AS account_number,
    COALESCE(pay_date, '') AS pay_date,
    withholding_tax,
    social_security,
    COALESCE(social_security_id, '') AS social_security_id,
    ot_rate,
    bonus_rate,
    leave_rights_year,
    leave_rights_sick,
    leave_rights_personal,
    COALESCE(color, '') AS color,
    COALESCE(password, '') AS password,
    role_id,
    is_active,
    last_login,
    created_at,
    updated_at,
    deleted_at
FROM WS_employees
WHERE id = ? AND deleted_at IS NULL;`

var SQL_CREATE_EMPLOYEES = `
INSERT INTO WS_employees (
	prefix, first_name, last_name, nickname, position_id, email, bank_name,
	bank_branch, account_number, pay_date, withholding_tax, social_security,
	social_security_id, ot_rate, leave_rights_year, leave_rights_sick,
	leave_rights_personal, color, password, role_id, is_active, last_login, created_at, updated_at,
	deleted_at, seat_rate, payment_channel, account_type

) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?
);`

var SQL_UPDATE_EMPLOYEES = `
UPDATE WS_employees 
SET prefix = ?, first_name = ?, last_name = ?, nickname = ?, gender = ?, start_date = ?, 
    position_id = ?, email = ?, bank_name = ?, bank_branch = ?, account_name = ?, 
    account_number = ?, pay_date = ?, withholding_tax = ?, social_security = ?, 
    social_security_id = ?, ot_rate = ?, bonus_rate = ?, leave_rights_year = ?, 
    leave_rights_sick = ?, leave_rights_personal = ?, color = ?, password = ?, role_id = ?, 
    is_active = ?, last_login = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;`

var SQL_DELETE_EMPLOYEES = `
UPDATE WS_employees 
SET deleted_at = ?
WHERE id = ?;`

var SQL_COUNT_EMPLOYEES = `
SELECT COUNT(*) 
FROM WS_employees
WHERE deleted_at IS NULL;`

var SQL_CHECK_EMPLOYEES = `
SELECT id, first_name, last_name, email
FROM WS_employees
WHERE id = ?;
`
