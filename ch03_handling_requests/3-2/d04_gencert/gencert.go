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

//生成SSL证书以及服务器私钥

//ssl证书是为了公证服务器公钥的

func main() {
	//生成证书序列号
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	//创建一个专有名称作为证书的标题
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	//对证书进行配置
	template := x509.Certificate{
		SerialNumber: serialNumber,//证书序列号
		Subject:      subject, //证书的标题
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), //证书的有效期设置为了1年
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, //表明这个X.509证书是用于进行服务器身份验证操作的
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},//表明这个X.509证书是用于进行服务器身份验证操作的
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")}, //程序将证书设置了只能在IP地址127.0.0.1之上运行
	}

	//生成一个RSA私钥
	//程序创建的RSA私钥的结构里面包含了一个能够公开访问的公钥,这个公钥在创建ssl证书的时候会用到
	pk, _ := rsa.GenerateKey(rand.Reader, 2048) //同时会返回一个公钥  pk.PublicKey
	//创建ssl证书
	//这个函数接受Certificate结构、公钥和私钥等多个参数，创建出一个经过DER编码格式的字节切片
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	//使用encoding/pem标准库将证书编码到cert.pem文件里面
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	//继续以PEM编码的方式把之前生成的秘钥编码并保存到key.pem文件里面
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
