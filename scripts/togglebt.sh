#!/bin/bash
LID_OPEN=$(ioreg -r -k AppleClamshellState | grep AppleClamshellState | cut -d \  -f8)

CONNECTED_TO_POWER=$(pmset -g batt | head -n 1 | cut -d \' -f2)

echo "${LID_OPEN}"
echo 
echo "${CONNECTED_TO_POWER}"
echo
if [[ "${LID_OPEN}" == "No" && "${CONNECTED_TO_POWER}" == "Battery Power" ]]
then
    blueutil -p 0
    echo "Lid is closed & bluetooth should be off" >> lidclosed.txt
else
    echo "Lid is open & bluetooth should be on" >> lidopen.txt
    blueutil -p 1
fi