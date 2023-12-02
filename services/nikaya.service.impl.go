package services

import (
	"context"
	"math/rand"
)

type NikayaServiceImpl struct {
	ctx context.Context
}

// TangChiBoKinhRand implements NikayaServive.
func (*NikayaServiceImpl) TangChiBoKinhRand() int {
	return randInt(1, 152)
}

// TieuBoKinhRand implements NikayaServive.
func (*NikayaServiceImpl) TieuBoKinhRand() int {
	return randInt(1, 152)
}

// TrungBoKinhRand implements NikayaServive.
func (*NikayaServiceImpl) TrungBoKinhRand() int {
	return randInt(1, 152)
}

// TruongBoKinhRand implements NikayaServive.
func (*NikayaServiceImpl) TruongBoKinhRand() int {
	// panic("unimplemented")
	return randInt(1, 34)
}

// TuongUngBoKinhRand implements NikayaServive.
func (*NikayaServiceImpl) TuongUngBoKinhRand() int {
	panic("unimplemented")
}

func NewNikayaService(ctx context.Context) NikayaServive {
	return &NikayaServiceImpl{
		ctx: ctx,
	}
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
