package go_wss

import (
	"testing"
	"sha1"
)
var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

func TestComputeAcceptKey(t *testing.T){
	str := computeAcceptKey("hello trans")
	t.Log(str)
}

func computeAcceptKey(challKey string)string{
	h := sha1.New()
	h.Write([]byte(challKey))
    h.Write(keyGUID)
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}