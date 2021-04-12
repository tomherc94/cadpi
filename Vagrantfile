# -*- mode: ruby -*-
# vi: set ft=ruby :

ENV['VAGRANT_NO_PARALLEL'] = 'yes'

Vagrant.configure(2) do |config|

  #config.vm.provision "shell", path: "bootstrap.sh"

  # Master Server
  config.vm.define "master" do |master|
    master.vm.box = "ubuntu/xenial64"
    master.vm.hostname = "master.example.com"
    master.vm.network "private_network", ip: "172.42.42.100"
    master.vm.provider "virtualbox" do |v|
      v.name = "master"
      v.memory = 1024
      v.cpus = 2
    end
    master.vm.provision "shell", path: "master.sh"
    master.vm.synced_folder "./masterInput", "/home/vagrant/input"
    master.vm.synced_folder "./masterOutput", "/home/vagrant/output"
  end

  NodeCount = 3

  # Worker Nodes
  (1..NodeCount).each do |i|
    config.vm.define "worker#{i}" do |workernode|
      workernode.vm.box = "ubuntu/xenial64"
      workernode.vm.hostname = "worker#{i}.example.com"
      workernode.vm.network "private_network", ip: "172.42.42.10#{i}"
      workernode.vm.provider "virtualbox" do |v|
        v.name = "worker#{i}"
        v.memory = 512
        v.cpus = 1
      end
      workernode.vm.provision "shell", path: "worker.sh"
    end
  end

end
