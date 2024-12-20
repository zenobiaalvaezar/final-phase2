# API Documentation

## Base URL
```
http://localhost:8080
```

---

## Authentication

### **Register**
**POST** `/api/v1/register`

**Request Body**
```json
{
  "email": "string",
  "password": "string"
}
```

**Response**
- **201 Created**
```json
{
  "message": "Registration successful"
}
```
- **400 Bad Request**
```json
{
  "message": "Error message"
}
```

### **Login**
**POST** `/api/v1/login`

**Request Body**
```json
{
  "email": "string",
  "password": "string"
}
```

**Response**
- **200 OK**
```json
{
  "token": "string",
  "user": {
    "id": "integer",
    "email": "string",
    "deposit_amount": "float"
  }
}
```
- **401 Unauthorized**
```json
{
  "message": "Invalid credentials"
}
```

---

## User

### **Get Profile**
**GET** `/api/v1/profile`

**Headers**
```
Authorization: Bearer <token>
```

**Response**
- **200 OK**
```json
{
  "id": "integer",
  "email": "string",
  "deposit_amount": "float",
  "created_at": "string"
}
```
- **404 Not Found**
```json
{
  "message": "User not found"
}
```

### **Top Up**
**POST** `/api/v1/topup`

**Headers**
```
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "amount": "float"
}
```

**Response**
- **200 OK**
```json
{
  "message": "Top up successful",
  "current_balance": "float"
}
```
- **400 Bad Request**
```json
{
  "message": "Error message"
}
```

---

## Cars

### **Get All Cars**
**GET** `/cars`

**Query Parameters**
- `category` (optional): Filter by car category.
- `available` (optional): Set to `true` to filter available cars.

**Response**
- **200 OK**
```json
{
  "data": [
    {
      "id": "integer",
      "name": "string",
      "stock_availability": "integer",
      "rental_costs": "float",
      "category": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ]
}
```

### **Get Car Detail**
**GET** `/cars/:id`

**Response**
- **200 OK**
```json
{
  "data": {
    "id": "integer",
    "name": "string",
    "stock_availability": "integer",
    "rental_costs": "float",
    "category": "string",
    "created_at": "string",
    "updated_at": "string"
  }
}
```
- **404 Not Found**
```json
{
  "message": "Car not found"
}
```

---

## Rentals

### **Create Rental**
**POST** `/api/v1/rentals`

**Headers**
```
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "car_id": "integer",
  "rental_start": "string", // Format: YYYY-MM-DD
  "rental_end": "string" // Format: YYYY-MM-DD
}
```

**Response**
- **201 Created**
```json
{
  "message": "Rental created, waiting for payment",
  "rental": {
    // Rental data
  },
  "payment": {
    "payment_url": "string",
    "amount": "float",
    "status": "string"
  }
}
```

### **Get User Rentals**
**GET** `/api/v1/rentals`

**Headers**
```
Authorization: Bearer <token>
```

**Response**
- **200 OK**
```json
{
  "data": [
    {
      // Rental history
    }
  ]
}
```

### **Return Car**
**POST** `/api/v1/rentals/:id/return`

**Headers**
```
Authorization: Bearer <token>
```

**Response**
- **200 OK**
```json
{
  "message": "Car returned successfully"
}
```
- **404 Not Found**
```json
{
  "message": "Rental not found"
}
```

---

## Payments

### **Get Payment History**
**GET** `/api/v1/payments`

**Headers**
```
Authorization: Bearer <token>
```

**Response**
- **200 OK**
```json
{
  "data": [
    {
      // Payment history
    }
  ]
}
```

### **Get Payment Detail**
**GET** `/api/v1/payments/:id`

**Headers**
```
Authorization: Bearer <token>
```

**Response**
- **200 OK**
```json
{
  "data": {
    // Payment detail
  }
}
```
- **404 Not Found**
```json
{
  "message": "Payment not found"
}
```

### **Webhook**
**POST** `/api/v1/payments/webhook`

**Request Body**
```json
{
  "external_id": "string",
  "status": "string",
  "amount": "float",
  "id": "string"
}
```

**Response**
- **200 OK**
```json
{
  "status": "success"
}
```
- **401 Unauthorized**
```json
{
  "message": "Invalid callback token"
}
```

