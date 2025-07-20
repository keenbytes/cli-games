package termui

import (
	"math"
	"strings"
	"unicode/utf8"
)

const (
	// NoSplit indicates that the pane is not split.
	NoSplit = iota

	// Horizontally indicates that the pane is split horizontally.
	Horizontally

	// Vertically indicates that the pane is split vertically.
	Vertically
)

const (
	_ = iota

	// LeftPane is used for vertical splits and means the left pane has a fixed size.
	LeftPane

	// RightPane is used for vertical splits and means the right pane has a fixed size.
	RightPane

	// TopPane is used for horizontal splits and means the top pane has a fixed size.
	TopPane

	// BottomPane is used for horizontal splits and means the bottom pane has a fixed size.
	BottomPane
)

const (
	_ = iota

	// Char indicates that the size is specified in characters.
	Char

	// Percent indicates that the size is specified as a percentage.
	Percent
)

// Pane represents a single pane on the screen.
type Pane struct {
	left            int
	top             int
	width           int
	height          int
	canvasLeft      int
	canvasTop       int
	canvasWidth     int
	canvasHeight    int
	minWidth        int
	minHeight       int
	tooSmall        bool
	splitType       int
	splitSizeTarget int
	splitSize       int
	splitUnit       int
	panes           [2]*Pane
	frame           FrameStyle
	ui              *TermUI
	Widget          Widget
}

// Split divides the current pane into two panes either horizontally or vertically,
// using the specified split type, target side, size, and unit.
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

// Write draws a UTF-8 string at the given coordinates inside the pane.
func (p *Pane) Write(x, y int, content string) {
	positionX, positionY := p.canvasLeft+x, p.canvasTop+y

	length := utf8.RuneCountInString(content)
	if length > p.canvasWidth {
		p.ui.Write(positionX, positionY, string([]rune(content)[:p.canvasWidth]))
	} else {
		p.ui.Write(positionX, positionY, content)
	}
}

// WriteNoFrame writes a UTF-8 string at the specified position, ignoring the pane's frame boundaries.
func (p *Pane) WriteNoFrame(x, y int, content string) {
	p.ui.Write(p.left+x, p.top+y, content)
}

// Clear resets the pane's canvas by writing space characters.
func (p *Pane) Clear() {
	for line := range p.canvasHeight {
		p.ui.Write(p.canvasLeft, p.canvasTop+line, strings.Repeat(" ", p.canvasWidth))
	}
}

// ClearNoFrame overwrites the whole pane, frame included, with space characters.
func (p *Pane) ClearNoFrame() {
	for line := range p.height {
		p.ui.Write(0, line, strings.Repeat(" ", p.width))
	}
}

// CanvasWidth returns the width of the pane's canvas.
func (p *Pane) CanvasWidth() int {
	return p.canvasWidth
}

// CanvasHeight returns the height of the pane's canvas.
func (p *Pane) CanvasHeight() int {
	return p.canvasHeight
}

// setWidth sets the pane's width, ensuring it's not smaller than the minimal allowed width.
// It also recursively updates the width of any nested panes.
func (p *Pane) setWidth(width int) {
	p.width = width
	if p.minWidth > 0 && p.width < p.minWidth {
		p.tooSmall = true

		return
	}

	p.tooSmall = false

	switch p.splitType {
	case Horizontally:
		p.panes[0].left, p.panes[1].left = p.left, p.left
		p.panes[0].setWidth(width)
		p.panes[1].setWidth(width)

	case Vertically:
		value1, value2, tooSmall := p.getSplitValues()
		if tooSmall {
			p.tooSmall = true

			return
		}

		p.tooSmall = false
		p.panes[0].left, p.panes[1].left = p.left, p.left+value1
		p.panes[0].setWidth(value1)
		p.panes[1].setWidth(value2)

	default:
		p.canvasLeft = p.left + p.frame.LeftFrameSize()
		p.canvasWidth = p.width - p.frame.LeftFrameSize() - p.frame.RightFrameSize()
	}
}

// setHeight sets the pane's height, ensuring it is not smaller than the minimal allowed height.
// It also recursively updates the height of any nested panes.
func (p *Pane) setHeight(height int) {
	p.height = height
	if p.minHeight > 0 && p.height < p.minHeight {
		p.tooSmall = true

		return
	}

	p.tooSmall = false

	switch p.splitType {
	case Vertically:
		p.panes[0].top = p.top
		p.panes[1].top = p.top
		p.panes[0].setHeight(height)
		p.panes[1].setHeight(height)

	case Horizontally:
		value1, value2, tooSmall := p.getSplitValues()
		if tooSmall {
			p.tooSmall = true

			return
		}

		p.tooSmall = false
		p.panes[0].top = p.top
		p.panes[1].top = p.top + value1
		p.panes[0].setHeight(value1)
		p.panes[1].setHeight(value2)

	default:
		p.canvasTop = p.top + p.frame.TopFrameSize()
		p.canvasHeight = p.height - p.frame.TopFrameSize() - p.frame.BottomFrameSize()
	}
}

// getSplitValues is used by Split functions to calculate the width and height of resulting panes.
// It takes the split type, split value, and its unit, and returns the sizes in characters.
// It also checks whether the calculated sizes are too small.
func (p *Pane) getSplitValues() (int, int, bool) {
	var (
		baseVal int
		calcVal int
	)

	switch p.splitType {
	case Vertically:
		baseVal = p.width

	case Horizontally:
		baseVal = p.height

	default:
		return 0, 0, false
	}

	switch p.splitUnit {
	case Percent:
		calcVal = int(math.Abs(float64(p.splitSize) / 100 * float64(baseVal)))

	case Char:
		calcVal = int(math.Abs(float64(p.splitSize)))

	default:
		return 0, 0, false
	}

	if calcVal >= baseVal || calcVal < 1 {
		return 0, 0, true
	}

	switch p.splitSizeTarget {
	case LeftPane, TopPane:
		size1 := calcVal
		size2 := baseVal - calcVal

		return size1, size2, false

	case RightPane, BottomPane:
		size1 := baseVal - calcVal
		size2 := calcVal

		return size1, size2, false

	default:
		return 0, 0, false
	}
}

func (p *Pane) render() {
	if p.tooSmall {
		if p.frame != nil {
			width := p.width - p.frame.LeftFrameSize() - p.frame.RightFrameSize()

			height := p.height - p.frame.TopFrameSize() - p.frame.BottomFrameSize()
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

	if p.Widget != nil {
		p.Widget.Render(p)
	}
}

func (p *Pane) renderFrame() {
	cornerChars := p.frame.CornerChars()

	// logic here actually works for 1 character frame only
	if p.frame.TopFrameSize() > 1 || p.frame.LeftFrameSize() > 1 || p.frame.RightFrameSize() > 1 ||
		p.frame.BottomFrameSize() > 1 {
		panic(
			"frame can have a width of 1 character only, functions must all return 0 or 1",
		)
	}

	// corners
	p.WriteNoFrame(0, 0, cornerChars[TopLeftChar])
	p.WriteNoFrame(0, p.height-1, cornerChars[BottomLeftChar])
	p.WriteNoFrame(p.width-1, 0, cornerChars[TopRightChar])
	p.WriteNoFrame(p.width-1, p.height-1, cornerChars[BottomRightChar])

	// top, bottom, left, right
	if p.frame.TopFrameSize() > 0 {
		p.WriteNoFrame(
			p.frame.LeftFrameSize(),
			0,
			strings.Repeat(cornerChars[TopChar], p.canvasWidth),
		)
	}

	if p.frame.BottomFrameSize() > 0 {
		p.WriteNoFrame(
			p.frame.LeftFrameSize(),
			p.height-1,
			strings.Repeat(cornerChars[TopChar], p.canvasWidth),
		)
	}

	if p.frame.LeftFrameSize() > 0 {
		for x := range p.canvasHeight {
			p.WriteNoFrame(0, p.frame.TopFrameSize()+x, cornerChars[LeftChar])
		}
	}

	if p.frame.RightFrameSize() > 0 {
		for x := range p.canvasHeight {
			p.WriteNoFrame(p.width-1, p.frame.TopFrameSize()+x, cornerChars[RightChar])
		}
	}
}

func (p *Pane) iterate() {
	if p.Widget != nil {
		p.Widget.Iterate(p)
	}
}
