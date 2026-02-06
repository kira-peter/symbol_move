package particleburst

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Config struct {
	BurstInterval float64 // 爆炸间隔（秒）
	ParticleCount int     // 每次爆炸的粒子数
	FPS           int
}

func DefaultConfig() *Config {
	return &Config{
		BurstInterval: 2.0,
		ParticleCount: 100,
		FPS:           30,
	}
}

type Particle struct {
	x, y     float64
	vx, vy   float64
	life     float64
	maxLife  float64
	color    tcell.Color
	char     rune
}

type ParticleBurst struct {
	screen    tcell.Screen
	config    *Config
	particles []*Particle
	width     int
	height    int
	rand      *rand.Rand
	timeSinceBurst float64
}

func New(screen tcell.Screen, config *Config) *ParticleBurst {
	if config == nil {
		config = DefaultConfig()
	}

	return &ParticleBurst{
		screen:    screen,
		config:    config,
		particles: make([]*Particle, 0),
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (p *ParticleBurst) Init() error {
	p.width, p.height = p.screen.Size()
	p.timeSinceBurst = 0
	return nil
}

func (p *ParticleBurst) createBurst() {
	// 随机爆炸位置
	centerX := float64(p.rand.Intn(p.width))
	centerY := float64(p.rand.Intn(p.height))

	colors := []tcell.Color{
		tcell.ColorRed, tcell.ColorYellow, tcell.ColorOrange,
		tcell.ColorPurple, tcell.ColorPink, tcell.ColorWhite,
	}

	chars := []rune{'*', '●', '○', '·', '+', '×'}

	for i := 0; i < p.config.ParticleCount; i++ {
		angle := p.rand.Float64() * 2 * math.Pi
		speed := 5 + p.rand.Float64()*10

		particle := &Particle{
			x:       centerX,
			y:       centerY,
			vx:      math.Cos(angle) * speed,
			vy:      math.Sin(angle) * speed,
			life:    1.0,
			maxLife: 1.0 + p.rand.Float64()*2,
			color:   colors[p.rand.Intn(len(colors))],
			char:    chars[p.rand.Intn(len(chars))],
		}
		p.particles = append(p.particles, particle)
	}
}

func (p *ParticleBurst) Update(deltaTime float64) {
	p.timeSinceBurst += deltaTime

	// 定期创建新爆炸
	if p.timeSinceBurst >= p.config.BurstInterval {
		p.createBurst()
		p.timeSinceBurst = 0
	}

	// 更新所有粒子
	alive := make([]*Particle, 0)
	for _, particle := range p.particles {
		particle.life -= deltaTime
		if particle.life <= 0 {
			continue
		}

		// 应用重力
		particle.vy += 5 * deltaTime

		// 更新位置
		particle.x += particle.vx * deltaTime
		particle.y += particle.vy * deltaTime

		// 边界检查
		if particle.x >= 0 && particle.x < float64(p.width) &&
			particle.y >= 0 && particle.y < float64(p.height) {
			alive = append(alive, particle)
		}
	}
	p.particles = alive
}

func (p *ParticleBurst) Render() {
	p.screen.Clear()

	for _, particle := range p.particles {
		x, y := int(particle.x), int(particle.y)
		if x >= 0 && x < p.width && y >= 0 && y < p.height {
			alpha := particle.life / particle.maxLife
			var color tcell.Color
			if alpha > 0.7 {
				color = particle.color
			} else if alpha > 0.3 {
				color = tcell.ColorGray
			} else {
				color = tcell.ColorDarkGray
			}

			style := tcell.StyleDefault.Foreground(color)
			p.screen.SetContent(x, y, particle.char, nil, style)
		}
	}

	p.screen.Show()
}

func (p *ParticleBurst) Run(quit <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(p.config.FPS))
	defer ticker.Stop()

	lastUpdate := time.Now()

	for {
		select {
		case <-quit:
			return nil
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(lastUpdate).Seconds()
			lastUpdate = now

			p.Update(deltaTime)
			p.Render()
		}
	}
}

func (p *ParticleBurst) Cleanup() error {
	return nil
}
