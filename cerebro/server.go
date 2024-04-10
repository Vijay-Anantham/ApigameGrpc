package cerebro

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "github.com/hellogrpc/game"
)

// TODO: All code for server governance
// TODO: 1. generate random score roll values
// TODO: 2. Update score pool
// TODO: 3. update alloted number
// INFO: Scoreboard: This marks the final data board for updating and tracking scores

// Server struct is central brain of the system
type Server struct {
	Scoreboard map[int32]Player
	alloted    map[int32]bool
}

// Player struct definition
type Player struct {
	Id      int32
	Name    string
	Score   int32
	alotted map[int32]interface{}
}

// Function initialise connection with server and kick start RPC connections
func (c *Server) InitServer() {

}

// This is an random number generator from the pool
func (c *Server) Makeroll(request *pb.RollRequest) (*pb.RollReply, error) {
	// Generate a random number
	var rolled = c.generateRoll()
	// Update the random number on allotment
	c.alloted[rolled] = true
	// Update the score on the scoreboard
	player, ok := c.Scoreboard[request.PlayerId]
	if !ok {
		return nil, fmt.Errorf("player with ID %d not found", request.PlayerId)
	}
	player.Score += rolled
	player.alotted[rolled] = true
	c.Scoreboard[request.PlayerId] = player

	var score, err = c.getScore(request.GetPlayerId())
	if err != nil {
		log.Printf("Error retreving score value of %s", request.GetName())
	}
	return &pb.RollReply{Rollvalue: 12, Score: score}, nil
}

func (c *Server) AddPlayer(player *pb.AddPlayerRequest) *pb.AddPlayerResponse {
	var playerId = c.generatePlayerId()
	var newbie = Player{
		Id:      playerId,
		Name:    player.GetName(),
		Score:   0,
		alotted: map[int32]interface{}{},
	}
	c.Scoreboard[playerId] = newbie
	return &pb.AddPlayerResponse{PlayerId: playerId}
}

func (c *Server) DeletePlayer(player *pb.DeletePlayerRequest) *pb.DeletePlayerResponse {
	var deleteplayer = player.GetPlayerId()
	// clean up the allotment maybe if needed
	delete(c.Scoreboard, deleteplayer)
	return &pb.DeletePlayerResponse{PlayerId: deleteplayer, StatusMessage: fmt.Sprintf("%d has been deleted", deleteplayer)}
}

// Function generates a random number
func (c *Server) generateRoll() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand := r.Int31n(30)
	if c.isAlloted(rand) {
		c.generateRoll()
	}
	return rand
}

func (c *Server) isAlloted(val int32) bool {
	return c.alloted[val]
}

// Function to update the latest code and update it on the `Scoreboard`
func (c *Server) getScore(playerId int32) (int32, error) {
	return 0, nil
}

// This checks the `Scoreboard` for the current score of the player
func (c *Server) getScoreRegistry(playerId int32) (int32, error) {
	return 0, nil
}
func (c *Server) generatePlayerId() int32 {
	var currentId = 0
	return int32(currentId + 1)
}

// Function updates score on the `scoreboard`
func (c *Server) updateScores(playerId int32) error {

	return nil
}
