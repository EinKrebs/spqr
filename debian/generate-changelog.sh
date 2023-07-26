#!/bin/bash
set -ex

cat - debian/changelog > tmp.txt <<EOF 
spqr (0.0.6) stable; urgency=low 

  $(git log --pretty="* %s" $(cat debian/changelog | head -n 1 | awk '{print(substr($2, 2, length($2) - 2))}')..$(git describe --tags --abbrev=0)) 

 -- Kirill Reshke <reshke@double.cloud>  $(date +'%a, %d %b %Y %H:%M:%S %z') 

EOF
mv tmp.txt debian/changelog
