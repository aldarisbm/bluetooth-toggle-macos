LID_OPEN=$(ioreg -r -k AppleClamshellState | grep AppleClamshellState | cut -d \= -f2)

CONNECTED_TO_POWER=$(pmset -g batt | head -n 1 | cut -d \' -f2)

if [[ "${LID_OPEN}" -eq "No" && "${CONNECTED_TO_POWER}" -eq "Battery Power" ]]
    blueutil -p 0
else
    blue
