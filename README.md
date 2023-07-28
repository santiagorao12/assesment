Transport Management System (TMS) Backend API

This repository contains the backend API for the Transport Management System (TMS), built using Golang and the Gin framework.
Getting Started

Follow these steps to set up and run the TMS backend API on your local machine:
Prerequisites

Before proceeding, ensure you have the following installed on your machine:

    Go programming language: Install Go
    Git: Install Git

Installation

    Clone the repository to your local machine using Git:

bash

git clone https://github.com/your-username/tms-backend.git

    Change directory to the project folder:

bash

cd tms-backend

    Install the necessary dependencies:

bash

go mod download

Running the TMS Backend API

Start the TMS backend API server with the following command:

bash

go run main.go

The server should now be running at http://localhost:8080.
Question 1: Create a New Docket

To create a new docket, send a POST request to the /docket endpoint with the required JSON payload. The API will generate a unique OrderNo and store the docket.

Example Request:

bash

curl -X POST -H "Content-Type: application/json" -d '{
  "Customer": "Lotus",
  "PickUpPoint": "Lotus Klang",
  "DeliveryPoint": "Lotus Kelantan",
  "Quantity": 13,
  "Volume": 2.34
}' http://localhost:8080/docket

Example Response:

json

{
  "OrderNo": "TDN0001",
  "Customer": "Lotus",
  "PickUpPoint": "Lotus Klang",
  "DeliveryPoint": "Lotus Kelantan",
  "Quantity": 13,
  "Volume": 2.34,
  "Status": "Created",
  "TruckNo": "",
  "LogsheetNo": ""
}

Question 2: Fetch a Docket

To fetch a docket based on its OrderNo, send a GET request to the /docket/{orderNo} endpoint.

Example Request:

bash

curl http://localhost:8080/docket/TDN0001

Example Response:

json

{
  "OrderNo": "TDN0001",
  "Customer": "Lotus",
  "PickUpPoint": "Lotus Klang",
  "DeliveryPoint": "Lotus Kelantan",
  "Quantity": 13,
  "Volume": 2.34,
  "Status": "Created",
  "TruckNo": "",
  "LogsheetNo": ""
}

Question 3: Fetch a List of Dockets

To fetch a list of all created dockets, send a GET request to the /docket endpoint.

Example Request:

bash

curl http://localhost:8080/docket

Example Response:

json

[
  {
    "OrderNo": "TDN0001",
    "Customer": "Lotus",
    "PickUpPoint": "Lotus Klang",
    "DeliveryPoint": "Lotus Kelantan",
    "Quantity": 13,
    "Volume": 2.34,
    "Status": "Created",
    "TruckNo": "",
    "LogsheetNo": ""
  },
  {
    "OrderNo": "TDN0002",
    "Customer": "AnotherCustomer",
    "PickUpPoint": "PickUpPoint2",
    "DeliveryPoint": "DeliveryPoint2",
    "Quantity": 10,
    "Volume": 1.5,
    "Status": "Created",
    "TruckNo": "",
    "LogsheetNo": ""
  }
]

Question 4: Create a New Logsheet

To create a new logsheet and update the dockets accordingly, send a POST request to the /logsheet endpoint with the required JSON payload.

Example Request:

bash

curl -X POST -H "Content-Type: application/json" -d '{
  "Dockets": ["TDN0001", "TDN0002"],
  "TruckNo": "BPR1234"
}' http://localhost:8080/logsheet

Example Response:

json

{
  "LogsheetNo": "DT0001",
  "Dockets": ["TDN0001", "TDN0002"],
  "TruckNo": "BPR1234"
}

Question 5: Fetch a Logsheet

To fetch a logsheet based on its LogsheetNo, send a GET request to the /logsheet/{logsheetNo} endpoint.

Example Request:

bash

curl http://localhost:8080/logsheet/DT0001

Example Response:

json

{
  "LogsheetNo": "DT0001",
  "Dockets": ["TDN0001", "TDN0002"],
  "TruckNo": "BPR1234"
}
