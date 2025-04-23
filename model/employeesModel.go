package model

import "time"

type (
	EmployeesPayload struct {
		ID        int     `json:"id"`
		FirstName *string `json:"firstName"`
		LastName  *string `json:"lastName"`
		Phone     *string `json:"phone"`
		Email     *string `json:"email"`
	}

	EmployeesPagination struct {
		TotalPages    int                `json:"totalPages"`
		EmployeesList []EmployeesPayload `json:"employeesList"`
	}

	CreateEmployeesPayload struct {
		Prefix         string  `json:"prefix"`
		FirstName      string  `json:"firstName"`
		LastName       string  `json:"lastName"`
		Nickname       string  `json:"nickname"`
		Gender         string  `json:"gender"`
		StartDate      string  `json:"startDate"`
		PositionID     int     `json:"positionId"`
		Phone          string  `json:"phone"`
		Email          string  `json:"email"`
		BankName       string  `json:"bankName"`
		BankBranch     string  `json:"bankBranch"`
		AccountName    string  `json:"accountName"`
		AccountNumber  string  `json:"accountNumber"`
		PayDate        string  `json:"payDate"`
		WithholdingTax float64 `json:"withholdingTax"`
		SocialSecurity bool    `json:"socialSecurity"`
		RoleID         int     `json:"roleId"`
		CreatedBy      int     `json:"createdBy"`
	}

	UpdateEmployeesPayload struct {
		ID             int     `json:"id"`
		Prefix         string  `json:"prefix"`
		FirstName      string  `json:"firstName"`
		LastName       string  `json:"lastName"`
		Nickname       string  `json:"nickname"`
		Gender         string  `json:"gender"`
		StartDate      string  `json:"startDate"`
		PositionID     int     `json:"positionId"`
		Phone          string  `json:"phone"`
		Email          string  `json:"email"`
		BankName       string  `json:"bankName"`
		BankBranch     string  `json:"bankBranch"`
		AccountName    string  `json:"accountName"`
		AccountNumber  string  `json:"accountNumber"`
		PayDate        string  `json:"payDate"`
		WithholdingTax float64 `json:"withholdingTax"`
		SocialSecurity bool    `json:"socialSecurity"`
		RoleID         int     `json:"roleId"`
		UpdatedBy      int     `json:"updatedBy"`
	}

	Employees struct {
		ID                  int        `json:"id"`
		Prefix              string     `json:"prefix"`
		FirstName           string     `json:"firstName"`
		LastName            string     `json:"lastName"`
		Nickname            string     `json:"nickname"`
		Gender              string     `json:"gender"`
		StartDate           *time.Time `json:"startDate"`
		PositionID          int        `json:"positionId"`
		Phone               string     `json:"phone"`
		Email               string     `json:"email"`
		BankName            string     `json:"bankName"`
		BankBranch          string     `json:"bankBranch"`
		AccountName         string     `json:"accountName"`
		AccountNumber       string     `json:"accountNumber"`
		PayDate             string     `json:"payDate"`
		WithholdingTax      float64    `json:"withholdingTax"`
		SocialSecurity      bool       `json:"socialSecurity"`
		SocialSecurityID    string     `json:"socialSecurityId"`
		OTRate              float64    `json:"otRate"`
		BonusRate           float64    `json:"bonusRate"`
		LeaveRightsYear     int        `json:"leaveRightsYear"`
		LeaveRightsSick     int        `json:"leaveRightsSick"`
		LeaveRightsPersonal int        `json:"leaveRightsPersonal"`
		Color               string     `json:"color"`
		Password            string     `json:"password"`
		RoleID              int        `json:"roleId"`
		IsActive            bool       `json:"isActive"`
		LastLogin           *time.Time `json:"lastLogin"`
		CreatedAt           *time.Time `json:"createdAt"`
		UpdatedAt           *time.Time `json:"updatedAt"`
		DeletedAt           *time.Time `json:"deletedAt"`
	}
)

var SQL_GET_EMPLOYEES = `
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
WHERE deleted_at IS NULL
ORDER BY id DESC
LIMIT ? OFFSET ?;`

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
    prefix, first_name, last_name, nickname, gender, start_date, position_id, phone, email, 
    bank_name, bank_branch, account_name, account_number, pay_date, withholding_tax, 
    social_security, social_security_id, ot_rate, bonus_rate, leave_rights_year, 
    leave_rights_sick, leave_rights_personal, color, password, role_id, is_active, 
    last_login, created_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

var SQL_UPDATE_EMPLOYEES = `
UPDATE WS_employees 
SET prefix = ?, first_name = ?, last_name = ?, nickname = ?, gender = ?, start_date = ?, 
    position_id = ?, phone = ?, email = ?, bank_name = ?, bank_branch = ?, account_name = ?, 
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
