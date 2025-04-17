package main

import (
	"net/url"
	"testing"
)

func TestSign(t *testing.T) {

	userID := "12251546"
	time := "1744899860814"

	// sign: dc78f718d7e896796b2120746c1e3051
	// fid: 12251546
	// time: 1744899860814

	data := url.Values{}
	data.Set("fid", userID)
	data.Set("time", time)

	signedData := appendSign(data)

	t.Log(signedData)
}
