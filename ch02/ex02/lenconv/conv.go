package lenconv

func MToF(m Meters) Feet { return Feet(m / Meters(FeetUnit)) }

func FToM(ft Feet) Meters { return Meters(ft * FeetUnit) }
