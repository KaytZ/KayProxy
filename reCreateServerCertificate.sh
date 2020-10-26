basepath=$(
  cd $(dirname "$0")
  pwd
)
echo "$basepath"
extFile="$basepath/extFile.txt"
serverCrt="$basepath/server.crt"
serverKey="$basepath/server.key"
serverCsr="$basepath/server.csr"
caCrt="$basepath/ca.crt"
caKey="$basepath/ca.key"
# 使用 CA 签发服务器证书
touch "${extFile}"
echo "basicConstraints=CA:FALSE
keyUsage=digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
extendedKeyUsage=serverAuth
subjectAltName=DNS:api.m.jd.com,DNS:m.jd.com,DNS:115.com" >"${extFile}"
openssl x509 -req -extfile "${extFile}" -days 825 -in "${serverCsr}" -CA "${caCrt}" -CAkey "${caKey}" -CAcreateserial -out "${serverCrt}"
rm -f "${extFile}"
