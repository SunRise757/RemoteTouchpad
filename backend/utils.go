package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net"
	"strings"
	"time"

	"github.com/piglig/go-qr/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate a tokenL %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func getLanIp() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	//	var ips []string
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err == nil && ip.To4() != nil && strings.HasPrefix(ip.String(), "192.168.") {
				return ip.String()
				//ips = append(ips, ip.String())
			}
		}
	}
	return ""
}

func generateQr(url string) {

	registerToken := generateToken(32)
	registerTokens[registerToken] = time.Now().Unix()

	code, err := qr.Encode(url + "?registerToken=" + registerToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(code)
	/*
		style := []qr.RenderOption{
			qr.WithScale(32),
			qr.WithModuleShape(qr.ModuleRounded),
			qr.WithFinderShape(qr.FinderRounded),
		}
		png, err := code.PNG(style...)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("../build/qr.png", png, 0o644); err != nil {
			log.Fatal(err)
		}
	*/
}
