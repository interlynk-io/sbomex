/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package http

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch %d", resp.StatusCode)
		return
	}
	bRes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response %v", err)
		return
	}

	fmt.Println(string(bRes))
}
