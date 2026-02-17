# DataDome Traefik Plugin

Bot protection middleware for Traefik with DataDome.
Detects and blocks automated threats, bots, and malicious traffic in real-time.

## Installation

### Prerequisites

- Traefik v2.3+
- A DataDome Server Side Key ([get your key](https://app.datadome.co))

### Adding the Plugin

Add the DataDome plugin to your Traefik static configuration:

**File (YAML)**

```yaml
experimental:
  plugins:
    datadome:
      moduleName: github.com/DataDome/module-traefik-package
      version: v1.0.0
```

**File (TOML)**

```toml
[experimental.plugins.datadome]
  moduleName = "github.com/DataDome/module-traefik-package"
  version = "v1.0.0"
```

**CLI**

```bash
--experimental.plugins.datadome.modulename=github.com/DataDome/module-traefik-package
--experimental.plugins.datadome.version=v1.0.0
```

### Configuring the Middleware

After adding the plugin, configure it as a middleware in your dynamic configuration:

**File (YAML)**

```yaml
http:
  middlewares:
    datadome-protection:
      plugin:
        datadome:
          serverSideKey: "YOUR_DATADOME_SERVER_SIDE_KEY"
          timeout: 300
```

**File (TOML)**

```toml
[http.middlewares]
  [http.middlewares.datadome-protection.plugin.datadome]
    serverSideKey = "YOUR_DATADOME_SERVER_SIDE_KEY"
    timeout = 300
```

**Docker Labels**

```yaml
labels:
  - "traefik.http.middlewares.datadome-protection.plugin.datadome.serverSideKey=YOUR_DATADOME_SERVER_SIDE_KEY"
  - "traefik.http.middlewares.datadome-protection.plugin.datadome.timeout=300"
```

### Applying the Middleware

Apply the middleware to your routes:

```yaml
http:
  routers:
    my-router:
      rule: "Host(`example.com`)"
      service: my-service
      middlewares:
        - datadome-protection
```

**Important:** Restart Traefik after adding or modifying plugins to load the changes.

## Configuration

| Field | Type | Required | Default | Description |
| ----- | ---- | -------- | ------- | ----------- |
| `serverSideKey` | string | Yes | - | Your DataDome Server Side Key. |
| `enableGraphQLSupport` | boolean | No | `false` | Enables the support of GraphQL requests. |
| `enableReferrerRestoration` | boolean | No | `false` | Restores original referrer after a challenge is passed. |
| `endpoint` | string | No | `api.datadome.co` | Host of the Protection API. |
| `maximumBodySize` | integer | No | 25 Kb | Maximum request body size (in bytes) to analyze. |
| `timeout` | integer | No | `150` | Timeout in milliseconds, after which the request will be allowed. |
| `urlPatternExclusion` | string | No | `(?i)\.(avi\|flv\|mka\|mkv\|mov\|mp4\|mpeg\|mpg\|mp3\|flac\|ogg\|ogm\|opus\|wav\|webm\|webp\|bmp\|gif\|ico\|jpeg\|jpg\|png\|svg\|svgz\|swf\|eot\|otf\|ttf\|woff\|woff2\|css\|less\|js\|map\|json)$` | Regex to match to exclude requests from being processed with the Protection API. If not defined, all requests will be processed. |
| `urlPatternInclusion` | string | No | - | Regex to match to process the request with the Protection API. If not defined, all requests that don't match `urlPatternExclusion` will be processed. |
| `useXForwardedHost` | boolean | No | `false` | Use the `X-Forwarded-Host` header instead of the `Host` header when the application is behind a reverse proxy/load balancer. |

### Configuration Example

```yaml
http:
  middlewares:
    datadome-protection:
      plugin:
        datadome:
          serverSideKey: "YOUR_DATADOME_SERVER_SIDE_KEY"
          timeout: 300
          enableGraphQLSupport: true
          useXForwardedHost: true
```

## How It Works

The DataDome plugin acts as a middleware in your Traefik routing pipeline:

1. **Request Interception**: All incoming HTTP requests pass through the DataDome middleware
2. **Bot Detection**: DataDome analyzes request patterns, headers, and behavior in real-time
3. **Decision**: The plugin receives a decision from DataDome's API (allow or block)
4. **Action**: Blocked requests receive a challenge page or are denied; legitimate requests continue to your service

## Documentation

[See the documentation here](https://docs.datadome.co/docs/traefik)

## Support

- **DataDome Support**: Contact [support@datadome.co](mailto:support@datadome.co)
- **Issues**: Report issues on [GitHub](https://github.com/DataDome/module-traefik-package/issues)
