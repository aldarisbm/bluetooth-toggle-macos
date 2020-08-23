LID_OPEN=$(ioreg -r -k AppleClamshellState | grep AppleClamshellState | cut -d \= -f2)

CONNECTED_TO_POWER=$(pmset -g batt | head -n 1 | cut -d \' -f2)

if [[ "${LID_OPEN}" -eq "No" && "${CONNECTED_TO_POWER}" -eq "Battery Power" ]]
then
    blueutil -p 0
    echo "Lid is closed & bluetooth should be off"
else
    echo "Lid is open & bluetooth should be resumed"
    blueutil -p 1
