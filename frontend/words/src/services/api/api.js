const request = async (url, method, body = null) => {
  const options = {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
  };

  if (body) {
    options.body = JSON.stringify(body);
  }

  try {
    const response = await fetch(`${url}`, options);

    if (!response.ok) {
      let errorData;
      try {
        errorData = await response.json();
      } catch (jsonError) {
        console.error(`Failed to parse error response: ${jsonError.message}`);
        return { success: false, error: `HTTP error! Status: ${response.status}, but failed to parse error message` };
      }
      return { success: false, error: `HTTP error! Status: ${response.status}, Message: ${errorData.message}` };
    }

    try {
      const data = await response.json();
      return { success: true, data };
    } catch (jsonError) {
      console.error(`Failed to parse success response: ${jsonError.message}`);
      return { success: false, error: `Request succeeded but failed to parse response: ${jsonError.message}` };
    }

  } catch (error) {
    console.error(`Request failed: ${error.message}`);
    return { success: false, error: `Request failed: ${error.message}` };
  }
};

export default request;
