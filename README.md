# OCPP KS

We will be checking how the communication using OCPP 1.6 works, how to create a client and test it agains our backend.

- Steve for demo: http://steve.zepp.cloud/steve/manager/home
- JSON Central System Address: ws://steve.zepp.cloud/steve/websocket/CentralSystemService/<YOUR_CHARGEPOINT_ID>

To run your own backend:
 - Clone this repository https://github.com/steve-community/steve
 - docker compose up

Install golang:
 - MacOS: brew install golang
 - Debian: apt-get install golang-go -y
 - Windows: [download](https://go.dev/dl/go1.21.6.windows-amd64.msi) and install


Run the sample code:
 - cd ocpp-client
 - go run *.go
