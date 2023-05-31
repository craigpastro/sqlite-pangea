# sqlite-pangea

An SQLite extension for calling Pangea services. This is a proof of concept
which I am "developing" in public. Not production ready :laughing:

## Usage

You will need a version of SQLite that allows loading of extensions. If you are
on a Mac, you can `brew install sqlite`.

```text
$ make 
$ sqlite3
> .load pangea.so
> insert into pangea_config (domain, token) values ('<pangea_domain>', '<pangea_token>');
> select redact('my phone number is 123-456-7890');
my phone number is <PHONE_NUMBER>
> .quit
```

For more examples, see [`test.ts`](./test.ts).

## Tests

Run `make test`.

## The roadmap maybe

- [x] [Redact](https://pangea.cloud/services/redact/)
- [ ] [Embargo](https://pangea.cloud/services/embargo-check/)
- [x] [IP Intel](https://pangea.cloud/services/ip-intel/reputation/)
- [ ] [File Intel](https://pangea.cloud/services/file-intel/)
- [ ] [Domain Intel](https://pangea.cloud/services/domain-intel/)
- [x] [URL Intel](https://pangea.cloud/services/url-intel/)
- [ ] [User Intel](https://pangea.cloud/services/user-intel/)
