# Banking System (Golang + Gin + GORM + PostgreSQL)

Comprehensive backend for a simple banking system. Implements multiple banks and branches, customer registration, savings accounts, transactions (deposit/withdraw), loans with a fixed interest rate of 12% per year, repayments (partial/full), reporting endpoints, and full persistence using GORM + PostgreSQL.

---

## Key Features

- Multi-tenant support: multiple `Bank` records, each with many `Branch` records
- Customer registration 
- Open savings accounts per customer + branch
- Deposit and withdraw operations with transactional updates and transaction records
- Loans with fixed interest rate (12% yearly) and ability to repay partially or fully
- Automated calculation of total interest and pending balance
- Endpoints to fetch account details, transaction history, loan details, interest for current year, and pending loan amount
- Input validation and clear JSON error responses
- GORM models with timestamps and relations; automatic migration

---

## Project Structure

- `cmd/main.go` — application entrypoint
- `initializers/database.go` — DB connection 
- `migrate/migrate.go` - migration
- `models/` — GORM models (Bank, Branch, Customer, Account, Transaction, Loan, LoanPayment)
- `services/` — business logic, DB transactions
- `controllers/` — Gin handlers / request validation
- `routes/` — route registration

```markdown
Assignment-2/
├── controllers/
│   ├── account_controller.go
│   ├── bank_controller.go
│   ├── branch_controller.go
│   ├── customer_controller.go
│   ├── loan_controller.go
│   ├── transaction_controller.go
├── initializers/
│   ├── database.go
│   ├── loadEnv.go
├── main.go
├── models/
│   ├── account.go
│   ├── bank.go
│   ├── branch.go
│   ├── customer.go
│   ├── loan.go
│   ├── loan_payment.go
│   ├── transaction.go
├── routes/
│   ├── routes.go
├── services/
│   ├── account_service.go
│   ├── loan_service.go
│   ├── transaction_service.go
└── utils/
    ├── utils.go
```

---

## Database Models (ER Overview)

All models use timestamps and GORM tags for relations and constraints. Primary keys are auto-increment `uint`.

- Bank: `id, name, code, created_at`
- Branch: `id, bank_id (FK), name, ifsc_code, address, created_at`
- Customer: `id, name, email, phone, address, created_at`
- Account: `id, customer_id (FK), branch_id (FK), account_number, type (savings), balance, status, created_at`
- Transaction: `id, account_id (FK), type (deposit|withdraw|loan_repay), amount, created_at`
- Loan: `id, customer_id (FK), branch_id (FK), principal_amount, interest_rate(12%), total_interest, amount_paid, remaining_amount, start_date, end_date, status, created_at`
- LoanPayment: `id, loan_id (FK), amount, payment_date, created_at`

Simplified relationship diagram:

Bank 1..* — 1..* Branch
Customer 1..* — 1..* Account
Account 1..* — 1..* Transaction
Customer 1..* — 1..* Loan
Loan 1..* — 1..* LoanPayment

---

## Business Rules / Important Logic

- Loan interest rate is fixed at 12% per year.
- Interest formula: Interest = Principal × Rate × Time (years)
- When applying a loan the system computes `TotalInterest` for the loan duration and sets `RemainingAmount = Principal + TotalInterest`.
- Withdrawals are prevented if the account `Balance` is insufficient.
- Deposit and withdraw operations are executed inside a DB transaction and each creates a `Transaction` record.
- Loan repayments reduce `RemainingAmount` and record a `LoanPayment` transaction in a DB transaction. If `RemainingAmount` reaches zero the loan `Status` becomes `closed` and `EndDate` is set.
- All inputs are validated; handlers return HTTP 4xx for bad requests and 5xx for server errors.

---

## Environment Variables

- `DB_HOST` (default: `localhost`)
- `DB_PORT` (default: `5432`)
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_SSLMODE` (default: `disable`)
- `PORT` (server port, default: `8080`)

Example export:

```bash
export DB_HOST=127.0.0.1
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=banking
export DB_SSLMODE=disable
export PORT=8080
```

---

## Install & Run

1. Ensure Go and PostgreSQL are installed.
2. Create the database in Postgres (e.g. `createdb bank`).
3. Set environment variables (see above).
4. Fetch dependencies and run:

```bash
go mod tidy
go run ./cmd
```

The server listens on `http://localhost:8080` by default.

---

## API Reference

Base path: `/api`

Bank
- POST `/api/banks`
  - Body: `{ "name": "Acme Bank", "code": "ACME001" }`
  - Response: created `Bank` JSON
- GET `/api/banks`
  - Response: list of banks
 - PUT `/api/banks/:id`
   - Body: `{ "name": "New Name", "code": "NEWCODE" }`
   - Response: updated `Bank` JSON
 - DELETE `/api/banks/:id`
   - Response: `{ "message": "bank deleted" }`

Branch
- POST `/api/branches`
  - Body: `{ "bank_id": 1, "name": "Main", "ifsc_code": "ACME0001", "address": "..." }`
  - Response: created `Branch`
- GET `/api/branches/:bankId`
  - Response: branches for bank
 - PUT `/api/branches/:id`
   - Body: `{ "name": "Updated", "ifsc_code": "NEWIFSC", "address": "..." }`
   - Response: updated `Branch`
 - DELETE `/api/branches/:id`
   - Response: `{ "message": "branch deleted" }`

Customer
- POST `/api/customers`
  - Body: `{ "name": "Alice", "email": "alice@example.com", "phone": "123", "address": "..." }`
  - Response: created `Customer`
- GET `/api/customers/:id`
  - Response: customer with `Accounts` and `Loans` preloaded
 - PUT `/api/customers/:id`
   - Body: `{ "name": "Alice New", "email": "alice@new.com", "phone": "321", "address": "..." }`
   - Response: updated `Customer`
 - DELETE `/api/customers/:id`
   - Response: `{ "message": "customer deleted" }`

Account
- POST `/api/accounts/open`
  - Body: `{ "customer_id": 1, "branch_id": 1 }`
  - Response: created `Account` (fields: account_number, balance = 0)
- GET `/api/accounts/:id`
  - Response: account details (preloaded transactions)
- GET `/api/accounts/customer/:customerId`
  - Response: list of accounts for customer
 - PUT `/api/accounts/:id`
   - Body: `{ "status": "active|suspended", "type": "savings" }`
   - Response: updated `Account`
 - DELETE `/api/accounts/:id`
   - Response: `{ "message": "account deleted" }`

Transactions
- POST `/api/accounts/deposit`
  - Body: `{ "account_id": 1, "amount": 500.0 }`
  - Behavior: transactionally updates account balance and creates `Transaction` record
  - Response: `{ "message": "deposit successful" }`
- POST `/api/accounts/withdraw`
  - Body: `{ "account_id": 1, "amount": 100.0 }`
  - Behavior: prevents overdraft; updates balance and creates `Transaction`
  - Response: `{ "message": "withdraw successful" }`
- GET `/api/accounts/:id/transactions`
  - Response: list of `Transaction` entries for the account

Loans
- POST `/api/loans/apply`
  - Body: `{ "customer_id":1, "branch_id":1, "principal_amount":10000.0, "duration_years":1 }`
  - Behavior: calculates `TotalInterest` = P × 12% × years, sets `RemainingAmount = P + TotalInterest` and stores loan
  - Response: created `Loan`
- GET `/api/loans/customer/:customerId`
  - Response: list of loans for the customer (with payments)
- GET `/api/loans/:loanId`
  - Response: loan detail (with payments)
 - PUT `/api/loans/:loanId`
   - Body: `{ "status": "closed" }` (only allowed updates; marking closed sets `EndDate` when fully repaid)
   - Response: updated `Loan`
 - DELETE `/api/loans/:loanId`
   - Behavior: only allowed when `remaining_amount == 0`
   - Response: `{ "message": "loan deleted" }`

Loan Payments
- POST `/api/loans/repay`
  - Body: `{ "loan_id": 1, "amount": 2000.0 }`
  - Behavior: transactionally reduces `RemainingAmount`, increments `AmountPaid`, stores `LoanPayment`, marks `Status` "closed" and sets `EndDate` when fully repaid
  - Response: `{ "message": "repayment successful" }`

Reporting
- GET `/api/loans/:loanId/interest-this-year`
  - Response: `{ "interest_this_year": <amount> }` — calculates pro-rated interest for up to one year since `StartDate`
- GET `/api/loans/:loanId/pending-amount`
  - Response: `{ "pending_amount": <amount> }`

---

## Sample curl Requests

Create a bank:

```bash
curl -X POST http://localhost:8080/api/banks \
  -H "Content-Type: application/json" \
  -d '{"name":"Acme Bank","code":"ACME001"}'
```

Open an account:

```bash
curl -X POST http://localhost:8080/api/accounts/open \
  -H "Content-Type: application/json" \
  -d '{"customer_id":1,"branch_id":1}'
```

Deposit money:

```bash
curl -X POST http://localhost:8080/api/accounts/deposit \
  -H "Content-Type: application/json" \
  -d '{"account_id":1,"amount":500.0}'
```

Apply for a loan:

```bash
curl -X POST http://localhost:8080/api/loans/apply \
  -H "Content-Type: application/json" \
  -d '{"customer_id":1,"branch_id":1,"principal_amount":10000.0,"duration_years":1}'
```

Repay a loan:

```bash
curl -X POST http://localhost:8080/api/loans/repay \
  -H "Content-Type: application/json" \
  -d '{"loan_id":1,"amount":2000.0}'
```

---

## SQL Schema (generated by GORM) — simplified examples

GORM `AutoMigrate` will create tables with columns and constraints similar to the following simplified SQL for PostgreSQL:

```sql
CREATE TABLE banks (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  code varchar(100) UNIQUE NOT NULL,
  created_at timestamptz
);

CREATE TABLE branches (
  id serial PRIMARY KEY,
  bank_id integer REFERENCES banks(id) ON DELETE CASCADE,
  name varchar(255) NOT NULL,
  ifsc_code varchar(100) UNIQUE NOT NULL,
  address text,
  created_at timestamptz
);

CREATE TABLE customers (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  phone varchar(50),
  address text,
  created_at timestamptz
);

CREATE TABLE accounts (
  id serial PRIMARY KEY,
  customer_id integer REFERENCES customers(id) ON DELETE CASCADE,
  branch_id integer REFERENCES branches(id) ON DELETE CASCADE,
  account_number varchar(100) UNIQUE NOT NULL,
  type varchar(50) NOT NULL,
  balance numeric DEFAULT 0,
  status varchar(50) DEFAULT 'active',
  created_at timestamptz
);

CREATE TABLE transactions (
  id serial PRIMARY KEY,
  account_id integer REFERENCES accounts(id) ON DELETE CASCADE,
  type varchar(50) NOT NULL,
  amount numeric NOT NULL,
  created_at timestamptz
);

CREATE TABLE loans (
  id serial PRIMARY KEY,
  customer_id integer REFERENCES customers(id) ON DELETE CASCADE,
  branch_id integer REFERENCES branches(id) ON DELETE CASCADE,
  principal_amount numeric NOT NULL,
  interest_rate numeric DEFAULT 12,
  total_interest numeric DEFAULT 0,
  amount_paid numeric DEFAULT 0,
  remaining_amount numeric DEFAULT 0,
  start_date timestamptz,
  end_date timestamptz,
  status varchar(50) DEFAULT 'active',
  created_at timestamptz
);

CREATE TABLE loan_payments (
  id serial PRIMARY KEY,
  loan_id integer REFERENCES loans(id) ON DELETE CASCADE,
  amount numeric NOT NULL,
  payment_date timestamptz,
  created_at timestamptz
);
```

---

## Validation and Errors

- Handlers use Gin's binding and return `400 Bad Request` when required fields are missing or invalid.
- Server and DB errors return `500 Internal Server Error` with an `error` message.
- Withdrawals return `400 Bad Request` with `insufficient balance` when funds are not available.
