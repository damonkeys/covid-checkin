// checkHTTPError checks a failes server-response and throws an error.
const checkHTTPError = (response) => {
  if (response.ok) {
    return response.json();
 } else {
    throw Error(response.status);
  }
};

export default checkHTTPError;
