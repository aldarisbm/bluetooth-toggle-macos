#!/bin/bash


BREW_INSTALLED=$(command -v brew)
BU_INSTALLED=$(command -v blueutil)

echo "This will install brew if it's not on your system already\n"
echo
read -p "If you're okay with this reply with \"Y/y\"): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
    echo
    echo "Continuing..."
else
    echo
    echo "Exiting..."
    exit 1
fi

if [[ -n "${BREW_INSTALLED}"  ]]
then
    echo
    echo "Brew is already installed in your system, proceeding..."
else
    echo
    echo "Installing homebrew"
    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
fi

if [[ -n "${BU_INSTALLED}" ]] 
then
    echo
    echo "Blueutil already installed, proceeding..."
else
    echo
    brew install blueutil
fi