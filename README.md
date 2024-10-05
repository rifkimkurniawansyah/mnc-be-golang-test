# Golang API for Merchant & Bank

## Description

This project is a simple API that connects customers with merchants for payment transactions. The API provides login, payment, and logout functionalities while logging all activities to a JSON file as a data storage simulation.

## Features

- **Login**: Authenticates users and verifies their existence.
- **Payment**: Allows logged-in customers to make payments without any transfer amount restrictions.
- **Logout**: Clears the customer's login status.
- **Logging**: Logs all activities to a JSON file.


## Prerequisites

Make sure you have [Go](https://golang.org/dl/) installed on your system.

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/rifkimkurniawansyah/mnc-be-golang-test
    ```

2. Navigate to the project directory:
    ```bash
    cd golang-api-test
    ```

## Running the Application

1. Run the application using the command:
    ```bash
    go run main.go
    ```

2. The API will run at `http://localhost:8080`.

## API Usage

### 1. Login

- **Endpoint**: `/login`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "username": "customer1",
        "password": "password1"
    }
    ```

- **Response**:
    - 200 OK: If login is successful.
    - 401 Unauthorized: If login fails.

### 2. Payment

- **Endpoint**: `/payment`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "username": "customer1",
        "amount": 100.50
    }
    ```

- **Response**:
    - 200 OK: If the payment is successful.
    - 401 Unauthorized: If the user is not logged in.

### 3. Logout

- **Endpoint**: `/logout`
- **Method**: `POST`
- **Body**:
    ```json
    {
        "username": "customer1"
    }
    ```

- **Response**:
    - 200 OK: If logout is successful.

## JSON Files

- **customers.json**: Stores customer data.


