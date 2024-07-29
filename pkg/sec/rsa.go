package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"goiot/pkg/conf"
	"goiot/pkg/logger"
	"math/big"
	"os"
	"strings"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	certPem    []byte
)

func InitRSAKeys() (err error) {
	// 从文件加载密钥
	caPath := strings.TrimRight(conf.Conf.GetString("ca.path"), "/") + "/"
	privateKeyPath := caPath + "private.pem"
	publicKeyPath := caPath + "public.pem"
	certPath := caPath + "certificate.pem"

	if _, err = os.Stat(caPath); os.IsNotExist(err) {
		logger.Logger.Info("CA directory does not exist, creating...")
		if err = os.MkdirAll(caPath, os.ModePerm); err != nil {
			logger.Logger.Errorf("create CA directory error: %v", err)
			return err
		}
	}

	// 尝试加载私钥
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		logger.Logger.Warnf("No existing private key found or read %s error: %v", privateKeyPath, err)
		// 生成新的密钥对
		privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			logger.Logger.Errorf("generate private key error: %v", err)
			return err
		}
		// 保存新的私钥和公钥
		if err = savePrivateKeyToFile(privateKey, privateKeyPath); err != nil {
			logger.Logger.Errorf("save private key error: %v", err)
			return err
		}
		if err = savePublicKeyToFile(&privateKey.PublicKey, publicKeyPath); err != nil {
			logger.Logger.Errorf("save public key error: %v", err)
			return err
		}
	} else {
		// 解析私钥
		privateKeyPem, _ := pem.Decode(privateKeyBytes)
		if privateKeyPem == nil {
			return errors.New("failed to decode PEM block containing the private key")
		}
		privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyPem.Bytes)
		if err != nil {
			return err
		}
	}

	// 尝试加载公钥证书
	certPem, err = os.ReadFile(certPath)
	if err != nil {
		logger.Logger.Warnf("No existing certificate found or read %s error: %v", certPath, err)
		// 生成并保存新证书
		if err = generateAndSaveCertificate(caPath, &privateKey.PublicKey); err != nil {
			return err
		}
	}
	return nil
}

func GetCertPem() []byte {
	return certPem
}

func GetPubPem() []byte {
	caPath := strings.TrimRight(conf.Conf.GetString("ca.path"), "/") + "/"
	publicKeyPath := caPath + "public.pem"
	publicKeyPem, _ := os.ReadFile(publicKeyPath)
	return publicKeyPem
}

func DecryptRSA(encryptedData string) (string, error) {
	// 解码
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		logger.Logger.Errorf("decode base64 data error: %v", err)
		return "", err
	}

	// 解密
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		logger.Logger.Errorf("decrypt RSA error: %v", err)
		return "", err
	}

	return string(decryptedBytes), nil
}

func loadKeysFromFile(privateKeyPath, publicKeyPath string) error {
	// 读取私钥
	privateBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}
	privatePem, _ := pem.Decode(privateBytes)
	if privatePem == nil {
		return fmt.Errorf("decode private key error")
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		return err
	}

	// 读取公钥
	pubBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}
	pubPem, _ := pem.Decode(pubBytes)
	if pubPem == nil {
		return errors.New("decode public key error")
	}
	_, err = x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return err
	}

	return nil
}

func savePrivateKeyToFile(privKey *rsa.PrivateKey, filename string) error {
	keyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	privateBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	// 使用PEM格式编码并保存
	return pem.Encode(keyFile, privateBlock)
}

func savePublicKeyToFile(pubKey *rsa.PublicKey, filename string) error {
	keyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return err
	}

	// 使用PEM格式编码并保存
	return pem.Encode(keyFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
}

func generateAndSaveCertificate(caPath string, publicKey *rsa.PublicKey) error {
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return err
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:         "Linkview Certificate",
			Country:            []string{"CN"},
			Province:           []string{"Shaanxi"},
			Locality:           []string{"Xi'an"},
			Organization:       []string{"Linkview"},
			OrganizationalUnit: []string{"IOT"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey, privateKey)
	if err != nil {
		logger.Logger.Errorf("create certificate error: %v", err)
		return err
	}

	// 编码证书到PEM格式并保存
	certPem = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	_ = os.WriteFile(caPath+"certificate.pem", certPem, 0644)

	return nil
}
