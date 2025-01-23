import { Configuration, AuthApi, TweetsApi } from "../api";

const getAuthToken = () => {
  if (typeof window !== "undefined") {
    const token = localStorage.getItem("token");
    return token ? `Bearer ${token}` : "";
  }
  return "";
};

const config = new Configuration({
  apiKey: getAuthToken,
});

export const authApi = new AuthApi(config);
export const tweetApi = new TweetsApi(config);
