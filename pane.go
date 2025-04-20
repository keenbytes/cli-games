package termui

import (
	"math"
	"strings"
	"unicode/utf8"
)

const (
	NoSplit = iota
	Horizontally
	Vertically
)

const (
	_ = iota
	Left
	Right
	Top
	Bottom
)

const (
	_ = iota
	Char
	Percent
)

type Pane struct {
	left int
	top int
	width int
	height int
	canvasLeft int
	canvasTop int
	canvasWidth int
	canvasHeight int
	minWidth int
	minHeight int
	tooSmall bool
	splitType int
	splitSizeTarget int
	splitSize int
	splitUnit int
	panes [2]*Pane
	frame FrameStyle
	ui *TermUI
}

// Split creates new two panes by splitting this pane either horizontally or vertically.
func (p *Pane) Split(typ int, sizeTarget int, size int, unit int) (*Pane, *Pane) {
	p.panes[0] = &Pane{
		ui: p.ui,
	}
	p.panes[1] = &Pane{
		ui: p.ui,
	}
	p.splitType = typ
	p.splitSizeTarget = sizeTarget
	p.splitSize = size
	p.splitUnit = unit
	return p.panes[0], p.panes[1]
}

func (p *Pane) Write(x, y int, content string) {
	cx, cy := p.canvasLeft+x, p.canvasTop+y
	length := utf8.RuneCountInString(content)
	if length > p.canvasWidth {
		p.ui.Write(cx, cy, string([]rune(content)[:p.canvasWidth]))
	} else {
		p.ui.Write(cx, cy, content)
	}
}

func (p *Pane) WriteNoFrame(x, y int, content string) {
	/*if p.splitType == Horizontally || p.splitType == Vertically {
		return
	}
	if p.tooSmall {
		return
	}*/
	p.ui.Write(p.left + x, p.top + y, content)
}

// setWidth sets width of pane, checks if it's not too small for the content (search for 'minimal width')
// and calls panes inside to set their width as well.
func (p *Pane) setWidth(w int) {
	p.width = w
	if p.minWidth > 0 && p.width < p.minWidth {
		p.tooSmall = true
		return
	}
	p.tooSmall = false

	switch p.splitType {
	case Horizontally:
		p.panes[0].left, p.panes[1].left = p.left, p.left
		p.panes[0].setWidth(w)
		p.panes[1].setWidth(w)
	case Vertically:
		v1, v2, tooSmall := p.getSplitValues()
		if tooSmall {
			p.tooSmall = true
			return
		}
		p.tooSmall = false
		p.panes[0].left, p.panes[1].left = p.left, p.left + v1
		p.panes[0].setWidth(v1)
		p.panes[1].setWidth(v2)
	default:
		p.canvasLeft = p.left + p.frame.L()
		p.canvasWidth = p.width - p.frame.L() - p.frame.R()
	}
}

// setHeight sets height of pane, checks if it's not too small for the content (search for 'minimal height')
// and calls panes inside to set their height as well.
func (p *Pane) setHeight(h int) {
	p.height = h
	if p.minHeight > 0 && p.height < p.minHeight {
		p.tooSmall = true
		return
	}
	p.tooSmall = false

	switch p.splitType {
	case Vertically:
		p.panes[0].top = p.top
		p.panes[1].top = p.top
		p.panes[0].setHeight(h)
		p.panes[1].setHeight(h)
	case Horizontally:
		v1, v2, tooSmall := p.getSplitValues()
		if tooSmall {
			p.tooSmall = true
			return
		}
		p.tooSmall = false
		p.panes[0].top = p.top
		p.panes[1].top = p.top + v1
		p.panes[0].setHeight(v1)
		p.panes[1].setHeight(v2)
	default:
		p.canvasTop = p.top + p.frame.T()
		p.canvasHeight = p.height - p.frame.T() - p.frame.B()
	}
}

// getSplitValues is used by Split functions to calculate the width
// and height of panes. It takes the split type, split value (and its unit)
// and calculates the size in number of characters. It also checks if the size
// is not too small as well.
func (p *Pane) getSplitValues() (size1 int, size2 int, tooSmall bool) {
	var baseVal int
	var calcVal int

	switch p.splitType {
	case Vertically:
		baseVal = p.width
	case Horizontally:
		baseVal = p.height
	default:
		return
	}

	switch p.splitUnit {
	case Percent:
		calcVal = int(math.Abs(float64(p.splitSize) / 100 * float64(baseVal)))
	case Char:
		calcVal = int(math.Abs(float64(p.splitSize)))
	default:
		return
	}

	if calcVal >= baseVal || calcVal < 1 {
		tooSmall = true
		return
	}

	switch p.splitSizeTarget {
	case Left, Top:
		size1 = calcVal
		size2 = baseVal - calcVal
	case Right, Bottom:
		size1 = baseVal - calcVal
		size2 = calcVal
	default:
		return
	}

	return
}

func (p *Pane) render() {
	if p.tooSmall {
		if p.frame != nil {
			width := p.width - p.frame.L() - p.frame.R()
			height := p.height - p.frame.T() - p.frame.B()
			if width > 0 && height > 0 {
				p.renderFrame()
				p.Write(0, 0, "!")
				return
			}
		}
		if p.width > 0 && p.height > 0 {
			p.WriteNoFrame(0, 0, "!")
			return
		}
	}

	if p.splitType == Horizontally || p.splitType == Vertically {
		p.panes[0].render()
		p.panes[1].render()
		return
	}

	p.renderFrame()
}

func (p *Pane) renderFrame() {
	c := p.frame.C()

	// TODO: logic here actually works for 1 character frame only
	if p.frame.T() > 1 || p.frame.L() > 1 || p.frame.R() > 1 || p.frame.B() > 1 {
		panic("frame can have a width of 1 character only, L(), R(), T(), B() must all return 0 or 1")
	}

	// corners
	p.WriteNoFrame(0, 0, c[NW])
	p.WriteNoFrame(0, p.height-1, c[SW])
	p.WriteNoFrame(p.width-1, 0, c[NE])
	p.WriteNoFrame(p.width-1, p.height-1, c[SE])

	// top, bottom, left, right
	if p.frame.T() > 0 {
		p.WriteNoFrame(p.frame.L(), 0, strings.Repeat(c[N], p.canvasWidth))
	}
	if p.frame.B() > 0 {
		p.WriteNoFrame(p.frame.L(), p.height-1, strings.Repeat(c[N], p.canvasWidth))
	}
	if p.frame.L() > 0 {
		for x := 0; x < p.canvasHeight; x++ {
			p.WriteNoFrame(0, p.frame.T() + x, c[W])
		}
	}
	if p.frame.R() > 0 {
		for x := 0; x < p.canvasHeight; x++ {
			p.WriteNoFrame(p.width-1, p.frame.T() + x, c[E])
		}
	}
}

func (p *Pane) iterate() {
}
