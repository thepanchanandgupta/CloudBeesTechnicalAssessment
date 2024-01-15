generate_proto_buf:
	protoc --go_out=. --go-grpc_out=. proto/assignment.proto
start_server:
	go run main.go
	
ListApi:
	grpcurl -plaintext localhost:8089 list

Book_Ticket_1:
	grpcurl -plaintext -d '{"from": "london", "to": "france","first_name": "f1","last_name": "l1","email": "email1"}' localhost:8089 assignment.Assignment/BookTicket

Book_Ticket_2:
	grpcurl -plaintext -d '{"from": "london", "to": "france","first_name": "f3","last_name": "l3","email": "email3"}' localhost:8089 assignment.Assignment/BookTicket

ShowSeatA:
	grpcurl -plaintext -d '{"seat": "A"}' localhost:8089 assignment.Assignment/TicketsByType

ShowSeatB:
	grpcurl -plaintext -d '{"seat": "B"}' localhost:8089 assignment.Assignment/TicketsByType

ShowAll:
	grpcurl -plaintext localhost:8089 assignment.Assignment/ShowAllTickets

Remove1:
	grpcurl -plaintext -d '{"id": "1"}' localhost:8089 assignment.Assignment/RemoveUser

Modify1:
	grpcurl -plaintext -d '{"id": "1"}' localhost:8089 assignment.Assignment/ModifyUser

