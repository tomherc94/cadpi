#!/bin/bash
mkdir /home/vagrant/workerInput
sudo chmod 777 workerInput
sudo apt-get update
sudo apt-get install default-jre -y


sudo sed -i 's/prohibit-password/yes/' /etc/ssh/sshd_config

sudo sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/' /etc/ssh/sshd_config

#habilita o SSH
sudo /etc/init.d/ssh restart

sudo usermod -p $(openssl passwd -1 '123') root
sudo usermod -p $(openssl passwd -1 '123') vagrant
