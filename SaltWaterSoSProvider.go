package ucnlsalinity

import (
	"errors"
	"math"
)

// CalcSoundSpeed calculates speed of sound in water according to Wilson formula
// See: https://rbr-global.com/speed-of-sound-in-water/
func CalcSoundSpeed(t, p, s float64) (float64, error) {
	if t < -4 || t > 30 {
		return 0, errors.New("t is out of range (-4 to 30)")
	}

	if p < 0.1 || p > 100 {
		return 0, errors.New("p is out of range (0.1 to 100 MPa)")
	}

	if s < 0 || s > 40 {
		return 0, errors.New("s is out of range (0 to 40)")
	}

	c0 := 1449.14
	Dct := 4.5721*t - 4.4532E-2*math.Pow(t, 2) - 2.6045E-4*math.Pow(t, 3) + 7.9851E-6*math.Pow(t, 4)
	Dcs := 1.39799*(s-35) - 1.69202E-3*math.Pow(s-35, 2)
	Dcp := 1.63432*p - 1.06768E-3*math.Pow(p, 2) + 3.73403E-6*math.Pow(p, 3) - 3.6332E-8*math.Pow(p, 4)
	Dcstp := (s-35)*(-1.1244E-2*t+7.7711E-7*math.Pow(t, 2)+7.85344E-4*p-
		1.3458E-5*math.Pow(p, 2)+3.2203E-7*p*t+1.6101E-8*math.Pow(t, 2)*p)+
		p*(-1.8974E-3*t+7.6287E-5*math.Pow(t, 2)+4.6176E-7*math.Pow(t, 3))+
		math.Pow(p, 2)*(-2.6301E-5*t+1.9302E-7*math.Pow(t, 2))+
		math.Pow(p, 3)*(-2.0831E-7*t)

	result := c0 + Dct + Dcs + Dcp + Dcstp

	return result, nil
}


