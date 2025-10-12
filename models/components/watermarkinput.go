package components

import (
	"encoding/json"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/utils"
)

// WatermarkInputType - Type of overlay (currently only supports 'watermark').
type WatermarkInputType string

const (
	WatermarkInputTypeWatermark WatermarkInputType = "watermark"
)

func (e WatermarkInputType) ToPointer() *WatermarkInputType {
	return &e
}
func (e *WatermarkInputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "watermark":
		*e = WatermarkInputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WatermarkInputType: %v", v)
	}
}

// XAlign - Horizontal alignment of the watermark.
type XAlign string

const (
	XAlignLeft   XAlign = "left"
	XAlignCenter XAlign = "center"
	XAlignRight  XAlign = "right"
)

func (e XAlign) ToPointer() *XAlign {
	return &e
}
func (e *XAlign) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "left":
		fallthrough
	case "center":
		fallthrough
	case "right":
		*e = XAlign(v)
		return nil
	default:
		return fmt.Errorf("invalid value for XAlign: %v", v)
	}
}

// YAlign - Vertical alignment of the watermark.
type YAlign string

const (
	YAlignTop    YAlign = "top"
	YAlignMiddle YAlign = "middle"
	YAlignBottom YAlign = "bottom"
)

func (e YAlign) ToPointer() *YAlign {
	return &e
}
func (e *YAlign) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "top":
		fallthrough
	case "middle":
		fallthrough
	case "bottom":
		*e = YAlign(v)
		return nil
	default:
		return fmt.Errorf("invalid value for YAlign: %v", v)
	}
}

type Placement struct {
	// Horizontal alignment of the watermark.
	XAlign *XAlign `json:"xAlign,omitempty"`
	// Horizontal margin from the edge of the video.
	XMargin *string `json:"xMargin,omitempty"`
	// Vertical alignment of the watermark.
	YAlign *YAlign `json:"yAlign,omitempty"`
	// Vertical margin from the edge of the video.
	YMargin *string `json:"yMargin,omitempty"`
}

func (p Placement) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Placement) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *Placement) GetXAlign() *XAlign {
	if p == nil {
		return nil
	}
	return p.XAlign
}

func (p *Placement) GetXMargin() *string {
	if p == nil {
		return nil
	}
	return p.XMargin
}

func (p *Placement) GetYAlign() *YAlign {
	if p == nil {
		return nil
	}
	return p.YAlign
}

func (p *Placement) GetYMargin() *string {
	if p == nil {
		return nil
	}
	return p.YMargin
}

type WatermarkInput struct {
	// Type of overlay (currently only supports 'watermark').
	Type *WatermarkInputType `json:"type,omitempty"`
	// URL of the watermark image.
	URL       *string    `json:"url,omitempty"`
	Placement *Placement `json:"placement,omitempty"`
	// Width of the watermark in percentage or pixels.
	Width *string `json:"width,omitempty"`
	// Height of the watermark in percentage or pixels.
	Height *string `json:"height,omitempty"`
	// Opacity of the watermark in percentage.
	Opacity *string `json:"opacity,omitempty"`
}

func (w WatermarkInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WatermarkInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WatermarkInput) GetType() *WatermarkInputType {
	if w == nil {
		return nil
	}
	return w.Type
}

func (w *WatermarkInput) GetURL() *string {
	if w == nil {
		return nil
	}
	return w.URL
}

func (w *WatermarkInput) GetPlacement() *Placement {
	if w == nil {
		return nil
	}
	return w.Placement
}

func (w *WatermarkInput) GetWidth() *string {
	if w == nil {
		return nil
	}
	return w.Width
}

func (w *WatermarkInput) GetHeight() *string {
	if w == nil {
		return nil
	}
	return w.Height
}

func (w *WatermarkInput) GetOpacity() *string {
	if w == nil {
		return nil
	}
	return w.Opacity
}
