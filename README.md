This is simple tools that can check token on <b>Binance Smart Chain</b> on golang languange.
Please pay attention this is beta version and still under development.
Because this is my first repo in golang language.

# installations
```bash
cd ~
git clone https://github.com/dbgid/EVMScannerGo.git
cd dbgidEVMScannerGo
go mod download
```
# Usage
Put your file wallet.txt in file on line <b>14</b> must full path if your file outside in your current directory.

if you want this script running like a thread, you can do it by editing on line <b>99</b> just add <b>go</b> in the front of
```golang
client.Fetch()
```
and after edit can be
```golang
go client.Fetch()
```
These method using golang runtime, not like threading by using CPU/Kernel manchine.
But I not recommended that you using these method.
Because your requests may fail and flag as flooding/spamming requests by remote server.
Any feedback are welcome, please write your idea in pull requests.
Regards,
Ajones Aka DBG ID
