# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://atlas.hashicorp.com/search.
  #config.vm.box = "centos7-cloud"
  #config.vm.box_url = "http://cloud.centos.org/centos/7/vagrant/x86_64/images/CentOS-7-x86_64-Vagrant-1711_01.VirtualBox.box"
  #config.vm.box = "centos/7"
  #config.vm.box_version = "1710.01"
  #config.vm.box = "fedora/25-cloud-base"
  #config.vm.box_version = "20161122"
  config.vm.box = "fedora26-cloud-base"
  config.vm.box_url = "https://download.fedoraproject.org/pub/fedora/linux/releases/26/CloudImages/x86_64/images/Fedora-Cloud-Base-Vagrant-26-1.5.x86_64.vagrant-virtualbox.box"
  #config.vm.box = "fedora27-cloud-base"
  #config.vm.box_url = "https://download.fedoraproject.org/pub/fedora/linux/releases/27/CloudImages/x86_64/images/Fedora-Cloud-Base-Vagrant-27-1.6.x86_64.vagrant-virtualbox.box"
  #config.vm.box = "ubuntu/xenial64"
  #config.vm.box_version = "20171212.0.0"
  #config.vm.box = "ubuntu16.04-cloud"
  #config.vm.box_url = "https://cloud-images.ubuntu.com/xenial/20171214/xenial-server-cloudimg-amd64-vagrant.box"

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # config.vm.network "forwarded_port", guest: 80, host: 8080
  #config.vm.network "forwarded_port", guest: 8080, host: 18080
  #config.vm.network "forwarded_port", guest: 443, host: 1443

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.network "private_network", ip: "172.17.4.59", auto_config: false
  #config.vm.network "private_network", ip: "10.64.33.82", auto_config: false
  config.vm.network "private_network", type: "dhcp", virtualbox_intnet: true, auto_config: false

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"
  config.vm.synced_folder ".", "/vagrant", disabled: true
  config.vm.synced_folder "/Users/fanhongling/Downloads", "/Users/fanhongling/Downloads"
  config.vm.synced_folder "/Users/fanhongling/go", "/Users/fanhongling/go"
  #config.vm.synced_folder "f:/", "/windows10_drive_f"
  #config.vm.synced_folder "g:/", "/windows10_drive_g"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.
  config.vm.provider "virtualbox" do |vb, override|
    vb.cpus = "1"
    vb.memory = "2048"

    #override.ssh.insert_key = false
    #override.vm.base_mac = "080027539284"
  end  

  config.ssh.insert_key = false

  # Define a Vagrant Push strategy for pushing to Atlas. Other push strategies
  # such as FTP and Heroku are also available. See the documentation at
  # https://docs.vagrantup.com/v2/push/atlas.html for more information.
  # config.push.define "atlas" do |push|
  #   push.app = "YOUR_ATLAS_USERNAME/YOUR_APPLICATION_NAME"
  # end

  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
  #   apt-get update
  #   apt-get install -y apache2
  # SHELL
end
