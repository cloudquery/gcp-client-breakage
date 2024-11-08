# Reproduction scenario for https://github.com/googleapis/google-cloud-go/issues/11089

1. Install [squid](https://www.squid-cache.org/Versions/)
2. Enable the access log configuration for `squid` by running:

```bash
echo "\naccess_log $(pwd)/squid_access.log squid" >> squid/squid.conf
```

3. Run `squid -f squid/squid.conf` to start the proxy
4. Run the following command from the root of the repository, see that it ends without errors

```bash
cd old && GOOGLE_CLOUD_PROJECT=<replace-with-project> https_proxy=http://localhost:3128 go run main.go
```

5. Run the following command from the root of the repository, see that it hangs since the proxy blocks IPs, but allows domains (you can see the access logs in `squid_access.log`)

```bash
cd new && GOOGLE_CLOUD_PROJECT=<replace-with-project> https_proxy=http://localhost:3128  go run main.go
```

The differences are that the `old` directory uses the `cloud.google.com/go/accessapproval@v1.8.0` that references `cloud.google.com/go/auth@v0.9.0`, and the `new` directory uses the `cloud.google.com/go/accessapproval@v1.8.1` that references `cloud.google.com/go/auth@v0.10.1`.

Setting `export GOOGLE_API_GO_EXPERIMENTAL_DISABLE_NEW_AUTH_LIB=true` will make the `new` directory work as expected.
