package snakeai

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 贪吃蛇AI配置
type Config struct {
	Speed float64 // 移动速度（步/秒）
	FPS   int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Speed: 5.0, // 5步/秒
		FPS:   30,
	}
}

// Point 坐标点
type Point struct {
	X, Y int
}

// Snake 蛇
type Snake struct {
	body      []Point
	direction Point
}

// SnakeAI 贪吃蛇AI特效
type SnakeAI struct {
	screen     tcell.Screen
	config     *Config
	snake      *Snake
	food       Point
	boardW     int
	boardH     int
	moveTimer  float64
	lastUpdate time.Time
	rand       *rand.Rand
	score      int
}

// New 创建贪吃蛇AI特效实例
func New(screen tcell.Screen, config *Config) *SnakeAI {
	if config == nil {
		config = DefaultConfig()
	}

	return &SnakeAI{
		screen: screen,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化贪吃蛇AI
func (s *SnakeAI) Init() error {
	width, height := s.screen.Size()
	s.boardW = width / 2
	s.boardH = height - 2

	s.reset()
	s.lastUpdate = time.Now()
	return nil
}

// reset 重置游戏
func (s *SnakeAI) reset() {
	// 初始化蛇（中间位置，长度3）
	centerX := s.boardW / 2
	centerY := s.boardH / 2

	s.snake = &Snake{
		body: []Point{
			{centerX, centerY},
			{centerX - 1, centerY},
			{centerX - 2, centerY},
		},
		direction: Point{1, 0}, // 向右
	}

	s.spawnFood()
	s.moveTimer = 0
	s.score = 0
}

// spawnFood 生成食物
func (s *SnakeAI) spawnFood() {
	// 找一个不在蛇身上的位置
	for {
		s.food = Point{
			X: s.rand.Intn(s.boardW),
			Y: s.rand.Intn(s.boardH),
		}

		// 检查食物是否在蛇身上
		onSnake := false
		for _, p := range s.snake.body {
			if p.X == s.food.X && p.Y == s.food.Y {
				onSnake = true
				break
			}
		}

		if !onSnake {
			break
		}
	}
}

// isSafe 检查位置是否安全
func (s *SnakeAI) isSafe(p Point) bool {
	// 边界检测
	if p.X < 0 || p.X >= s.boardW || p.Y < 0 || p.Y >= s.boardH {
		return false
	}

	// 蛇身碰撞检测（不检查尾巴，因为尾巴会移动）
	for i := 0; i < len(s.snake.body)-1; i++ {
		if s.snake.body[i].X == p.X && s.snake.body[i].Y == p.Y {
			return false
		}
	}

	return true
}

// findPath BFS寻找到食物的路径
func (s *SnakeAI) findPath() Point {
	// BFS从蛇头到食物
	start := s.snake.body[0]
	queue := []Point{start}
	visited := make(map[Point]bool)
	parent := make(map[Point]Point)

	visited[start] = true

	directions := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// 找到食物
		if current.X == s.food.X && current.Y == s.food.Y {
			// 回溯找第一步
			path := []Point{}
			for current != start {
				path = append([]Point{current}, path...)
				current = parent[current]
			}

			if len(path) > 0 {
				// 返回第一步的方向
				return Point{
					X: path[0].X - start.X,
					Y: path[0].Y - start.Y,
				}
			}
		}

		// 遍历4个方向
		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}

			if !visited[next] && s.isSafe(next) {
				visited[next] = true
				parent[next] = current
				queue = append(queue, next)
			}
		}
	}

	// 找不到路径，随机移动到安全的方向
	return s.getRandomSafeDirection()
}

// getRandomSafeDirection 获取随机安全方向
func (s *SnakeAI) getRandomSafeDirection() Point {
	directions := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// 过滤掉反方向
	safeDirections := []Point{}
	for _, dir := range directions {
		// 不能向反方向移动
		if dir.X == -s.snake.direction.X && dir.Y == -s.snake.direction.Y {
			continue
		}

		next := Point{
			X: s.snake.body[0].X + dir.X,
			Y: s.snake.body[0].Y + dir.Y,
		}

		if s.isSafe(next) {
			safeDirections = append(safeDirections, dir)
		}
	}

	if len(safeDirections) > 0 {
		return safeDirections[s.rand.Intn(len(safeDirections))]
	}

	// 没有安全方向，保持当前方向
	return s.snake.direction
}

// Update 更新贪吃蛇AI状态
func (s *SnakeAI) Update(deltaTime float64) {
	s.moveTimer += deltaTime

	if s.moveTimer >= 1.0/s.config.Speed {
		s.moveTimer = 0

		// AI决策
		nextDir := s.findPath()
		s.snake.direction = nextDir

		// 移动蛇
		newHead := Point{
			X: s.snake.body[0].X + s.snake.direction.X,
			Y: s.snake.body[0].Y + s.snake.direction.Y,
		}

		// 碰撞检测
		if !s.isSafe(newHead) {
			s.reset() // 死亡重启
			return
		}

		// 吃食物
		if newHead.X == s.food.X && newHead.Y == s.food.Y {
			s.snake.body = append([]Point{newHead}, s.snake.body...)
			s.spawnFood()
			s.score++
		} else {
			// 正常移动
			s.snake.body = append([]Point{newHead}, s.snake.body[:len(s.snake.body)-1]...)
		}
	}
}

// Render 渲染贪吃蛇AI
func (s *SnakeAI) Render() {
	s.screen.Clear()

	width, height := s.screen.Size()
	offsetX := (width - s.boardW*2) / 2
	offsetY := (height - s.boardH) / 2

	// 绘制边框
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	for y := 0; y <= s.boardH+1; y++ {
		if y == 0 || y == s.boardH+1 {
			for x := 0; x <= s.boardW*2+1; x++ {
				s.screen.SetContent(offsetX+x-1, offsetY+y-1, '─', nil, borderStyle)
			}
		} else {
			s.screen.SetContent(offsetX-1, offsetY+y-1, '│', nil, borderStyle)
			s.screen.SetContent(offsetX+s.boardW*2, offsetY+y-1, '│', nil, borderStyle)
		}
	}

	// 绘制蛇
	for i, p := range s.snake.body {
		var char rune
		var color tcell.Color

		if i == 0 {
			// 蛇头
			char = '◉'
			color = tcell.ColorYellow
		} else {
			// 蛇身
			char = '●'
			color = tcell.ColorGreen
		}

		style := tcell.StyleDefault.Foreground(color).Bold(true)
		s.screen.SetContent(offsetX+p.X*2, offsetY+p.Y, char, nil, style)
		s.screen.SetContent(offsetX+p.X*2+1, offsetY+p.Y, ' ', nil, style)
	}

	// 绘制食物
	foodStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true)
	s.screen.SetContent(offsetX+s.food.X*2, offsetY+s.food.Y, '♥', nil, foodStyle)
	s.screen.SetContent(offsetX+s.food.X*2+1, offsetY+s.food.Y, ' ', nil, foodStyle)

	s.screen.Show()
}

// Run 运行贪吃蛇AI特效
func (s *SnakeAI) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(s.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(s.lastUpdate).Seconds()
			s.lastUpdate = now

			s.Update(deltaTime)
			s.Render()
		}
	}
}

// Cleanup 清理资源
func (s *SnakeAI) Cleanup() error {
	return nil
}
