
# Docker Setup Demo

- encore app create
- encore build docker hello       
- docker run -e PORT=8081 -p 8081:8081 hello
- docker image ls
- docker image save hello -o hello.tar
- docker image load --input hello.tar
- docker run -e PORT=8081 -p 8081:8081 hello
                                            
