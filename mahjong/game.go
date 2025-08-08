package mahjong

import (
	"github.com/kevin-chtw/tw_game_svr/game"
)

type IGame interface {
	CreatePlay() IPlay
	OnStart()
	OnReqMsg(seat int32, data []byte)
}

type Game struct {
	IGame
	*game.Table
	timer       *Timer
	CurState    IState
	nextState   IState
	rule        *Rule
	players     []*Player
	increasedID int32   // 当前请求ID
	requestIDs  []int32 // 记录每个玩家的请求ID

	Play IPlay
}

func NewGame(subGame IGame, t *game.Table) *Game {
	g := &Game{
		IGame:       subGame,
		Table:       t,
		timer:       NewTimer(),
		rule:        NewRule(),
		players:     make([]*Player, t.GetPlayerCount()),
		increasedID: 1,
		requestIDs:  make([]int32, t.GetPlayerCount()),
	}
	g.rule.LoadRule(t.GetGameRule(), Service.GetDefaultRules())
	for i := int32(0); i < t.GetPlayerCount(); i++ {
		g.players[i] = NewPlayer(g, t.GetGamePlayer(i))
	}
	g.Play = subGame.CreatePlay()
	return g
}

func (g *Game) OnGameBegin() {
	g.IGame.OnStart()
	g.enterNextState()
}

func (g *Game) OnPlayerMsg(player *game.Player, data []byte) {
	seat := player.Seat
	if !g.IsValidSeat(seat) {
		return
	}

	g.IGame.OnReqMsg(seat, data)
	g.enterNextState()
}

func (g *Game) OnGameTimer() {
	g.timer.OnTick()
	g.enterNextState()
}

func (g *Game) GetTimer() *Timer {
	return g.timer
}

func (g *Game) GetRule() *Rule {
	return g.rule
}

func (g *Game) GetPlay() IPlay {
	return g.Play
}

func (g *Game) GetPlayer(seat int32) *Player {
	if g.IsValidSeat(seat) {
		return g.players[seat]
	}
	return nil
}

func (g *Game) SetNextState(newFn func(IGame, ...any) IState, args ...any) {
	g.nextState = newFn(g.IGame, args...)
}

func (g *Game) enterNextState() {
	for g.nextState != nil {
		g.CurState = g.nextState
		g.nextState = nil
		g.timer.Cancel()
		g.CurState.OnEnter()
	}
}

func (g *Game) GetRequestID(seat int32) int32 {
	g.increasedID++
	if g.IsValidSeat(seat) {
		g.requestIDs[seat] = g.increasedID
	} else {
		for i := range g.requestIDs {
			g.requestIDs[i] = g.increasedID
		}
	}
	return g.increasedID
}

func (g *Game) IsRequestID(seat, id int32) bool {
	if !g.IsValidSeat(seat) {
		return false
	}
	return g.requestIDs[seat] == id
}
