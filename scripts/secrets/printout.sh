#!/bin/zsh

DD=/squaaat/squaaat-api/alpha/env

region=""
project="squaaat"
app=""
environment="alpha"
outDir="./"

usage() {
  echo "
Description: Printout AWS System Store Manager Parameter
Usage: $(basename $0)
  -r region (default: ap-northeast-2)
  -a app (default: squaaat-api)
  -e environment (default: alpha)
  -o outDir (default: ./)
  [-h help]

Example:
  ./scripts/secrets/printout.sh -r ap-northeast-2 -a squaaat-api -e alpha
"
exit 1;
}

while getopts 'r:a:e:o:h' optname; do
  case "${optname}" in
    h) usage;;
    r) region=${OPTARG};;
    a) app=${OPTARG};;
    e) environment=${OPTARG};;
    o) outDir=${OPTARG};;
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
echo "${YML}" > ${outDir}./env.${environment}.yml
echo "${outDir}./env.${environment}.yml)"
