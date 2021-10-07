package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

/**
 *  GetPersonalSSLCaAndPriKey
 *  @Description: 生成个人使用的SSL证书和私钥
 *  @Notice!:
 **/
func GetPersonalSSLCaAndPriKey() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co"},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		// 需要将 生成的随机 Int64 转为 *big.Int
		SerialNumber: serialNumber,		// 记录 CA 的唯一标识码
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)	// 创建一个私钥

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template,
		&pk.PublicKey, pk)

	// 将证书进行编码
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTFICATE", Bytes: derBytes})
	certOut.Close()

	// 将密钥进行编码
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}

func main() {
	GetPersonalSSLCaAndPriKey()
}
