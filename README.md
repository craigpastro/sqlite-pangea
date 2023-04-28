# sqlite-pangea

An SQLite extension for calling Pangea services. This is a proof of concept which I am "developing" in public. Not production ready :laughing:

## Usage

You will need a token for a project that is in us-west-1 (i.e., the domain is `aws.us.pangea.cloud`), enabled for the services that you would like to use.

```text
$ make 
$ sqlite3
> .load pangea.so
> select redact(PANGEA_TOKEN, 'my phone number is 123-456-7890');
my phone number is <PHONE_NUMBER>
> .quit
```

## Tests

Run `make test`.

## The roadmap maybe

[x] [Redact](https://pangea.cloud/services/redact/)
[ ] [Embargo](https://pangea.cloud/services/embargo-check/)
[ ] [IP Intel](https://pangea.cloud/services/ip-intel/reputation/)
[ ] [File Intel](https://pangea.cloud/services/file-intel/)
[ ] [Domain Intel](https://pangea.cloud/services/domain-intel/)
[ ] [URL Intel](https://pangea.cloud/services/url-intel/)
[ ] [User Intel](https://pangea.cloud/services/user-intel/)
