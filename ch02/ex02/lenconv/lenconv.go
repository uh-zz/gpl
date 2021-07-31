package lenconv

import "fmt"

type Meters float64
type Feet float64

const (
	// 便宜上、小数点第一位まで
	FeetUnit Feet = 0.3
)

// %gは浮動小数点
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
func (ft Feet) String() string  { return fmt.Sprintf("%g ft", ft) }
