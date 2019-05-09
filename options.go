package main

import (
	"flag"

	"github.com/spf13/viper"
)

const bucketName = "couchbase.bucket"
const username = "couchbase.username"
const password = "couchbase.password"
const connStr = "couchbase.connString"
const useCertAuth = "couchbase.useCertAuth"
const useAnalytics = "couchbase.useAnalytics"
const n1qlFallback = "couchbase.n1qlFallback"

type Options struct {
	ConnStr         string
	Username        string
	Password        string
	BucketName      string
	UseCertAuth     bool
	UseAnalytics    bool
	UseN1QLFallback bool
}

func (opt *Options) AddFlags(flagSet *flag.FlagSet) {
}

func (opt *Options) InitFromViper(v *viper.Viper) {
	opt.ConnStr = v.GetString(connStr)
	opt.Username = v.GetString(username)
	opt.Password = v.GetString(password)
	opt.BucketName = v.GetString(bucketName)
	opt.UseCertAuth = v.GetBool(useCertAuth)
	opt.UseAnalytics = v.GetBool(useAnalytics)
	opt.UseN1QLFallback = v.GetBool(n1qlFallback)
}
