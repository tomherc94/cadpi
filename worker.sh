#!/bin/bash
#criar sistema de arquivos
mkdir /home/vagrant/workerInput
mkdir /home/vagrant/workerOutput
sudo chmod 777 workerInput
sudo chmod 777 workerOutput
sudo chmod 777 workerApp.jar
sudo chmod 777 workerCopy.jar
sudo chmod 777 executeWorkerApp.sh
sudo chmod 777 executeWorkerCopy.sh
sudo chmod 777 clearWorker.sh

#instalar dependências
sudo apt-get update
sudo apt-get install default-jre -y

#SSHPASS
sudo apt-get install sshpass -y

#configurar autenticação SSH
sudo sed -i 's/prohibit-password/yes/' /etc/ssh/sshd_config
sudo sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/' /etc/ssh/sshd_config

#resetar o SSH
sudo /etc/init.d/ssh restart

#configurar senha dos usuários
sudo usermod -p $(openssl passwd -1 '123') root
sudo usermod -p $(openssl passwd -1 '123') vagrant
