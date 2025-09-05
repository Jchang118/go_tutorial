package encryption_test

import (
    "go_tutorial/go_basic/encryption"
    "fmt"
    "testing"
)

func TestECC(t *testing.T) {
    prvKey, err := encryption.GenPrivateKey()
    if err != nil {
        t.Fatalf("genPrivateKey fail: %s\n", err)
    }
    pubKey := prvKey.PublicKey
    plain := "因为我们没有什么不同"
    cipher, err := encryption.ECCEncrypt(plain, pubKey)
    if err != nil {
        t.Fatalf("ECCEnrypt fail: %s\n", err)
    }
    plain, err = encryption.ECCDecrypt(cipher, prvKey)
    if err != nil {
        t.Fatalf("ECCDecrypt fail: %s\n", err)
    }
    fmt.Printf("明文: %s\n", plain)
}

// go test -v ./encryption -run=^TestECC$ -count=1
