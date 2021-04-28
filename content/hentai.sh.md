# hentai.sh

Automatically downloads hentai from reddit and shows you using feh. You need to have curl, jq, feh and mpv installed. Run it using bash.

[Download link](https://marekmaskarinec.github.io/files/hentai.sh)

Raw code:

```
#!/bin/bash
JSON=~/.cache/hentai.json
if [ -f $JSON ]; then if test `find $JSON -mmin +60`; then curl -s -A "hentai.sh v0.1 - linux" "https://www.reddit.com/r/hentai.json?limit=100" | sed 's/\\n/ /g' > $JSON; fi   
else curl -s -A "hentai.sh v0.1 - linux" "https://www.reddit.com/r/hentai.json?limit=100" | sed 's/\\n/ /g' > $JSON; fi
R="$(($RANDOM % $(cat $JSON | jq '.data.children | length')+1))"
URL=$(cat $JSON | jq ".data.children[$R].data.url" | sed 's/\"//g')
echo $URL
cat $JSON | jq ".data.children[$R].data.title" --raw-output
if [[ "$URL" == *".gif"* ]]; then mpv --loop $URL; exit 0; fi 
curl -sA "hentai.sh v0.1" $URL | feh -F -
```
