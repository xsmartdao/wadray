package wadray

import "github.com/holiman/uint256"

var (
	wad     = uint256.NewInt(1).Exp(uint256.NewInt(10), uint256.NewInt(18))
	halfWad = new(uint256.Int).Div(wad, uint256.NewInt(2))

	ray     = uint256.NewInt(1).Exp(uint256.NewInt(10), uint256.NewInt(27))
	halfRay = new(uint256.Int).Div(ray, uint256.NewInt(2))

	wadRayRatio = uint256.NewInt(1).Exp(uint256.NewInt(10), uint256.NewInt(9))
)

func Wad() *uint256.Int {
	return wad
}

func HalfWad() *uint256.Int {
	return halfWad
}

func Ray() *uint256.Int {
	return ray
}

func HalfRay() *uint256.Int {
	return halfRay
}

func WadMul(a, b *uint256.Int) *uint256.Int {
	return new(uint256.Int).Div(new(uint256.Int).Add(halfWad, new(uint256.Int).Mul(a, b)), wad)
}

func WadDiv(a, b *uint256.Int) *uint256.Int {
	halfB := new(uint256.Int).Div(b, uint256.NewInt(2))
	return new(uint256.Int).Div(new(uint256.Int).Add(halfB, new(uint256.Int).Mul(a, wad)), b)
}

func RayMul(a, b *uint256.Int) *uint256.Int {
	return new(uint256.Int).Div(new(uint256.Int).Add(halfRay, new(uint256.Int).Mul(a, b)), ray)
}

func RayDiv(a, b *uint256.Int) *uint256.Int {
	halfB := new(uint256.Int).Div(b, uint256.NewInt(2))
	return new(uint256.Int).Div(new(uint256.Int).Add(halfB, new(uint256.Int).Mul(a, ray)), b)
}

func RayToWad(a *uint256.Int) *uint256.Int {
	halfRatio := new(uint256.Int).Div(wadRayRatio, uint256.NewInt(2))
	return new(uint256.Int).Div(new(uint256.Int).Add(halfRatio, a), wadRayRatio)
}

func WadToRay(a *uint256.Int) *uint256.Int {
	return new(uint256.Int).Mul(a, wadRayRatio)
}

func RayPow(x, n *uint256.Int) (z *uint256.Int) {
	z = new(uint256.Int)
	if !new(uint256.Int).Mod(n, uint256.NewInt(2)).Eq(uint256.NewInt(0)) {
		z = x
	} else {
		z = ray
	}
	for n.Div(n, uint256.NewInt(2)); !n.Eq(uint256.NewInt(0)); n.Div(n, uint256.NewInt(2)) {
		x = RayMul(x, x)
		if !new(uint256.Int).Mod(z, x).Eq(uint256.NewInt(0)) {
			z = RayMul(z, x)
		}
	}
	return
}
