package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

func getKeyFile(sshPrivateKeyPath string) (key ssh.Signer, err error) {
	//usr, _ := user.Current()
	// file := "Path to your key file(.pem)"
	buf, err := ioutil.ReadFile(sshPrivateKeyPath)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}
func SCPFileToRemote(sshIpAddress, sshPort, sshPrivateKeyPath, sshUser, srcFilePath string) error {
	//	key, err := ioutil.ReadFile(sshPrivateKeyPath)
	key, err := getKeyFile(sshPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("Unable to read private key:%s", err)
	}
	/*signer, err := ssh.ParsePrivateKey(key)
	{
		if err != nil {
			return fmt.Errorf("Unable to read private key:%s", err)
		}

	}
	fmt.Println("signer is %s", signer)
	*/
	IpAddressAndPort := sshIpAddress + ":" + sshPort
	client, err := ssh.Dial("tcp", IpAddressAndPort, &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return fmt.Errorf("Failed to Dial %w", err)
	}
	fmt.Println("client is %s", client)
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("Failed to create session %w", err)
	}
	src, err := os.Open(srcFilePath)
	if err != nil {

		return fmt.Errorf("Failed to Open src file %w", err)
	}
	src.Close()
	err1 := scp.CopyPath(src.Name(), src.Name(), session)
	if err1 != nil {

		return fmt.Errorf("Failed to Open src file %w", err1)
	}
	return nil
}
func main() {
	err := SCPFileToRemote("65.1.93.147", "22", "TestVM-Renitha.ppk", "ec2-user", "/home/ec2-user/scp.go")
	if err != nil {
		fmt.Println("Error in SCP:", err)
		return
	}
}
