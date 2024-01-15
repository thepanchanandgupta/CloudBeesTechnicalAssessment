# CloudBeesTechnicalAssessment
# Train Ticket Management System

This is a simple train ticket management system implemented in GoLang using gRPC and Protobuf.

## Getting Started

To get started with this project, follow the steps below:

### Prerequisites

Make sure you have the following installed:a

- Go 
- Protocol Buffers compiler (`protoc`)
- gRPC for Go
- To install other dependencies `go mod tidy`
- Run using `go run main.go`
- Server will start on port `8089`
- [grpcurl](https://github.com/fullstorydev/grpcurl) cli interface to interact with grpc api's


Features: 
- Purchase Train Ticket by proving User's `first_name`, `last_name`, and `email`
- SeatType Allocation randomly : 'A' or 'B'
- View Receipts
- View Receipts and User's Details
- Remove User


Generate Client Code using `protoc -I proto/ proto/assessment.proto --go_out=plugins=grpc:proto`

Makefile has been attached for easy accessibility.
- To run the project, run `make start_server`
- To call different APIs, use the respective commands:
    1. To Book Tickets: `	grpcurl -plaintext -d '{"from": "london", "to": "france","first_name": "f3","last_name": "l3","email": "email3"}' localhost:8089 assignment.Assignment/BookTicket`    {Enter details as per requirement}
 
    2. To Show All Booked Tickets: `grpcurl -plaintext localhost:8089 assignment.Assignment/ShowAllTickets`
 
    3. To Show All Booked Tickets of SectionA:  `grpcurl -plaintext -d '{"seat": "A"}' localhost:8089 assignment.Assignment/TicketsByType`
 
    4. To Show All Booked Tickets of SectionB:  `grpcurl -plaintext -d '{"seat": "B"}' localhost:8089 assignment.Assignment/TicketsByType`
 
    5. To Remove a user: `grpcurl -plaintext -d '{"id": "1"}' localhost:8089 assignment.Assignment/RemoveUser`
