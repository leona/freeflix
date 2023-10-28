#!/bin/sh
# For provisioning a fresh Vultr Alpine instance with block storage attached

apk update && apk upgrade
apk add docker docker-compose fail2ban rsync
mkdir /root/freeflix && mkdir /root/freeflix/config && mkdir /root/freeflix/data
rc-update add fail2ban
rc-update add docker boot
service docker start
service fail2ban start
parted -s /dev/vdb mklabel gpt
parted -s /dev/vdb unit mib mkpart primary 0% 100%
mkfs.ext4 /dev/vdb1
mkdir /mnt/blockstorage
echo >> /etc/fstab
echo /dev/vdb1               /mnt/blockstorage       ext4    defaults,noatime,nofail 0 0 >> /etc/fstab
mount /mnt/blockstorage