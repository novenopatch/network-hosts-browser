package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/ini.v1"
)

type Hosts struct {
	IP  string
	MAC string
}

type Config struct {
	ServerMAC string
}

func main() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("arp", "-a")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("ip", "neigh", "show")
	} else {
		fmt.Println("Le système d'exploitation n'est pas pris en charge.")
		return
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Erreur en exécutant la commande:", err)
		return
	}
	hosts := getHosts(string(output))
	if hosts == nil {
		return
	}
	config, err := loadConfig("conf.ini")
	if err != nil {
		fmt.Println("Erreur en chargeant la configuration:", err)
		return
	}

	serverIP := findIPByMAC(hosts, config.ServerMAC)
	if serverIP == "" {
		fmt.Println("L'adresse IP du serveur n'a pas été trouvée. MAC:", config.ServerMAC)
		return
	}

	err = openInBrowser(serverIP)
	if err != nil {
		fmt.Println("Erreur en ouvrant l'adresse IP dans le navigateur:", err)
		return
	}
	fmt.Println("L'adresse IP du serveur a été ouverte dans votre navigateur par défaut.")
	fmt.Println("Lien :", "http://"+serverIP)

	fmt.Println("Appuyez sur Entrée pour quitter...")
	fmt.Scanln()
}
func getHosts(output string) []Hosts {
	var hosts []Hosts

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		var ip, mac string
		if runtime.GOOS == "windows" {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				ip = fields[0]
				mac = fields[1]
			}
		} else if runtime.GOOS == "linux" {
			fields := strings.Fields(line)
			if len(fields) >= 5 {
				ip = fields[0]
				mac = fields[4]
			}
		}
		if ip != "" && mac != "" {
			mac = normalizeMAC(mac)
			hosts = append(hosts, Hosts{IP: ip, MAC: mac})
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur en analysant la sortie de la commande:", err)
		return nil
	}
	return hosts

}
func findIPByMAC(hosts []Hosts, mac string) string {
	for _, host := range hosts {
		if host.MAC == mac {
			return host.IP
		}
	}
	return ""
}

func loadConfig(filename string) (*Config, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}
	serverMAC := cfg.Section("server").Key("server_mac").String()
    if serverMAC == "" {
        return nil, fmt.Errorf("clé 'server_mac' manquante dans le fichier de configuration")
    }
    
    return &Config{ServerMAC: serverMAC}, nil
}

func openInBrowser(ip string)  error{
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", "http://"+ip)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "http://"+ip)
	case "linux":
		cmd = exec.Command("xdg-open", "http://"+ip)
	default:
		fmt.Println("Système d'exploitation non pris en charge.")
	}
	return cmd.Run()
}

func normalizeMAC(mac string) string {
	if runtime.GOOS == "windows" {
		return strings.Replace(mac, "-", ":", -1)
	}
	return mac
}


func printHosts(hosts []Hosts) {
	for i, Hosts := range hosts {
		fmt.Printf("Hosts %d - IP: %s, MAC: %s\n", i+1, Hosts.IP, Hosts.MAC)
	}
}