/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package utils

import "math/rand"

func RandomPick(min int, max int) int {
	return min + rand.Intn(max-min)
}
