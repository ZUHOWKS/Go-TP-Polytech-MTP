package main

import "math"

type vec2D32 struct {
	x, y float32
}

func (v vec2D32) add(v2 vec2D32) vec2D32 {
	return vec2D32{x: v.x + v2.x, y: v.y + v2.y}
}

func (v vec2D32) sub(v2 vec2D32) vec2D32 {
	return vec2D32{x: v.x - v2.x, y: v.y - v2.y}
}

func (v vec2D32) norm() float32 {
	return float32(math.Sqrt(float64(v.x*v.x + v.y*v.y)))
}

func (v vec2D32) dot(v2 vec2D32) float32 {
	return v.x*v2.x + v.y*v2.y
}

func (v vec2D32) scalar(v2 vec2D32) vec2D32 {
	return vec2D32{x: v.x * v2.x, y: v.y * v2.y}
}

func (v vec2D32) product(v2 vec2D32) float32 {
	return v.x*v2.y - v.y*v2.x
}
