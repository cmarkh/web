package web

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/cmarkh/errs"
)

// UserAgent for using with HTTP requests
var UserAgent = "Mozilla/5.0 (Windows NT 4.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2049.0 Safari/537.36"

// HTMLHeader is my prefered defaults for sending HTML emails
func HTMLHeader() string {
	return `
<html>
	<head>
		<style>
			table, th, td {
				border: 1px solid black;
				padding-left: 5px;
				padding-right: 5px;
				border-collapse: collapse;
			}
			h2 {
		    margin: 0;
			}
			a {
  			color: black;
			}
		</style>
	</head>
	<body>
		`
}

// MD5Token returns a MD5 hash timestamp (for creating unique IDs)
func MD5Token() (string, error) {
	crutime := time.Now().Unix()
	h := md5.New()
	_, err := io.WriteString(h, strconv.FormatInt(crutime, 10))
	if err != nil {
		errs.Log(err)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func PrintResp(resp *http.Response) (err error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errs.WrapErr(err)
		return
	}
	fmt.Println(string(body))

	return
}
