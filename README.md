```
# Network Hosts Browser

This is a simple Go script that retrieves network hosts' information, such as IP addresses and MAC addresses, and opens the IP address of a specified server in the default browser.

## Requirements

- Go (Golang) installed on your system.
- Supported operating systems: Windows, Linux.

## Installation

1. Clone or download this repository to your local machine.

```bash
git clone https://github.com/novenopatch/network-hosts-browser.git
```

2. Navigate to the project directory.

```bash
cd network-hosts-browser
```

3. Run the Go script.

```bash
go run main.go
```

## Configuration

Before running the script, make sure to set up the `conf.ini` file with the configuration details:

- `server_mac`: MAC address of the server whose IP address you want to open in the browser.

Example `conf.ini` file:

```
[server]
server_mac = "00:11:22:33:44:55"
```

## Usage

Once you've set up the `conf.ini` file, simply run the script, and it will retrieve the network hosts' information, find the IP address of the specified server, and open it in the default browser.

## Support

If you encounter any issues or have any questions, feel free to [open an issue](https://github.com/yourusername/network-hosts-browser/issues) on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```