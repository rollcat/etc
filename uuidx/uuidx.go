// Package uuidx partially implements the revised RFC4122
// (draft-ietf-uuidrev-rfc4122bis-07).
//
// It provides a couple alternative implementations for UUIDv8 that
// use a lower-resolution, "verbose" timestamp. This timestamp is
// human-readable, when printed in the standard hex representation,
// and similar to an ISO8601 datetime (with hours and minutes).

package uuidx

import (
	"crypto/rand"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var hostname string
var hostUUID uuid.UUID

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		// TODO: maybe derive from something else as a last resort?
		panic(err)
	}
	hostUUID = uuid.NewSHA1(uuid.NameSpaceDNS, []byte(hostname))
}

// UUID v8 "time + random" layout
//  0                   1                   2                   3
//  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                           timestamp                           |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           timestamp           |  ver  |         rand          |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |var|                         rand                              |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                             rand                              |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

func NewUUID8TimeRandom() (u uuid.UUID, err error) {
	// YYYYMMDD-HHMM-VRRR-VRRR-RRRRRRRRRRRR

	// Set the random bits
	_, err = io.ReadFull(rand.Reader, u[6:16])
	if err != nil {
		return uuid.UUID{}, err
	}

	// Set the human-readable date bits
	//          "%Y%m%d%H%M"
	layout := "200601021504"
	now := time.Now().UTC()
	t, _ := strconv.ParseInt(now.Format(layout), 16, 64)
	u[0] = byte(0xff & (t >> 40))
	u[1] = byte(0xff & (t >> 32))
	u[2] = byte(0xff & (t >> 24))
	u[3] = byte(0xff & (t >> 16))
	u[4] = byte(0xff & (t >> 8))
	u[5] = byte(0xff & t)

	// Set the version and variant bits
	u[6] = (u[6] & 0x0f) | 0x80 // Version 8
	u[8] = (u[8] & 0x3f) | 0x80 // Variant 0b10
	return
}

// UUID v8 "time + node + random" layout.
//
// The node bits follow the timestamp and are derived from the DNS
// host name.
//
//  0                   1                   2                   3
//  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                           timestamp                           |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           timestamp           |  ver  |         node          |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |var|           node            |             rand              |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                             rand                              |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

func NewUUID8TimeNodeRandom() (u uuid.UUID, err error) {
	// YYYYMMDD-HHMM-VNNN-VNNN-RRRRRRRRRRRR

	// Set the random bits
	_, err = io.ReadFull(rand.Reader, u[10:16])
	if err != nil {
		return uuid.UUID{}, err
	}

	// Get the node UUID bits
	u[6] = hostUUID[6]
	u[7] = hostUUID[7]
	u[8] = hostUUID[8]
	u[9] = hostUUID[9]

	// Set the human-readable date bits
	//          "%Y%m%d%H%M"
	layout := "200601021504"
	now := time.Now().UTC()
	t, _ := strconv.ParseInt(now.Format(layout), 16, 64)
	u[0] = byte(0xff & (t >> 40))
	u[1] = byte(0xff & (t >> 32))
	u[2] = byte(0xff & (t >> 24))
	u[3] = byte(0xff & (t >> 16))
	u[4] = byte(0xff & (t >> 8))
	u[5] = byte(0xff & t)

	// Set the version and variant bits
	u[6] = (u[6] & 0x0f) | 0x80 // Version 8
	u[8] = (u[8] & 0x3f) | 0x80 // Variant 0b10
	return
}

// func main() {
// 	var u uuid.UUID
// 	u, _ = NewUUID8TimeRandom()
// 	u, _ = NewUUID8TimeNodeRandom()
// 	println(u.String())
// }
