import { toast } from "react-toastify";

const request = async (url, method, body = null) => {
  const options = {
    method,
    headers: {
      "Content-Type": "application/json",
    },
  };

  if (body) {
    options.body = JSON.stringify(body);
  }

  try {
    const response = await fetch(url, options);
    if (!response.ok) {
      const errorMessage = await response.text();
      toast.error(errorMessage);
    }
    return response;
  } catch (error) {
    toast.error(`Request failed: ${error.message}`);
  }
};

export default request;
