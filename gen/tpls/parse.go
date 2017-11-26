package tpls

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func parseInt(s []byte) int64 {
	i, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		log.Printf("Failed parsing %s to int", string(s))
	}
	return i
}

func parseUInt(s []byte) uint64 {
	i, err := strconv.ParseUint(string(s), 10, 64)
	if err != nil {
		log.Printf("Failed parsing %s to uint", string(s))
	}
	return i
}

func parseFloat(s []byte) float64 {
	i, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		log.Printf("Failed parsing %s to float", string(s))
	}
	return i
}

func parseTime(s []byte, precision int) time.Time {
	format := "2006-01-02 15:04:05"
	if precision > 0 {
		format += "." + strings.Repeat("0", precision)
	}
	t, err := time.Parse(format, string(s))
	if err != nil {
		log.Printf("Failed parsing '%s' to time.Time", string(s))
	}
	return t
}

func parseBool(s []byte) bool {
	return s[0] != 0
}
