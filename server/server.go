package main

import (
	pb "chatpb"
)

func main() {
	
}

func (s *pb.Server) CreateStream(pconn *pb.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
	   stream: stream,
	   id: pconn.User.Id,
	   active: true,
	   err : make(chan error),
	}
	s.Connection = append(s.Connection, conn)
 
	return <-conn.err
 }
