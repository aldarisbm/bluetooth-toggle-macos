#!/bin/bash


BREW_INSTALLED=$(command -v brew)
BU_INSTALLED=$(command -v blueutil)


if [[ -n "${BREW_INSTALLED}" ]]
then
    echo "Brew is already installed in your system, proceeding..."
else
    echo "Installing homebrew"
    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
fi

if [[ -n "${BU_INSTALLED}" ]] 
then
    echo "Blueutil already installed, proceeding..."
else
    brew install blueutil
fi

if ! grep -qF "/mnt/dev" usr/bin/crontab; then
      echo "/dev/sda1 /mnt/dev ext4 defaults 0 0" | sudo tee -a /etc/fstab
fi
