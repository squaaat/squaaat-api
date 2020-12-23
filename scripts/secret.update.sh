#!/bin/zsh

DD=/squaaat/squaaat-api/alpha/env

region=""
project="squaaat"
app=""
environment=""

usage() {
  echo "
Description: Create AWS System Store Manager Parameter
Usage: $(basename $0)
  -r region (default: ap-northeast-2)
  -a app (default: squaaat-api)
  -e stage (default: alpha)
  [-h help]
"
exit 1;
}

while getopts 'r:a:e:h' optname; do
  case "${optname}" in
    h) usage;;
    r) region=${OPTARG};;
    a) app=${OPTARG};;
    e) environment=${OPTARG};;
    *) usage;;
  esac
done

[ -z "${app}" ] && >&2 echo "Error: -n app required" && usage
[ -z "${environment}" ] && >&2 echo "Error: -m environment required" && usage

echo "/${project}/${app}/${environment}/env"

YML="$(cat ./env.${environment}.yml)"
echo "${YML}"

echo "- Output -------------------------------"

aws ssm put-parameter \
  --region ${region} \
  --name "/${project}/${app}/${environment}/env" \
  --type "SecureString" \
  --value "${YML}" \
  --overwrite | jq
