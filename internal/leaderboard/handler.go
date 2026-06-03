package leaderboard

import (
	"slices"
	"sync"

	"github.com/labstack/echo/v5"
)

type LeaderboardHandler struct {
	Board *Leaderboard
	mu    *sync.RWMutex
}

func NewHandler() *LeaderboardHandler {
	return &LeaderboardHandler{
		Board: NewLeaderboard(),
		mu:    &sync.RWMutex{},
	}
}

func (h *LeaderboardHandler) Setup(g *echo.Group) {
	sub := g.Group("/leaderboard")
	sub.GET("", h.leaderboard)
	sub.POST("", h.AddEntry)
}

type Rank struct {
	Rank  int    `json:"rank"`
	Score int    `json:"score"`
	Name  string `json:"name"`
}

type LeaderBoardResponse struct {
	Items []Rank `json:"items"`
}

func (h *LeaderboardHandler) leaderboard(c *echo.Context) error {
	h.mu.RLock()
	cpyBoard := make([]LeaderboardItems, len(h.Board.Items))
	copy(cpyBoard, h.Board.Items)
	h.mu.RUnlock()
	c.Logger().Info("", map[string]interface{}{"Data": cpyBoard})
	slices.SortFunc(cpyBoard, func(a, b LeaderboardItems) int {
		if a.Score > b.Score {
			return 1
		}
		if b.Score > a.Score {
			return -1
		}
		return 0
	})

	resp := LeaderBoardResponse{Items: make([]Rank, 0)}
	for i, it := range cpyBoard {
		resp.Items = append(resp.Items, Rank{
			Rank:  i + 1,
			Score: it.Score,
			Name:  it.Name,
		})
	}
	return c.JSON(200, resp)
}

type AddEntryBody struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (h *LeaderboardHandler) AddEntry(c *echo.Context) error {
	var body AddEntryBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "cant add entry",
			"cause": err,
			"code":  400,
		})
	}

	h.mu.Lock()

	h.Board.Items = append(h.Board.Items, LeaderboardItems{
		Name:  body.Name,
		Score: body.Score,
	})
	h.mu.Unlock()

	c.Logger().Info("%v", h.Board.Items)

	return c.JSON(201, map[string]interface{}{
		"status": "success",
	})
}
