#! bin/bash 
ls -lapt --time=atime --group-directories-first --ignore=".*" | awk '{ printf "%s%s", $9, ($9 ~ /\/$/) ? "" : "," }' 