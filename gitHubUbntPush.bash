echo "pingpong"

git status

git add .
#git add -A
#git add -f .

echo Write your commit:
read cmmt
echo You commit: $cmmt
git commit -m "$cmmt </> $(date +"%a, %Y-%m-%d, %H:%M:%S %Z %j")"

git pull origin 925
git push origin 925
#git push -f origin 925

git status

$SHELL