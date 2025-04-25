package model

import "time"

type (
	CustomerPayload struct {
		CustomerId *string `json:"customerId"`
		Name       *string `json:"name"`
		Phone      *string `json:"phone"`
		Email      *string `json:"email"`
	}

	CustomerPagination struct {
		TotalPages   int               `json:"TotalPages"`
		CustomerList []CustomerPayload `json:"CustomerList"`
	}

	// CreateCustomerPayload struct {
	// 	Name  string `json:"name"`
	// 	Phone string `json:"phone"`
	// 	Email string `json:"email"`
	// }

	// ใช้ตอนสร้างลูกค้าใหม่
	CreateCustomerPayload struct {
		CustomerID string `json:"customerId"` // ต้องเตรียมค่ามาให้ก่อน เช่น UUID
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		CreatedBy  int    `json:"createdBy"`
	}

	// ใช้ตอนอัปเดตลูกค้า
	UpdateCustomerPayload struct {
		CustomerID string `json:"customerId"`
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
	}

	// struct สำหรับเก็บข้อมูลลูกค้าแบบเต็ม ๆ (map ตรงกับ WS_customers)
	Customer struct {
		ID         int        `json:"id"`
		CustomerID string     `json:"customerId"`
		Name       string     `json:"name"`
		Phone      string     `json:"phone"`
		Email      string     `json:"email"`
		CreatedBy  int        `json:"createdBy"`
		CreatedAt  *time.Time `json:"createdAt"`
		UpdatedAt  *time.Time `json:"updatedAt"`
		DeletedAt  *time.Time `json:"deletedAt"`
	}
)

var SQL_GET_CUSTOMERS = `
SELECT 
    customer_id,
    COALESCE(name, '') AS name,
    COALESCE(email, '') AS email
FROM WS_customers
WHERE deleted_at IS NULL
ORDER BY customer_id DESC
LIMIT ? OFFSET ?;`

var SQL_GET_CUSTOMER_BY_ID = `
SELECT 
    customer_id,
    COALESCE(name, '') AS name,
    COALESCE(email, '') AS email
FROM WS_customers
WHERE customer_id = ? AND deleted_at IS NULL;`

var SQL_CREATE_CUSTOMER = `
INSERT INTO WS_customers (customer_id, name, phone, email, created_by, created_at)
VALUES (?, ?, ?, ?, ?, NOW());
`

var SQL_UPDATE_CUSTOMER = `
UPDATE WS_customers 
SET name = ?, email = ?, updated_at = ?
WHERE customer_id = ? AND deleted_at IS NULL;`

var SQL_DELETE_CUSTOMER = `
UPDATE WS_customers 
SET deleted_at = ?
WHERE customer_id = ?;`

var SQL_COUNT_CUSTOMER = `
SELECT COUNT(*) 
FROM WS_customers
WHERE deleted_at IS NULL;`

var SQL_CHECK_CUSTOMER = `
SELECT customer_id, name, phone, email
FROM WS_customers
WHERE customer_id = ? AND deleted_at IS NULL;
`
