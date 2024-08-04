workdir=$1
if [ "$workdir" == "" ]; then workdir="/mnt/us"; fi
./anchordav-linux-arm -d="$workdir" -http=:8080 -user=password -password=user > /dev/null
