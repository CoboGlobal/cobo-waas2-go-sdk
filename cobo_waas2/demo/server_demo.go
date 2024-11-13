package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
	"github.com/gin-gonic/gin"
)

var pubKeys = map[string]string{
	"DEV":  "a04ea1d5fa8da71f1dcfccf972b9c4eba0a2d8aba1f6da26f49977b08a0d2718",
	"PROD": "8d4a482641adb2a34b726f05827dba9a9653e5857469b8749052bf4458a86729",
}

// DEV or PROD
var env = "DEV"

func verifySignature(publicKey ed25519.PublicKey, signature, message string) bool {
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	doubleHash := crypto.Hash256x2(message)
	return ed25519.Verify(publicKey, doubleHash, signatureBytes)
}

func verifyRequest(c *gin.Context) {
	signature := c.GetHeader("BIZ_RESP_SIGNATURE")
	bizTimestamp := c.GetHeader("BIZ_TIMESTAMP")

	if signature == "" || bizTimestamp == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing signature or timestamp"})
		return
	}

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	log.Printf("Request Body: %s\n", string(rawBody))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	message := fmt.Sprintf("%s|%s", rawBody, bizTimestamp)
	publicKeyHex, exists := pubKeys[env]
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Environment key not found"})
		return
	}

	publicKey, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode public key"})
		return
	}

	if !verifySignature(publicKey, signature, message) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Signature verification failed"})
		return
	}
}

func handleWebhook(c *gin.Context) {
	verifyRequest(c)
	// Add your biz logic here
	c.JSON(http.StatusOK, gin.H{})
}

func handleCallback(c *gin.Context) {
	verifyRequest(c)
	// Add your callback logic here
	c.String(http.StatusOK, "ok")
}

func main() {
	r := gin.Default()
	r.POST("/api/webhook", handleWebhook)
	r.POST("/api/callback", handleCallback)
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
