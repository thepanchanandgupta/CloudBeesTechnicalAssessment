package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"

	"github.com/thepanchanandgupta/CloudBeesTechnicalAssessment/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// AssignmentServer ...
type AssignmentServer struct {
	proto.UnimplementedAssignmentServer
	mu         sync.Mutex
	AllTickets []*proto.Ticket
}

// func generateRandomSectionType return string a or b
func generateRandomSectionType() string {
	// golang is range exclusive means if max is 3 it will never give 3
	min := 1
	max := 3
	index := rand.Intn(max-min) + min

	if index == 1 {
		return "A"
	}
	return "B"
}

// BookTicket ..
// S1. Create API where you can submit a purchase for a ticket
func (s *AssignmentServer) BookTicket(ctx context.Context, req *proto.Ticket) (*proto.Ticket, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Assign a sequential ID
	req.Id = int32(len(s.AllTickets) + 1)

	//Assign Price
	req.Price = 20

	// Assign Seat Type
	// S2. Allot seat Section either A or B
	req.Seat = generateRandomSectionType()
	// Save the product in memory
	s.AllTickets = append(s.AllTickets, req)

	return req, nil
}

// ShowAllTickets ...
// S3 An API that shows the details of the receipt for the user
func (s *AssignmentServer) ShowAllTickets(req *proto.NoParam, stream proto.Assignment_ShowAllTicketsServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ticket := range s.AllTickets {
		if err := stream.Send(ticket); err != nil {
			return err
		}
	}

	return nil
}

// TicketsByType ...
// S4 An API that lets you view the users and seat they are allocated by the requested section
func (s *AssignmentServer) TicketsByType(req *proto.SeatSectionRequest, stream proto.Assignment_TicketsByTypeServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ticket := range s.AllTickets {
		if ticket.Seat == req.Seat {
			if err := stream.Send(ticket); err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveUser ...
// S5 An API to remove a user from the train
// func (s AssignmentServer) RemoveUser(ctx context.Context, req *proto.UserIDRequest) (*proto.StringResponse, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	var deleted bool
// 	for i, ticket := range s.AllTickets {
// 		id, _ := strconv.ParseInt(req.Id, 10, 32)
// 		if ticket.Id == int32(id) {
// 			s.AllTickets = append(s.AllTickets[:i], s.AllTickets[i+1:]...)
// 			deleted = true
// 			fmt.Printf("%+v", ticket)
// 			break
// 		}
// 	}
// 	if deleted {
// 		return &proto.StringResponse{Message: "Delete"}, nil
// 	}
// 	return &proto.StringResponse{Message: "UserID not found"}, nil
// }

func (s *AssignmentServer) RemoveUser(ctx context.Context, req *proto.UserIDRequest) (*proto.StringResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	idToRemove, err := strconv.ParseInt(req.Id, 10, 32)
	if err != nil {
		return nil, err
	}

	var deleted bool
	for i, ticket := range s.AllTickets {
		if ticket.Id == int32(idToRemove) {
			s.AllTickets = append(s.AllTickets[:i], s.AllTickets[i+1:]...)
			deleted = true
			fmt.Printf("Deleted Ticket: %+v\n", ticket)
			break
		}
	}

	if deleted {
		return &proto.StringResponse{Message: "Ticket deleted"}, nil
	}
	return &proto.StringResponse{Message: "Ticket not found"}, nil
}

// ModifyUser ...
// S6 An API to modify a user's details by user ID
func (s *AssignmentServer) ModifyUser(ctx context.Context, req *proto.UserIDRequest) (*proto.StringResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	idToModify, err := strconv.ParseInt(req.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	var modified bool

	for _, ticket := range s.AllTickets {
		if ticket.Id == int32(idToModify) {
			// Modify user details

			// ticket.first_name = req.Ticket.first_name
			// ticket.last_name = req.Ticket.last_name
			// ticket.Email = req.Ticket.Email

			ticket.Email = "macintoshOS@gmail.com"

			modified = true
			fmt.Printf("Modified Ticket: %+v\n", ticket)
			break
		}
	}

	if modified {
		return &proto.StringResponse{Message: "User details modified"}, nil
	}
	return &proto.StringResponse{Message: "User not found"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &AssignmentServer{}
	proto.RegisterAssignmentServer(serverRegistrar, service)
	reflection.Register(serverRegistrar)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
