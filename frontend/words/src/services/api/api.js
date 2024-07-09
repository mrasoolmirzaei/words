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
    const response = await fetch(`${url}`, options);
    return response;
  } catch (error) {
    toast.error(`Request failed: ${error.message}`);
    throw error;
  }
};

export default request;
