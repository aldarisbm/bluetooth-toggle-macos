#!/bin/bash
IS_LID_CLOSED=$(ioreg -r -k AppleClamshellState | grep AppleClamshellState | cut -d \  -f8)

CONNECTED_TO_POWER=$(pmset -g batt | head -n 1 | cut -d \' -f2)

echo "${IS_LID_CLOSED}"
echo 
echo "${CONNECTED_TO_POWER}"
echo
if [[ "${IS_LID_CLOSED}" == "Yes" ]] && [[ "${CONNECTED_TO_POWER}" == "Battery Power" ]]
then
    blueutil -p 0
    echo "Lid is closed & bluetooth should be off"
else
    echo "Lid is open & bluetooth should be on"
    blueutil -p 1
fi