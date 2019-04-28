# yaml-merge-cli
Merges yaml files from command line using [mergo](https://github.com/imdario/mergo)

## Installation
```
go install github.com/ericwenn/yaml-merge-cli
```

## Usage
```
usage: yaml-merge [<flags>] <files>...

merge yaml files using mergo

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Args:
  <files>  Files to merge, from left to right
```

## Example
```sh
yaml-merge-cli examples/app.yaml examples/app.override.yaml > examples/out.yaml
```

*examples/app.yaml*
```yml
runtime: python27
api_version: 1
threadsafe: true
env_variables:
  SOME_SECRET: 'some-secret'
  CORS_ALLOW_ORIGIN: '*'

handlers:
  - url: /
    script: home.app

  - url: /index\.html
    script: home.app

  - url: /stylesheets
    static_dir: stylesheets

  - url: /(.*\.(gif|png|jpg))$
    static_files: static/\1
    upload: static/.*\.(gif|png|jpg)$

  - url: /admin/.*
    script: admin.app
    login: admin

  - url: /.*
    script: not_found.app
```

*examples/app.override.yaml*
```yml
env_variables:
  CORS_ALLOW_ORIGIN: 'example.com'
```

*examples/out.yaml*
```yml
api_version: 1
env_variables:
  CORS_ALLOW_ORIGIN: example.com
  SOME_SECRET: some-secret
handlers:
- script: home.app
  url: /
- script: home.app
  url: /index\.html
- static_dir: stylesheets
  url: /stylesheets
- static_files: static/\1
  upload: static/.*\.(gif|png|jpg)$
  url: /(.*\.(gif|png|jpg))$
- login: admin
  script: admin.app
  url: /admin/.*
- script: not_found.app
  url: /.*
runtime: python27
threadsafe: true

```
