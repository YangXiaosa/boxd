#!/bin/bash -e

ERROR_COUNT=0
while read -r file
do
	case "$(head -1 "${file}")" in
		*"Copyright (c) "*" ContentBox Authors.")
			# everything's cool
			;;
		*)
			echo "$file is missing license header."
			(( ERROR_COUNT++ ))
			;;
	esac
done < <(git ls-files "*\.go" | grep -v ".pb.go" | grep -v ".pb.gw.go")

exit $ERROR_COUNT