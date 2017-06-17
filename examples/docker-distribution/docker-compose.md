# Setting up Docker Registry2

## Tables of content


### Link

https://github.com/docker/distribution/blob/master/docs/configuration.md

https://github.com/docker/docker.github.io/blob/master/registry/insecure.md

https://github.com/docker/docker.github.io/blob/master/registry/deploying.md#native-basic-auth

## Deploy with embeded htpasswd






## Compse with BasicAuth Server

__Inspired from__

https://the.binbashtheory.com/creating-private-docker-registry-2-0-with-token-authentication-service/

__and__

https://github.com/cesanta/docker_auth

__See also__

https://stackoverflow.com/questions/10175812/how-to-create-a-self-signed-certificate-with-openssl

https://stackoverflow.com/questions/21488845/how-can-i-generate-a-self-signed-certificate-with-subjectaltname-using-openssl

https://crsr.net/Notes/SSL.html

### Prerequistes

Image - _Fedora23_
```
[vagrant@localhost ~]$ docker pull registry:2
Trying to pull repository docker.io/library/registry ... 
2: Pulling from docker.io/library/registry
79650cf9cc01: Pull complete 
70ce42745103: Pull complete 
77edd1a7fa4d: Pull complete 
432773976ace: Pull complete 
3234a47fe5a9: Pull complete 
Digest: sha256:a3551c422521617e86927c3ff57e05edf086f1648f4d8524633216ca363d06c2
Status: Downloaded newer image for docker.io/registry:2
[vagrant@localhost ~]$ docker pull cesanta/docker_auth:1
Trying to pull repository docker.io/cesanta/docker_auth ... 
1: Pulling from docker.io/cesanta/docker_auth
7520415ce762: Pull complete 
1377a09f71ba: Pull complete 
0f1847d4f768: Pull complete 
Digest: sha256:2426c23c28c18a5d68522123883bdb8ff0cc85cff8d836925fd632072b4c163a
Status: Downloaded newer image for docker.io/cesanta/docker_auth:1
```

Self-signed certificates

ssl config - _Fedora23_
```
[vagrant@localhost ~]$ [[ -f /etc/pki/tls/openssl.cnf ]] && [[ -n $(grep -e '^x509_extensions\s=\sv3_ca' /etc/pki/tls/openssl.cnf) ]] && sed "s/\[ v3_ca \]/[ alternate_names ]\n\nDNS.1 = dockerregistry2.default.svc.cluster.local\n\DNS.2 = dockerauth.default.svc.cluster.local\nIP.1 = 172.17.4.50\nIP.2 = 172.17.4.200\nIP.3 = 10.3.240.20\nIP.4 = 10.123.240.20\n\n&\n\nsubjectAltName = @alternate_names\n/" /etc/pki/tls/openssl.cnf | tee /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/openssl.cnf
#
# OpenSSL example configuration file.
# This is mostly being used for generation of certificate requests.
#

# This definition stops the following lines choking if HOME isn't
# defined.
HOME			= .
RANDFILE		= $ENV::HOME/.rnd

# Extra OBJECT IDENTIFIER info:
#oid_file		= $ENV::HOME/.oid
oid_section		= new_oids

# To use this configuration file with the "-extfile" option of the
# "openssl x509" utility, name here the section containing the
# X.509v3 extensions to use:
# extensions		= 
# (Alternatively, use a configuration file that has only
# X.509v3 extensions in its main [= default] section.)

[ new_oids ]

# We can add new OIDs in here for use by 'ca', 'req' and 'ts'.
# Add a simple OID like this:
# testoid1=1.2.3.4
# Or use config file substitution like this:
# testoid2=${testoid1}.5.6

# Policies used by the TSA examples.
tsa_policy1 = 1.2.3.4.1
tsa_policy2 = 1.2.3.4.5.6
tsa_policy3 = 1.2.3.4.5.7

####################################################################
[ ca ]
default_ca	= CA_default		# The default ca section

####################################################################
[ CA_default ]

dir		= /etc/pki/CA		# Where everything is kept
certs		= $dir/certs		# Where the issued certs are kept
crl_dir		= $dir/crl		# Where the issued crl are kept
database	= $dir/index.txt	# database index file.
#unique_subject	= no			# Set to 'no' to allow creation of
					# several ctificates with same subject.
new_certs_dir	= $dir/newcerts		# default place for new certs.

certificate	= $dir/cacert.pem 	# The CA certificate
serial		= $dir/serial 		# The current serial number
crlnumber	= $dir/crlnumber	# the current crl number
					# must be commented out to leave a V1 CRL
crl		= $dir/crl.pem 		# The current CRL
private_key	= $dir/private/cakey.pem# The private key
RANDFILE	= $dir/private/.rand	# private random number file

x509_extensions	= usr_cert		# The extentions to add to the cert

# Comment out the following two lines for the "traditional"
# (and highly broken) format.
name_opt 	= ca_default		# Subject Name options
cert_opt 	= ca_default		# Certificate field options

# Extension copying option: use with caution.
# copy_extensions = copy

# Extensions to add to a CRL. Note: Netscape communicator chokes on V2 CRLs
# so this is commented out by default to leave a V1 CRL.
# crlnumber must also be commented out to leave a V1 CRL.
# crl_extensions	= crl_ext

default_days	= 365			# how long to certify for
default_crl_days= 30			# how long before next CRL
default_md	= sha256		# use SHA-256 by default
preserve	= no			# keep passed DN ordering

# A few difference way of specifying how similar the request should look
# For type CA, the listed attributes must be the same, and the optional
# and supplied fields are just that :-)
policy		= policy_match

# For the CA policy
[ policy_match ]
countryName		= match
stateOrProvinceName	= match
organizationName	= match
organizationalUnitName	= optional
commonName		= supplied
emailAddress		= optional

# For the 'anything' policy
# At this point in time, you must list all acceptable 'object'
# types.
[ policy_anything ]
countryName		= optional
stateOrProvinceName	= optional
localityName		= optional
organizationName	= optional
organizationalUnitName	= optional
commonName		= supplied
emailAddress		= optional

####################################################################
[ req ]
default_bits		= 2048
default_md		= sha256
default_keyfile 	= privkey.pem
distinguished_name	= req_distinguished_name
attributes		= req_attributes
x509_extensions	= v3_ca	# The extentions to add to the self signed cert

# Passwords for private keys if not present they will be prompted for
# input_password = secret
# output_password = secret

# This sets a mask for permitted string types. There are several options. 
# default: PrintableString, T61String, BMPString.
# pkix	 : PrintableString, BMPString (PKIX recommendation before 2004)
# utf8only: only UTF8Strings (PKIX recommendation after 2004).
# nombstr : PrintableString, T61String (no BMPStrings or UTF8Strings).
# MASK:XXXX a literal mask value.
# WARNING: ancient versions of Netscape crash on BMPStrings or UTF8Strings.
string_mask = utf8only

# req_extensions = v3_req # The extensions to add to a certificate request

[ req_distinguished_name ]
countryName			= Country Name (2 letter code)
countryName_default		= XX
countryName_min			= 2
countryName_max			= 2

stateOrProvinceName		= State or Province Name (full name)
#stateOrProvinceName_default	= Default Province

localityName			= Locality Name (eg, city)
localityName_default		= Default City

0.organizationName		= Organization Name (eg, company)
0.organizationName_default	= Default Company Ltd

# we can do this but it is not needed normally :-)
#1.organizationName		= Second Organization Name (eg, company)
#1.organizationName_default	= World Wide Web Pty Ltd

organizationalUnitName		= Organizational Unit Name (eg, section)
#organizationalUnitName_default	=

commonName			= Common Name (eg, your name or your server\'s hostname)
commonName_max			= 64

emailAddress			= Email Address
emailAddress_max		= 64

# SET-ex3			= SET extension number 3

[ req_attributes ]
challengePassword		= A challenge password
challengePassword_min		= 4
challengePassword_max		= 20

unstructuredName		= An optional company name

[ usr_cert ]

# These extensions are added when 'ca' signs a request.

# This goes against PKIX guidelines but some CAs do it and some software
# requires this to avoid interpreting an end user certificate as a CA.

basicConstraints=CA:FALSE

# Here are some examples of the usage of nsCertType. If it is omitted
# the certificate can be used for anything *except* object signing.

# This is OK for an SSL server.
# nsCertType			= server

# For an object signing certificate this would be used.
# nsCertType = objsign

# For normal client use this is typical
# nsCertType = client, email

# and for everything including object signing:
# nsCertType = client, email, objsign

# This is typical in keyUsage for a client certificate.
# keyUsage = nonRepudiation, digitalSignature, keyEncipherment

# This will be displayed in Netscape's comment listbox.
nsComment			= "OpenSSL Generated Certificate"

# PKIX recommendations harmless if included in all certificates.
subjectKeyIdentifier=hash
authorityKeyIdentifier=keyid,issuer

# This stuff is for subjectAltName and issuerAltname.
# Import the email address.
# subjectAltName=email:copy
# An alternative to produce certificates that aren't
# deprecated according to PKIX.
# subjectAltName=email:move

# Copy subject details
# issuerAltName=issuer:copy

#nsCaRevocationUrl		= http://www.domain.dom/ca-crl.pem
#nsBaseUrl
#nsRevocationUrl
#nsRenewalUrl
#nsCaPolicyUrl
#nsSslServerName

# This is required for TSA certificates.
# extendedKeyUsage = critical,timeStamping

[ v3_req ]

# Extensions to add to a certificate request

basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment

[ alternate_names ]

DNS.1 = dockerregistry2.default.svc.cluster.local
DNS.2 = dockerauth.default.svc.cluster.local
IP.1 = 172.17.4.50
IP.2 = 172.17.4.200
IP.3 = 10.3.240.20
IP.4 = 10.123.240.20

[ v3_ca ]

subjectAltName = @alternate_names



# Extensions for a typical CA


# PKIX recommendation.

subjectKeyIdentifier=hash

authorityKeyIdentifier=keyid:always,issuer

# This is what PKIX recommends but some broken software chokes on critical
# extensions.
#basicConstraints = critical,CA:true
# So we do this instead.
basicConstraints = CA:true

# Key usage: this is typical for a CA certificate. However since it will
# prevent it being used as an test self-signed certificate it is best
# left out by default.
# keyUsage = cRLSign, keyCertSign

# Some might want this also
# nsCertType = sslCA, emailCA

# Include email address in subject alt name: another PKIX recommendation
# subjectAltName=email:copy
# Copy issuer details
# issuerAltName=issuer:copy

# DER hex encoding of an extension: beware experts only!
# obj=DER:02:03
# Where 'obj' is a standard or added object
# You can even override a supported extension:
# basicConstraints= critical, DER:30:03:01:01:FF

[ crl_ext ]

# CRL extensions.
# Only issuerAltName and authorityKeyIdentifier make any sense in a CRL.

# issuerAltName=issuer:copy
authorityKeyIdentifier=keyid:always

[ proxy_cert_ext ]
# These extensions should be added when creating a proxy certificate

# This goes against PKIX guidelines but some CAs do it and some software
# requires this to avoid interpreting an end user certificate as a CA.

basicConstraints=CA:FALSE

# Here are some examples of the usage of nsCertType. If it is omitted
# the certificate can be used for anything *except* object signing.

# This is OK for an SSL server.
# nsCertType			= server

# For an object signing certificate this would be used.
# nsCertType = objsign

# For normal client use this is typical
# nsCertType = client, email

# and for everything including object signing:
# nsCertType = client, email, objsign

# This is typical in keyUsage for a client certificate.
# keyUsage = nonRepudiation, digitalSignature, keyEncipherment

# This will be displayed in Netscape's comment listbox.
nsComment			= "OpenSSL Generated Certificate"

# PKIX recommendations harmless if included in all certificates.
subjectKeyIdentifier=hash
authorityKeyIdentifier=keyid,issuer

# This stuff is for subjectAltName and issuerAltname.
# Import the email address.
# subjectAltName=email:copy
# An alternative to produce certificates that aren't
# deprecated according to PKIX.
# subjectAltName=email:move

# Copy subject details
# issuerAltName=issuer:copy

#nsCaRevocationUrl		= http://www.domain.dom/ca-crl.pem
#nsBaseUrl
#nsRevocationUrl
#nsRenewalUrl
#nsCaPolicyUrl
#nsSslServerName

# This really needs to be in place for it to be a proxy certificate.
proxyCertInfo=critical,language:id-ppl-anyLanguage,pathlen:3,policy:foo

####################################################################
[ tsa ]

default_tsa = tsa_config1	# the default TSA section

[ tsa_config1 ]

# These are used by the TSA reply generation only.
dir		= ./demoCA		# TSA root directory
serial		= $dir/tsaserial	# The current serial number (mandatory)
crypto_device	= builtin		# OpenSSL engine to use for signing
signer_cert	= $dir/tsacert.pem 	# The TSA signing certificate
					# (optional)
certs		= $dir/cacert.pem	# Certificate chain to include in reply
					# (optional)
signer_key	= $dir/private/tsakey.pem # The TSA private key (optional)

default_policy	= tsa_policy1		# Policy if request did not specify it
					# (optional)
other_policies	= tsa_policy2, tsa_policy3	# acceptable policies (optional)
digests		= sha1, sha256, sha384, sha512	# Acceptable message digests (mandatory)
accuracy	= secs:1, millisecs:500, microsecs:100	# (optional)
clock_precision_digits  = 0	# number of digits after dot. (optional)
ordering		= yes	# Is ordering defined for timestamps?
				# (optional, default: no)
tsa_name		= yes	# Must the TSA name be included in the reply?
				# (optional, default: no)
ess_cert_id_chain	= no	# Must the ESS cert id chain be included?
				# (optional, default: no)
```

Generate
```
[vagrant@localhost ~]$ cert_dir=/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution; openssl req -new -newkey rsa:4096 -days 3650 -nodes -x509 -subj "/CN=172.17.4.50:5000/O=stackdocker" -config "${cert_dir}/openssl.cnf" -keyout "${cert_dir}/server.key" -out "${cert_dir}/server.cert"
Generating a 4096 bit RSA private key
................................................................................................................................++
...............................................................................................................................................................................................................................................++
writing new private key to '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/server.key'
-----
```

Print
```
[vagrant@localhost ~]$ openssl x509 -in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/server.cert -text -noout
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 10079960323688096441 (0x8be336822b75e2b9)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=172.17.4.50:5000, O=stackdocker
        Validity
            Not Before: Jun 16 09:25:56 2017 GMT
            Not After : Jun 14 09:25:56 2027 GMT
        Subject: CN=172.17.4.50:5000, O=stackdocker
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (4096 bit)
                Modulus:
                    00:be:5e:cc:f5:91:86:2f:59:6f:8c:4e:44:ac:cd:
                    e0:8c:38:48:0e:57:bc:96:c6:de:37:df:ee:a2:75:
                    c2:99:12:31:44:dc:b0:4d:fa:ee:60:20:c2:d4:de:
                    8d:8b:46:74:7d:03:fd:e9:eb:e8:bb:72:f9:b6:e9:
                    44:6f:13:a8:9a:33:5c:f5:1c:71:a5:85:92:a2:9c:
                    2b:86:08:bb:4e:5a:78:e3:7e:5e:94:5f:ab:63:89:
                    2a:f4:af:12:d9:40:4d:63:9f:ba:6c:0d:21:03:ae:
                    28:3b:d6:1d:79:2d:28:ce:3b:60:9e:38:78:53:89:
                    36:0c:1f:ed:65:81:22:92:43:e4:43:87:6d:f0:55:
                    95:09:ac:8a:13:d2:4e:c3:1a:ac:77:af:b6:7d:a8:
                    e6:74:50:70:21:d7:5f:8b:ee:de:fd:85:c4:eb:32:
                    d3:6d:d9:8c:e3:f7:c8:a8:bb:20:01:86:10:0d:76:
                    ea:2a:16:dd:f1:30:14:91:b3:e9:7b:c9:4b:b5:48:
                    d6:1b:c3:a6:f7:33:81:b9:01:c2:c8:71:8d:07:e8:
                    e7:4c:d8:35:b2:56:0a:e5:97:fc:ed:65:af:5b:d9:
                    d9:35:d2:c7:5a:de:d1:3a:dc:53:3c:5b:21:b8:65:
                    96:90:74:fe:c2:6b:b7:01:73:49:5b:a8:3e:9b:d4:
                    f4:39:45:55:c7:e1:04:06:e7:fd:6f:44:a0:f5:ad:
                    49:da:5d:0c:24:89:91:d5:48:13:18:ef:0f:c7:32:
                    10:7f:9c:d9:27:6e:81:b6:53:62:0f:8e:0a:18:4a:
                    cd:1b:9d:5d:d2:0c:b7:e5:b6:dd:47:a3:e0:29:f5:
                    9c:98:c2:36:96:83:4c:c8:44:cf:1e:2a:99:74:28:
                    8f:43:64:83:e9:87:04:cc:c6:57:70:69:26:c4:76:
                    6a:6d:79:f2:b2:65:bc:87:4e:96:2e:4e:b6:31:bc:
                    02:4a:88:4b:3f:40:55:e7:a6:05:bf:5c:ca:d7:44:
                    53:7e:b7:40:b3:ff:ee:f5:30:9e:54:5a:ad:05:f4:
                    f3:d0:84:b3:8d:85:f8:66:13:b5:1f:1b:18:7b:04:
                    ac:69:46:72:2f:d6:bf:ac:46:a4:45:58:ad:72:8b:
                    ac:ec:5c:3e:a8:50:64:88:fe:36:73:7c:8b:c2:3a:
                    33:d4:82:fd:69:18:50:17:e5:77:40:e7:1d:ee:fc:
                    1f:1d:cd:a8:ab:12:ff:68:2e:74:77:d7:a1:3f:26:
                    48:7f:72:f7:fe:5f:04:6d:76:92:9c:52:a5:dc:81:
                    23:ff:c3:3c:e0:7c:c8:51:a2:65:a7:a2:1e:67:26:
                    35:8d:49:c1:fc:08:79:b8:69:84:3b:0b:d5:04:c5:
                    5f:7a:1b
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Subject Alternative Name: 
                DNS:dockerregistry2.default.svc.cluster.local, DNS:dockerauth.default.svc.cluster.local, IP Address:172.17.4.50, IP Address:172.17.4.200, IP Address:10.3.240.20, IP Address:10.123.240.20
            X509v3 Subject Key Identifier: 
                40:44:6B:2C:50:E8:2B:7D:06:10:7C:6C:AA:5C:58:B0:EF:ED:78:0C
            X509v3 Authority Key Identifier: 
                keyid:40:44:6B:2C:50:E8:2B:7D:06:10:7C:6C:AA:5C:58:B0:EF:ED:78:0C

            X509v3 Basic Constraints: 
                CA:TRUE
    Signature Algorithm: sha256WithRSAEncryption
         a1:b3:76:f5:26:8a:d5:b5:48:b3:9d:44:37:08:62:1d:d2:ac:
         9b:2f:49:7f:30:6c:b0:df:bd:94:35:d3:22:c2:cd:cb:71:8d:
         38:d1:83:fa:f9:10:75:5c:3b:b3:10:27:ce:9f:34:83:65:ba:
         93:41:51:cb:68:5a:85:63:d2:df:19:9b:39:f2:4a:1f:d5:e3:
         fa:4b:32:36:6e:80:99:85:14:85:a5:71:2c:20:da:b9:1a:a2:
         3f:81:83:28:14:95:09:11:43:f4:9f:22:ae:72:67:a1:fd:0c:
         ee:dc:73:db:ca:41:0b:aa:6a:7f:85:b2:83:4c:3b:4d:2d:0c:
         bf:cb:6a:84:70:33:20:80:07:55:cc:b5:5c:ea:16:4e:89:72:
         f8:ca:2c:8d:73:70:89:19:58:ca:09:ba:ed:a6:da:61:73:0f:
         a0:04:b6:b2:ea:e6:4e:e8:71:6b:65:3d:df:57:09:33:55:32:
         f1:9c:d9:14:52:e9:63:9b:ef:a4:4c:be:73:af:0b:6e:6c:35:
         18:fd:2a:fe:94:dc:7e:91:df:c4:c4:a4:d7:95:12:7b:78:27:
         03:f3:f8:5f:72:63:1c:16:0d:c5:db:e0:94:7b:a7:d6:1d:9a:
         b3:81:ab:56:75:4b:93:c5:06:6b:6c:7d:bb:63:ec:71:3f:b2:
         dc:5a:44:e9:48:73:24:48:6e:d3:92:83:69:23:ad:86:71:a8:
         07:6a:b3:8d:15:e8:c2:50:6b:57:66:56:28:66:a8:cd:18:9e:
         23:ed:1d:de:d5:be:c7:62:5d:26:d9:ac:0a:99:70:f2:d3:f9:
         22:bc:fe:1a:2f:04:14:35:b2:6d:dd:78:b4:47:df:00:13:80:
         77:f3:ca:26:f4:a6:93:f3:9a:2e:2c:b8:8e:6b:4b:81:35:6a:
         79:e5:87:df:10:90:69:3c:e3:3d:b8:d2:c7:1a:42:e7:72:93:
         33:ae:3c:b9:cc:c9:75:ee:21:0d:2b:db:f7:d4:67:2f:51:2a:
         5a:9d:92:dc:6c:e4:98:f5:db:7c:85:b6:30:e4:c0:1b:df:1d:
         c4:b6:bf:49:1c:f8:73:39:62:d8:b9:60:22:db:b1:e1:39:ab:
         9d:91:34:f7:88:69:fe:ec:90:99:eb:ee:5b:98:ae:37:fd:7f:
         93:0d:97:85:bf:7a:5e:e5:8c:a6:b5:09:fd:72:b7:3e:7e:15:
         b8:f9:0f:32:95:cd:eb:d5:cc:0c:f2:e0:08:ba:ba:b4:64:aa:
         c5:b4:75:3f:a3:06:fb:50:7c:e4:8c:a2:57:0d:bf:d2:11:06:
         95:36:b9:1b:f2:7e:3a:d3:21:7d:25:a7:8e:71:3e:60:be:fe:
         9e:d2:f9:81:93:e3:e8:5b
```

Enable
```
[vagrant@localhost ~]$ sudo cp /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/server.cert /etc/docker/certs.d/172.17.4.50\:5000/ca.crt
[vagrant@localhost ~]$ sudo systemctl restart docker.service
```

Registry config
```
[vagrant@localhost ~]$ curl -jkSL https://raw.githubusercontent.com/docker/distribution/master/cmd/registry/config-example.yml -o /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/registry_config.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   295  100   295    0     0    211      0  0:00:01  0:00:01 --:--:--   211
```

Compose
```
[vagrant@localhost ~]$ docker-compose --project-name registry2 --file /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/docker-compose.yml up
Recreating registry2_dockerauth_1 ... 
Starting registry2_registry_1 ... 
Recreating registry2_dockerauth_1
Recreating registry2_dockerauth_1 ... done
Attaching to registry2_registry_1, registry2_dockerauth_1
registry_1    | time="2017-06-16T11:24:11Z" level=warning msg="No HTTP secret provided - generated random secret. This may cause problems with uploads if multiple registries are behind a load-balancer. To provide a shared secret, fill in http.secret in the configuration file or set the REGISTRY_HTTP_SECRET environment variable." go.version=go1.7.3 instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | time="2017-06-16T11:24:11Z" level=info msg="redis not configured" go.version=go1.7.3 instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | time="2017-06-16T11:24:11Z" level=info msg="Starting upload purge in 24m0s" go.version=go1.7.3 instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | time="2017-06-16T11:24:11Z" level=info msg="using inmemory blob descriptor cache" go.version=go1.7.3 instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | time="2017-06-16T11:24:11Z" level=info msg="listening on [::]:5000, tls" go.version=go1.7.3 instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | time="2017-06-16T11:24:54Z" level=warning msg="error authorizing context: authorization token required" go.version=go1.7.3 http.request.host="172.17.4.50:5000" http.request.id=90fe6dbb-67b7-4be0-b6d0-491c73b18094 http.request.method=GET http.request.remoteaddr="172.17.4.50:35096" http.request.uri="/v2/" http.request.useragent="docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64" instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | 172.17.4.50 - - [16/Jun/2017:11:24:54 +0000] "GET /v2/ HTTP/1.1" 401 87 "" "docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64"
registry_1    | time="2017-06-16T11:24:54Z" level=info msg="response completed" go.version=go1.7.3 http.request.host="172.17.4.50:5000" http.request.id=50b0be4e-738e-46d8-ba1a-ea6c39598a52 http.request.method=GET http.request.remoteaddr="172.17.4.50:35100" http.request.uri="/v2/" http.request.useragent="docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64" http.response.contenttype="application/json; charset=utf-8" http.response.duration=1.370898ms http.response.status=200 http.response.written=2 instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | 172.17.4.50 - - [16/Jun/2017:11:24:54 +0000] "GET /v2/ HTTP/1.1" 200 2 "" "docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64"
registry_1    | time="2017-06-16T11:25:30Z" level=warning msg="error authorizing context: authorization token required" go.version=go1.7.3 http.request.host="172.17.4.50:5000" http.request.id=ec35d10c-3580-4575-ba45-a26a30f3c06c http.request.method=GET http.request.remoteaddr="172.17.4.50:35102" http.request.uri="/v2/" http.request.useragent="docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64" instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | 172.17.4.50 - - [16/Jun/2017:11:25:30 +0000] "GET /v2/ HTTP/1.1" 401 87 "" "docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64"
registry_1    | time="2017-06-16T11:25:30Z" level=info msg="response completed" go.version=go1.7.3 http.request.host="172.17.4.50:5000" http.request.id=018001fd-e5f7-4654-b035-44ab6f6801fb http.request.method=GET http.request.remoteaddr="172.17.4.50:35106" http.request.uri="/v2/" http.request.useragent="docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64" http.response.contenttype="application/json; charset=utf-8" http.response.duration=1.114731ms http.response.status=200 http.response.written=2 instance.id=7502f8dc-964b-4c10-80b2-6680e393ae21 version=v2.6.1 
registry_1    | 172.17.4.50 - - [16/Jun/2017:11:25:30 +0000] "GET /v2/ HTTP/1.1" 200 2 "" "docker/1.10.3 go/go1.5.4 kernel/4.2.3-300.fc23.x86_64 os/linux arch/amd64"
```

Login
```
[vagrant@localhost ~]$ docker login 172.17.4.50\:5000
Username (tangfx): tangfx
Password: 
WARNING: login credentials saved in /home/vagrant/.docker/config.json
Login Succeeded
[vagrant@localhost ~]$ docker login 172.17.4.50:5000
Username (tangfx): 
Password: 
WARNING: login credentials saved in /home/vagrant/.docker/config.json
Login Succeeded
[vagrant@localhost ~]$ docker login 172.17.4.50:5000
Username (tangfx): admin
Password: 
Email: 
WARNING: login credentials saved in /home/vagrant/.docker/config.json
Login Succeeded
[vagrant@localhost ~]$ docker login 172.17.4.50:5000
Username (admin): test
Password: 
Email: 
WARNING: login credentials saved in /home/vagrant/.docker/config.json
Login Succeeded
[vagrant@localhost ~]$ docker login 172.17.4.50:5000
Username (test): tangfx
Password: 
Email: 
WARNING: login credentials saved in /home/vagrant/.docker/config.json
Login Succeeded
```

Curl
```
[vagrant@localhost ~]$ curl --cacert /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/server.cert --cert /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/server.cert --key /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/server.key --user admin:password -SL https://172.17.4.50:5001/auth
{"token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik1WRDQ6QjZCQjo0SFc3OktSUU46T1JMWDpQTE5IOkRQMkI6MkZGRjpUU1BCOlZHWlE6RzQ1VjpWT0tNIn0.eyJpc3MiOiJBdXRoIFNlcnZpY2UiLCJzdWIiOiJhZG1pbiIsImF1ZCI6IiIsImV4cCI6MTQ5NzYxNTQ4NSwibmJmIjoxNDk3NjE0NTc1LCJpYXQiOjE0OTc2MTQ1ODUsImp0aSI6IjY2NjMzNzA3MzA4NTA1NzQwODUiLCJhY2Nlc3MiOltdfQ.lTTIKYJirRsY9EN8lfui-Scgg0dbl0F9qn37tzNB793uebfviakemugYz3M3F3JrupIcaps01CVTtXmVQJJN7pA7nsTH-EVyG3w2HFh5pkXEEKFBGCT0s-HyikIMqxZeZyIJDcFBIxiS32PFTszRhJOsaBDOYNbzF3V2C_iubjo-NGZX0sWC-xeHG9Avng87CGVzoZ2S2NPypzz0wiPCmwRTjoRnNkVDcLJI_5b-VDBz70O4UYmzp1nGcXfYi7dVdXL7iqZPraxV2K7tCdrg-JajOZbzTT1SVShMPPqpJTm0VjmgcVXLVI-fvB5ta6e3DfCdUUyrTsrPgd5W-UkkEsnAskH4J-76fRBj7e9W2Cbti52B5-QsapUbTwNd7Ydn5KqOmdUAB-P7yFxicehoNLyZqrOEL4Z1DE6AIpd8nf0ut2gi0Doxtry8v97Ux23Ksj3Q41WYCj_f11_yV_PX1-Esl-KpSB6T2t45EDUXyTQdtF2JR6YkJmzXOKRPxPNBX9LEWTfbgLTKJqNcw7DR-H6XReZpGl77cNzP43Pa-rHTyWVxF7hVDyWoNd57e-g573gLLL2_b0U964bzytDWRQQ6OcYvBr_QhbyvCJGJvLgUfET3Snt7bcdyNWJb9ELCj6Q0lMoLcKsLQOLkPo6kFJlM6mLUE6pNITnUiQyIBh4"}
[vagrant@localhost ~]$ curl --cacert /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-distribution/server.cert --header "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik1WRDQ6QjZCQjo0SFc3OktSUU46T1JMWDpQTE5IOkRQMkI6MkZGRjpUU1BCOlZHWlE6RzQ1VjpWT0tNIn0.eyJpc3MiOiJBdXRoIFNlcnZpY2UiLCJzdWIiOiJhZG1pbiIsImF1ZCI6IiIsImV4cCI6MTQ5NzYxNTQ4NSwibmJmIjoxNDk3NjE0NTc1LCJpYXQiOjE0OTc2MTQ1ODUsImp0aSI6IjY2NjMzNzA3MzA4NTA1NzQwODUiLCJhY2Nlc3MiOltdfQ.lTTIKYJirRsY9EN8lfui-Scgg0dbl0F9qn37tzNB793uebfviakemugYz3M3F3JrupIcaps01CVTtXmVQJJN7pA7nsTH-EVyG3w2HFh5pkXEEKFBGCT0s-HyikIMqxZeZyIJDcFBIxiS32PFTszRhJOsaBDOYNbzF3V2C_iubjo-NGZX0sWC-xeHG9Avng87CGVzoZ2S2NPypzz0wiPCmwRTjoRnNkVDcLJI_5b-VDBz70O4UYmzp1nGcXfYi7dVdXL7iqZPraxV2K7tCdrg-JajOZbzTT1SVShMPPqpJTm0VjmgcVXLVI-fvB5ta6e3DfCdUUyrTsrPgd5W-UkkEsnAskH4J-76fRBj7e9W2Cbti52B5-QsapUbTwNd7Ydn5KqOmdUAB-P7yFxicehoNLyZqrOEL4Z1DE6AIpd8nf0ut2gi0Doxtry8v97Ux23Ksj3Q41WYCj_f11_yV_PX1-Esl-KpSB6T2t45EDUXyTQdtF2JR6YkJmzXOKRPxPNBX9LEWTfbgLTKJqNcw7DR-H6XReZpGl77cNzP43Pa-rHTyWVxF7hVDyWoNd57e-g573gLLL2_b0U964bzytDWRQQ6OcYvBr_QhbyvCJGJvLgUfET3Snt7bcdyNWJb9ELCj6Q0lMoLcKsLQOLkPo6kFJlM6mLUE6pNITnUiQyIBh4" https://172.17.4.50:5000/v2/_catalog
{"errors":[{"code":"UNAUTHORIZED","message":"authentication required","detail":[{"Type":"registry","Class":"","Name":"catalog","Action":"*"}]}]}
```

Push
```
[vagrant@localhost ~]$ docker login 172.17.4.50:5000
Username (tangfx): admin
Password: 
Email: 
WARNING: login credentials saved in /home/vagrant/.docker/config.json
Login Succeeded
[vagrant@localhost ~]$ docker push 172.17.4.50:5000/tangfx/osospringbootapp
The push refers to a repository [172.17.4.50:5000/tangfx/osospringbootapp]
d40d824227e3: Layer already exists 
5f70bf18a086: Layer already exists 
ee439127ef0d: Layer already exists 
3ec1888d4c01: Layer already exists 
3752804e1d78: Layer already exists 
cb752db84dc2: Layer already exists 
0e7bfacd004c: Layer already exists 
f72f4c12a35d: Layer already exists 
f59b7e59ceaa: Layer already exists 
latest: digest: sha256:2ac98a59fcbd65c3a8f88c502e3335e9da021404fee38b1e4aec018ec3919f53 size: 6508
```