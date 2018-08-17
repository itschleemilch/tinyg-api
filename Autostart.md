# Autostart on Debian / Raspberry PI

1. Install screen: `sudo apt-get install screen`
2. Add screen session on start. Modify /etc/rc.local, add `su - pi -c "screen -dm -S pistartup ~/tinyg.sh"` before `exit 0`.
3. Create `/home/pi/tinyg.sh`, set executable (`chmod +x tinyg.sh`).
4. Set script content:

```
#!/bin/bash
cd ~/go/src/github.com/itschleemilch/tinyg-api/v0/cmd/tinyg-control
./tinyg-control
```
