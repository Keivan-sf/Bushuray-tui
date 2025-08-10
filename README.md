<img width="1297" height="778" alt="image" src="https://github.com/user-attachments/assets/721fcb9b-08f0-475b-a164-48ed681710fa" />

## Bushuray-tui
Bushuray-tui is a keyboard-driven Xray client built for your terminal. It uses [Bushuray-core](https://github.com/Keivan-sf/Bushuray-core) as its back-end and It's written in Go with [bubbletea](https://github.com/charmbracelet/bubbletea)
### How to use
Download and extract the latest version from releases section then run
```
./bushuray
```

### Configuration
Bushuray will create its configuration file in `~/.config/bushuray/config.json` with the following template:
```json
{
  "socks-port": 3090,
  "http-port": 3091,
  "test-port-range": {
    "start": 3096,
    "end": 30120
  },
  "no-background": false 
}
```
- `socks-port`: Exposed local socks5 port
- `http-port`: Exposed local http port
- `test-port-range`: Port range used for profile testing
- `no-background`: Whether or not tui should have a background. Use this if you want your own terminal background or you have a transparent terminal

### Tun mode
To use tun mode, simply connect to a client then press `v`. Tun mode is experimental at this time but should work. If it doesn't, please create an issue. Running as root raises security concerns for the current version, see [this issue](https://github.com/Keivan-sf/Bushuray-core/issues/10).
### Debugging
If something is not working as expected, you can examine `debug.log` and `core-debug.log`. For example:
```
tail -f core-debug.log
```

