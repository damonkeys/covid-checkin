// checkHTTPError checks a failes server-response and throws an error.
const checkHTTPError = (response): Promise<string> | Error => {
  if (response.ok) {
    return Promise.resolve(response.json());
  } else {
    throw Error(response.status);
  }
};

export default checkHTTPError;
