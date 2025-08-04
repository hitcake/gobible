package main

func FToM(f Feet) Meter {
	return Meter(f * 0.3048)
}

func MToF(m Meter) Feet {
	return Feet(m / 0.3048)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.4536)
}

func KToP(k Kilogram) Pound {
	return Pound(k / 0.4536)
}
