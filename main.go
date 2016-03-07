package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

const (
	onosIP = "10.0.0.1"
)

func main() {

	ip := flag.String("ip", "", "IP address of the switch")
	host := flag.String("hostname", "", "Hoatname of the switch")
	dpid := flag.String("dpid", "", "DPID of the switch")
	user := flag.String("user", "root", "Username for the switch login")
	password := flag.String("password", "onl", "Password for the switch login")

	flag.Parse()

	var buf bytes.Buffer

	logger := log.New(&buf, "AUTOCONFIG: ", log.Ltime)

	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.Password(*password),
		},
	}
	client, err := ssh.Dial("tcp", *ip+":22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	cmd1 := "Working on... Hostname: " + *host + " with DPID: " + *dpid + " IP: " + *ip
	logger.Println(cmd1)
	scpCmd := "scp"

	cmdRC := "echo dpkg -i --force-overwrite /mnt/flash2/ofdpa-i.12.1.1_12.1.1+accton1.7-1_amd64.deb > /etc/rc.local"
	hostnameString := fmt.Sprintf("hostname %s", *host)
	cmdRChost := "echo " + hostnameString + " >> /etc/rc.local"
	cmdRCexit := "echo exit 0 >> /etc/rc.local"

	connect := "brcm-indigo-ofdpa-ofagent --dpid=" + *dpid + " --controller=" + onosIP

	cmds := []string{"test -e /etc/.configured && echo 'found' || echo 'notFound'",
		"persist /etc/network/interfaces",
		"savepersist",
		scpCmd,
		"service ofdpa stop",
		"dpkg -i --force-overwrite /mnt/flash2/ofdpa-i.12.1.1_12.1.1+accton1.7-1_amd64.deb",
		"service ofdpa restart",
		"persist /etc/accton/ofdpa.conf",
		"savepersist",
		cmdRC,
		cmdRChost,
		cmdRCexit,
		"persist /etc/rc.local",
		"savepersist",
		connect,
		"touch /etc/.configured",
		"persist /etc/.configured",
		"savepersist",
	}

	for cmdNumber, cmd := range cmds {

		session, err := client.NewSession()
		if err != nil {
			panic("Failed to create session: " + err.Error())
		}
		defer session.Close()

		var b bytes.Buffer
		session.Stdout = &b

		if cmd == scpCmd {

			src := "ofdpa-i.12.1.1_12.1.1+accton1.7-1_amd64.deb"
			dst := "/mnt/flash2/" + src

			err = scp.CopyPath(src, dst, session)
			if _, err := os.Stat(src); os.IsNotExist(err) {
				logger.Printf("no such file or directory: %s", src)
				panic(err)
			} else {
				logger.Println("SCP Success")
				continue
			}

		}

		logger.Println(" RUNNING: " + cmd)
		if cmd == connect {
			session.Run(cmd)

		} else {
			if err := session.Run(cmd); err != nil {
				logger.Println("Failed to run cmd: " + cmd + " ERROR: " + err.Error())
			}

		}

		rpl := b.String()

		if cmdNumber < 1 {
			fmt.Println(rpl[:5])
			if rpl[:5] == "found" {
				logger.Println("Switch is already configured!")
				break
			}

		}

	}

}
