// General error handler to log error details.
class ErrorHandler {
  // Instruct the error handler how to log errors
  constructor(logErrorDetails) {
    this.logErrorDetails = logErrorDetails;
  }

  // Logs response error details if isError is true.
  logError(isError, res, tags = {}) {
    if (!isError) return;

    // the Traceparent header is a W3C Trace Context
    const traceparentHeader = res.request.headers['Traceparent'];

    // Add any other useful information
    const errorData = Object.assign(
        {
          url: res.url,
          status: res.status,
          error_code: res.error_code,
          traceparent: traceparentHeader && traceparentHeader.toString(),
        },
        tags,
    );
    this.logErrorDetails(errorData);
  }
}

// Set up the error handler to log errors to the console
const errorHandler = new ErrorHandler((error) => {
  console.error(error);
});

export default errorHandler;