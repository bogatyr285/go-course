ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key 
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub


ssh-keygen -t ed25519 -f -m PEM ed25519.key
ssh-keygen -e -m PEM -f ed25519.key > ed25519.key.pub13 