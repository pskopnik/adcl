// Package maybe contains data types which may or not be set.
// Its content is generated from the generic/ sub-package.
package maybe

//go:generate sh -c "genny -in=generic/maybe.go gen 'Type=BUILTINS' | sed s/Maybe//g > gen-builtins.go"
//go:generate sh -c "genny -in=generic/maybe.go gen 'Type=*encoding.Base32Value' | sed s/MaybeEncodingBase32Value/Base32Value/g > gen-base32value.go"
//go:generate sh -c "genny -in=generic/maybe.go gen 'Type=net.IP' | sed s/MaybeNetIP/IP/g > gen-ip.go"
