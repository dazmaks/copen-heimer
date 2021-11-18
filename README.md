## copen-heimer
Now I am become Death, the destroyer of worlds.

## How to use
1. [Build](#building) the project
2. Find masscan list or create a new one:
```
sudo apt install masscan
wget https://raw.githubusercontent.com/robertdavidgraham/masscan/master/data/exclude.conf
sudo masscan -p25565 0.0.0.0/0 --max-rate <maxrate> --excludefile exclude.conf -oL masscan.txt
```
3. Run script

## Building
```
git clone https://github.com/dazmaks/copen-heimer copen-heimer
cd copen-heimer
go build .
```

## TODO
- Remove timeout

## License
[MIT](LICENSE)
