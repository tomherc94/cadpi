#!/bin/bash

#Instalando o Golang
wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.13.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/vagrant/.profile
source /home/vagrant/.profile

#sudo ssh-keygen


sudo ssh-keygen -t rsa -b 4096 -C "comment" -P "examplePassphrase" -f "keyName" -q
yes "yes" | sudo ssh-copy-id -f -i keyName 172.42.42.101
yes "yes" | sudo ssh-copy-id -f -i keyName 172.42.42.102
yes "yes" | sudo ssh-copy-id -f -i keyName 172.42.42.103



#Listando IP's da rede
#sudo apt-get update
#sudo apt-get install nmap -y
#sudo nmap -sS 172.42.42.0/24 | grep "Nmap scan" | cut -d " " -f 5 | grep 172* > ~/ips.txt
