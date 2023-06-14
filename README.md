# pingpong-go
Various flavour of ping pong game


1. Local ping pong
   - the code resides in pingpong.go
   - run it as go run ping[ong.go
   - code courtsey: https://riptutorial.com/go/example/6056/ping-pong-with-two-goroutines
2. Web ping pong running on local http server
    - pingpong running on local http server
    - two endpoints are exposed on local webserver
    - ping is calling pong ednpoint of same webserver
    - code courtesy: https://github.com/nkratzke/pingpong/tree/master/pingpong-go/src/pingpong
3. Distributed ping pong <!ToDo>
    - It is imagined to be a variance of #2
    - In this worls there will be multiple webserver running ping and pong endpoints
    - A ping now can call any of the available pong endpoints
    - Instead of repreated 'o' character, pingers/pongers can add their own hash/special word and forward it to some random ping
    - A ponger may decide to stop forwarding and return the call after processing say 10 words/hashes (block_chain inspirartion)
