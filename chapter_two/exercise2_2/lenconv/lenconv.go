package lenconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m)}
func (f Foot) String() string { return fmt.Sprintf("%gft", f)}

func MToFt(m Meter) Foot { return Foot(m / 0.3048) }
func FtToM(f Foot) Meter { return Meter(f * 0.3048) }