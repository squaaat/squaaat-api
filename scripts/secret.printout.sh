#!/bin/zsh

DD=/squaaat/squaaat-api/alpha/env

region=""
project="squaaat"
app=""
environment="alpha"


usage() {
  echo "
Description: Printout AWS System Store Manager Parameter
Usage: $(basename $0)
  -r region (default: ap-northeast-2)
  -a app (default: squaaat-api)
  -e environment (default: environment)
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

YML=$( \
aws ssm get-parameter \
  --region ${region} \
  --name "/${project}/${app}/${environment}/env" \
  --with-decryption \
  | jq -crM ".Parameter.Value" \
)

echo "${YML}"
echo "${YML}" > env.${environment}.yml
