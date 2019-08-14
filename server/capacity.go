// Copyright 2019 RHK Development <dev@rosskeen.house>. All rights reserved.
package capacity;

import (
  "github.com/rosskeenhouse/capacity/tolerancemodel"
)

// Cacpacity server

type Server struct {
  model ToleranceModel
  probes Probes
  
  // Services
  rest RestServer
  grpc GrpcServer
}


func New() *Server {
  return Server{}
}

func (s *Server) Start() {
  s.rest.Accept()
  s.grpc.Accept()
}


