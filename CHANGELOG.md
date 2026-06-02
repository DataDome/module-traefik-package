# DataDome Traefik plugin

## v1.2.0 (2026-06-02)

- Remove hard-coded status codes in favor of DataDome Protection API response headers, enabling seamless support for upcoming features
- Increase length limit from 128 to 512 characters for DataDome cookie to support upcoming features
- Reduce log verbosity in `httpClient`

## v1.1.1 (2026-05-13)

- Reduce log verbosity in `client.handler`

## v1.1.0 (2026-05-05)

- Collect [Sec-Fetch-Storage-Access header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Sec-Fetch-Storage-Access) from requests

## v1.0.1 (2026-02-23)

- Fix package naming in code files to be referenced in the Traefik Plugin catalog

## v1.0.0 (2026-02-16)

- Initial release
