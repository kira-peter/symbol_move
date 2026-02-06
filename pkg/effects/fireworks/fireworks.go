package fireworks

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Config 烟花配置
type Config struct {
	LaunchInterval    float64 // 发射间隔（秒）
	ParticlesPerBurst int     // 每次爆炸的粒子数
	Gravity           float64 // 重力加速度
	FPS               int     // 帧率
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		LaunchInterval:    1.5,  // 1.5秒发射一次
		ParticlesPerBurst: 50,   // 每次爆炸50个粒子
		Gravity:           20.0, // 重力
		FPS:               30,
	}
}

// Particle 粒子
type Particle struct {
	x, y   float64
	vx, vy float64
	life   float64     // 生命值（0-1）
	decay  float64     // 衰减速度
	color  tcell.Color
}

// Firework 烟花
type Firework struct {
	x, y      float64
	vy        float64 // 上升速度
	stage     int     // 0=上升, 1=爆炸
	targetY   float64 // 目标高度
	particles []*Particle
	color     tcell.Color
}

// Fireworks 烟花绽放特效
type Fireworks struct {
	screen         tcell.Screen
	config         *Config
	fireworks      []*Firework
	width          int
	height         int
	timeSinceLaunch float64
	lastUpdate     time.Time
	rand           *rand.Rand
}

// New 创建烟花绽放特效实例
func New(screen tcell.Screen, config *Config) *Fireworks {
	if config == nil {
		config = DefaultConfig()
	}

	return &Fireworks{
		screen:    screen,
		config:    config,
		fireworks: make([]*Firework, 0),
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Init 初始化烟花绽放
func (f *Fireworks) Init() error {
	f.width, f.height = f.screen.Size()
	f.timeSinceLaunch = 0
	f.lastUpdate = time.Now()
	return nil
}

// launch 发射新烟花
func (f *Fireworks) launch() {
	colors := []tcell.Color{
		tcell.ColorRed,
		tcell.ColorGreen,
		tcell.ColorBlue,
		tcell.ColorYellow,
		tcell.ColorPurple,
		tcell.ColorTeal,
	}

	fw := &Firework{
		x:       float64(f.rand.Intn(f.width)),
		y:       float64(f.height),
		vy:      -40.0 - f.rand.Float64()*20.0, // 上升速度
		stage:   0,
		targetY: float64(f.height/4) + f.rand.Float64()*float64(f.height/4),
		color:   colors[f.rand.Intn(len(colors))],
	}
	f.fireworks = append(f.fireworks, fw)
}

// explode 爆炸
func (fw *Firework) explode(config *Config) {
	fw.stage = 1
	fw.particles = make([]*Particle, config.ParticlesPerBurst)

	for i := 0; i < config.ParticlesPerBurst; i++ {
		angle := rand.Float64() * 2 * math.Pi
		speed := 10.0 + rand.Float64()*15.0

		fw.particles[i] = &Particle{
			x:     fw.x,
			y:     fw.y,
			vx:    speed * math.Cos(angle),
			vy:    speed * math.Sin(angle),
			life:  1.0,
			decay: 0.5 + rand.Float64()*0.5, // 0.5-1.0
			color: fw.color,
		}
	}
}

// Update 更新烟花绽放状态
func (f *Fireworks) Update(deltaTime float64) {
	// 发射新烟花
	f.timeSinceLaunch += deltaTime
	if f.timeSinceLaunch >= f.config.LaunchInterval {
		f.launch()
		f.timeSinceLaunch = 0
	}

	// 更新现有烟花
	activeFireworks := make([]*Firework, 0)
	for _, fw := range f.fireworks {
		if fw.stage == 0 {
			// 上升阶段
			fw.vy += f.config.Gravity * deltaTime
			fw.y += fw.vy * deltaTime

			// 检测是否到达目标高度
			if fw.y <= fw.targetY {
				fw.explode(f.config)
			}
			activeFireworks = append(activeFireworks, fw)
		} else {
			// 爆炸阶段 - 更新粒子
			hasAlive := false
			for _, p := range fw.particles {
				if p.life > 0 {
					p.vy += f.config.Gravity * deltaTime
					p.x += p.vx * deltaTime
					p.y += p.vy * deltaTime
					p.life -= p.decay * deltaTime
					hasAlive = true
				}
			}
			if hasAlive {
				activeFireworks = append(activeFireworks, fw)
			}
		}
	}
	f.fireworks = activeFireworks
}

// Render 渲染烟花绽放
func (f *Fireworks) Render() {
	f.screen.Clear()

	for _, fw := range f.fireworks {
		if fw.stage == 0 {
			// 绘制上升中的烟花
			x := int(fw.x)
			y := int(fw.y)
			if x >= 0 && x < f.width && y >= 0 && y < f.height {
				style := tcell.StyleDefault.Foreground(fw.color).Bold(true)
				f.screen.SetContent(x, y, '●', nil, style)
			}
		} else {
			// 绘制爆炸粒子
			for _, p := range fw.particles {
				if p.life <= 0 {
					continue
				}

				x := int(p.x)
				y := int(p.y)
				if x >= 0 && x < f.width && y >= 0 && y < f.height {
					// 根据生命值选择字符
					var char rune
					if p.life > 0.7 {
						char = '*'
					} else if p.life > 0.4 {
						char = '·'
					} else {
						char = '.'
					}

					style := tcell.StyleDefault.Foreground(p.color)
					f.screen.SetContent(x, y, char, nil, style)
				}
			}
		}
	}

	f.screen.Show()
}

// Run 运行烟花绽放特效
func (f *Fireworks) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(f.config.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(f.lastUpdate).Seconds()
			f.lastUpdate = now

			f.Update(deltaTime)
			f.Render()
		}
	}
}

// Cleanup 清理资源
func (f *Fireworks) Cleanup() error {
	return nil
}
