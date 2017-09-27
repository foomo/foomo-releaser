# Foomo Releaser Application

Used to pbulish releases to github repositories according to the foomo solution

## Homebrew

To install from homebrew

```
$ brew install foomo/foomo-releaser/foomo-releaser
```


## Ubuntu / CI

```
export RELEASER_VERSION = {VERSION}
curl "https://github.com/foomo/foomo-releaser/releases/download/${RELEASER_VERSION}/foomo-releaser_${RELEASER_VERSION}_linux_amd64.tar.gz" | tar xz

```