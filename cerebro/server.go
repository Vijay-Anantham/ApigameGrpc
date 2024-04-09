package cerebro

// TODO: All code for server governance
// TODO: 1. generate random score roll values
// TODO: 2. Update score pool
// TODO: 3. update alloted number
// INFO: Scoreboard: This marks the final data board for updating and tracking scores

// Server struct is central brain of the system
type server struct {
	Scoreboard map[string]interface{}
}

// Function initialise connection with server and kick start RPC connections
func (c *server) InitServer() {
}

// This is an random number generator from the pool
func (c *server) Makeroll() {
}

func (c *server) isAlloted() {

}
func (c *server) UpdateScores() {

}
