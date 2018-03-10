#!/bin/bash

FUNC_NAME="MentionNotifier"
CRON_NAME="every-minute"
ROLE_NAME="lambda_esscreener"

usage(){
	echo "Usage: "
	exit 1
}

has(){
	type ${1} >/dev/null 2>&1 || { echo "Command not found: ${1}" 1>&2; exit 1; }
}

create_function(){
	local ROLE=$(aws iam list-roles | jq -r ".Roles[] | select(.RoleName == \"${ROLE_NAME}\") | .Arn")
	cat env.json | sed -e "s/${FUNC_NAME}/Variables/" > env.tmp.json
	aws lambda create-function \
		--function-name "${FUNC_NAME}" \
		--runtime "go1.x" \
		--role "${ROLE}" \
		--handler "main" \
		--zip-file fileb://main.zip \
		--environment file://env.tmp.json
	rm env.tmp.json

	event-enable
}

event_enable(){
	add-permission
	local ARN=$(aws lambda list-functions | jq -r ".Functions[] | select(.FunctionName == \"${FUNC_NAME}\") | .FunctionArn")
	aws events enable-rule --name "every-minute"
	aws events put-targets --rule every-minute --targets "Id"="lowply-${FUNC_NAME}","Arn"="${ARN}"
}

event_disable(){
	remove-permission
	aws events disable-rule --name "every-minute"
	aws events remove-targets --rule "every-minute" --ids "lowply-${FUNC_NAME}"
}

list_all(){
	aws events list-targets-by-rule --rule every-minute | jq .
	aws lambda list-functions | jq ".Functions[] | select(.FunctionName == \"${FUNC_NAME}\")"

}

update(){
	cat env.json | sed -e "s/${FUNC_NAME}/Variables/" > env.tmp.json
	aws lambda update-function-configuration --function-name ${FUNC_NAME} --environment file://env.tmp.json
	rm env.tmp.json
	aws lambda update-function-code --function-name ${FUNC_NAME} --zip-file fileb://main.zip
}

run(){
	aws lambda invoke --function-name ${FUNC_NAME} local.log
}

build(){
	make lambda
}

sam-local(){
	build
	aws-sam-local local invoke "${FUNC_NAME}" \
		--env-vars env.json \
		--event event.json \
		--log-file local.log
}

add-permission(){
	local ARN=$(aws events list-rules | jq -r ".Rules[] | select(.Name == \"${CRON_NAME}\") | .Arn")
	aws lambda add-permission \
		--function-name "${FUNC_NAME}" \
		--statement-id "lowply-${FUNC_NAME}" \
		--action "lambda:InvokeFunction" \
		--principal events.amazonaws.com \
		--source-arn "${ARN}"
}

remove-permission(){
	aws lambda remove-permission \
		--function-name "${FUNC_NAME}" \
		--statement-id "lowply-${FUNC_NAME}"
}

main(){
	has aws
	has aws-sam-local
	HANDLER="${1}"
	shift

	case ${HANDLER} in
		"create")
			create_function
		;;
		"enable")
			event_enable
		;;
		"disable")
			event_disable
		;;
		"list")
			list_all
		;;
		"update")
			update
		;;
		"run")
			run
		;;
		"local")
			sam-local
		;;
		"add-permission")
			add-permission
		;;
		"remove-permission")
			remove-permission
		;;
		*)
			usage
		;;
	esac
}

main $@
