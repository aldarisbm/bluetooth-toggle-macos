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

if ! crontab -l | grep -qF "togglebt.sh"; then
    crontab -l | { cat; echo "0 0 * * * $(pwd)/togglebt.sh"; } | crontab -
fi
