## Register User

**Endpoint:** `POST /register`

**Description:** This endpoint allows a new user to register by providing their details. If the registration is successful, an activation email is sent to the provided email address.

**Request Body:**
```json
{
  "name": "string",
  "username": "string",
  "email": "string",
  "password": "string"
}
```

**Responses:**

  - 201 Created: User successfully registered, activation email sent.
  - 400 Bad Request: Invalid input or user already exists.
  - 500 Internal Server Error: An error occurred while processing the request.

## User Login

**Endpoint:** `POST /login`

**Description:** This endpoint allows a registered user to log in by providing their email and password. Upon successful authentication, a token is generated and returned for the user to access protected routes.

**Request Body:**
```json
{
  "email": "string",
  "password": "string"
}
```

**Responses**

  - 200 OK: Login successful, token returned.
  - 400 Bad Request: Invalid input.
  - 401 Unauthorized: Incorrect email or password.
  - 500 Internal Server Error: An error occurred while processing the request.

## Get User Profile

**Endpoint:** `GET /profile`

**Description:** This endpoint retrieves the profile information of the authenticated user. The user must be logged in, and the request must include a valid token.

**Authentication:** This endpoint requires authentication via a token.

**Response:**
- **200 OK:** Returns the user's profile data.
- **401 Unauthorized:** Authentication failed or token is missing/invalid.
- **500 Internal Server Error:** An error occurred while processing the request.

**Example Request:**
```bash
curl -X GET "http://localhost:8000/profile" \
-H "Authorization: Bearer <your_token>"


```

## Activate User Account

**Endpoint:** `GET /activate/:userID/:token`

**Description:** This endpoint is used to activate a user's account. The user needs to click on the activation link sent to their email, which includes the `userID` and `token` as URL parameters.

**Parameters:**
- **userID**: The unique identifier of the user.
- **token**: The activation token sent to the user's email.

**Response:**
- **200 OK:** The account is successfully activated.
- **400 Bad Request:** The activation link is invalid or expired.
- **500 Internal Server Error:** An error occurred while processing the request.

**Example Request:**
```bash
curl -X GET "http://localhost:8000/activate/{userID}/{token}"
```


## Forget Password

**Endpoint:** `GET /forget-password`

**Description:** This endpoint is used to initiate the password reset process. When a user requests to reset their password, they will receive an email with a link to reset it.

**Parameters:** None

**Response:**
- **200 OK:** Password reset email has been sent.
- **400 Bad Request:** Email is not associated with any account.
- **500 Internal Server Error:** An error occurred while processing the request.

**Example Request:**
```bash
curl -X GET "http://localhost:8000/forget-password"
```

## Reset Password

**Endpoint:** `PUT /reset/:userid/:tokentime/:token`

**Description:** This endpoint allows a user to reset their password by providing their user ID, the token time, and the reset token. It validates the token and allows the user to set a new password.

**Parameters:**
- `userid` (path): The unique identifier of the user.
- `tokentime` (path): The timestamp of when the token was generated.
- `token` (path): The reset token sent to the user's email.

**Request Body:**
- `password` (string, required): The new password to be set.

**Response:**
- **200 OK:** Password successfully reset.
- **400 Bad Request:** Invalid token or token expired.
- **500 Internal Server Error:** An error occurred during the password reset process.

**Example Request:**
```bash
curl -X PUT "http://localhost:8000/reset/12345/1623072021/sometoken" \
-H "Content-Type: application/json" \
-d '{"password": "newpassword123"}'
```
## Get All Users

**Endpoint:** `GET /all-users`

**Description:** Retrieves a list of all users in the system. This endpoint is protected and can only be accessed by admin users.

**Middleware:**
- `AuthMiddleware()`: Ensures the user is authenticated.
- `AdminMiddleware()`: Ensures the user has admin privileges.

**Response:**
- **200 OK:** Returns a list of users.
- **401 Unauthorized:** If the user is not authenticated.
- **403 Forbidden:** If the user is not an admin.
- **500 Internal Server Error:** If an error occurs while retrieving users.

**Example Request:**
```bash
curl -X GET "http://localhost:8000/all-users" \
-H "Authorization: Bearer your_token_here"
```

## Delete User Endpoint

### Method
`GET`

### Endpoint
`/delete/:id`

### Description
This endpoint allows an authenticated admin user to delete a specific user from the system. The user to be deleted is identified by the `id` provided in the URL path.

### Path Parameters
- **`id`** (string, required): The unique identifier of the user to be deleted.

### Middleware
- **`AuthMiddleware()`**: Ensures the request is made by an authenticated user. If the user is not authenticated, access is denied.
- **`AdminMiddleware()`**: Ensures the authenticated user has administrative privileges. If the user is not an admin, access is denied.

### Request
- **Method**: `GET`
- **URL**: `/delete/:id`

### Responses

#### Success
- **HTTP 200 OK**
- **Content**: A JSON response confirming the successful deletion of the user.
- **Example**:
```json
  {
    "message": "User successfully deleted."
  }
```
### Endpoint: Apply for Loan

- **URL**: `/`
- **Method**: `POST`
- **Middleware**: `AuthMiddleware`

#### Description
This endpoint allows an authenticated user to apply for a loan. The user must be authenticated using a JWT token, which is validated by the `AuthMiddleware`. The loan application is processed and saved in the database.

#### Request Headers
- **Authorization**: `Bearer <token>`
  - A valid JWT token is required to authenticate the user.

#### Request Body
The body should be a JSON object representing the loan application:

```json
{
  "amount": <loan_amount>
}
```

### Endpoint: View Loan Status

- **URL**: `/loan-status/:loanid`
- **Method**: `GET`
- **Middleware**: `AuthMiddleware`

#### Description
This endpoint allows an authenticated user to view the status of a specific loan. The user must provide a valid loan ID and be authenticated using a JWT token, which is validated by the `AuthMiddleware`.

#### Request Headers
- **Authorization**: `Bearer <token>`
  - A valid JWT token is required to authenticate the user.

#### URL Parameters
- **loanid** (string): The unique identifier of the loan whose status you want to view.

#### Responses

- **200 OK**: The loan status was successfully retrieved.
  - **Response Body**:
    ```json
    {
      "status": "approved"
    }
    ```
  - The status can be "pending", "approved", or "rejected".

- **401 Unauthorized**: The user is not authenticated or the token is invalid.
  - **Response Body**:
    ```json
    {
      "message": "authorization header is missing"
    }
    ```
    
    Or
    
    ```json
    {
      "message": "unauthorized access for user"
    }
    ```

- **404 Not Found**: The loan with the specified ID does not exist.
  - **Response Body**:
    ```json
    {
      "message": "loan not found"
    }
    ```

#### Example

**Request**:
```bash
curl -X GET http://localhost:8080/loan-status/12345 \
-H "Authorization: Bearer <token>"
```

### Endpoint: View All Loans

- **URL**: `/all-loans`
- **Method**: `GET`
- **Middleware**: `AuthMiddleware`, `AdminMiddleware`

#### Description
This endpoint allows an authenticated admin user to view all loans associated with a specific user. The user must be authenticated using a JWT token, which is validated by the `AuthMiddleware`, and must have admin privileges, verified by the `AdminMiddleware`.

#### Request Headers
- **Authorization**: `Bearer <token>`
  - A valid JWT token is required to authenticate the user.

#### Responses

- **200 OK**: A list of loans was successfully retrieved.
  - **Response Body**:
    ```json
    [
      {
        "loanid": "12345",
        "userid": "user123",
        "amount": 50000,
        "status": "approved",
        "createdat": "2023-08-29T14:00:00Z",
        "updatedat": "2023-08-30T10:00:00Z"
      },
      {
        "loanid": "67890",
        "userid": "user456",
        "amount": 25000,
        "status": "pending",
        "createdat": "2023-08-28T14:00:00Z",
        "updatedat": "2023-08-29T10:00:00Z"
      }
    ]
    ```

- **401 Unauthorized**: The user is not authenticated or the token is invalid.
  - **Response Body**:
    ```json
    {
      "message": "authorization header is missing"
    }
    ```
    
    Or
    
    ```json
    {
      "message": "unauthorized access for user"
    }
    ```

- **403 Forbidden**: The user does not have admin privileges.
  - **Response Body**:
    ```json
    {
      "message": "unauthorized access of admin"
    }
    ```

#### Example

**Request**:
```bash
curl -X GET http://localhost:8080/all-loans \
-H "Authorization: Bearer <token>"
```

### Endpoint: Approve Loan Status

- **URL**: `/approve-status/:loanid`
- **Method**: `PUT`
- **Middleware**: `AuthMiddleware`

#### Description
This endpoint allows an authenticated user to approve the status of a specific loan. The user must provide a valid loan ID and be authenticated using a JWT token, which is validated by the `AuthMiddleware`.

#### Request Headers
- **Authorization**: `Bearer <token>`
  - A valid JWT token is required to authenticate the user.

#### URL Parameters
- **loanid** (string): The unique identifier of the loan whose status you want to approve.

#### Responses

- **200 OK**: The loan status was successfully approved.
  - **Response Body**:
    ```json
    {
      "message": "Loan status approved"
    }
    ```

- **401 Unauthorized**: The user is not authenticated or the token is invalid.
  - **Response Body**:
    ```json
    {
      "message": "authorization header is missing"
    }
    ```
    
    Or
    
    ```json
    {
      "message": "unauthorized access for user"
    }
    ```

- **404 Not Found**: The loan with the specified ID does not exist.
  - **Response Body**:
    ```json
    {
      "message": "loan not found"
    }
    ```

#### Example

**Request**:
```bash
curl -X PUT http://localhost:8080/approve-status/12345 \
-H "Authorization: Bearer <token>"

### Endpoint: Reject Loan Status

- **URL**: `/reject-status/:loanid`
- **Method**: `PUT`
- **Middleware**: `AuthMiddleware`

#### Description
This endpoint allows an authenticated user to reject the status of a specific loan. The user must provide a valid loan ID and be authenticated using a JWT token, which is validated by the `AuthMiddleware`.

#### Request Headers
- **Authorization**: `Bearer <token>`
  - A valid JWT token is required to authenticate the user.

#### URL Parameters
- **loanid** (string): The unique identifier of the loan whose status you want to reject.

#### Responses

- **200 OK**: The loan status was successfully rejected.
  - **Response Body**:
    ```json
    {
      "message": "Loan status rejected"
    }
    ```

- **401 Unauthorized**: The user is not authenticated or the token is invalid.
  - **Response Body**:
    ```json
    {
      "message": "authorization header is missing"
    }
    ```
    
    Or
    
    ```json
    {
      "message": "unauthorized access for user"
    }
    ```

- **404 Not Found**: The loan with the specified ID does not exist.
  - **Response Body**:
    ```json
    {
      "message": "loan not found"
    }
    ```

#### Example

**Request**:
```bash
curl -X PUT http://localhost:8080/reject-status/12345 \
-H "Authorization: Bearer <token>"
