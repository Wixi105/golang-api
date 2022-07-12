!#/bin/bash

echo "# golang-api" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/Wixi105/golang-api.git
git push -u origin main
