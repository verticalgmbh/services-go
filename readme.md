## Services

This package provides general structures used in vertical golang services like error handling and configuration

### Configuration

The following statement gets a configuration value from environment variables or returns the provided default value if the environment variable is not set
```
cfg.EnvConfig(<key>, <default>)
```

### Errors

**errors** contains two structures used to handle service errors. *ErrorResponse* is a structure which is returned by every service on any error. *APIError* wraps an
*ErrorResponse* and adds a statuscode which is written as http status.

Controllers should use

```
errors.WriteApiError(http.ResponseWriter,*APIError)
```

or any of the other *errors.WriteXXX* methods to return errors as json to the requesting client